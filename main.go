package main

import (
	"fmt"
	"math/big"
	fbHelp "multihop_gobot/HelperFunctions/flashbotsHelperFunctions"
	mainHelp "multihop_gobot/HelperFunctions/mainHelperFunctions"
	qryHelp "multihop_gobot/HelperFunctions/queryHelperFunctions"
	uniHelp "multihop_gobot/HelperFunctions/uniswapHelperFunctions"
	"multihop_gobot/global"
	"time"
)


func main() {

	start := time.Now()
	
	fmt.Println("=================== PART ONE : Find All Paths ===============================")

	// Client and Chain data
	client := mainHelp.GetClient()
 	chainID, gasPrice := mainHelp.GetChainParams(client)

	// Contract instances
	UniFlashQueryInstance := mainHelp.GetUniFlashQueryInstance(client)
	UniswapV2PoolInstance := mainHelp.GetUniV2PoolInstance(client)
	BundleExecutorInstance := mainHelp.GetExecutorInstance(client)
	//QuoterInst := mainHelp.GetQuoterInstance(client)

	// Get the filtered pools (Uniswap V2 Forks)
	filteredPathsLength2 := mainHelp.GetFilteredPools2(UniFlashQueryInstance)
	fmt.Println("Number Paths of Length 2 (Univ2 Forks):",len(filteredPathsLength2))

	// Get UniV3 data
	uniV3PoolStructs := mainHelp.GetUniV3Structs(client, global.QueryOrLoadUniV3Structs)

	// Find Paths between UniV2 and Univ3 pools
	allPathsLength2 := mainHelp.GetPathsUniV2UniV3(filteredPathsLength2, uniV3PoolStructs)
	fmt.Println("Number Paths of Length 2 (Univ2-Forks and Univ3):",len(allPathsLength2))

	elapsedStart := time.Since(start)
	mid := time.Now()

	fmt.Println("=================== PART TWO : Find Profitable Paths ===============================")
	/*
	INSERT CODE TO SUBSCRIBE TO NEW BLOCKS SO THAT WE UPDATE DATA EVERY BLOCK
	NOTE : THIS FEATURE ISN'T FREE. EITHER PAY FOR IT, OR RUN YOUR OWN NODE.
	newBlockChan := make(chan<- *types.Header)
	_, errNewBlock := client.SubscribeNewHead(context.Background(),  newBlockChan)
	*/
	curBlockNumber := *big.NewInt(1)

	// Update the reserves for paths of length 2
	qryHelp.UpdateReserves2(filteredPathsLength2, UniFlashQueryInstance, curBlockNumber)

	// Check for arbitrage amount the set of paths defined in part one
	uniHelp.GetArbAmountsForAllPaths2(filteredPathsLength2)	

	// Filter the paths that have a positive arbitrage value, and get best  arbitrage
	var profitablePathsLength2 []global.PathsLength2
	bestArb2 := qryHelp.FilterProfitablePathsLength2(filteredPathsLength2, &profitablePathsLength2)

	//uniHelp.PrintArbAmountsForAllPaths2(profitablePathsLength2)
	fmt.Println("Best Arbitrage Path \n", bestArb2)
	fmt.Println("Best Arbitrage Path PnL \n", bestArb2.ArbDustPnl)

	// Assuming we have an arbitrage path we want to profit on, we generate the payloads using the Flashbots contract
	fmt.Println("=================== PART THREE : Submit Profitable Paths To Flashbots ===============================")
	
	// Get the payloads needed for the Executor contract
	flashbotTargets, flashbotPayload := fbHelp.GetPayloadTargets(UniswapV2PoolInstance, bestArb2)

	// Get the payloads of the Executor contract - this goes in the flashbots bundle
	executorPayloadBytes := fbHelp.GetBundleExectorPayload(client, BundleExecutorInstance, flashbotTargets, flashbotPayload, bestArb2)
	
	// Build Transaction with parameters for the flashbotBundle
	flashbotBundles := fbHelp.GetFBBundles(client, executorPayloadBytes, chainID, gasPrice)

	// Sign and Simulate the flashbots bundle
	fbHelp.SimulateFBBundle(flashbotBundles)

	// Simulate, Sign, and Submit the bundle to flashbots
	// fbHelp.ExecuteFBBundle(flashbotBundles)


	fmt.Println("=================== TIME TO EXECUTE ===============================")
	elapsed2 := time.Since(mid)
	fmt.Println("Code PART ONE execution took", elapsedStart)
	fmt.Println("Code PART TWO and THREE execution took", elapsed2)

}
