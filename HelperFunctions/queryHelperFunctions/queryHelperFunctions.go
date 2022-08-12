package queryHelperFunctions

import (
	"fmt"
	"log"
	"math/big"
	"multihop_gobot/global"
	"sync"

	FlashBotsUniswapQuery "multihop_gobot/contracts/Query"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)


func PrintReserves2(allPathsLength2 []global.PathsLength2) {
	// All Reserves should be updated by now - Print
	for _, elt := range allPathsLength2{
		fmt.Println("Path2, Pool1, Token1 Reserve : ",elt.Tuple1ReserveToken1)
		fmt.Println("Path2, Pool1, Token2 Reserve : ",elt.Tuple1ReserveToken2)
		fmt.Println("Path2, Pool2, Token1 Reserve : ",elt.Tuple2ReserveToken1)
		fmt.Println("Path2, Pool2, Token2 Reserve : ",elt.Tuple2ReserveToken2)
		fmt.Println("Path2, Lastest Update : ",elt.LastUpdate)
		fmt.Println("____")
		}

	}

func FilterPathsLength2(allPathsLength2 []global.PathsLength2, filteredPathsLength2 *[]global.PathsLength2){
	numPaths := len(allPathsLength2)
	for ii:=0; ii<numPaths; ii++ {
		var WETHReserve1 *big.Int
		var WETHReserve2 *big.Int

		// Get Reserves of WETH token - Since path of length 2, can just check both ends of the path
		if allPathsLength2[ii].Tuple1[0] == global.WETH_ADDRESS {
			WETHReserve1 = allPathsLength2[ii].Tuple1ReserveToken1
		} else {
			WETHReserve1 = allPathsLength2[ii].Tuple1ReserveToken2
		}

		if allPathsLength2[ii].Tuple2[0] == global.WETH_ADDRESS {
			WETHReserve2 = allPathsLength2[ii].Tuple2ReserveToken1
		} else {
			WETHReserve2 = allPathsLength2[ii].Tuple2ReserveToken2
		}

		// The pool has more WETH than the threshold limit
		if WETHReserve1.Cmp(global.FilterWETHThreshold) > 0  && WETHReserve2.Cmp(global.FilterWETHThreshold) > 0 {
			*filteredPathsLength2 = append(*filteredPathsLength2, allPathsLength2[ii])
		} else {
			//fmt.Println("Not Enough WETH in pool, ", WETHReserve1, WETHReserve2, global.FilterWETHThreshold)
			//fmt.Println(allPathsLength2[ii].ID)
		}
	}

	if len(*filteredPathsLength2) == 0 {
		fmt.Println("FilterPathsLength2 has length 0 - you have to pools to ARB!")
	}
}


func FilterProfitablePathsLength2(filteredPathsLength2 []global.PathsLength2, profitablePathsLength2 *[]global.PathsLength2) global.PathsLength2 {
	maxPnL := big.NewFloat(-1)
	maxPnlIND := 0
	numPaths := len(filteredPathsLength2)
	for ii:=0; ii<numPaths; ii++ {
		if filteredPathsLength2[ii].ArbDustPnl.Cmp(big.NewFloat( global.ArbPercentageThreshold ))==1 {
			*profitablePathsLength2 = append(*profitablePathsLength2, filteredPathsLength2[ii])
		}
		// Index to get best arb path
		if filteredPathsLength2[ii].ArbDustPnl.Cmp( maxPnL )==1 {
			maxPnlIND = ii
			maxPnL = filteredPathsLength2[ii].ArbDustPnl
		}
	}

	if len(*profitablePathsLength2) == 0 {
		fmt.Println("FilterPathsLength2 has length 0 - you have to pools to ARB!")
	}

	return filteredPathsLength2[maxPnlIND]
}


func UpdateReserves2(allPathsLength2 []global.PathsLength2, UniFlashQueryInstance *FlashBotsUniswapQuery.FlashBotUniswapQuery, curBlockNumber big.Int){

	// Get batch sizes
	numResBatch, numResBatchLeftOver := global.GetBatchSizes(len(allPathsLength2))
	
	// Query concurrently
	wgReserve := sync.WaitGroup{}

	for ii:=0; ii< numResBatch+1; ii++ {
		// Get the batch indices
		var startInd int
		var endInd int
		if numResBatch > ii{
			startInd = ii*global.NumPerBatch
			endInd = (ii+1)*global.NumPerBatch

		} else {
			startInd = ii*global.NumPerBatch
			endInd = ii*global.NumPerBatch + numResBatchLeftOver
		}
	
		// Update the reserves
		wgReserve.Add(1)
		go func(){
			QueryReserves2(startInd, endInd, UniFlashQueryInstance, allPathsLength2, &curBlockNumber)
			wgReserve.Done()
		}()
	}
	wgReserve.Wait()
	
}

