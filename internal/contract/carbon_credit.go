// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
	_ = abi.ConvertType
)

// CarbonCreditMetaData contains all meta data concerning the CarbonCredit contract.
var CarbonCreditMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"minter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinter\",\"inputs\":[{\"name\":\"_minter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinterUpdated\",\"inputs\":[{\"name\":\"oldMinter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newMinter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608034620003fe57601f90601f19906001600160401b03620012343881900385810185168401838111858210176200030257849282916040528339602093849181010312620003fe57516001600160a01b038082169490929091859003620003fe576200006b62000403565b95600d87526c10d85c989bdb8810dc99591a5d609a1b858801526200008f62000403565b966002885261434360f01b86890152805191838311620003025760039283546001938482811c92168015620003f3575b8a831014620003dd57818484931162000387575b508990848311600114620003245760009262000318575b505060001982861b1c191690831b1783555b8851938411620003025760049889548381811c91168015620002f7575b89821014620002e25782811162000297575b50879185116001146200022e578495509084929160009562000222575b50501b92600019911b1c19161784555b6005549160018060a01b0319923384821617600555604051923391167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a38315620001e25783808460065416176006556040519060007f1cf2de25c5bf439ac0287061c3a0fa69b3b02867d0ccfd2ded34e42577050b738180a3610e109081620004248239f35b62461bcd60e51b825281850152601960248201527f436172626f6e4372656469743a207a65726f206d696e74657200000000000000604482015260649150fd5b01519350388062000148565b9291948416928960005284886000209460005b8a898383106200027f575050501062000264575b50505050811b01845562000158565b01519060f884600019921b161c191690553880808062000255565b86860151895590970196948501948893500162000241565b8a600052886000208380880160051c8201928b8910620002d8575b0160051c019084905b828110620002cb5750506200012b565b60008155018490620002bb565b92508192620002b2565b60228b634e487b7160e01b6000525260246000fd5b90607f169062000119565b634e487b7160e01b600052604160045260246000fd5b015190503880620000ea565b908886941691876000528b6000209260005b8d82821062000370575050841162000357575b505050811b018355620000fc565b015160001983881b60f8161c1916905538808062000349565b838501518655899790950194938401930162000336565b90915085600052896000208480850160051c8201928c8610620003d3575b918791869594930160051c01915b828110620003c3575050620000d3565b60008155859450879101620003b3565b92508192620003a5565b634e487b7160e01b600052602260045260246000fd5b91607f1691620000bf565b600080fd5b60408051919082016001600160401b03811183821017620003025760405256fe60806040908082526004918236101561001757600080fd5b600092833560e01c92836306fdde0314610862575082630754617214610839578263095ea7b31461080f57826318160ddd146107f057826323b872dd146107b3578263313ce56714610797578263395093511461074757826340c10f191461066e57826370a0823114610637578263715018a6146105da5782638da5cb5b146105b157826395d89b41146104925782639dc29fac14610356578263a457c2d7146102b3578263a9059cbb14610282578263dd62ed3e14610235578263f2fde38b1461016457505063fca3b5aa146100ed57600080fd5b3461016157602036600319011261016157610106610983565b61010e6109b4565b6001600160a01b0390811690610125821515610d37565b816006549182167f1cf2de25c5bf439ac0287061c3a0fa69b3b02867d0ccfd2ded34e42577050b738580a36001600160a01b0319161760065580f35b80fd5b9091503461023157602036600319011261023157610180610983565b906101896109b4565b6001600160a01b039182169283156101df575050600554826bffffffffffffffffffffffff60a01b821617600555167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b906020608492519162461bcd60e51b8352820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152fd5b8280fd5b83823461027e578060031936011261027e5780602092610253610983565b61025b61099e565b6001600160a01b0391821683526001865283832091168252845220549051908152f35b5080fd5b83823461027e578060031936011261027e576020906102ac6102a2610983565b6024359033610a2f565b5160018152f35b83346101615782600319360112610161576102cc610983565b918360243592338152600160205281812060018060a01b0386168252602052205490828210610305576020856102ac8585038733610b9d565b608490602086519162461bcd60e51b8352820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152fd5b83903461027e578260031936011261027e57610370610983565b600654602435916001600160a01b039161038d9083163314610d83565b610398833383610c9f565b16918215610445578284528360205284842054908282106103f757508184957fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef936020938688528785520381872055816002540360025551908152a380f35b608490602087519162461bcd60e51b8352820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152fd5b608490602086519162461bcd60e51b8352820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152fd5b83823461027e578160031936011261027e57805190828454600181811c908083169283156105a7575b6020938484108114610594578388529081156105785750600114610523575b505050829003601f01601f191682019267ffffffffffffffff841183851017610510575082918261050c92528261093a565b0390f35b634e487b7160e01b815260418552602490fd5b8787529192508591837f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b5b83851061056457505050508301018580806104da565b80548886018301529301928490820161054e565b60ff1916878501525050151560051b84010190508580806104da565b634e487b7160e01b895260228a52602489fd5b91607f16916104bb565b83823461027e578160031936011261027e5760055490516001600160a01b039091168152602090f35b83346101615780600319360112610161576105f36109b4565b600580546001600160a01b0319811690915581906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a380f35b83823461027e57602036600319011261027e5760209181906001600160a01b0361065f610983565b16815280845220549051908152f35b915034610231578060031936011261023157610688610983565b60065460243592916001600160a01b03916106a69083163314610d83565b169283156106b48115610d37565b61070557506020827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef926106eb8795600254610a0c565b60025585855284835280852082815401905551908152a380f35b6020606492519162461bcd60e51b8352820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152fd5b83823461027e578060031936011261027e576102ac60209261079061076a610983565b338352600186528483206001600160a01b03821684528652918490205460243590610a0c565b9033610b9d565b83823461027e578160031936011261027e576020905160128152f35b83823461027e57606036600319011261027e576020906102ac6107d4610983565b6107dc61099e565b604435916107eb833383610c9f565b610a2f565b83823461027e578160031936011261027e576020906002549051908152f35b83823461027e578060031936011261027e576020906102ac61082f610983565b6024359033610b9d565b83823461027e578160031936011261027e5760065490516001600160a01b039091168152602090f35b8491923461023157826003193601126102315782600354600181811c90808316928315610930575b60209384841081146105945783885290811561057857506001146108da57505050829003601f01601f191682019267ffffffffffffffff841183851017610510575082918261050c92528261093a565b600387529192508591837fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b83851061091c57505050508301018580806104da565b805488860183015293019284908201610906565b91607f169161088a565b6020808252825181830181905290939260005b82811061096f57505060409293506000838284010152601f8019910116010190565b81810186015184820160400152850161094d565b600435906001600160a01b038216820361099957565b600080fd5b602435906001600160a01b038216820361099957565b6005546001600160a01b031633036109c857565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b91908201809211610a1957565b634e487b7160e01b600052601160045260246000fd5b6001600160a01b03908116918215610b4a5716918215610af957600082815280602052604081205491808310610aa557604082827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef958760209652828652038282205586815220818154019055604051908152a3565b60405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608490fd5b60405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608490fd5b60405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608490fd5b6001600160a01b03908116918215610c4e5716918215610bfe5760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925918360005260018252604060002085600052825280604060002055604051908152a3565b60405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608490fd5b60405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608490fd5b9060018060a01b0380831660005260016020526040600020908216600052602052604060002054926000198403610cd7575b50505050565b808410610cf257610ce9930391610b9d565b38808080610cd1565b60405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606490fd5b15610d3e57565b60405162461bcd60e51b815260206004820152601a60248201527f436172626f6e4372656469743a207a65726f20616464726573730000000000006044820152606490fd5b15610d8a57565b60405162461bcd60e51b815260206004820152602260248201527f436172626f6e4372656469743a2063616c6c6572206973206e6f74206d696e7460448201526132b960f11b6064820152608490fdfea26469706673582212203c5be1c6530c730a845f2f3998184753f997b9657b129939d3a7f4548acdf95264736f6c63430008130033",
}

