package uniswapHelperFunctions

import (
	"fmt"
	"log"
	"math/big"
	"multihop_gobot/global"

	qryHelp "multihop_gobot/HelperFunctions/queryHelperFunctions"
	"multihop_gobot/contracts/UniswapV2Pool"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

func GetArbAmountsForAllPaths2(allPathsLength2 []global.PathsLength2) {
	// Checking if dust is profitable.
	numAllPathsLength2 := len(allPathsLength2)
	amountToTrade := global.AmtDustIn
	for ii:=0; ii<numAllPathsLength2;ii++ {
		amount2Out, amount1Out, err := GetArbAmountForPath2(allPathsLength2[ii], amountToTrade)
		if err != nil {log.Fatal(err)}
		// Get PnL %
		pnl, err := GetPnl(amountToTrade, amount2Out)
		if err != nil { log.Fatal(err)}
		// Save that tasty data
		allPathsLength2[ii].AmtDustOut = amount2Out
		allPathsLength2[ii].AmtDustMiddle = amount1Out
		allPathsLength2[ii].ArbDustPnl = pnl
		allPathsLength2[ii].AmtToTrade = amountToTrade
	}

}

func PrintArbAmountsForAllPaths2(allPathsLength2 []global.PathsLength2) {
	fmt.Println("Finished Checking For Arbs")
	for _,paths := range allPathsLength2 {
		fmt.Println("ID :",paths.ID)
		fmt.Println("AmtDustIn :",global.AmtDustIn)
		fmt.Println("AmtDustOut :",paths.AmtDustOut)
		fmt.Println("ArbDustPnl :",paths.ArbDustPnl)
		fmt.Println("Tuple1ReserveToken1 :",paths.Tuple1ReserveToken1)
		fmt.Println("Tuple1ReserveToken2 :",paths.Tuple1ReserveToken2)
		fmt.Println("Tuple2ReserveToken1 :",paths.Tuple2ReserveToken1)
		fmt.Println("Tuple2ReserveToken2 :",paths.Tuple2ReserveToken2)
		fmt.Println("--------")
	}

}


func GetArbAmountForPath2(path2 global.PathsLength2, amountToTrade *big.Int) (*big.Int, *big.Int, error){
	// This function calculates how much you would expect to get out from the
	// entire path if you trade amountToTrade.
	// Note : Neglects gas. Assumes a 0.3% fee.

	// Price Token1 -> Token2. Pool1.
	tokenPool1In := path2.Token1
	tokenPool1Out := path2.Token2
	amount1In := amountToTrade
	amount1Out, err := GetTokensOutPath2(path2, tokenPool1In , tokenPool1Out , amount1In , 1)
	if err != nil {
		log.Fatal(err)
		return big.NewInt(0), big.NewInt(0), err
	}
		
	// Price Token2 -> Token1. Pool2.
	tokenPool2In := tokenPool1Out
	tokenPool2Out := tokenPool1In
	amount2In := amount1Out
	amount2Out, err := GetTokensOutPath2(path2, tokenPool2In , tokenPool2Out , amount2In , 2)
	if err != nil {
		log.Fatal(err)
		return big.NewInt(0), big.NewInt(0), err
	}	
	return amount2Out, amount1Out, nil
}



func GetPnl(amountIn *big.Int, amountOut *big.Int) (*big.Float, error){
	var Profit *big.Float = big.NewFloat(0)
	Profit.Quo(new(big.Float).SetInt(amountOut),  new(big.Float).SetInt(amountIn))
	Profit.Sub(Profit, big.NewFloat(1))
	return Profit, nil
}

func GetTokensOutPath2(pathLength2 global.PathsLength2, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, poolNumber int) (*big.Int, error) {
	var reserveTokenIn *big.Int
	var reserveTokenOut *big.Int

	if poolNumber==1 {
		// Query the first pool in the path
		if tokenIn == pathLength2.Tuple1[0] {
			reserveTokenIn = pathLength2.Tuple1ReserveToken1
			reserveTokenOut = pathLength2.Tuple1ReserveToken2
		} else if tokenIn == pathLength2.Tuple1[1] {
			reserveTokenIn = pathLength2.Tuple1ReserveToken2
			reserveTokenOut = pathLength2.Tuple1ReserveToken1
		} else {
			panic("ERROR : GetTokensIn mismatch in parameters of tokenIn, tokenOut, pool1,..")
			//log.Fatal("ERROR : GetTokensIn mismatch in parameters of tokenIn, tokenOut, pool1,..")
			//return big.NewInt(0), errors.New("ERROR: Mismatch between pool and tokens in GetTokensInPath2")
		}
	} else if poolNumber==2{
		// Query the second pool in the path
		if tokenIn == pathLength2.Tuple2[0] {
			reserveTokenIn = pathLength2.Tuple2ReserveToken1
			reserveTokenOut = pathLength2.Tuple2ReserveToken2
		} else if tokenIn == pathLength2.Tuple2[1] {
			reserveTokenIn = pathLength2.Tuple2ReserveToken2
			reserveTokenOut = pathLength2.Tuple2ReserveToken1
		} else {
			panic("ERROR : GetTokensIn mismatch in parameters of tokenIn, tokenOut, pool1,..")
			//log.Fatal("ERROR : GetTokensIn mismatch in parameters of tokenIn, tokenOut, pool2,..")
			//return big.NewInt(0), errors.New("ERROR: Mismatch between pool and tokens in GetTokensInPath2")
		}
	}

	amountInWithFee := new(big.Int).Mul(amountIn, big.NewInt(997))
    numerator := new(big.Int).Mul(amountInWithFee, reserveTokenOut);
    denominator := new(big.Int).Mul(reserveTokenIn, big.NewInt(1000))
	denominator.Add(denominator, amountInWithFee);
	return new(big.Int).Div(numerator, denominator), nil

  }


  ////////////////////

func SellTokensToNextMarket(UniswapV2PoolInst *UniswapV2Pool.UniswapV2Pool, tokenIn common.Address,  tokenOut common.Address, amountIn *big.Int, amountOut *big.Int, pathLength2 global.PathsLength2, poolNumber int) (global.TargetPayload, error) {
	
	var targetPayload global.TargetPayload
	var err error

	amount1Out := big.NewInt(0)
	amount2Out := big.NewInt(0)
	recipient := global.ADDRESS_BOT_WALLET

	if tokenIn == pathLength2.Token1 {
		amount2Out, err = GetTokensOutPath2(pathLength2, tokenIn, tokenOut, amountIn , poolNumber)
	} else if tokenIn == pathLength2.Token2 {
		amount1Out, err = GetTokensOutPath2(pathLength2, tokenIn, tokenOut, amountIn , poolNumber)
	} else {
		panic("invalid token in SellTokensToNextMarket")
	}
	if err!=nil {panic("Error in SellTokensToNextMarket(..)")}

	// Generate the transaction payload from the interface?
	payload := buildTxSwapPayloadBytes(amount1Out, amount2Out, recipient)

	// create struct to hold & return data
	targetPayload.Payload = payload
	targetPayload.Target = qryHelp.GetPool(pathLength2, poolNumber)

	return targetPayload, nil
}


func buildTxSwapPayloadBytes(amount1Out *big.Int, amount2Out *big.Int, recipient common.Address ) []byte {
	
	var data []byte
	var emptyData []byte

	funcSignature := []byte("swap(uint256,uint256,address,bytes)")
	
	hash := sha3.NewLegacyKeccak256()
    hash.Write(funcSignature)
    methodID := hash.Sum(nil)[:4]
	var param1Bytes = common.LeftPadBytes(amount1Out.Bytes(), 32)
	var param2Bytes = common.LeftPadBytes(amount2Out.Bytes(), 32)
	var param3Bytes = common.LeftPadBytes(recipient.Bytes(), 32)
	var param4Bytes = common.LeftPadBytes(emptyData, 32)

	data = append(data, methodID...)
	data = append(data, param1Bytes...)
	data = append(data, param2Bytes...)
	data = append(data, param3Bytes...)
	data = append(data, param4Bytes...)

	return data
	
	/*
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
    signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
    if err != nil {log.Fatal(err)}
    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {log.Fatal(err)}
    fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) 
	*/

}
