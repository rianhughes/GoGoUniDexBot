package mainHelperFunctions

import (
	"context"
	"fmt"
	"log"
	"math/big"
	csvHelp "multihop_gobot/HelperFunctions/csvHelperFunctions"
	qryHelp "multihop_gobot/HelperFunctions/queryHelperFunctions"
	uniV3Help "multihop_gobot/HelperFunctions/uniV3HelperFunctions"
	"multihop_gobot/contracts/BundleExecutor"
	FlashBotUniswapQuery "multihop_gobot/contracts/Query"
	"multihop_gobot/contracts/UniswapV2Pool"
	uniV3Quoter "multihop_gobot/contracts/uniV3IQuoter"
	uniV3Pool "multihop_gobot/contracts/uniV3Pool"
	"multihop_gobot/global"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

////// CLIENTS
func GetClient()  *ethclient.Client {
	// connect to the client
	client, err := ethclient.Dial(global.ETHEREUM_CLIENT)
	if err != nil {log.Fatal(err); panic(err)}
	return client
}

func GetChainParams(client *ethclient.Client ) (*big.Int, *big.Int) {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {log.Fatal(err); panic(err)}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {log.Fatal(err); panic(err)}
	
	return chainID, gasPrice	
}

////// CONTRACTS
func GetUniFlashQueryInstance(client *ethclient.Client) *FlashBotUniswapQuery.FlashBotUniswapQuery {
	address := global.ADDRESS_UNISWAP_FLASH_QUERY
	UniFlashQueryInstance, err := FlashBotUniswapQuery.NewFlashBotUniswapQuery(address, client)
	if err != nil {		log.Fatal(err); panic(err)	}
	return UniFlashQueryInstance
}

func GetUniV2PoolInstance(client *ethclient.Client) *UniswapV2Pool.UniswapV2Pool {
	// create instance of UniV2-Pool contract
	UniswapV2PoolInst, err := UniswapV2Pool.NewUniswapV2Pool(global.WETH_ADDRESS, client)
	if err != nil {		log.Fatal(err); panic(err)	}
	return UniswapV2PoolInst
}

func GetExecutorInstance(client *ethclient.Client) *BundleExecutor.BundleExecutor{
	// Create instance of the transaction EXECUTOR contract
	addressBE := global.ADDRESS_BUNDLE_EXECUTOR
	BundleExecutorInst, err := BundleExecutor.NewBundleExecutor(addressBE, client)
	if err != nil {		log.Fatal(err); panic(err)	}
	return BundleExecutorInst
}


func GetQuoterInstance(client *ethclient.Client) *uniV3Quoter.IQuoter{
	// Create instance of the transaction EXECUTOR contract
	address := global.ADDRESS_UNIV3_QUOTER
	QuoterInst, err := uniV3Quoter.NewIQuoter(address, client)
	if err != nil {		log.Fatal(err); panic(err)	}
	return QuoterInst
}

func GetUniV3PoolInstance(client *ethclient.Client, poolAddress common.Address) *uniV3Pool.UniV3Pool{
	// Create instance of the transaction EXECUTOR contract
	UniV3PoolInst, err := uniV3Pool.NewUniV3Pool(poolAddress, client)
	if err != nil {		log.Fatal(err); panic(err)	}
	return UniV3PoolInst
}


////// Uniswap V2 forks - Pools and Paths

func GetFilteredPools2(UniFlashQueryInstance *FlashBotUniswapQuery.FlashBotUniswapQuery) []global.PathsLength2 {
	
	// pools -> paths => filtered paths

	// Query the raw pools, update the reserves, and filter
	if global.QueryOrLoadFilteredPools == "Query" {
		// Get all Raw Paths
		allPathsLength2, _ := GetAllPaths(UniFlashQueryInstance)
	
		// Filter the paths by using the reserves
		qryHelp.UpdateReserves2(allPathsLength2, UniFlashQueryInstance, *big.NewInt(0))
		filteredPathsLength2 := FilterPaths2(allPathsLength2)

		// Save the filtered path
		csvHelp.WriteToCSVFilteredPools(global.FilenameToSaveFilteredPool, filteredPathsLength2)

		return filteredPathsLength2

	// Load the pre-filtered data
	} else if global.QueryOrLoadFilteredPools == "Load" {
		filteredPathsLength2 := csvHelp.LoadSavedFilteredPaths(global.FilenameToSaveFilteredPool)
		return filteredPathsLength2

	} else {
		panic("Could not load filtered pools in GetFilteredPools2")
	}

}

func GetQueriedAddresses(action string, UniFlashQueryInstance *FlashBotUniswapQuery.FlashBotUniswapQuery) [][3]common.Address {
	if action=="Query"{
		queriedAddresses := QueryAndSaveAddresses(UniFlashQueryInstance)
		return queriedAddresses
	} else if action == "Load" {
		queriedAddresses := csvHelp.LoadSavedQueriedPools(global.FilenameToSaveRawPool)
		return queriedAddresses
	} else {
		panic("Incorrect action in GetQueriedAddresses. Can only Query or Load.")
	}
	
}


func QueryAndSaveAddresses(UniFlashQueryInstance *FlashBotUniswapQuery.FlashBotUniswapQuery) [][3]common.Address{

	// Get Factory addresses
	FACTORY_ADDRESSES := global.GetFactoryAddress()

	// create channel to hold the queried data (tuples) we want
	var storeQueryAddresses = make(chan [3]common.Address, global.NumBatches*global.NumPerBatch*len(FACTORY_ADDRESSES))

	// QUERY contract
	fmt.Println("Starting to query")
	wgFactory := sync.WaitGroup{}
	for _, fAddress := range FACTORY_ADDRESSES {
		wgFactory.Add(1)
		go func(fAddressTMP common.Address, storeQueryAddresses chan [3]common.Address) {
			qryHelp.QueryFactory(fAddressTMP, UniFlashQueryInstance, storeQueryAddresses)
			wgFactory.Done()
		}(fAddress, storeQueryAddresses)		
	}
	wgFactory.Wait()
	close(storeQueryAddresses)

	// Get all values that were added to the channel
	queriedAddresses := [][3]common.Address{}
	fmt.Println("Finished Query.")
	for qTuple := range storeQueryAddresses{
		queriedAddresses = append(queriedAddresses, qTuple)
	}

	// Save the data
	csvHelp.WriteToCSV(global.FilenameToSaveRawPool, queriedAddresses)

	return queriedAddresses
}

func GetDistinctTokens(queriedAddresses [][3]common.Address) map[common.Address]bool{
	// Get a list of the distinct tokens. queriedAddresses[0:1] need to be the tokens
	distinctTokens := make(map[common.Address]bool)
	for ii := range queriedAddresses{
		token1 := queriedAddresses[ii][0]
		token2 := queriedAddresses[ii][1]
		if !distinctTokens[token1] {distinctTokens[token1]=true}
		if !distinctTokens[token2] {distinctTokens[token2]=true}		
	}
	return distinctTokens
}



func GetDistinctTokensFormStructs(queriedAddresses []global.UniPool) map[common.Address]bool{
	// Get a list of the distinct tokens. queriedAddresses[0:1] need to be the tokens
	distinctTokens := make(map[common.Address]bool)
	for ii := range queriedAddresses{
		token1 := queriedAddresses[ii].Token1
		token2 := queriedAddresses[ii].Token2
		if !distinctTokens[token1] {distinctTokens[token1]=true}
		if !distinctTokens[token2] {distinctTokens[token2]=true}		
	}
	return distinctTokens
}


func GetDistinctTokensPools(distinctTokens map[common.Address]bool, queriedAddresses [][3]common.Address) map[common.Address][][3]common.Address {

	// Group by each distinct token. groupByToken[distinctToken] = array of tuples containing distinctToken
	groupByToken := map[common.Address][][3]common.Address{} 
	for key, _ := range distinctTokens{
		for _,elt := range queriedAddresses {		
			tokenKey, _, err := qryHelp.GetOrderedTokens(elt[0], elt[1], key)
			if err {continue}
			groupByToken[tokenKey] = append(groupByToken[tokenKey], elt )
		}
	}
	return groupByToken
}

func GetDistinctTokensPoolsFromStruct(distinctTokens map[common.Address]bool, allQueriedPoolsStruct []global.UniPool) map[common.Address][]global.UniPool  {

	// Group by each distinct token. groupByToken[distinctToken] = array of tuples containing distinctToken
	groupByToken := map[common.Address][]global.UniPool{}
	for distinctToken, _ := range distinctTokens{
		for _,elt := range allQueriedPoolsStruct {		
			tokenKey, _, err := qryHelp.GetOrderedTokens(elt.Token1, elt.Token2, distinctToken)
			if err {continue}//{fmt.Println("Here");fmt.Println(err); continue}
			groupByToken[tokenKey] = append(groupByToken[tokenKey], elt )
		}
	}
	return groupByToken
}

func GetAllPaths(UniFlashQueryInstance *FlashBotUniswapQuery.FlashBotUniswapQuery) ( []global.PathsLength2,  []global.PathsLength3){
	// First get the set of all pools. Then you need to get all paths through those pools.

	// List of all Univ2 (token0, token1, pool) addresses
	queriedAddresses := GetQueriedAddresses(global.QueryOrLoadRawPools, UniFlashQueryInstance)

	// Map of distinct tokens to bool 
	distinctTokens := GetDistinctTokens(queriedAddresses)

	// Map of distinct tokens to collection of pools involving that token
	groupByToken := GetDistinctTokensPools(distinctTokens, queriedAddresses)
	
	// Print some information about the quried data
	fmt.Println("Number of distinct tuples queried :",len(queriedAddresses))
	fmt.Println("Number of distinct tokens :",len(distinctTokens))
	fmt.Println("Number of pools involving WETH :",len(groupByToken[global.WETH_ADDRESS]))

	// Get all paths of length 2 - and print
	allPathsLength2 := qryHelp.GetAllPathsLength2(global.StartAllPathsWithToken, queriedAddresses, groupByToken)

	// Get all paths of length 3 - and print
	allPathsLength3 := qryHelp.GetAllPathsLength3(global.StartAllPathsWithToken, queriedAddresses, groupByToken)

	return allPathsLength2, allPathsLength3
}

func FilterPaths2(allPathsLength2 []global.PathsLength2) []global.PathsLength2{
	// Get filtered paths
	var filteredPathsLength2 []global.PathsLength2
	qryHelp.FilterPathsLength2(allPathsLength2, &filteredPathsLength2)

	// Get the set of pools from filteredPathsLength2
	var filteredTuples [][3]common.Address
	for i,_ := range filteredPathsLength2{
		filteredTuples = append(filteredTuples, filteredPathsLength2[i].Tuple1)
		filteredTuples = append(filteredTuples, filteredPathsLength2[i].Tuple2)
	}

	return filteredPathsLength2
}

////// Uniswap V3 forks - Pools and Paths

func QueryUniV3Structs(client *ethclient.Client) []global.UniV3Pool {

	var UniV3Structs []global.UniV3Pool

	// Load all the pool addresses
	poolAddresses := csvHelp.LoadUniV3PoolAddresses(global.FilenameToLoadUniV3Pools)

	// Fill in the structs
	for i, _ := range poolAddresses {
	//for i:=0; i < len(); i++ {
		poolAddress := poolAddresses[i]
		UniV3PoolInst:=GetUniV3PoolInstance(client, poolAddress)
		poolStruct := uniV3Help.GetUniV3PoolParams(client, poolAddress, UniV3PoolInst)
		UniV3Structs = append(UniV3Structs, poolStruct)
	}

	return UniV3Structs
}

func GetUniV3Structs(client *ethclient.Client, QueryOrLoad string) []global.UniV3Pool {
	if QueryOrLoad=="Query"{
		// Query for the structs
		uniV3PoolStructs := QueryUniV3Structs(client)
		// Save the structs
		csvHelp.SaveUniV3Structs(uniV3PoolStructs)
		return uniV3PoolStructs
	} else if QueryOrLoad=="Load" {
		uniV3PoolStructs := csvHelp.LoadUniV3Structs(global.FilenameToSaveUniV3Structs)
		return uniV3PoolStructs
	} else{
		panic("Improper QueryorLoad in GetUniV3Structs()")
	}
}


////// UniV2 and Uni V3
func GetPathsUniV2UniV3(filteredPathsLength2 []global.PathsLength2, uniV3PoolStructs []global.UniV3Pool)  []global.PathsLength2 {

	// Put all the pools into one [][3]common.Address variable 
	var allQueriedPoolsStruct []global.UniPool

	for i,_:= range filteredPathsLength2{
		var pool1 global.UniPool
		pool1.Token1 = filteredPathsLength2[i].Tuple1[0]
		pool1.Token2 = filteredPathsLength2[i].Tuple1[1]
		pool1.Pool = filteredPathsLength2[i].Tuple1[2]
		pool1.Fee = filteredPathsLength2[i].FeePool1
		pool1.ID = pool1.Token1.String() + pool1.Token2.String() + pool1.Pool.String()
		allQueriedPoolsStruct = append(allQueriedPoolsStruct, pool1)
		var pool2 global.UniPool
		pool2.Token1 = filteredPathsLength2[i].Tuple2[0]
		pool2.Token2 = filteredPathsLength2[i].Tuple2[1]
		pool2.Pool = filteredPathsLength2[i].Tuple2[2]
		pool2.Fee = filteredPathsLength2[i].FeePool2
		pool2.ID = pool2.Token1.String() + pool2.Token2.String() + pool2.Pool.String()
		allQueriedPoolsStruct = append(allQueriedPoolsStruct, pool2)
	}

	for i,_:= range uniV3PoolStructs{
		var pool1 global.UniPool
		pool1.Token1 = uniV3PoolStructs[i].Token1
		pool1.Token2 = uniV3PoolStructs[i].Token2
		pool1.Pool = uniV3PoolStructs[i].Pool
		pool1.Fee = uniV3PoolStructs[i].Fee
		pool1.ID = pool1.Token1.String() + pool1.Token2.String() + pool1.Pool.String()
		allQueriedPoolsStruct = append(allQueriedPoolsStruct, pool1)
	}

	// List of distinct token
	allDistinctTokens := GetDistinctTokensFormStructs(allQueriedPoolsStruct)

	// Group by distinct tokens.
	allGroupedTokenPoolsStruct := GetDistinctTokensPoolsFromStruct(allDistinctTokens, allQueriedPoolsStruct)

	// Get paths
	allPathsLength2 := qryHelp.GetAllPathsLength2FromStruct(global.StartAllPathsWithToken, allQueriedPoolsStruct, allGroupedTokenPoolsStruct)
	
	return allPathsLength2

}

/*

var tokens []common.Address
	var fees []*big.Int
	tokens = append(tokens, global.WETH_ADDRESS)
	fees = append(fees, big.NewInt(3000))
	tokens = append(tokens, global.USDC_ADDRESS)
	fees = append(fees, big.NewInt(500))
	tokens = append(tokens, global.WETH_ADDRESS)

	path := uniV3Help.EncodePath(tokens, fees)
	amountOut := uniV3Help.GetQuote(QuoterInst, path, big.NewInt(10000), chainID)
	fmt.Println("amountOut: ", amountOut)

	*/