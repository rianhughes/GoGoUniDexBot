// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BundleExecutor

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

// BundleExecutorMetaData contains all meta data concerning the BundleExecutor contract.
var BundleExecutorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"call\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_wethAmountToFirstMarket\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ethAmountToCoinbase\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_targets\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_payloads\",\"type\":\"bytes[]\"}],\"name\":\"uniswapWeth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c0604052604051610e83380380610e838339818101604052810190610025919061012d565b3373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff1660601b815250508073ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1660601b8152505060003411156101125773c02aaa39b223fe8d0a0e5c4f27ead9083c756cc273ffffffffffffffffffffffffffffffffffffffff1663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b1580156100f857600080fd5b505af115801561010c573d6000803e3d6000fd5b50505050505b5061019f565b60008151905061012781610188565b92915050565b60006020828403121561013f57600080fd5b600061014d84828501610118565b91505092915050565b600061016182610168565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b61019181610156565b811461019c57600080fd5b50565b60805160601c60a05160601c610cbc6101c7600039806101a552508060895250610cbc6000f3fe60806040526004361061002d5760003560e01c80636dbf2fa014610039578063ecd494b31461006957610034565b3661003457005b600080fd5b610053600480360381019061004e91906107b9565b610085565b6040516100609190610a3a565b60405180910390f35b610083600480360381019061007e9190610877565b6101a3565b005b60607f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146100df57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16141561011957600080fd5b600060608673ffffffffffffffffffffffffffffffffffffffff168686866040516101459291906109c6565b60006040518083038185875af1925050503d8060008114610182576040519150601f19603f3d011682016040523d82523d6000602084013e610187565b606091505b50915091508161019657600080fd5b8092505050949350505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101fb57600080fd5b805182511461020957600080fd5b600073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc273ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161025891906109f6565b60206040518083038186803b15801561027057600080fd5b505afa158015610284573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102a8919061084e565b905073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc273ffffffffffffffffffffffffffffffffffffffff1663a9059cbb846000815181106102e757fe5b6020026020010151876040518363ffffffff1660e01b815260040161030d929190610a11565b602060405180830381600087803b15801561032757600080fd5b505af115801561033b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061035f9190610825565b5060005b835181101561041b576000606085838151811061037c57fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff168584815181106103a657fe5b60200260200101516040516103bb91906109df565b6000604051808303816000865af19150503d80600081146103f8576040519150601f19603f3d011682016040523d82523d6000602084013e6103fd565b606091505b50915091508161040c57600080fd5b50508080600101915050610363565b50600073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc273ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161046b91906109f6565b60206040518083038186803b15801561048357600080fd5b505afa158015610497573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104bb919061084e565b905084820181116104cb57600080fd5b60008514156104db5750506105b5565b60004790508581101561056a5773c02aaa39b223fe8d0a0e5c4f27ead9083c756cc273ffffffffffffffffffffffffffffffffffffffff16632e1a7d4d8288036040518263ffffffff1660e01b81526004016105379190610a5c565b600060405180830381600087803b15801561055157600080fd5b505af1158015610565573d6000803e3d6000fd5b505050505b4173ffffffffffffffffffffffffffffffffffffffff166108fc879081150290604051600060405180830381858888f193505050501580156105b0573d6000803e3d6000fd5b505050505b50505050565b6000813590506105ca81610c2a565b92915050565b6000813590506105df81610c41565b92915050565b600082601f8301126105f657600080fd5b813561060961060482610aa4565b610a77565b9150818183526020840193506020810190508385602084028201111561062e57600080fd5b60005b8381101561065e578161064488826105bb565b845260208401935060208301925050600181019050610631565b5050505092915050565b600082601f83011261067957600080fd5b813561068c61068782610acc565b610a77565b9150818183526020840193506020810190508360005b838110156106d257813586016106b8888261073b565b8452602084019350602083019250506001810190506106a2565b5050505092915050565b6000815190506106eb81610c58565b92915050565b60008083601f84011261070357600080fd5b8235905067ffffffffffffffff81111561071c57600080fd5b60208301915083600182028301111561073457600080fd5b9250929050565b600082601f83011261074c57600080fd5b813561075f61075a82610af4565b610a77565b9150808252602083016020830185838301111561077b57600080fd5b610786838284610bd7565b50505092915050565b60008135905061079e81610c6f565b92915050565b6000815190506107b381610c6f565b92915050565b600080600080606085870312156107cf57600080fd5b60006107dd878288016105d0565b94505060206107ee8782880161078f565b935050604085013567ffffffffffffffff81111561080b57600080fd5b610817878288016106f1565b925092505092959194509250565b60006020828403121561083757600080fd5b6000610845848285016106dc565b91505092915050565b60006020828403121561086057600080fd5b600061086e848285016107a4565b91505092915050565b6000806000806080858703121561088d57600080fd5b600061089b8782880161078f565b94505060206108ac8782880161078f565b935050604085013567ffffffffffffffff8111156108c957600080fd5b6108d5878288016105e5565b925050606085013567ffffffffffffffff8111156108f257600080fd5b6108fe87828801610668565b91505092959194509250565b61091381610ba1565b82525050565b61092281610b47565b82525050565b60006109348385610b3c565b9350610941838584610bd7565b82840190509392505050565b600061095882610b20565b6109628185610b2b565b9350610972818560208601610be6565b61097b81610c19565b840191505092915050565b600061099182610b20565b61099b8185610b3c565b93506109ab818560208601610be6565b80840191505092915050565b6109c081610b97565b82525050565b60006109d3828486610928565b91508190509392505050565b60006109eb8284610986565b915081905092915050565b6000602082019050610a0b600083018461090a565b92915050565b6000604082019050610a266000830185610919565b610a3360208301846109b7565b9392505050565b60006020820190508181036000830152610a54818461094d565b905092915050565b6000602082019050610a7160008301846109b7565b92915050565b6000604051905081810181811067ffffffffffffffff82111715610a9a57600080fd5b8060405250919050565b600067ffffffffffffffff821115610abb57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115610ae357600080fd5b602082029050602081019050919050565b600067ffffffffffffffff821115610b0b57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b6000610b5282610b77565b9050919050565b6000610b6482610b77565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610bac82610bb3565b9050919050565b6000610bbe82610bc5565b9050919050565b6000610bd082610b77565b9050919050565b82818337600083830152505050565b60005b83811015610c04578082015181840152602081019050610be9565b83811115610c13576000848401525b50505050565b6000601f19601f8301169050919050565b610c3381610b47565b8114610c3e57600080fd5b50565b610c4a81610b59565b8114610c5557600080fd5b50565b610c6181610b6b565b8114610c6c57600080fd5b50565b610c7881610b97565b8114610c8357600080fd5b5056fea26469706673582212208cef7dba03ba7dcc957532087be9f2292ac4746ec8356f433cb614452ab7772364736f6c634300060c0033",
}

// BundleExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use BundleExecutorMetaData.ABI instead.
var BundleExecutorABI = BundleExecutorMetaData.ABI

// BundleExecutorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BundleExecutorMetaData.Bin instead.
var BundleExecutorBin = BundleExecutorMetaData.Bin

// DeployBundleExecutor deploys a new Ethereum contract, binding an instance of BundleExecutor to it.
func DeployBundleExecutor(auth *bind.TransactOpts, backend bind.ContractBackend, _executor common.Address) (common.Address, *types.Transaction, *BundleExecutor, error) {
	parsed, err := BundleExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BundleExecutorBin), backend, _executor)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BundleExecutor{BundleExecutorCaller: BundleExecutorCaller{contract: contract}, BundleExecutorTransactor: BundleExecutorTransactor{contract: contract}, BundleExecutorFilterer: BundleExecutorFilterer{contract: contract}}, nil
}

// BundleExecutor is an auto generated Go binding around an Ethereum contract.
type BundleExecutor struct {
	BundleExecutorCaller     // Read-only binding to the contract
	BundleExecutorTransactor // Write-only binding to the contract
	BundleExecutorFilterer   // Log filterer for contract events
}

// BundleExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BundleExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundleExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BundleExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundleExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BundleExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BundleExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BundleExecutorSession struct {
	Contract     *BundleExecutor   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BundleExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BundleExecutorCallerSession struct {
	Contract *BundleExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BundleExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BundleExecutorTransactorSession struct {
	Contract     *BundleExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BundleExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BundleExecutorRaw struct {
	Contract *BundleExecutor // Generic contract binding to access the raw methods on
}

// BundleExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BundleExecutorCallerRaw struct {
	Contract *BundleExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// BundleExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BundleExecutorTransactorRaw struct {
	Contract *BundleExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBundleExecutor creates a new instance of BundleExecutor, bound to a specific deployed contract.
func NewBundleExecutor(address common.Address, backend bind.ContractBackend) (*BundleExecutor, error) {
	contract, err := bindBundleExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BundleExecutor{BundleExecutorCaller: BundleExecutorCaller{contract: contract}, BundleExecutorTransactor: BundleExecutorTransactor{contract: contract}, BundleExecutorFilterer: BundleExecutorFilterer{contract: contract}}, nil
}

// NewBundleExecutorCaller creates a new read-only instance of BundleExecutor, bound to a specific deployed contract.
func NewBundleExecutorCaller(address common.Address, caller bind.ContractCaller) (*BundleExecutorCaller, error) {
	contract, err := bindBundleExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BundleExecutorCaller{contract: contract}, nil
}

// NewBundleExecutorTransactor creates a new write-only instance of BundleExecutor, bound to a specific deployed contract.
func NewBundleExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*BundleExecutorTransactor, error) {
	contract, err := bindBundleExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BundleExecutorTransactor{contract: contract}, nil
}

// NewBundleExecutorFilterer creates a new log filterer instance of BundleExecutor, bound to a specific deployed contract.
func NewBundleExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*BundleExecutorFilterer, error) {
	contract, err := bindBundleExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BundleExecutorFilterer{contract: contract}, nil
}

