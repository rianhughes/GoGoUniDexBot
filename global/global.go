package global

import (
	"math/big"
	"multihop_gobot/PrivateData"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

////////////////////////////////
/////////////// PARAMETERS
////////////////////////////////

var NumBatches = 10
var NumPerBatch = 200
const QueryOrLoadRawPools = "Load" // "Query", "Load"
const QueryOrLoadFilteredPools = "Load" // "Query", "Load"
const QueryOrLoadUniV3Structs = "Load" // "Query", "Load"

const FilenameToSaveRawPool = "Data/Raw_Pools.csv"
const FilenameToSaveFilteredPool = "Data/Filtered_Pools.csv"
const FilenameToLoadUniV3Pools = "Data/UniV3_Pools.csv"
const FilenameToSaveUniV3Structs = "Data/UniV3_Structs.csv"

var PRIVATE_KEY = PrivateData.PrivateKey
var ADDRESS_BUNDLE_EXECUTOR = PrivateData.ADDRESS_BUNDLE_EXECUTOR

var ArbPercentageThreshold = 0.0 
var ChainID = big.NewInt(1)
var GasLimit = uint64(21000 + 21000)
var GasPrice = uint64(1)
var MinerReward = big.NewInt(1000)

var StartAllPathsWithToken = WETH_ADDRESS
// NOTE: Be sure your units here correspond to the token you're trading!
var AmtDustIn = new(big.Int).Div(big.NewInt(params.Ether), big.NewInt(1))
var FilterWETHThreshold =big.NewInt(params.Ether)

////////////////////////////////
/////////////// CLIENT, ADDRESSES ETC
////////////////////////////////


// Mainnet
var ETHEREUM_CLIENT = "https://rpc.ankr.com/eth"

var ADDRESS_FB_SIGN_WALLET = common.HexToAddress("0x") // NEED TO UPDATE THIS
var ADDRESS_BOT_WALLET = common.HexToAddress("0x")	   // NEED TO UPDATE THIS

var ADDRESS_UNIV3_QUOTER = common.HexToAddress("0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6") 
var ADDRESS_UNISWAP_FLASH_QUERY = common.HexToAddress("0x5EF1009b9FCD4fec3094a5564047e190D72Bd511")

var WETH_ADDRESS = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
var USDC_ADDRESS = common.HexToAddress("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")


var FLASHBOTSMULTICALL_ABI = `[{"inputs":[{"internalType":"address","name":"_executor","type":"address"}],"stateMutability":"payable","type":"constructor"},{"inputs":[{"internalType":"address payable","name":"_to","type":"address"},{"internalType":"uint256","name":"_value","type":"uint256"},{"internalType":"bytes","name":"_data","type":"bytes"}],"name":"call","outputs":[{"internalType":"bytes","name":"","type":"bytes"}],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"uint256","name":"_wethAmountToFirstMarket","type":"uint256"},{"internalType":"uint256","name":"_ethAmountToCoinbase","type":"uint256"},{"internalType":"address[]","name":"_targets","type":"address[]"},{"internalType":"bytes[]","name":"_payloads","type":"bytes[]"}],"name":"uniswapWeth","outputs":[],"stateMutability":"payable","type":"function"},{"stateMutability":"payable","type":"receive"}]`

func GetFactoryAddress() [2]common.Address{
	var FACTORY_ADDRESS_UNISWAP = common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")		// 77708
	var FACTORY_ADDRESS_SUSHISWAP = common.HexToAddress("0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac")	// 2939

	var FACTORY_ADDRESSES [2]common.Address

	FACTORY_ADDRESSES[0] = FACTORY_ADDRESS_UNISWAP
	FACTORY_ADDRESSES[1] = FACTORY_ADDRESS_SUSHISWAP
	return FACTORY_ADDRESSES
}

func GetFactoryAddress_Goerli() [2]common.Address{
	var FACTORY_ADDRESS_UNISWAP = common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")		// 9408
	var FACTORY_ADDRESS_SUSHISWAP = common.HexToAddress("0xc35dadb65012ec5796536bd9864ed8773abc74c4")	// 900

	var FACTORY_ADDRESSES [2]common.Address

	FACTORY_ADDRESSES[0] = FACTORY_ADDRESS_UNISWAP
	FACTORY_ADDRESSES[1] = FACTORY_ADDRESS_SUSHISWAP
	return FACTORY_ADDRESSES

}




func GetBatchSizes(lengthAllPaths int) (int,int){
	// Index math. Want to batch by size NumPerBatch, but lenAllPathsLength2%NumPerBatch might not be zero.
	lenAllPathsLength2 := lengthAllPaths
	numResBatch := int(lenAllPathsLength2/NumPerBatch)
	numResBatchLeftOver := lenAllPathsLength2%NumPerBatch -1 // Start from index 0
	return numResBatch, numResBatchLeftOver
}

////////////////////////////////
/////////////// STRUCTS
////////////////////////////////

type FBBundle struct {
	Signer	types.Signer
	Transaction *types.Transaction 
}

// Struct for the {target, payload} needed for FB bundles
type TargetPayload struct {
	Target	common.Address
	Payload []byte
}

type UniV3Pool struct {
	Fee *big.Int
	Token1 common.Address
	Token2 common.Address
	Pool common.Address
}

type UniPool struct {
	ID string // pool1+pool2
	Fee *big.Int
	Version string // "V2", "V3"
	Token1 common.Address
	Token2 common.Address
	Pool common.Address
}


// Struct for Paths of Length 2
type PathsLength2 struct {
	ID string // pool1+pool2
	FeePool1 *big.Int
	FeePool2 *big.Int
	Version string // V2, V3
	Tuple1 [3]common.Address
	Tuple2 [3]common.Address
	Pool1 common.Address
	Pool2 common.Address
	Token1 common.Address
	Token2 common.Address
	Tuple1ReserveToken1 *big.Int 
	Tuple1ReserveToken2 *big.Int
	Tuple2ReserveToken1 *big.Int
	Tuple2ReserveToken2 *big.Int
	LastUpdate *big.Int
	AmtToTrade *big.Int
	AmtDustMiddle *big.Int
	AmtDustOut *big.Int
	ArbDustPnl *big.Float
}

// Struct for Paths of Length 3
type PathsLength3 struct {
	ID string // pool1+pool2+pool3. Note: no need to include fee here as it's implicit in the pool
	Fee *big.Int
	Version string // V2, V3
	Tuple1 [3]common.Address
	Tuple2 [3]common.Address
	Tuple3 [3]common.Address
	Pool1 common.Address
	Pool2 common.Address
	Pool3 common.Address
	Token1 common.Address
	Token2 common.Address
	Token3 common.Address
	Tuple1ReserveToken1 *big.Int 
	Tuple1ReserveToken2 *big.Int
	Tuple2ReserveToken1 *big.Int
	Tuple2ReserveToken2 *big.Int
	Tuple3ReserveToken1 *big.Int
	Tuple3ReserveToken2 *big.Int
	LastUpdate *big.Int
	ArbAmtOnDust *big.Int
}



