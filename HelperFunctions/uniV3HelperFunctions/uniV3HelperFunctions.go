package uniV3HelperFunctions

import (
	"fmt"
	"math/big"
	IQuoter "multihop_gobot/contracts/uniV3IQuoter"
	"multihop_gobot/contracts/uniV3Pool"
	"multihop_gobot/global"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// The pools are deinfed by (token0, token1, fee)
// 1. Get a list of all pools / tokens : https://info.uniswap.org/#/pools
// 2. Call the Quoter (uniswapV3Callback) to find profitable opportunities -> Mainnet fork?
// 3. Call swap on the pool address - if profitable : https://docs.uniswap.org/protocol/reference/core/UniswapV3Pool#swap
// 4. Flashswaps?

// Note on paths : https://docs.uniswap.org/protocol/guides/swaps/multihop-swaps
//				 : https://soliditydeveloper.com/uniswap3
// uniV3Poolinstances.Token0, uniV3Poolinstances.Token1

// path is token - fee - token , token - fee - token ,..

// QuoterInstance.uniswapV3SwapCallBack(amount0In, amount1In, path)


func GetUniV3PoolParams(client *ethclient.Client, poolAddress common.Address, UniV3PoolInst *uniV3Pool.UniV3Pool) global.UniV3Pool{

	// Data structure to hold that tasty tast data
	var v3Pools global.UniV3Pool
	
	// Get the pools paramters
	token0, err := UniV3PoolInst.Token0(&bind.CallOpts{})
	if err!=nil{ fmt.Println("Failed to get Token0 in GetUniV3PoolParams()"); panic(err)}
	token1, err := UniV3PoolInst.Token1(&bind.CallOpts{})
	if err!=nil{ fmt.Println("Failed to get Token1 in GetUniV3PoolParams()"); panic(err)}
	fee, err := UniV3PoolInst.Fee(&bind.CallOpts{})
	if err!=nil{ fmt.Println("Failed to get Fee in GetUniV3PoolParams()"); panic(err)}

	v3Pools.Token1 = token0
	v3Pools.Token2 = token1
	v3Pools.Pool = poolAddress
	v3Pools.Fee = fee

	return v3Pools
}


func GetQuote(QuoterInst *IQuoter.IQuoter, path []byte, amountIn *big.Int, chainID *big.Int) *types.Transaction {
	privateKey, err := crypto.HexToECDSA(global.PRIVATE_KEY)
	if err!=nil{ fmt.Println("Failed to get Private Key in GetQuote()"); panic(err)}

	TransactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err!=nil{ fmt.Println("Failed to get the Transaction Options in GetQuote()"); panic(err)}
	TransactOpts.GasLimit = 21000 + 21000

	// This is going to actually transact, then revert, but still cost gas?
	tx, err := QuoterInst.QuoteExactInput( TransactOpts ,path, amountIn )
	if err!=nil{ fmt.Println("Failed to get the quote in GetQuote()"); panic(err)}
	return tx
}

func EncodePath(tokens []common.Address, fees []*big.Int) []byte {
	// 20byte for token address
	// 3 byte for fee
	var path []byte

	numElts := len(tokens)
	for i:=0; i<numElts-1; i++ {
		feeBytes := fees[i].Bytes()
		feeBytes = common.LeftPadBytes(feeBytes, 3)
		path = append(path, tokens[i].Bytes()...)
		path = append(path, feeBytes...	)
	}

	path = append(path, tokens[numElts-1].Bytes()...)
	
	numBytes := 20*len(tokens) + 3*len(fees)

	if len(path)==numBytes{
		return path
	} else {
		fmt.Println(len(path))
		panic("length of path is not correct, ")
	}
}


func EncodePathSimple(token0 common.Address, fee *big.Int, token1 common.Address) []byte {
	// 20byte for token address
	// 3 byte for fee
	var path []byte
	token0Bytes := token0.Bytes()
	token1Bytes := token1.Bytes()
	feeBytes := fee.Bytes()
	feeBytes = common.LeftPadBytes(feeBytes, 3)	
	path = append(path, token0Bytes...)
	path = append(path, feeBytes...)
	path = append(path, token1Bytes...)
	if len(path)==43{
		return path
	} else {
		panic("length of path is not correct")
	}
}