func QueryReserves2(startInd int, endInd int, UniFlashQueryInstance *FlashBotsUniswapQuery.FlashBotUniswapQuery,
		 allPathsLength2 []global.PathsLength2,
		 curBlockNumber *big.Int){

	var callData []common.Address
	for ii:=startInd; ii<=endInd; ii++{
		callData = append(callData,allPathsLength2[ii].Tuple1[2])	
		callData = append(callData,allPathsLength2[ii].Tuple2[2])	
	}
	
	resp, errResp := UniFlashQueryInstance.GetReservesByPairs( &bind.CallOpts{}, callData)
	if errResp != nil {   fmt.Println("ERROR: Query Reserves Failed.")}

	p1Index := 0 
	p2Index := 1
	for ii:=startInd; ii<=endInd; ii++{
		allPathsLength2[ii].Tuple1ReserveToken1 = resp[p1Index][0]
		allPathsLength2[ii].Tuple1ReserveToken2 = resp[p1Index][1]
		allPathsLength2[ii].Tuple2ReserveToken1 = resp[p2Index][0]
		allPathsLength2[ii].Tuple2ReserveToken2 = resp[p2Index][1]
		allPathsLength2[ii].LastUpdate = curBlockNumber
		p1Index += 2
		p2Index += 2
	}


}


func GetAllPathsLength2FromStruct(startToken common.Address, allQueriedPoolsStruct []global.UniPool, allGroupedTokenPoolsStruct map[common.Address][]global.UniPool) ([]global.PathsLength2) {
	// Data structure to store results
	var allPathsLength2 []global.PathsLength2
	// We don't want to keep checking pools we've already accounted for. Effectively, a catepiller double-loop.
	alreadyCheckedThisTuple := map[string]bool{} // 
	for _, UniPool1 := range allGroupedTokenPoolsStruct[startToken]{
		
		// Mark this tuple as accounted for
		alreadyCheckedThisTuple[UniPool1.ID] = true

		// start with tokenStart1 (eg ETH), End with tokenEnd1 (eg nonWETH)
		tokenStart1, tokenEnd1, err := GetOrderedTokens(UniPool1.Token1, UniPool1.Token2, startToken)
		if err {fmt.Println("SERIOUS ERROR : Trying to sort tuple by no-existent token. getAllPathsLength2 - 1 ");break}

		for _, UniPool2 := range allGroupedTokenPoolsStruct[tokenEnd1] {

			// This tuple has already been accounted for - skip it
			if alreadyCheckedThisTuple[UniPool2.ID] {continue}

			// start with tokenStart2 (eg nonWETH, tokenEnd1), End with tokenEnd2 (eg WETH, tokenStart1)
			tokenStart2, tokenEnd2, err:= GetOrderedTokens(UniPool2.Token1, UniPool2.Token2, tokenEnd1)
			if err {fmt.Println("SERIOUS ERROR : Trying to sort tuple by no-existent token. getAllPathsLength2 - 2 ");break}


			// perform (double) check : the path should be consistent. A->B, B->A. Not A->B, C->A, etc
			if tokenStart1!=tokenEnd2 || tokenEnd1!=tokenStart2 { continue }
			if tokenStart1==tokenEnd2 && tokenEnd1==tokenStart2 { 
			} else {
				{fmt.Println("SERIOUS ERROR : PATH BROKEN getAllPathsLength2  ");break}
			}
			if UniPool1.Pool==UniPool2.Pool  {
				fmt.Println("SERIOUS ERROR : USING SAME POOL MORE THAN ONCE getAllPathsLength2  ");break
			}
			

			// create struct to hold the path
			var foundPath2 global.PathsLength2 
			//foundPath2.Tuple1 = tuple1
			//foundPath2.Tuple2 = tuple2
			foundPath2.Pool1 = UniPool1.Pool
			foundPath2.Pool2 = UniPool2.Pool
			foundPath2.Token1 = tokenStart1
			foundPath2.Token2 = tokenEnd1
			foundPath2.FeePool1 = big.NewInt(3000)
			foundPath2.FeePool2 = big.NewInt(3000)
			foundPath2.ID = UniPool1.ID + UniPool2.ID 
			foundPath2.LastUpdate = big.NewInt(-1)
			foundPath2.AmtToTrade = global.AmtDustIn
			//fmt.Println("Found a Path!!", tuple1, tuple2)

			allPathsLength2 = append(allPathsLength2, foundPath2)
		}
	}
	return allPathsLength2
}

