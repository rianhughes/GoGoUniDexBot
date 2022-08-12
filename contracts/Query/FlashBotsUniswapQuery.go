// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FlashBotUniswapQuery

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FlashBotUniswapQueryMetaData contains all meta data concerning the FlashBotUniswapQuery contract.
var FlashBotUniswapQueryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractUniswapV2Factory\",\"name\":\"_uniswapFactory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stop\",\"type\":\"uint256\"}],\"name\":\"getPairsByIndexRange\",\"outputs\":[{\"internalType\":\"address[3][]\",\"name\":\"\",\"type\":\"address[3][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIUniswapV2Pair[]\",\"name\":\"_pairs\",\"type\":\"address[]\"}],\"name\":\"getReservesByPairs\",\"outputs\":[{\"internalType\":\"uint256[3][]\",\"name\":\"\",\"type\":\"uint256[3][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610da7806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80634dbf0f391461003b578063ab2217e41461006b575b600080fd5b61005560048036038101906100509190610795565b61009b565b6040516100629190610b23565b60405180910390f35b61008560048036038101906100809190610803565b610263565b6040516100929190610b01565b60405180910390f35b6060808383905067ffffffffffffffff811180156100b857600080fd5b506040519080825280602002602001820160405280156100f257816020015b6100df61064b565b8152602001906001900390816100d75790505b50905060005b848490508110156102585784848281811061010f57fe5b905060200201602081019061012491906107da565b73ffffffffffffffffffffffffffffffffffffffff16630902f1ac6040518163ffffffff1660e01b815260040160606040518083038186803b15801561016957600080fd5b505afa15801561017d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101a19190610852565b826dffffffffffffffffffffffffffff169250816dffffffffffffffffffffffffffff1691508063ffffffff1690508484815181106101dc57fe5b60200260200101516000600381106101f057fe5b6020020185858151811061020057fe5b602002602001015160016003811061021457fe5b6020020186868151811061022457fe5b602002602001015160026003811061023857fe5b6020020183815250838152508381525050505080806001019150506100f8565b508091505092915050565b606060008473ffffffffffffffffffffffffffffffffffffffff1663574f2ba36040518163ffffffff1660e01b815260040160206040518083038186803b1580156102ad57600080fd5b505afa1580156102c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102e591906108a1565b9050808311156102f3578092505b83831015610336576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032d90610b45565b60405180910390fd5b6000848403905060608167ffffffffffffffff8111801561035657600080fd5b5060405190808252806020026020018201604052801561039057816020015b61037d61066d565b8152602001906001900390816103755790505b50905060005b8281101561063d5760008873ffffffffffffffffffffffffffffffffffffffff16631e3dd18b838a016040518263ffffffff1660e01b81526004016103db9190610b65565b60206040518083038186803b1580156103f357600080fd5b505afa158015610407573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061042b919061076c565b90508073ffffffffffffffffffffffffffffffffffffffff16630dfe16816040518163ffffffff1660e01b815260040160206040518083038186803b15801561047357600080fd5b505afa158015610487573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ab919061076c565b8383815181106104b757fe5b60200260200101516000600381106104cb57fe5b602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1663d21220a76040518163ffffffff1660e01b815260040160206040518083038186803b15801561054857600080fd5b505afa15801561055c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610580919061076c565b83838151811061058c57fe5b60200260200101516001600381106105a057fe5b602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050808383815181106105e457fe5b60200260200101516002600381106105f857fe5b602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050508080600101915050610396565b508093505050509392505050565b6040518060600160405280600390602082028036833780820191505090505090565b6040518060600160405280600390602082028036833780820191505090505090565b60008151905061069e81610ce7565b92915050565b60008083601f8401126106b657600080fd5b8235905067ffffffffffffffff8111156106cf57600080fd5b6020830191508360208202830111156106e757600080fd5b9250929050565b6000813590506106fd81610cfe565b92915050565b60008135905061071281610d15565b92915050565b60008151905061072781610d2c565b92915050565b60008135905061073c81610d43565b92915050565b60008151905061075181610d43565b92915050565b60008151905061076681610d5a565b92915050565b60006020828403121561077e57600080fd5b600061078c8482850161068f565b91505092915050565b600080602083850312156107a857600080fd5b600083013567ffffffffffffffff8111156107c257600080fd5b6107ce858286016106a4565b92509250509250929050565b6000602082840312156107ec57600080fd5b60006107fa848285016106ee565b91505092915050565b60008060006060848603121561081857600080fd5b600061082686828701610703565b93505060206108378682870161072d565b92505060406108488682870161072d565b9150509250925092565b60008060006060848603121561086757600080fd5b600061087586828701610718565b935050602061088686828701610718565b925050604061089786828701610757565b9150509250925092565b6000602082840312156108b357600080fd5b60006108c184828501610742565b91505092915050565b60006108d6838361092a565b60208301905092915050565b60006108ee8383610939565b60608301905092915050565b60006109068383610a4c565b60608301905092915050565b600061091e8383610ae3565b60208301905092915050565b61093381610c5d565b82525050565b61094281610bb4565b61094c8184610c14565b925061095782610b80565b8060005b8381101561098857815161096f87826108ca565b965061097a83610be0565b92505060018101905061095b565b505050505050565b600061099b82610bbf565b6109a58185610c1f565b93506109b083610b8a565b8060005b838110156109e15781516109c888826108e2565b97506109d383610bed565b9250506001810190506109b4565b5085935050505092915050565b60006109f982610bca565b610a038185610c30565b9350610a0e83610b9a565b8060005b83811015610a3f578151610a2688826108fa565b9750610a3183610bfa565b925050600181019050610a12565b5085935050505092915050565b610a5581610bd5565b610a5f8184610c41565b9250610a6a82610baa565b8060005b83811015610a9b578151610a828782610912565b9650610a8d83610c07565b925050600181019050610a6e565b505050505050565b6000610ab0602083610c4c565b91507f73746172742063616e6e6f7420626520686967686572207468616e2073746f706000830152602082019050919050565b610aec81610ccd565b82525050565b610afb81610ccd565b82525050565b60006020820190508181036000830152610b1b8184610990565b905092915050565b60006020820190508181036000830152610b3d81846109ee565b905092915050565b60006020820190508181036000830152610b5e81610aa3565b9050919050565b6000602082019050610b7a6000830184610af2565b92915050565b6000819050919050565b6000819050602082019050919050565b6000819050602082019050919050565b6000819050919050565b600060039050919050565b600081519050919050565b600081519050919050565b600060039050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b6000602082019050919050565b600081905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000610c6882610cad565b9050919050565b6000610c7a82610c5d565b9050919050565b6000610c8c82610c5d565b9050919050565b60006dffffffffffffffffffffffffffff82169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600063ffffffff82169050919050565b610cf081610c5d565b8114610cfb57600080fd5b50565b610d0781610c6f565b8114610d1257600080fd5b50565b610d1e81610c81565b8114610d2957600080fd5b50565b610d3581610c93565b8114610d4057600080fd5b50565b610d4c81610ccd565b8114610d5757600080fd5b50565b610d6381610cd7565b8114610d6e57600080fd5b5056fea264697066735822122069e8a69e8ff8f31a80b87a9ee1e6da39b49723342a0439c480906696fc7130db64736f6c634300060c0033",
}

