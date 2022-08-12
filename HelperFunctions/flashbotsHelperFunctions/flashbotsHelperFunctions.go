package flashbotsHelperFunctions

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	uniHelp "multihop_gobot/HelperFunctions/uniswapHelperFunctions"
	"multihop_gobot/PrivateData"
	"multihop_gobot/contracts/BundleExecutor"
	"multihop_gobot/contracts/UniswapV2Pool"
	"multihop_gobot/global"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/lmittmann/flashbots"
	"github.com/lmittmann/w3"
)


func GetPayloadTargets(UniswapV2PoolInst *UniswapV2Pool.UniswapV2Pool, bestArb global.PathsLength2) ([]common.Address,  [][]byte) {

	var flashbotTargets []common.Address
	var flashbotPayload [][]byte

	tokenIn := global.WETH_ADDRESS
	amountIn := global.AmtDustIn
	
	for i:=0; i<2;i++{
		tokenOut := GetComplementToken(tokenIn, bestArb)
		amountOut, _ := uniHelp.GetTokensOutPath2(bestArb, tokenIn, tokenOut, amountIn, i+1)
		targetPayload, err := uniHelp.SellTokensToNextMarket(UniswapV2PoolInst, tokenIn, tokenOut, amountIn, amountOut , bestArb , i+1);
		if err!=nil {log.Fatal(err); break}
		flashbotTargets = append(flashbotTargets, targetPayload.Target)
		flashbotPayload = append(flashbotPayload, targetPayload.Payload)
		tokenIn = tokenOut
		amountIn = amountOut
	}
	return flashbotTargets, flashbotPayload
	
}

func GetBundleExectorPayload(client *ethclient.Client, BundleExecutorInst *BundleExecutor.BundleExecutor, targets []common.Address, payloads [][]byte,
	 					  bestArb global.PathsLength2)  []byte {

	// This is the transaction payload 
	contractAbi, err := abi.JSON(strings.NewReader(string(BundleExecutor.BundleExecutorABI)))
	if err != nil {fmt.Println("Error in getting payload  in GetBundleExectorPayload");panic(err)}
	input, err := contractAbi.Pack("uniswapWeth",  bestArb.AmtToTrade, global.MinerReward, targets, payloads)
	if err != nil {fmt.Println("Error in packing payload  in GetBundleExectorPayload");panic(err)}
	return input

}

func SendBundleExecutionNonFlashBots(client *ethclient.Client, BundleExecutorInst *BundleExecutor.BundleExecutor, targets []common.Address, payloads [][]byte,
	bestArb global.PathsLength2) *types.Transaction {
	
	// NOTE: THIS ACTUALLY SENDS THE TRANSACTION 
	privateKey, err := crypto.HexToECDSA(global.PRIVATE_KEY)
	if err != nil {fmt.Println("Error in getting privateKey in GetBundleExectorResp");panic(err)}
	
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {fmt.Println("Error in getting gasPrice in GetBundleExectorResp");panic(err)}

	transactionOptions,err := bind.NewKeyedTransactorWithChainID(privateKey, global.ChainID)
	if err != nil { fmt.Println("Error in NewKeyedTransactorWithChainID()"); panic(err)    }
	transactionOptions.GasPrice = gasPrice
	transactionOptions.GasLimit = global.GasLimit
	//nonce, err := cl.NonceAt(ctx, addr, nil)

	tx, err := BundleExecutorInst.UniswapWeth(transactionOptions, bestArb.AmtToTrade, global.MinerReward, targets, payloads)
	if err!=nil {fmt.Println("Error in UniswapWeth()"); panic(err)}
	return tx
	
}

func BuildExecutorPayload(UniswapV2PoolInst *UniswapV2Pool.UniswapV2Pool, bestArb global.PathsLength2, targets []common.Address, payloads [][]byte) []byte {
		
	var executorPayload []byte

	fnSelectorBytes := []byte("uniswapWeth(uint256,uint256,address[],bytes[])")
	fnMethodID := crypto.Keccak256Hash(fnSelectorBytes)
	fnMethodIDHEX := fnMethodID.Hex()[2:10]
	fnMethodIDBytes := fnMethodID.Bytes()[0:4]
	fnMethodIDHexDecore,_ := hex.DecodeString(fnMethodIDHEX)
	fmt.Println("fnMethodIDHEX",fnMethodIDHEX)
	fmt.Println("fnMethodIDHEX",fnMethodIDBytes)
	fmt.Println("fnMethodIfnMethodIDHexDecoreDHEX",fnMethodIDHexDecore)
	
	amountWETHIn := bestArb.AmtToTrade
	ethAmountToCoinbase := global.MinerReward

	var param0Bytes = fnMethodIDBytes
	var param1Bytes = common.LeftPadBytes(amountWETHIn.Bytes(), 32)
	var param2Bytes = common.LeftPadBytes(ethAmountToCoinbase.Bytes(), 32)
	var param3Bytes []byte
	var param4Bytes []byte
	for i:=0; i<len(targets); i++ {
		target := common.LeftPadBytes(targets[i].Bytes(), 32)
		payload := common.LeftPadBytes(payloads[i], 32)
	 	param3Bytes = append(param3Bytes, target...)
		param4Bytes = append(param4Bytes, payload...)
	}

	executorPayload = append(executorPayload, param0Bytes...)
	executorPayload = append(executorPayload, param1Bytes...)
	executorPayload = append(executorPayload, param2Bytes...)
	executorPayload = append(executorPayload, param3Bytes...)
	executorPayload = append(executorPayload, param4Bytes...)

	return executorPayload
}

