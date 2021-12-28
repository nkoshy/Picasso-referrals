// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrappers

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

// PicassoReferalMetaData contains all meta data concerning the PicassoReferal contract.
var PicassoReferalMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"referalCode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"influencerName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"influencerAddress\",\"type\":\"address\"}],\"name\":\"createInfluencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"referalCode\",\"type\":\"string\"}],\"name\":\"enableRefereals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"feeReceipients\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getFeeRecepient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"referalCode\",\"type\":\"string\"}],\"name\":\"getInfluencerName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"picassoadmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"referalCode\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"referUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"referalCodeMapping\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"referalEnabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"influencerAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"referalCode\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a9f273e5": "createInfluencer(string,string,address)",
		"275a21ca": "enableRefereals(string)",
		"d1031e7c": "feeReceipients(address)",
		"fe70c44e": "getFeeRecepient(address)",
		"4cccd8fb": "getInfluencerName(string)",
		"155c6dc3": "picassoadmin()",
		"6fbfe945": "referUser(string,address)",
		"4f326a2a": "referalCodeMapping(string)",
	},
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b03191633179055610bf0806100326000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80636fbfe9451161005b5780636fbfe94514610115578063a9f273e514610128578063d1031e7c1461013b578063fe70c44e1461016457610088565b8063155c6dc31461008d578063275a21ca146100bd5780634cccd8fb146100d25780634f326a2a146100f2575b600080fd5b6000546100a0906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100d06100cb366004610968565b610190565b005b6100e56100e0366004610968565b6101fa565b6040516100b49190610aa8565b610105610100366004610968565b6102ac565b6040516100b49493929190610abb565b6100d06101233660046109a3565b6103ff565b6100d06101363660046109ef565b610724565b6100a0610149366004610947565b6002602052600090815260409020546001600160a01b031681565b6100a0610172366004610947565b6001600160a01b039081166000908152600260205260409020541690565b6000546001600160a01b031633146101c35760405162461bcd60e51b81526004016101ba90610b02565b60405180910390fd5b600180826040516101d49190610a8c565b908152604051908190036020019020600101805491151560ff1990921691909117905550565b606060018260405161020c9190610a8c565b908152604051908190036020019020805461022690610b69565b80601f016020809104026020016040519081016040528092919081815260200182805461025290610b69565b801561029f5780601f106102745761010080835404028352916020019161029f565b820191906000526020600020905b81548152906001019060200180831161028257829003601f168201915b505050505090505b919050565b80516020818301810180516001825292820191909301209152805481906102d290610b69565b80601f01602080910402602001604051908101604052809291908181526020018280546102fe90610b69565b801561034b5780601f106103205761010080835404028352916020019161034b565b820191906000526020600020905b81548152906001019060200180831161032e57829003601f168201915b505050506001830154600284018054939460ff8316946101009093046001600160a01b031693509161037c90610b69565b80601f01602080910402602001604051908101604052809291908181526020018280546103a890610b69565b80156103f55780601f106103ca576101008083540402835291602001916103f5565b820191906000526020600020905b8154815290600101906020018083116103d857829003601f168201915b5050505050905084565b6000546001600160a01b031633146104295760405162461bcd60e51b81526004016101ba90610b02565b8160018160405161043a9190610a8c565b90815260405190819003602001902060019081015460ff161515146104a15760405162461bcd60e51b815260206004820152601b60248201527f5265666572616c20636f6465206973206e6f7420456e61626c6564000000000060448201526064016101ba565b6000546001600160a01b0316331461050a5760405162461bcd60e51b815260206004820152602660248201527f4f6e6c79207069636173736f61646d696e2063616e2061646420696e666c75656044820152653731b2b9399760d11b60648201526084016101ba565b6001600160a01b0382811660009081526002602052604090205416156105725760405162461bcd60e51b815260206004820152601c60248201527f546865207573657220697320616c726561647920726566657265642e0000000060448201526064016101ba565b60006001846040516105849190610a8c565b90815260200160405180910390206040518060800160405290816000820180546105ad90610b69565b80601f01602080910402602001604051908101604052809291908181526020018280546105d990610b69565b80156106265780601f106105fb57610100808354040283529160200191610626565b820191906000526020600020905b81548152906001019060200180831161060957829003601f168201915b5050509183525050600182015460ff81161515602083015261010090046001600160a01b0316604082015260028201805460609092019161066690610b69565b80601f016020809104026020016040519081016040528092919081815260200182805461069290610b69565b80156106df5780601f106106b4576101008083540402835291602001916106df565b820191906000526020600020905b8154815290600101906020018083116106c257829003601f168201915b505050919092525050506040908101516001600160a01b03948516600090815260026020529190912080546001600160a01b0319169490911693909317909255505050565b6000546001600160a01b0316331461074e5760405162461bcd60e51b81526004016101ba90610b02565b60408051608081018252838152600060208201526001600160a01b03831681830152606081018590529051600190610787908690610a8c565b908152602001604051809103902060008201518160000190805190602001906107b1929190610810565b5060208281015160018301805460408601516001600160a01b031661010002610100600160a81b031993151560ff199092169190911792909216919091179055606083015180516108089260028501920190610810565b505050505050565b82805461081c90610b69565b90600052602060002090601f01602090048101928261083e5760008555610884565b82601f1061085757805160ff1916838001178555610884565b82800160010185558215610884579182015b82811115610884578251825591602001919060010190610869565b50610890929150610894565b5090565b5b808211156108905760008155600101610895565b80356001600160a01b03811681146102a757600080fd5b600082601f8301126108d0578081fd5b813567ffffffffffffffff808211156108eb576108eb610ba4565b604051601f8301601f19908116603f0116810190828211818310171561091357610913610ba4565b8160405283815286602085880101111561092b578485fd5b8360208701602083013792830160200193909352509392505050565b600060208284031215610958578081fd5b610961826108a9565b9392505050565b600060208284031215610979578081fd5b813567ffffffffffffffff81111561098f578182fd5b61099b848285016108c0565b949350505050565b600080604083850312156109b5578081fd5b823567ffffffffffffffff8111156109cb578182fd5b6109d7858286016108c0565b9250506109e6602084016108a9565b90509250929050565b600080600060608486031215610a03578081fd5b833567ffffffffffffffff80821115610a1a578283fd5b610a26878388016108c0565b94506020860135915080821115610a3b578283fd5b50610a48868287016108c0565b925050610a57604085016108a9565b90509250925092565b60008151808452610a78816020860160208601610b39565b601f01601f19169290920160200192915050565b60008251610a9e818460208701610b39565b9190910192915050565b6000602082526109616020830184610a60565b600060808252610ace6080830187610a60565b85151560208401526001600160a01b03851660408401528281036060840152610af78185610a60565b979650505050505050565b6020808252601a908201527f43616c6c6572206973206e6f74207069636173736f61646d696e000000000000604082015260600190565b60005b83811015610b54578181015183820152602001610b3c565b83811115610b63576000848401525b50505050565b600281046001821680610b7d57607f821691505b60208210811415610b9e57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fdfea264697066735822122028ed75bdbe1eb66e4f1a11624baf4475542fed05293d6ddb05149da8ce6598a664736f6c63430008020033",
}