func GetAllPathsLength2(startToken common.Address, queriedAddresses [][3]common.Address, groupByToken map[common.Address][][3]common.Address) ([]global.PathsLength2) {
	// Data structure to store results
	var allPathsLength2 []global.PathsLength2
	// We don't want to keep checking pools we've already accounted for. Effectively, a catepiller double-loop.
	alreadyCheckedThisTuple := map[[3]common.Address]bool{} // 
	for _, tuple1 := range groupByToken[startToken]{
		
		// Mark this tuple as accounted for
		alreadyCheckedThisTuple[tuple1] = true

		// start with tokenStart1 (eg ETH), End with tokenEnd1 (eg nonWETH)
		tokenStart1, tokenEnd1, err := GetOrderedTokens(tuple1[0], tuple1[1], startToken)
		if err {fmt.Println("SERIOUS ERROR : Trying to sort tuple by no-existent token. getAllPathsLength2 - 1 ");break}

		for _, tuple2 := range groupByToken[tokenEnd1] {

			// This tuple has already been accounted for - skip it
			if alreadyCheckedThisTuple[tuple2] {continue}

			// start with tokenStart2 (eg nonWETH, tokenEnd1), End with tokenEnd2 (eg WETH, tokenStart1)
			tokenStart2, tokenEnd2, err:= GetOrderedTokens(tuple2[0], tuple2[1], tokenEnd1)
			if err {fmt.Println("SERIOUS ERROR : Trying to sort tuple by no-existent token. getAllPathsLength2 - 2 ");break}


			// perform (double) check : the path should be consistent. A->B, B->A. Not A->B, C->A, etc
			if tokenStart1!=tokenEnd2 || tokenEnd1!=tokenStart2 { continue }
			if tokenStart1==tokenEnd2 && tokenEnd1==tokenStart2 { 
			} else {
				{fmt.Println("SERIOUS ERROR : PATH BROKEN getAllPathsLength2  ");break}
			}
			if tuple1[2]==tuple2[2]  {
				fmt.Println("SERIOUS ERROR : USING SAME POOL MORE THAN ONCE getAllPathsLength2  ");break
			}
			

			// create struct to hold the path
			var foundPath2 global.PathsLength2 
			foundPath2.Tuple1 = tuple1
			foundPath2.Tuple2 = tuple2
			foundPath2.Token1 = tokenStart1
			foundPath2.Token2 = tokenEnd1
			foundPath2.FeePool1 = big.NewInt(3000)
			foundPath2.FeePool2 = big.NewInt(3000)
			foundPath2.ID = tuple1[2].String() + tuple2[2].String()
			foundPath2.LastUpdate = big.NewInt(-1)
			foundPath2.AmtToTrade = global.AmtDustIn
			//fmt.Println("Found a Path!!", tuple1, tuple2)

			allPathsLength2 = append(allPathsLength2, foundPath2)
		}
	}
	return allPathsLength2
}