// CarbonCreditABI is the input ABI used to generate the binding from.
// Deprecated: Use CarbonCreditMetaData.ABI instead.
var CarbonCreditABI = CarbonCreditMetaData.ABI

// CarbonCreditBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CarbonCreditMetaData.Bin instead.
var CarbonCreditBin = CarbonCreditMetaData.Bin

// DeployCarbonCredit deploys a new Ethereum contract, binding an instance of CarbonCredit to it.
func DeployCarbonCredit(auth *bind.TransactOpts, backend bind.ContractBackend, _minter common.Address) (common.Address, *types.Transaction, *CarbonCredit, error) {
	parsed, err := CarbonCreditMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CarbonCreditBin), backend, _minter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CarbonCredit{CarbonCreditCaller: CarbonCreditCaller{contract: contract}, CarbonCreditTransactor: CarbonCreditTransactor{contract: contract}, CarbonCreditFilterer: CarbonCreditFilterer{contract: contract}}, nil
}

// CarbonCredit is an auto generated Go binding around an Ethereum contract.
type CarbonCredit struct {
	CarbonCreditCaller     // Read-only binding to the contract
	CarbonCreditTransactor // Write-only binding to the contract
	CarbonCreditFilterer   // Log filterer for contract events
}

// CarbonCreditCaller is an auto generated read-only Go binding around an Ethereum contract.
type CarbonCreditCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarbonCreditTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CarbonCreditTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarbonCreditFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CarbonCreditFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarbonCreditSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CarbonCreditSession struct {
	Contract     *CarbonCredit     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CarbonCreditCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CarbonCreditCallerSession struct {
	Contract *CarbonCreditCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CarbonCreditTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CarbonCreditTransactorSession struct {
	Contract     *CarbonCreditTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CarbonCreditRaw is an auto generated low-level Go binding around an Ethereum contract.
type CarbonCreditRaw struct {
	Contract *CarbonCredit // Generic contract binding to access the raw methods on
}

// CarbonCreditCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CarbonCreditCallerRaw struct {
	Contract *CarbonCreditCaller // Generic read-only contract binding to access the raw methods on
}

// CarbonCreditTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CarbonCreditTransactorRaw struct {
	Contract *CarbonCreditTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCarbonCredit creates a new instance of CarbonCredit, bound to a specific deployed contract.
func NewCarbonCredit(address common.Address, backend bind.ContractBackend) (*CarbonCredit, error) {
	contract, err := bindCarbonCredit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CarbonCredit{CarbonCreditCaller: CarbonCreditCaller{contract: contract}, CarbonCreditTransactor: CarbonCreditTransactor{contract: contract}, CarbonCreditFilterer: CarbonCreditFilterer{contract: contract}}, nil
}

// NewCarbonCreditCaller creates a new read-only instance of CarbonCredit, bound to a specific deployed contract.
func NewCarbonCreditCaller(address common.Address, caller bind.ContractCaller) (*CarbonCreditCaller, error) {
	contract, err := bindCarbonCredit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CarbonCreditCaller{contract: contract}, nil
}

// NewCarbonCreditTransactor creates a new write-only instance of CarbonCredit, bound to a specific deployed contract.
func NewCarbonCreditTransactor(address common.Address, transactor bind.ContractTransactor) (*CarbonCreditTransactor, error) {
	contract, err := bindCarbonCredit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CarbonCreditTransactor{contract: contract}, nil
}

// NewCarbonCreditFilterer creates a new log filterer instance of CarbonCredit, bound to a specific deployed contract.
func NewCarbonCreditFilterer(address common.Address, filterer bind.ContractFilterer) (*CarbonCreditFilterer, error) {
	contract, err := bindCarbonCredit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CarbonCreditFilterer{contract: contract}, nil
}

// bindCarbonCredit binds a generic wrapper to an already deployed contract.
func bindCarbonCredit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CarbonCreditMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CarbonCredit *CarbonCreditRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CarbonCredit.Contract.CarbonCreditCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CarbonCredit *CarbonCreditRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CarbonCredit.Contract.CarbonCreditTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CarbonCredit *CarbonCreditRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CarbonCredit.Contract.CarbonCreditTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CarbonCredit *CarbonCreditCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CarbonCredit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CarbonCredit *CarbonCreditTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CarbonCredit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CarbonCredit *CarbonCreditTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CarbonCredit.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CarbonCredit *CarbonCreditCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CarbonCredit.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CarbonCredit *CarbonCreditSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CarbonCredit.Contract.Allowance(&_CarbonCredit.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CarbonCredit *CarbonCreditCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CarbonCredit.Contract.Allowance(&_CarbonCredit.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CarbonCredit *CarbonCreditCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CarbonCredit.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CarbonCredit *CarbonCreditSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CarbonCredit.Contract.BalanceOf(&_CarbonCredit.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CarbonCredit *CarbonCreditCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CarbonCredit.Contract.BalanceOf(&_CarbonCredit.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CarbonCredit *CarbonCreditCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CarbonCredit.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CarbonCredit *CarbonCreditSession) Decimals() (uint8, error) {
	return _CarbonCredit.Contract.Decimals(&_CarbonCredit.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CarbonCredit *CarbonCreditCallerSession) Decimals() (uint8, error) {
	return _CarbonCredit.Contract.Decimals(&_CarbonCredit.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CarbonCredit *CarbonCreditCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CarbonCredit.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CarbonCredit *CarbonCreditSession) Minter() (common.Address, error) {
	return _CarbonCredit.Contract.Minter(&_CarbonCredit.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_CarbonCredit *CarbonCreditCallerSession) Minter() (common.Address, error) {
	return _CarbonCredit.Contract.Minter(&_CarbonCredit.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CarbonCredit *CarbonCreditCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CarbonCredit.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CarbonCredit *CarbonCreditSession) Name() (string, error) {
	return _CarbonCredit.Contract.Name(&_CarbonCredit.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CarbonCredit *CarbonCreditCallerSession) Name() (string, error) {
	return _CarbonCredit.Contract.Name(&_CarbonCredit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CarbonCredit *CarbonCreditCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CarbonCredit.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CarbonCredit *CarbonCreditSession) Owner() (common.Address, error) {
	return _CarbonCredit.Contract.Owner(&_CarbonCredit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CarbonCredit *CarbonCreditCallerSession) Owner() (common.Address, error) {
	return _CarbonCredit.Contract.Owner(&_CarbonCredit.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CarbonCredit *CarbonCreditCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CarbonCredit.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CarbonCredit *CarbonCreditSession) Symbol() (string, error) {
	return _CarbonCredit.Contract.Symbol(&_CarbonCredit.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CarbonCredit *CarbonCreditCallerSession) Symbol() (string, error) {
	return _CarbonCredit.Contract.Symbol(&_CarbonCredit.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CarbonCredit *CarbonCreditCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CarbonCredit.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CarbonCredit *CarbonCreditSession) TotalSupply() (*big.Int, error) {
	return _CarbonCredit.Contract.TotalSupply(&_CarbonCredit.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CarbonCredit *CarbonCreditCallerSession) TotalSupply() (*big.Int, error) {
	return _CarbonCredit.Contract.TotalSupply(&_CarbonCredit.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CarbonCredit *CarbonCreditTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CarbonCredit *CarbonCreditSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.Contract.Approve(&_CarbonCredit.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CarbonCredit *CarbonCreditTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.Contract.Approve(&_CarbonCredit.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_CarbonCredit *CarbonCreditTransactor) Burn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.contract.Transact(opts, "burn", from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_CarbonCredit *CarbonCreditSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.Contract.Burn(&_CarbonCredit.TransactOpts, from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_CarbonCredit *CarbonCreditTransactorSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.Contract.Burn(&_CarbonCredit.TransactOpts, from, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_CarbonCredit *CarbonCreditTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_CarbonCredit *CarbonCreditSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.Contract.DecreaseAllowance(&_CarbonCredit.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_CarbonCredit *CarbonCreditTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.Contract.DecreaseAllowance(&_CarbonCredit.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_CarbonCredit *CarbonCreditTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_CarbonCredit *CarbonCreditSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.Contract.IncreaseAllowance(&_CarbonCredit.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_CarbonCredit *CarbonCreditTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.Contract.IncreaseAllowance(&_CarbonCredit.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_CarbonCredit *CarbonCreditTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_CarbonCredit *CarbonCreditSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.Contract.Mint(&_CarbonCredit.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_CarbonCredit *CarbonCreditTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.Contract.Mint(&_CarbonCredit.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CarbonCredit *CarbonCreditTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CarbonCredit.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CarbonCredit *CarbonCreditSession) RenounceOwnership() (*types.Transaction, error) {
	return _CarbonCredit.Contract.RenounceOwnership(&_CarbonCredit.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CarbonCredit *CarbonCreditTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CarbonCredit.Contract.RenounceOwnership(&_CarbonCredit.TransactOpts)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_CarbonCredit *CarbonCreditTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _CarbonCredit.contract.Transact(opts, "setMinter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_CarbonCredit *CarbonCreditSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _CarbonCredit.Contract.SetMinter(&_CarbonCredit.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_CarbonCredit *CarbonCreditTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _CarbonCredit.Contract.SetMinter(&_CarbonCredit.TransactOpts, _minter)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_CarbonCredit *CarbonCreditTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_CarbonCredit *CarbonCreditSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.Contract.Transfer(&_CarbonCredit.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_CarbonCredit *CarbonCreditTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.Contract.Transfer(&_CarbonCredit.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_CarbonCredit *CarbonCreditTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_CarbonCredit *CarbonCreditSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.Contract.TransferFrom(&_CarbonCredit.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_CarbonCredit *CarbonCreditTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CarbonCredit.Contract.TransferFrom(&_CarbonCredit.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CarbonCredit *CarbonCreditTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CarbonCredit.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CarbonCredit *CarbonCreditSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CarbonCredit.Contract.TransferOwnership(&_CarbonCredit.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CarbonCredit *CarbonCreditTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CarbonCredit.Contract.TransferOwnership(&_CarbonCredit.TransactOpts, newOwner)
}

// CarbonCreditApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CarbonCredit contract.
type CarbonCreditApprovalIterator struct {
	Event *CarbonCreditApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CarbonCreditApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CarbonCreditApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CarbonCreditApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CarbonCreditApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CarbonCreditApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CarbonCreditApproval represents a Approval event raised by the CarbonCredit contract.
type CarbonCreditApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CarbonCredit *CarbonCreditFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CarbonCreditApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CarbonCredit.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CarbonCreditApprovalIterator{contract: _CarbonCredit.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CarbonCredit *CarbonCreditFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CarbonCreditApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CarbonCredit.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CarbonCreditApproval)
				if err := _CarbonCredit.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CarbonCredit *CarbonCreditFilterer) ParseApproval(log types.Log) (*CarbonCreditApproval, error) {
	event := new(CarbonCreditApproval)
	if err := _CarbonCredit.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CarbonCreditMinterUpdatedIterator is returned from FilterMinterUpdated and is used to iterate over the raw logs and unpacked data for MinterUpdated events raised by the CarbonCredit contract.
type CarbonCreditMinterUpdatedIterator struct {
	Event *CarbonCreditMinterUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CarbonCreditMinterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CarbonCreditMinterUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CarbonCreditMinterUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CarbonCreditMinterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CarbonCreditMinterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CarbonCreditMinterUpdated represents a MinterUpdated event raised by the CarbonCredit contract.
type CarbonCreditMinterUpdated struct {
	OldMinter common.Address
	NewMinter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinterUpdated is a free log retrieval operation binding the contract event 0x1cf2de25c5bf439ac0287061c3a0fa69b3b02867d0ccfd2ded34e42577050b73.
//
// Solidity: event MinterUpdated(address indexed oldMinter, address indexed newMinter)
func (_CarbonCredit *CarbonCreditFilterer) FilterMinterUpdated(opts *bind.FilterOpts, oldMinter []common.Address, newMinter []common.Address) (*CarbonCreditMinterUpdatedIterator, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _CarbonCredit.contract.FilterLogs(opts, "MinterUpdated", oldMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return &CarbonCreditMinterUpdatedIterator{contract: _CarbonCredit.contract, event: "MinterUpdated", logs: logs, sub: sub}, nil
}

// WatchMinterUpdated is a free log subscription operation binding the contract event 0x1cf2de25c5bf439ac0287061c3a0fa69b3b02867d0ccfd2ded34e42577050b73.
//
// Solidity: event MinterUpdated(address indexed oldMinter, address indexed newMinter)
func (_CarbonCredit *CarbonCreditFilterer) WatchMinterUpdated(opts *bind.WatchOpts, sink chan<- *CarbonCreditMinterUpdated, oldMinter []common.Address, newMinter []common.Address) (event.Subscription, error) {

	var oldMinterRule []interface{}
	for _, oldMinterItem := range oldMinter {
		oldMinterRule = append(oldMinterRule, oldMinterItem)
	}
	var newMinterRule []interface{}
	for _, newMinterItem := range newMinter {
		newMinterRule = append(newMinterRule, newMinterItem)
	}

	logs, sub, err := _CarbonCredit.contract.WatchLogs(opts, "MinterUpdated", oldMinterRule, newMinterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CarbonCreditMinterUpdated)
				if err := _CarbonCredit.contract.UnpackLog(event, "MinterUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinterUpdated is a log parse operation binding the contract event 0x1cf2de25c5bf439ac0287061c3a0fa69b3b02867d0ccfd2ded34e42577050b73.
//
// Solidity: event MinterUpdated(address indexed oldMinter, address indexed newMinter)
func (_CarbonCredit *CarbonCreditFilterer) ParseMinterUpdated(log types.Log) (*CarbonCreditMinterUpdated, error) {
	event := new(CarbonCreditMinterUpdated)
	if err := _CarbonCredit.contract.UnpackLog(event, "MinterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CarbonCreditOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CarbonCredit contract.
type CarbonCreditOwnershipTransferredIterator struct {
	Event *CarbonCreditOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CarbonCreditOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CarbonCreditOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CarbonCreditOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CarbonCreditOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CarbonCreditOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CarbonCreditOwnershipTransferred represents a OwnershipTransferred event raised by the CarbonCredit contract.
type CarbonCreditOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CarbonCredit *CarbonCreditFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CarbonCreditOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CarbonCredit.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CarbonCreditOwnershipTransferredIterator{contract: _CarbonCredit.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CarbonCredit *CarbonCreditFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CarbonCreditOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CarbonCredit.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CarbonCreditOwnershipTransferred)
				if err := _CarbonCredit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CarbonCredit *CarbonCreditFilterer) ParseOwnershipTransferred(log types.Log) (*CarbonCreditOwnershipTransferred, error) {
	event := new(CarbonCreditOwnershipTransferred)
	if err := _CarbonCredit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CarbonCreditTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CarbonCredit contract.
type CarbonCreditTransferIterator struct {
	Event *CarbonCreditTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CarbonCreditTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CarbonCreditTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CarbonCreditTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CarbonCreditTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CarbonCreditTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CarbonCreditTransfer represents a Transfer event raised by the CarbonCredit contract.
type CarbonCreditTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CarbonCredit *CarbonCreditFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CarbonCreditTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CarbonCredit.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CarbonCreditTransferIterator{contract: _CarbonCredit.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CarbonCredit *CarbonCreditFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CarbonCreditTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CarbonCredit.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CarbonCreditTransfer)
				if err := _CarbonCredit.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CarbonCredit *CarbonCreditFilterer) ParseTransfer(log types.Log) (*CarbonCreditTransfer, error) {
	event := new(CarbonCreditTransfer)
	if err := _CarbonCredit.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