// PicassoReferalABI is the input ABI used to generate the binding from.
// Deprecated: Use PicassoReferalMetaData.ABI instead.
var PicassoReferalABI = PicassoReferalMetaData.ABI

// Deprecated: Use PicassoReferalMetaData.Sigs instead.
// PicassoReferalFuncSigs maps the 4-byte function signature to its string representation.
var PicassoReferalFuncSigs = PicassoReferalMetaData.Sigs

// PicassoReferalBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PicassoReferalMetaData.Bin instead.
var PicassoReferalBin = PicassoReferalMetaData.Bin

// DeployPicassoReferal deploys a new Ethereum contract, binding an instance of PicassoReferal to it.
func DeployPicassoReferal(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PicassoReferal, error) {
	parsed, err := PicassoReferalMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PicassoReferalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PicassoReferal{PicassoReferalCaller: PicassoReferalCaller{contract: contract}, PicassoReferalTransactor: PicassoReferalTransactor{contract: contract}, PicassoReferalFilterer: PicassoReferalFilterer{contract: contract}}, nil
}

// PicassoReferal is an auto generated Go binding around an Ethereum contract.
type PicassoReferal struct {
	PicassoReferalCaller     // Read-only binding to the contract
	PicassoReferalTransactor // Write-only binding to the contract
	PicassoReferalFilterer   // Log filterer for contract events
}