// FlashBotUniswapQueryABI is the input ABI used to generate the binding from.
// Deprecated: Use FlashBotUniswapQueryMetaData.ABI instead.
var FlashBotUniswapQueryABI = FlashBotUniswapQueryMetaData.ABI

// FlashBotUniswapQueryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FlashBotUniswapQueryMetaData.Bin instead.
var FlashBotUniswapQueryBin = FlashBotUniswapQueryMetaData.Bin

// DeployFlashBotUniswapQuery deploys a new Ethereum contract, binding an instance of FlashBotUniswapQuery to it.
func DeployFlashBotUniswapQuery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FlashBotUniswapQuery, error) {
	parsed, err := FlashBotUniswapQueryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FlashBotUniswapQueryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FlashBotUniswapQuery{FlashBotUniswapQueryCaller: FlashBotUniswapQueryCaller{contract: contract}, FlashBotUniswapQueryTransactor: FlashBotUniswapQueryTransactor{contract: contract}, FlashBotUniswapQueryFilterer: FlashBotUniswapQueryFilterer{contract: contract}}, nil
}

// FlashBotUniswapQuery is an auto generated Go binding around an Ethereum contract.
type FlashBotUniswapQuery struct {
	FlashBotUniswapQueryCaller     // Read-only binding to the contract
	FlashBotUniswapQueryTransactor // Write-only binding to the contract
	FlashBotUniswapQueryFilterer   // Log filterer for contract events
}

// FlashBotUniswapQueryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlashBotUniswapQueryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashBotUniswapQueryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlashBotUniswapQueryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashBotUniswapQueryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlashBotUniswapQueryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashBotUniswapQuerySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlashBotUniswapQuerySession struct {
	Contract     *FlashBotUniswapQuery // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FlashBotUniswapQueryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlashBotUniswapQueryCallerSession struct {
	Contract *FlashBotUniswapQueryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// FlashBotUniswapQueryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlashBotUniswapQueryTransactorSession struct {
	Contract     *FlashBotUniswapQueryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// FlashBotUniswapQueryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlashBotUniswapQueryRaw struct {
	Contract *FlashBotUniswapQuery // Generic contract binding to access the raw methods on
}

// FlashBotUniswapQueryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlashBotUniswapQueryCallerRaw struct {
	Contract *FlashBotUniswapQueryCaller // Generic read-only contract binding to access the raw methods on
}

// FlashBotUniswapQueryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlashBotUniswapQueryTransactorRaw struct {
	Contract *FlashBotUniswapQueryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlashBotUniswapQuery creates a new instance of FlashBotUniswapQuery, bound to a specific deployed contract.
func NewFlashBotUniswapQuery(address common.Address, backend bind.ContractBackend) (*FlashBotUniswapQuery, error) {
	contract, err := bindFlashBotUniswapQuery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FlashBotUniswapQuery{FlashBotUniswapQueryCaller: FlashBotUniswapQueryCaller{contract: contract}, FlashBotUniswapQueryTransactor: FlashBotUniswapQueryTransactor{contract: contract}, FlashBotUniswapQueryFilterer: FlashBotUniswapQueryFilterer{contract: contract}}, nil
}

// NewFlashBotUniswapQueryCaller creates a new read-only instance of FlashBotUniswapQuery, bound to a specific deployed contract.
func NewFlashBotUniswapQueryCaller(address common.Address, caller bind.ContractCaller) (*FlashBotUniswapQueryCaller, error) {
	contract, err := bindFlashBotUniswapQuery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlashBotUniswapQueryCaller{contract: contract}, nil
}

// NewFlashBotUniswapQueryTransactor creates a new write-only instance of FlashBotUniswapQuery, bound to a specific deployed contract.
func NewFlashBotUniswapQueryTransactor(address common.Address, transactor bind.ContractTransactor) (*FlashBotUniswapQueryTransactor, error) {
	contract, err := bindFlashBotUniswapQuery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlashBotUniswapQueryTransactor{contract: contract}, nil
}

// NewFlashBotUniswapQueryFilterer creates a new log filterer instance of FlashBotUniswapQuery, bound to a specific deployed contract.
func NewFlashBotUniswapQueryFilterer(address common.Address, filterer bind.ContractFilterer) (*FlashBotUniswapQueryFilterer, error) {
	contract, err := bindFlashBotUniswapQuery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlashBotUniswapQueryFilterer{contract: contract}, nil
}

// bindFlashBotUniswapQuery binds a generic wrapper to an already deployed contract.
func bindFlashBotUniswapQuery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FlashBotUniswapQueryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashBotUniswapQuery *FlashBotUniswapQueryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashBotUniswapQuery.Contract.FlashBotUniswapQueryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashBotUniswapQuery *FlashBotUniswapQueryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashBotUniswapQuery.Contract.FlashBotUniswapQueryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashBotUniswapQuery *FlashBotUniswapQueryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashBotUniswapQuery.Contract.FlashBotUniswapQueryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashBotUniswapQuery *FlashBotUniswapQueryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashBotUniswapQuery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashBotUniswapQuery *FlashBotUniswapQueryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashBotUniswapQuery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashBotUniswapQuery *FlashBotUniswapQueryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashBotUniswapQuery.Contract.contract.Transact(opts, method, params...)
}