func GetAllPathsLength3(startToken common.Address, queriedAddresses [][3]common.Address, groupByToken map[common.Address][][3]common.Address) ([]global.PathsLength3) {
	
	// Data structure to store results
	var allPathsLength3 []global.PathsLength3

	// We don't want to keep checking pools we've already accounted for. Effectively, a catepiller double-loop. Twice.
	alreadyCheckedThisTuple1 := map[[3]common.Address]bool{}
	alreadyCheckedThisTuple2 := map[[3]common.Address]bool{}

	for _, tuple1 := range groupByToken[startToken]{
		
		// Mark this tuple as accounted for
		alreadyCheckedThisTuple1[tuple1] = true

		// start with tokenStart1 (eg ETH), End with tokenEnd1 (eg nonWETH)
		tokenStart1, tokenEnd1, err := GetOrderedTokens(tuple1[0], tuple1[1], startToken)
		if err {fmt.Println("SERIOUS ERROR : Trying to sort tuple by no-existent token. getAllPathsLength3 - 1 ");break}

		for _, tuple2 := range groupByToken[tokenEnd1] {

			// Mark this tuple as accounted for
			if alreadyCheckedThisTuple1[tuple2] {continue}

			// Skip this pool, since alread accounted for
			alreadyCheckedThisTuple2[tuple2] = true

			// start with tokenStart2 (eg nonWETH, tokenEnd1), End with tokenEnd2 (eg WETH, tokenStart1)
			tokenStart2, tokenEnd2, err:= GetOrderedTokens(tuple2[0], tuple2[1], tokenEnd1)
			if err {fmt.Println("SERIOUS ERROR : Trying to sort tuple by no-existent token. getAllPathsLength3 - 2 ");break}

			for _, tuple3 := range groupByToken[tokenEnd2] {

				// Skip this pool, since alread accounted for
				if alreadyCheckedThisTuple2[tuple3] {continue}

				// start with tokenStart2 (eg nonWETH, tokenEnd1), End with tokenEnd2 (eg WETH, tokenStart1)
				tokenStart3, tokenEnd3, err:= GetOrderedTokens(tuple3[0], tuple3[1], tokenEnd2)
				if err {fmt.Println("SERIOUS ERROR : Trying to sort tuple by no-existent token. getAllPathsLength3 - 3 ");break}

				// perform checks : the path should be consistent. A->B, B->A. Not A->B, C->A, etc
				if tokenStart1!=tokenEnd3 || tokenEnd1!=tokenStart2 || tokenEnd2!=tokenStart3 { continue }
				if tokenStart1==tokenEnd3 && tokenEnd1==tokenStart2 && tokenEnd2==tokenStart3 {
				}else {fmt.Println("SERIOUS ERROR : PATH BROKEN getAllPathsLength3  ");break}
				if tuple1[2]==tuple2[2] || tuple2[2]==tuple3[2] || tuple1[2]==tuple3[2] {
					fmt.Println("SERIOUS ERROR : USING SAME POOL MORE THAN ONCE getAllPathsLength3  ");break
				}

				// create struct to hold the path
				var foundPath3 global.PathsLength3
				foundPath3.Tuple1 = tuple1
				foundPath3.Tuple2 = tuple2
				foundPath3.Tuple3 = tuple3
				foundPath3.Token1 = tokenStart1
				foundPath3.Token2 = tokenEnd1
				foundPath3.Token3 = tokenEnd2
				foundPath3.Fee = big.NewInt(3000)
				foundPath3.ID = tuple1[2].String() + tuple2[2].String() + tuple3[2].String()
				foundPath3.LastUpdate = big.NewInt(-1)
				//fmt.Println("Found a 3Path!!", tuple1, tuple2, tuple3)
				allPathsLength3 = append(allPathsLength3, foundPath3)

			}
		}
	}
	return allPathsLength3
}






func PrintAll3Paths(allPathsLength3  []global.PathsLength3){
	fmt.Println("Number of paths with length 3 :",len(allPathsLength3), "\n   Start token:", global.StartAllPathsWithToken)
	for ii := range allPathsLength3 {
		fmt.Println(allPathsLength3[ii].Tuple1)
		fmt.Println(allPathsLength3[ii].Tuple2)
		fmt.Println(allPathsLength3[ii].Tuple3)
		fmt.Println("__")
	}
}


func PrintAll2Paths(allPathsLength2  []global.PathsLength2){
	fmt.Println("Number of paths with length 2 :",len(allPathsLength2), "\n   Start token:", global.StartAllPathsWithToken)
	for ii := range allPathsLength2 {
		fmt.Println(allPathsLength2[ii].Tuple1)
		fmt.Println(allPathsLength2[ii].Tuple2)
		fmt.Println("__")
	}
	
}
func GetOrderedTokens(token1 common.Address, token2 common.Address, tokenToPutFirst common.Address) (common.Address, common.Address, bool){

	if token1 == tokenToPutFirst {
		return token1, token2, false
	} else if token2 == tokenToPutFirst {
		return token2, token1, false
	} else {
		return token1, token2, true
	}


}

// Query the contract in batches
func QueryFactory(fAddress common.Address, UniFlashQueryInstance *FlashBotsUniswapQuery.FlashBotUniswapQuery, storeQueryAddresses chan<- [3]common.Address){
		for iter:=0; iter<global.NumBatches; iter++ {
			
			storeQueryAddressesTMP, err := UniFlashQueryInstance.GetPairsByIndexRange( &bind.CallOpts{}, fAddress, 
													big.NewInt(int64(iter*global.NumPerBatch)), big.NewInt(int64((iter+1)*global.NumPerBatch)) )			
			if err != nil {   fmt.Println("Error in Query", big.NewInt(int64(iter*global.NumPerBatch)), big.NewInt(int64((iter+1)*global.NumPerBatch)) ); log.Fatal(err)	}
	
			// Check WETH is one of the tokens
			for _, qTuple := range storeQueryAddressesTMP{
				storeQueryAddresses<-qTuple			
			}
		} 
}

func GetPool(path2 global.PathsLength2, poolNumber int) common.Address {
	if poolNumber==1{
		return path2.Tuple1[2]
	} else if poolNumber==2 {
		return path2.Tuple2[2]
	} else {
		panic("Bad parameters in GetPool")
	}
}