// PicassoReferalCaller is an auto generated read-only Go binding around an Ethereum contract.
type PicassoReferalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PicassoReferalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PicassoReferalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PicassoReferalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PicassoReferalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PicassoReferalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PicassoReferalSession struct {
	Contract     *PicassoReferal   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PicassoReferalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PicassoReferalCallerSession struct {
	Contract *PicassoReferalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PicassoReferalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PicassoReferalTransactorSession struct {
	Contract     *PicassoReferalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PicassoReferalRaw is an auto generated low-level Go binding around an Ethereum contract.
type PicassoReferalRaw struct {
	Contract *PicassoReferal // Generic contract binding to access the raw methods on
}

// PicassoReferalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PicassoReferalCallerRaw struct {
	Contract *PicassoReferalCaller // Generic read-only contract binding to access the raw methods on
}

// PicassoReferalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PicassoReferalTransactorRaw struct {
	Contract *PicassoReferalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPicassoReferal creates a new instance of PicassoReferal, bound to a specific deployed contract.
func NewPicassoReferal(address common.Address, backend bind.ContractBackend) (*PicassoReferal, error) {
	contract, err := bindPicassoReferal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PicassoReferal{PicassoReferalCaller: PicassoReferalCaller{contract: contract}, PicassoReferalTransactor: PicassoReferalTransactor{contract: contract}, PicassoReferalFilterer: PicassoReferalFilterer{contract: contract}}, nil
}

// NewPicassoReferalCaller creates a new read-only instance of PicassoReferal, bound to a specific deployed contract.
func NewPicassoReferalCaller(address common.Address, caller bind.ContractCaller) (*PicassoReferalCaller, error) {
	contract, err := bindPicassoReferal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PicassoReferalCaller{contract: contract}, nil
}

// NewPicassoReferalTransactor creates a new write-only instance of PicassoReferal, bound to a specific deployed contract.
func NewPicassoReferalTransactor(address common.Address, transactor bind.ContractTransactor) (*PicassoReferalTransactor, error) {
	contract, err := bindPicassoReferal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PicassoReferalTransactor{contract: contract}, nil
}

// NewPicassoReferalFilterer creates a new log filterer instance of PicassoReferal, bound to a specific deployed contract.
func NewPicassoReferalFilterer(address common.Address, filterer bind.ContractFilterer) (*PicassoReferalFilterer, error) {
	contract, err := bindPicassoReferal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PicassoReferalFilterer{contract: contract}, nil
}

// bindPicassoReferal binds a generic wrapper to an already deployed contract.
func bindPicassoReferal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PicassoReferalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PicassoReferal *PicassoReferalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PicassoReferal.Contract.PicassoReferalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PicassoReferal *PicassoReferalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PicassoReferal.Contract.PicassoReferalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PicassoReferal *PicassoReferalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PicassoReferal.Contract.PicassoReferalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PicassoReferal *PicassoReferalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PicassoReferal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PicassoReferal *PicassoReferalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PicassoReferal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PicassoReferal *PicassoReferalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PicassoReferal.Contract.contract.Transact(opts, method, params...)
}

// FeeReceipients is a free data retrieval call binding the contract method 0xd1031e7c.
//
// Solidity: function feeReceipients(address ) view returns(address)
func (_PicassoReferal *PicassoReferalCaller) FeeReceipients(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _PicassoReferal.contract.Call(opts, &out, "feeReceipients", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceipients is a free data retrieval call binding the contract method 0xd1031e7c.
//
// Solidity: function feeReceipients(address ) view returns(address)
func (_PicassoReferal *PicassoReferalSession) FeeReceipients(arg0 common.Address) (common.Address, error) {
	return _PicassoReferal.Contract.FeeReceipients(&_PicassoReferal.CallOpts, arg0)
}

// FeeReceipients is a free data retrieval call binding the contract method 0xd1031e7c.
//
// Solidity: function feeReceipients(address ) view returns(address)
func (_PicassoReferal *PicassoReferalCallerSession) FeeReceipients(arg0 common.Address) (common.Address, error) {
	return _PicassoReferal.Contract.FeeReceipients(&_PicassoReferal.CallOpts, arg0)
}

// GetFeeRecepient is a free data retrieval call binding the contract method 0xfe70c44e.
//
// Solidity: function getFeeRecepient(address userAddress) view returns(address)
func (_PicassoReferal *PicassoReferalCaller) GetFeeRecepient(opts *bind.CallOpts, userAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _PicassoReferal.contract.Call(opts, &out, "getFeeRecepient", userAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeRecepient is a free data retrieval call binding the contract method 0xfe70c44e.
//
// Solidity: function getFeeRecepient(address userAddress) view returns(address)
func (_PicassoReferal *PicassoReferalSession) GetFeeRecepient(userAddress common.Address) (common.Address, error) {
	return _PicassoReferal.Contract.GetFeeRecepient(&_PicassoReferal.CallOpts, userAddress)
}

// GetFeeRecepient is a free data retrieval call binding the contract method 0xfe70c44e.
//
// Solidity: function getFeeRecepient(address userAddress) view returns(address)
func (_PicassoReferal *PicassoReferalCallerSession) GetFeeRecepient(userAddress common.Address) (common.Address, error) {
	return _PicassoReferal.Contract.GetFeeRecepient(&_PicassoReferal.CallOpts, userAddress)
}

// GetInfluencerName is a free data retrieval call binding the contract method 0x4cccd8fb.
//
// Solidity: function getInfluencerName(string referalCode) view returns(string)
func (_PicassoReferal *PicassoReferalCaller) GetInfluencerName(opts *bind.CallOpts, referalCode string) (string, error) {
	var out []interface{}
	err := _PicassoReferal.contract.Call(opts, &out, "getInfluencerName", referalCode)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetInfluencerName is a free data retrieval call binding the contract method 0x4cccd8fb.
//
// Solidity: function getInfluencerName(string referalCode) view returns(string)
func (_PicassoReferal *PicassoReferalSession) GetInfluencerName(referalCode string) (string, error) {
	return _PicassoReferal.Contract.GetInfluencerName(&_PicassoReferal.CallOpts, referalCode)
}

// GetInfluencerName is a free data retrieval call binding the contract method 0x4cccd8fb.
//
// Solidity: function getInfluencerName(string referalCode) view returns(string)
func (_PicassoReferal *PicassoReferalCallerSession) GetInfluencerName(referalCode string) (string, error) {
	return _PicassoReferal.Contract.GetInfluencerName(&_PicassoReferal.CallOpts, referalCode)
}

// Picassoadmin is a free data retrieval call binding the contract method 0x155c6dc3.
//
// Solidity: function picassoadmin() view returns(address)
func (_PicassoReferal *PicassoReferalCaller) Picassoadmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PicassoReferal.contract.Call(opts, &out, "picassoadmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Picassoadmin is a free data retrieval call binding the contract method 0x155c6dc3.
//
// Solidity: function picassoadmin() view returns(address)
func (_PicassoReferal *PicassoReferalSession) Picassoadmin() (common.Address, error) {
	return _PicassoReferal.Contract.Picassoadmin(&_PicassoReferal.CallOpts)
}

// Picassoadmin is a free data retrieval call binding the contract method 0x155c6dc3.
//
// Solidity: function picassoadmin() view returns(address)
func (_PicassoReferal *PicassoReferalCallerSession) Picassoadmin() (common.Address, error) {
	return _PicassoReferal.Contract.Picassoadmin(&_PicassoReferal.CallOpts)
}

// ReferalCodeMapping is a free data retrieval call binding the contract method 0x4f326a2a.
//
// Solidity: function referalCodeMapping(string ) view returns(string name, bool referalEnabled, address influencerAddress, string referalCode)
func (_PicassoReferal *PicassoReferalCaller) ReferalCodeMapping(opts *bind.CallOpts, arg0 string) (struct {
	Name              string
	ReferalEnabled    bool
	InfluencerAddress common.Address
	ReferalCode       string
}, error) {
	var out []interface{}
	err := _PicassoReferal.contract.Call(opts, &out, "referalCodeMapping", arg0)

	outstruct := new(struct {
		Name              string
		ReferalEnabled    bool
		InfluencerAddress common.Address
		ReferalCode       string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.ReferalEnabled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.InfluencerAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.ReferalCode = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// ReferalCodeMapping is a free data retrieval call binding the contract method 0x4f326a2a.
//
// Solidity: function referalCodeMapping(string ) view returns(string name, bool referalEnabled, address influencerAddress, string referalCode)
func (_PicassoReferal *PicassoReferalSession) ReferalCodeMapping(arg0 string) (struct {
	Name              string
	ReferalEnabled    bool
	InfluencerAddress common.Address
	ReferalCode       string
}, error) {
	return _PicassoReferal.Contract.ReferalCodeMapping(&_PicassoReferal.CallOpts, arg0)
}

// ReferalCodeMapping is a free data retrieval call binding the contract method 0x4f326a2a.
//
// Solidity: function referalCodeMapping(string ) view returns(string name, bool referalEnabled, address influencerAddress, string referalCode)
func (_PicassoReferal *PicassoReferalCallerSession) ReferalCodeMapping(arg0 string) (struct {
	Name              string
	ReferalEnabled    bool
	InfluencerAddress common.Address
	ReferalCode       string
}, error) {
	return _PicassoReferal.Contract.ReferalCodeMapping(&_PicassoReferal.CallOpts, arg0)
}

// CreateInfluencer is a paid mutator transaction binding the contract method 0xa9f273e5.
//
// Solidity: function createInfluencer(string referalCode, string influencerName, address influencerAddress) returns()
func (_PicassoReferal *PicassoReferalTransactor) CreateInfluencer(opts *bind.TransactOpts, referalCode string, influencerName string, influencerAddress common.Address) (*types.Transaction, error) {
	return _PicassoReferal.contract.Transact(opts, "createInfluencer", referalCode, influencerName, influencerAddress)
}

// CreateInfluencer is a paid mutator transaction binding the contract method 0xa9f273e5.
//
// Solidity: function createInfluencer(string referalCode, string influencerName, address influencerAddress) returns()
func (_PicassoReferal *PicassoReferalSession) CreateInfluencer(referalCode string, influencerName string, influencerAddress common.Address) (*types.Transaction, error) {
	return _PicassoReferal.Contract.CreateInfluencer(&_PicassoReferal.TransactOpts, referalCode, influencerName, influencerAddress)
}

// CreateInfluencer is a paid mutator transaction binding the contract method 0xa9f273e5.
//
// Solidity: function createInfluencer(string referalCode, string influencerName, address influencerAddress) returns()
func (_PicassoReferal *PicassoReferalTransactorSession) CreateInfluencer(referalCode string, influencerName string, influencerAddress common.Address) (*types.Transaction, error) {
	return _PicassoReferal.Contract.CreateInfluencer(&_PicassoReferal.TransactOpts, referalCode, influencerName, influencerAddress)
}

// EnableRefereals is a paid mutator transaction binding the contract method 0x275a21ca.
//
// Solidity: function enableRefereals(string referalCode) returns()
func (_PicassoReferal *PicassoReferalTransactor) EnableRefereals(opts *bind.TransactOpts, referalCode string) (*types.Transaction, error) {
	return _PicassoReferal.contract.Transact(opts, "enableRefereals", referalCode)
}

// EnableRefereals is a paid mutator transaction binding the contract method 0x275a21ca.
//
// Solidity: function enableRefereals(string referalCode) returns()
func (_PicassoReferal *PicassoReferalSession) EnableRefereals(referalCode string) (*types.Transaction, error) {
	return _PicassoReferal.Contract.EnableRefereals(&_PicassoReferal.TransactOpts, referalCode)
}

// EnableRefereals is a paid mutator transaction binding the contract method 0x275a21ca.
//
// Solidity: function enableRefereals(string referalCode) returns()
func (_PicassoReferal *PicassoReferalTransactorSession) EnableRefereals(referalCode string) (*types.Transaction, error) {
	return _PicassoReferal.Contract.EnableRefereals(&_PicassoReferal.TransactOpts, referalCode)
}

// ReferUser is a paid mutator transaction binding the contract method 0x6fbfe945.
//
// Solidity: function referUser(string referalCode, address userAddress) returns()
func (_PicassoReferal *PicassoReferalTransactor) ReferUser(opts *bind.TransactOpts, referalCode string, userAddress common.Address) (*types.Transaction, error) {
	return _PicassoReferal.contract.Transact(opts, "referUser", referalCode, userAddress)
}

// ReferUser is a paid mutator transaction binding the contract method 0x6fbfe945.
//
// Solidity: function referUser(string referalCode, address userAddress) returns()
func (_PicassoReferal *PicassoReferalSession) ReferUser(referalCode string, userAddress common.Address) (*types.Transaction, error) {
	return _PicassoReferal.Contract.ReferUser(&_PicassoReferal.TransactOpts, referalCode, userAddress)
}

// ReferUser is a paid mutator transaction binding the contract method 0x6fbfe945.
//
// Solidity: function referUser(string referalCode, address userAddress) returns()
func (_PicassoReferal *PicassoReferalTransactorSession) ReferUser(referalCode string, userAddress common.Address) (*types.Transaction, error) {
	return _PicassoReferal.Contract.ReferUser(&_PicassoReferal.TransactOpts, referalCode, userAddress)
}