func GetFBBundles(client *ethclient.Client, executorPayloadBytes []byte, chainID *big.Int, gasPrice *big.Int) []global.FBBundle {
	// signer is a wallet object used to sign the arbitrage transactions....
	// fbBundle is {tx, signer_wallet_real_pk}

	// Build the TX we want to sign
	//nonce, err := client.PendingNonceAt(context.Background(), global.ADDRESS_BOT_WALLET)
	nonce, err := client.PendingNonceAt(context.Background(), GetAddressFromPrivateKey(PrivateData.PrivateKey))
	if err != nil {fmt.Println("Error in nonce");panic(err)}
	toAddress := global.ADDRESS_BUNDLE_EXECUTOR
	value := big.NewInt(0)
	gasLimit := uint64(21000 + 11000)
	data := executorPayloadBytes


	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// Get the privateKey to sign the transaction
	cred, err := crypto.HexToECDSA(global.PRIVATE_KEY)
    if err != nil {fmt.Println("Error in getting privateKey in GetFBBundles");panic(err)}

	// Get the signer (?) and the sign the Tx with it and the privateKey
	signer := types.LatestSignerForChainID(chainID)
    signedTx, err := types.SignTx(tx, signer, cred)
    if err != nil {fmt.Println("Error in signing Tx in GetFBBundles");  panic(err)}

	// Create the set of flashbots bundle (note: bundles = [bundle, bundle,..])
	var fbBundle1 global.FBBundle
	fbBundle1.Signer = signer
	fbBundle1.Transaction = signedTx
	
	// Create the flashbots bundles
	var fbBundles []global.FBBundle
	fbBundles = append(fbBundles, fbBundle1)

	return fbBundles

}

func GetAddressFromPrivateKey(privateKeyHex string) common.Address{

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
  		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
  		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return fromAddress
}


func PrintFBBundles(fbBundles []global.FBBundle){

	for i:=0; i<len(fbBundles); i++{
		fbBundle := fbBundles[i]
		fmt.Println("fbBundles[i].Signer", fbBundle.Signer)
		fmt.Println("fbBundles[i].Transaction", fbBundle.Transaction)
		fmt.Println("fbBundles[i].Transaction.Data()", fbBundle.Transaction.Data())
		fmt.Println("Hex(Data)", hex.EncodeToString(fbBundle.Transaction.Data()))
		fmt.Println("fbBundles[i].Transaction.Nonce()", fbBundle.Transaction.Nonce())
		fmt.Println("fbBundles[i].Transaction.To()", fbBundle.Transaction.To())
		fmt.Println("fbBundles[i].Transaction.Value()", fbBundle.Transaction.Value())
		fmt.Println("fbBundles[i].Transaction.Gas()", fbBundle.Transaction.Gas())
	}
	
}

func SimulateFBBundle(flashbotBundles []global.FBBundle){

	// Private key for request authentication
	//var FlashbotsPrivKey *ecdsa.PrivateKey
	FlashbotsPrivKey, _ := crypto.GenerateKey()

	// Connect to relay
	rpcClient, err := rpc.DialHTTPWithClient(
		"https://relay.flashbots.net",
		&http.Client{
			Transport: flashbots.AuthTransport(FlashbotsPrivKey),
		},
	)
	if err != nil {		fmt.Printf("Failed to connect to Flashbots relay: "); panic(err)	}

	client := w3.NewClient(rpcClient)
	defer client.Close()

	// Call/Simulate bundle
	var bundle types.Transactions // list of signed transactions
	var	callBundleResp flashbots.CallBundleResponse

	// append the signed transactions to the bundle
	bundle = append(bundle, flashbotBundles[0].Transaction)

	//CallBundle(r *CallBundleRequest) core.CallerFactory[CallBundleResponse]
	if err := client.Call(
		flashbots.CallBundle(&flashbots.CallBundleRequest{
			Transactions: bundle,
			BlockNumber:  big.NewInt(999_999_999),
		}).Returns(&callBundleResp),
	);

	// NOTE : returning just so I can time the overall execution time
	err != nil {		fmt.Println("Error in call bundle to Flashbots relay:",err); return }//panic(err)}
	fmt.Println("Called bundle successfully:", callBundleResp)


}


func GetComplementToken(tokenIn common.Address, path2 global.PathsLength2) common.Address {

	if tokenIn == path2.Token1 {
		return path2.Token2
	} else if tokenIn == path2.Token2 {
		return path2.Token1
	} else {
		panic("Bad parameters in GetComplementToken")
	}

}