// GetPairsByIndexRange is a free data retrieval call binding the contract method 0xab2217e4.
//
// Solidity: function getPairsByIndexRange(address _uniswapFactory, uint256 _start, uint256 _stop) view returns(address[3][])
func (_FlashBotUniswapQuery *FlashBotUniswapQueryCaller) GetPairsByIndexRange(opts *bind.CallOpts, _uniswapFactory common.Address, _start *big.Int, _stop *big.Int) ([][3]common.Address, error) {
	var out []interface{}
	err := _FlashBotUniswapQuery.contract.Call(opts, &out, "getPairsByIndexRange", _uniswapFactory, _start, _stop)

	if err != nil {
		return *new([][3]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([][3]common.Address)).(*[][3]common.Address)

	return out0, err

}

// GetPairsByIndexRange is a free data retrieval call binding the contract method 0xab2217e4.
//
// Solidity: function getPairsByIndexRange(address _uniswapFactory, uint256 _start, uint256 _stop) view returns(address[3][])
func (_FlashBotUniswapQuery *FlashBotUniswapQuerySession) GetPairsByIndexRange(_uniswapFactory common.Address, _start *big.Int, _stop *big.Int) ([][3]common.Address, error) {
	return _FlashBotUniswapQuery.Contract.GetPairsByIndexRange(&_FlashBotUniswapQuery.CallOpts, _uniswapFactory, _start, _stop)
}

// GetPairsByIndexRange is a free data retrieval call binding the contract method 0xab2217e4.
//
// Solidity: function getPairsByIndexRange(address _uniswapFactory, uint256 _start, uint256 _stop) view returns(address[3][])
func (_FlashBotUniswapQuery *FlashBotUniswapQueryCallerSession) GetPairsByIndexRange(_uniswapFactory common.Address, _start *big.Int, _stop *big.Int) ([][3]common.Address, error) {
	return _FlashBotUniswapQuery.Contract.GetPairsByIndexRange(&_FlashBotUniswapQuery.CallOpts, _uniswapFactory, _start, _stop)
}

// GetReservesByPairs is a free data retrieval call binding the contract method 0x4dbf0f39.
//
// Solidity: function getReservesByPairs(address[] _pairs) view returns(uint256[3][])
func (_FlashBotUniswapQuery *FlashBotUniswapQueryCaller) GetReservesByPairs(opts *bind.CallOpts, _pairs []common.Address) ([][3]*big.Int, error) {
	var out []interface{}
	err := _FlashBotUniswapQuery.contract.Call(opts, &out, "getReservesByPairs", _pairs)

	if err != nil {
		return *new([][3]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][3]*big.Int)).(*[][3]*big.Int)

	return out0, err

}

// GetReservesByPairs is a free data retrieval call binding the contract method 0x4dbf0f39.
//
// Solidity: function getReservesByPairs(address[] _pairs) view returns(uint256[3][])
func (_FlashBotUniswapQuery *FlashBotUniswapQuerySession) GetReservesByPairs(_pairs []common.Address) ([][3]*big.Int, error) {
	return _FlashBotUniswapQuery.Contract.GetReservesByPairs(&_FlashBotUniswapQuery.CallOpts, _pairs)
}

// GetReservesByPairs is a free data retrieval call binding the contract method 0x4dbf0f39.
//
// Solidity: function getReservesByPairs(address[] _pairs) view returns(uint256[3][])
func (_FlashBotUniswapQuery *FlashBotUniswapQueryCallerSession) GetReservesByPairs(_pairs []common.Address) ([][3]*big.Int, error) {
	return _FlashBotUniswapQuery.Contract.GetReservesByPairs(&_FlashBotUniswapQuery.CallOpts, _pairs)
}