// bindBundleExecutor binds a generic wrapper to an already deployed contract.
func bindBundleExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BundleExecutorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BundleExecutor *BundleExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BundleExecutor.Contract.BundleExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BundleExecutor *BundleExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BundleExecutor.Contract.BundleExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BundleExecutor *BundleExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BundleExecutor.Contract.BundleExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BundleExecutor *BundleExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BundleExecutor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BundleExecutor *BundleExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BundleExecutor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BundleExecutor *BundleExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BundleExecutor.Contract.contract.Transact(opts, method, params...)
}

// Call is a paid mutator transaction binding the contract method 0x6dbf2fa0.
//
// Solidity: function call(address _to, uint256 _value, bytes _data) payable returns(bytes)
func (_BundleExecutor *BundleExecutorTransactor) Call(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _BundleExecutor.contract.Transact(opts, "call", _to, _value, _data)
}

// Call is a paid mutator transaction binding the contract method 0x6dbf2fa0.
//
// Solidity: function call(address _to, uint256 _value, bytes _data) payable returns(bytes)
func (_BundleExecutor *BundleExecutorSession) Call(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _BundleExecutor.Contract.Call(&_BundleExecutor.TransactOpts, _to, _value, _data)
}

// Call is a paid mutator transaction binding the contract method 0x6dbf2fa0.
//
// Solidity: function call(address _to, uint256 _value, bytes _data) payable returns(bytes)
func (_BundleExecutor *BundleExecutorTransactorSession) Call(_to common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _BundleExecutor.Contract.Call(&_BundleExecutor.TransactOpts, _to, _value, _data)
}

// UniswapWeth is a paid mutator transaction binding the contract method 0xecd494b3.
//
// Solidity: function uniswapWeth(uint256 _wethAmountToFirstMarket, uint256 _ethAmountToCoinbase, address[] _targets, bytes[] _payloads) payable returns()
func (_BundleExecutor *BundleExecutorTransactor) UniswapWeth(opts *bind.TransactOpts, _wethAmountToFirstMarket *big.Int, _ethAmountToCoinbase *big.Int, _targets []common.Address, _payloads [][]byte) (*types.Transaction, error) {
	return _BundleExecutor.contract.Transact(opts, "uniswapWeth", _wethAmountToFirstMarket, _ethAmountToCoinbase, _targets, _payloads)
}

// UniswapWeth is a paid mutator transaction binding the contract method 0xecd494b3.
//
// Solidity: function uniswapWeth(uint256 _wethAmountToFirstMarket, uint256 _ethAmountToCoinbase, address[] _targets, bytes[] _payloads) payable returns()
func (_BundleExecutor *BundleExecutorSession) UniswapWeth(_wethAmountToFirstMarket *big.Int, _ethAmountToCoinbase *big.Int, _targets []common.Address, _payloads [][]byte) (*types.Transaction, error) {
	return _BundleExecutor.Contract.UniswapWeth(&_BundleExecutor.TransactOpts, _wethAmountToFirstMarket, _ethAmountToCoinbase, _targets, _payloads)
}

// UniswapWeth is a paid mutator transaction binding the contract method 0xecd494b3.
//
// Solidity: function uniswapWeth(uint256 _wethAmountToFirstMarket, uint256 _ethAmountToCoinbase, address[] _targets, bytes[] _payloads) payable returns()
func (_BundleExecutor *BundleExecutorTransactorSession) UniswapWeth(_wethAmountToFirstMarket *big.Int, _ethAmountToCoinbase *big.Int, _targets []common.Address, _payloads [][]byte) (*types.Transaction, error) {
	return _BundleExecutor.Contract.UniswapWeth(&_BundleExecutor.TransactOpts, _wethAmountToFirstMarket, _ethAmountToCoinbase, _targets, _payloads)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BundleExecutor *BundleExecutorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BundleExecutor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BundleExecutor *BundleExecutorSession) Receive() (*types.Transaction, error) {
	return _BundleExecutor.Contract.Receive(&_BundleExecutor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BundleExecutor *BundleExecutorTransactorSession) Receive() (*types.Transaction, error) {
	return _BundleExecutor.Contract.Receive(&_BundleExecutor.TransactOpts)
}
