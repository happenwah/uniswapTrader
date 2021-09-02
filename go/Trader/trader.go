// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trader

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DEXTraderABI is the input ABI used to generate the binding from.
const DEXTraderABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETHAddress\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"convertETHtoWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAddressWithData\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amounts\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"}],\"name\":\"executeTrade\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// DEXTraderFuncSigs maps the 4-byte function signature to its string representation.
var DEXTraderFuncSigs = map[string]string{
	"2de9800a": "convertETHtoWETH(uint256)",
	"cec43477": "executeTrade(uint256,uint256,address,uint256)",
	"41c0e1b5": "kill()",
	"f14210a6": "withdrawETH(uint256)",
	"89476069": "withdrawToken(address)",
}

// DEXTraderBin is the compiled bytecode used for deploying new contracts.
var DEXTraderBin = "0x60c0604052604051610f57380380610f578339810160408190526100229161003e565b33606090811b6080521b6001600160601b03191660a05261006e565b60006020828403121561005057600080fd5b81516001600160a01b038116811461006757600080fd5b9392505050565b60805160601c60a05160601c610e6a6100ed6000396000818161017e01528181610250015281816108ef01526109c601526000818161013f015281816101fe0152818161033c015281816103710152818161042e0152818161053b0152818161096b01528181610ad301528181610b800152610bf40152610e6a6000f3fe60806040526004361061004e5760003560e01c80632de9800a146100ac57806341c0e1b5146100cc57806389476069146100e1578063cec4347714610101578063f14210a61461011457610055565b3661005557005b34801561006157600080fd5b5036156100aa5760405162461bcd60e51b815260206004820152601260248201527111905313109050d2d7d5149251d1d154915160721b60448201526064015b60405180910390fd5b005b3480156100b857600080fd5b506100aa6100c7366004610c7b565b610134565b3480156100d857600080fd5b506100aa6101f3565b3480156100ed57600080fd5b506100aa6100fc366004610c59565b610366565b6100aa61010f366004610cad565b610530565b34801561012057600080fd5b506100aa61012f366004610c7b565b610960565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461017c5760405162461bcd60e51b81526004016100a190610d6c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b1580156101d757600080fd5b505af11580156101eb573d6000803e3d6000fd5b505050505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461023b5760405162461bcd60e51b81526004016100a190610d6c565b6040516370a0823160e01b81523060048201527f0000000000000000000000000000000000000000000000000000000000000000906000906001600160a01b038316906370a082319060240160206040518083038186803b15801561029f57600080fd5b505afa1580156102b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102d79190610c94565b9050801561033a57604051632e1a7d4d60e01b8152600481018290526001600160a01b03831690632e1a7d4d90602401600060405180830381600087803b15801561032157600080fd5b505af1158015610335573d6000803e3d6000fd5b505050505b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316ff5b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103ae5760405162461bcd60e51b81526004016100a190610d6c565b6040516370a0823160e01b81523060048201526000906001600160a01b038316906370a082319060240160206040518083038186803b1580156103f057600080fd5b505afa158015610404573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104289190610c94565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b039081166024830152604480830185905283518084039091018152606490920183526020820180516001600160e01b031663a9059cbb60e01b1790529151929350600092918516916104a69190610d16565b6000604051808303816000865af19150503d80600081146104e3576040519150601f19603f3d011682016040523d82523d6000602084013e6104e8565b606091505b505090508061052b5760405162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b60448201526064016100a1565b505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105785760405162461bcd60e51b81526004016100a190610d6c565b60a884901c4381146105bd5760405162461bcd60e51b815260206004820152600e60248201526d57524f4e475f424c4f434b5f4e4f60901b60448201526064016100a1565b600060a0866001600160a81b0316901c6001146105db5760006105de565b60015b604080516001600160a01b0389811660248301819052604480840189905284518085039091018152606490930184526020830180516001600160e01b031663a9059cbb60e01b17905292519394509192608089901c926fffffffffffffffffffffffffffffffff8a1692600092918a16916106599190610d16565b6000604051808303816000865af19150503d8060008114610696576040519150601f19603f3d011682016040523d82523d6000602084013e61069b565b606091505b50509050806106de5760405162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b60448201526064016100a1565b84156107d857604080516000808252602082019092526001600160a01b0386169063022c0d9f906107189084908890309060448101610d32565b6040516020818303038152906040529060e01b6020820180516001600160e01b0383818316178352505050506040516107519190610d16565b6000604051808303816000865af19150503d806000811461078e576040519150601f19603f3d011682016040523d82523d6000602084013e610793565b606091505b50509050806107d25760405162461bcd60e51b815260206004820152600b60248201526a14d5d05417d1905253115160aa1b60448201526064016100a1565b506108c8565b604080516000808252602082019092526001600160a01b0386169063022c0d9f9061080c9087908590309060448101610d90565b6040516020818303038152906040529060e01b6020820180516001600160e01b0383818316178352505050506040516108459190610d16565b6000604051808303816000865af19150503d8060008114610882576040519150601f19603f3d011682016040523d82523d6000602084013e610887565b606091505b50509050806108c65760405162461bcd60e51b815260206004820152600b60248201526a14d5d05417d1905253115160aa1b60448201526064016100a1565b505b34156108d357610954565b811561095457604051632e1a7d4d60e01b8152600481018390527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632e1a7d4d90602401600060405180830381600087803b15801561093b57600080fd5b505af115801561094f573d6000803e3d6000fd5b505050505b50505050505050505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146109a85760405162461bcd60e51b81526004016100a190610d6c565b4780821115610be7576040516370a0823160e01b81523060048201527f0000000000000000000000000000000000000000000000000000000000000000906000906001600160a01b038316906370a082319060240160206040518083038186803b158015610a1557600080fd5b505afa158015610a29573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a4d9190610c94565b905083610a5a8483610dc3565b10610b24576001600160a01b038216632e1a7d4d610a788587610ddb565b6040518263ffffffff1660e01b8152600401610a9691815260200190565b600060405180830381600087803b158015610ab057600080fd5b505af1158015610ac4573d6000803e3d6000fd5b50506040516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925086156108fc02915086906000818181858888f19350505050158015610b1e573d6000803e3d6000fd5b50610be1565b604051632e1a7d4d60e01b8152600481018290526001600160a01b03831690632e1a7d4d90602401600060405180830381600087803b158015610b6657600080fd5b505af1158015610b7a573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166108fc8285610bb79190610dc3565b6040518115909202916000818181858888f19350505050158015610bdf573d6000803e3d6000fd5b505b50505050565b6040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169083156108fc029084906000818181858888f1935050505015801561052b573d6000803e3d6000fd5b80356001600160a01b0381168114610c5457600080fd5b919050565b600060208284031215610c6b57600080fd5b610c7482610c3d565b9392505050565b600060208284031215610c8d57600080fd5b5035919050565b600060208284031215610ca657600080fd5b5051919050565b60008060008060808587031215610cc357600080fd5b8435935060208501359250610cda60408601610c3d565b9396929550929360600135925050565b60008151808452610d02816020860160208601610df2565b601f01601f19169290920160200192915050565b60008251610d28818460208701610df2565b9190910192915050565b60ff8516815283602082015260018060a01b0383166040820152608060608201526000610d626080830184610cea565b9695505050505050565b6020808252600a908201526927a7262cafa7aba722a960b11b604082015260600190565b84815260ff841660208201526001600160a01b0383166040820152608060608201819052600090610d6290830184610cea565b60008219821115610dd657610dd6610e1e565b500190565b600082821015610ded57610ded610e1e565b500390565b60005b83811015610e0d578181015183820152602001610df5565b83811115610be15750506000910152565b634e487b7160e01b600052601160045260246000fdfea264697066735822122037b6970fee2c107affd806eb54c094c4b810193150fe952dc1a7379e2d90de0764736f6c63430008060033"

// DeployDEXTrader deploys a new Ethereum contract, binding an instance of DEXTrader to it.
func DeployDEXTrader(auth *bind.TransactOpts, backend bind.ContractBackend, _WETHAddress common.Address) (common.Address, *types.Transaction, *DEXTrader, error) {
	parsed, err := abi.JSON(strings.NewReader(DEXTraderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DEXTraderBin), backend, _WETHAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DEXTrader{DEXTraderCaller: DEXTraderCaller{contract: contract}, DEXTraderTransactor: DEXTraderTransactor{contract: contract}, DEXTraderFilterer: DEXTraderFilterer{contract: contract}}, nil
}

// DEXTrader is an auto generated Go binding around an Ethereum contract.
type DEXTrader struct {
	DEXTraderCaller     // Read-only binding to the contract
	DEXTraderTransactor // Write-only binding to the contract
	DEXTraderFilterer   // Log filterer for contract events
}

// DEXTraderCaller is an auto generated read-only Go binding around an Ethereum contract.
type DEXTraderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DEXTraderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DEXTraderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DEXTraderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DEXTraderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DEXTraderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DEXTraderSession struct {
	Contract     *DEXTrader        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DEXTraderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DEXTraderCallerSession struct {
	Contract *DEXTraderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DEXTraderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DEXTraderTransactorSession struct {
	Contract     *DEXTraderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DEXTraderRaw is an auto generated low-level Go binding around an Ethereum contract.
type DEXTraderRaw struct {
	Contract *DEXTrader // Generic contract binding to access the raw methods on
}

// DEXTraderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DEXTraderCallerRaw struct {
	Contract *DEXTraderCaller // Generic read-only contract binding to access the raw methods on
}

// DEXTraderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DEXTraderTransactorRaw struct {
	Contract *DEXTraderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDEXTrader creates a new instance of DEXTrader, bound to a specific deployed contract.
func NewDEXTrader(address common.Address, backend bind.ContractBackend) (*DEXTrader, error) {
	contract, err := bindDEXTrader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DEXTrader{DEXTraderCaller: DEXTraderCaller{contract: contract}, DEXTraderTransactor: DEXTraderTransactor{contract: contract}, DEXTraderFilterer: DEXTraderFilterer{contract: contract}}, nil
}

// NewDEXTraderCaller creates a new read-only instance of DEXTrader, bound to a specific deployed contract.
func NewDEXTraderCaller(address common.Address, caller bind.ContractCaller) (*DEXTraderCaller, error) {
	contract, err := bindDEXTrader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DEXTraderCaller{contract: contract}, nil
}

// NewDEXTraderTransactor creates a new write-only instance of DEXTrader, bound to a specific deployed contract.
func NewDEXTraderTransactor(address common.Address, transactor bind.ContractTransactor) (*DEXTraderTransactor, error) {
	contract, err := bindDEXTrader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DEXTraderTransactor{contract: contract}, nil
}

// NewDEXTraderFilterer creates a new log filterer instance of DEXTrader, bound to a specific deployed contract.
func NewDEXTraderFilterer(address common.Address, filterer bind.ContractFilterer) (*DEXTraderFilterer, error) {
	contract, err := bindDEXTrader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DEXTraderFilterer{contract: contract}, nil
}

// bindDEXTrader binds a generic wrapper to an already deployed contract.
func bindDEXTrader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DEXTraderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DEXTrader *DEXTraderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DEXTrader.Contract.DEXTraderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DEXTrader *DEXTraderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DEXTrader.Contract.DEXTraderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DEXTrader *DEXTraderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DEXTrader.Contract.DEXTraderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DEXTrader *DEXTraderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DEXTrader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DEXTrader *DEXTraderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DEXTrader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DEXTrader *DEXTraderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DEXTrader.Contract.contract.Transact(opts, method, params...)
}

// ConvertETHtoWETH is a paid mutator transaction binding the contract method 0x2de9800a.
//
// Solidity: function convertETHtoWETH(uint256 amount) returns()
func (_DEXTrader *DEXTraderTransactor) ConvertETHtoWETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _DEXTrader.contract.Transact(opts, "convertETHtoWETH", amount)
}

// ConvertETHtoWETH is a paid mutator transaction binding the contract method 0x2de9800a.
//
// Solidity: function convertETHtoWETH(uint256 amount) returns()
func (_DEXTrader *DEXTraderSession) ConvertETHtoWETH(amount *big.Int) (*types.Transaction, error) {
	return _DEXTrader.Contract.ConvertETHtoWETH(&_DEXTrader.TransactOpts, amount)
}

// ConvertETHtoWETH is a paid mutator transaction binding the contract method 0x2de9800a.
//
// Solidity: function convertETHtoWETH(uint256 amount) returns()
func (_DEXTrader *DEXTraderTransactorSession) ConvertETHtoWETH(amount *big.Int) (*types.Transaction, error) {
	return _DEXTrader.Contract.ConvertETHtoWETH(&_DEXTrader.TransactOpts, amount)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0xcec43477.
//
// Solidity: function executeTrade(uint256 poolAddressWithData, uint256 amounts, address from, uint256 inputAmount) payable returns()
func (_DEXTrader *DEXTraderTransactor) ExecuteTrade(opts *bind.TransactOpts, poolAddressWithData *big.Int, amounts *big.Int, from common.Address, inputAmount *big.Int) (*types.Transaction, error) {
	return _DEXTrader.contract.Transact(opts, "executeTrade", poolAddressWithData, amounts, from, inputAmount)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0xcec43477.
//
// Solidity: function executeTrade(uint256 poolAddressWithData, uint256 amounts, address from, uint256 inputAmount) payable returns()
func (_DEXTrader *DEXTraderSession) ExecuteTrade(poolAddressWithData *big.Int, amounts *big.Int, from common.Address, inputAmount *big.Int) (*types.Transaction, error) {
	return _DEXTrader.Contract.ExecuteTrade(&_DEXTrader.TransactOpts, poolAddressWithData, amounts, from, inputAmount)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0xcec43477.
//
// Solidity: function executeTrade(uint256 poolAddressWithData, uint256 amounts, address from, uint256 inputAmount) payable returns()
func (_DEXTrader *DEXTraderTransactorSession) ExecuteTrade(poolAddressWithData *big.Int, amounts *big.Int, from common.Address, inputAmount *big.Int) (*types.Transaction, error) {
	return _DEXTrader.Contract.ExecuteTrade(&_DEXTrader.TransactOpts, poolAddressWithData, amounts, from, inputAmount)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_DEXTrader *DEXTraderTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DEXTrader.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_DEXTrader *DEXTraderSession) Kill() (*types.Transaction, error) {
	return _DEXTrader.Contract.Kill(&_DEXTrader.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_DEXTrader *DEXTraderTransactorSession) Kill() (*types.Transaction, error) {
	return _DEXTrader.Contract.Kill(&_DEXTrader.TransactOpts)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_DEXTrader *DEXTraderTransactor) WithdrawETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _DEXTrader.contract.Transact(opts, "withdrawETH", amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_DEXTrader *DEXTraderSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _DEXTrader.Contract.WithdrawETH(&_DEXTrader.TransactOpts, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_DEXTrader *DEXTraderTransactorSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _DEXTrader.Contract.WithdrawETH(&_DEXTrader.TransactOpts, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address token) returns()
func (_DEXTrader *DEXTraderTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _DEXTrader.contract.Transact(opts, "withdrawToken", token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address token) returns()
func (_DEXTrader *DEXTraderSession) WithdrawToken(token common.Address) (*types.Transaction, error) {
	return _DEXTrader.Contract.WithdrawToken(&_DEXTrader.TransactOpts, token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address token) returns()
func (_DEXTrader *DEXTraderTransactorSession) WithdrawToken(token common.Address) (*types.Transaction, error) {
	return _DEXTrader.Contract.WithdrawToken(&_DEXTrader.TransactOpts, token)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_DEXTrader *DEXTraderTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _DEXTrader.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_DEXTrader *DEXTraderSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DEXTrader.Contract.Fallback(&_DEXTrader.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_DEXTrader *DEXTraderTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _DEXTrader.Contract.Fallback(&_DEXTrader.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DEXTrader *DEXTraderTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DEXTrader.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DEXTrader *DEXTraderSession) Receive() (*types.Transaction, error) {
	return _DEXTrader.Contract.Receive(&_DEXTrader.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DEXTrader *DEXTraderTransactorSession) Receive() (*types.Transaction, error) {
	return _DEXTrader.Contract.Receive(&_DEXTrader.TransactOpts)
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"a9059cbb": "transfer(address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, owner)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// IWETHABI is the input ABI used to generate the binding from.
const IWETHABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IWETHFuncSigs maps the 4-byte function signature to its string representation.
var IWETHFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"d0e30db0": "deposit()",
	"a9059cbb": "transfer(address,uint256)",
	"2e1a7d4d": "withdraw(uint256)",
}

// IWETH is an auto generated Go binding around an Ethereum contract.
type IWETH struct {
	IWETHCaller     // Read-only binding to the contract
	IWETHTransactor // Write-only binding to the contract
	IWETHFilterer   // Log filterer for contract events
}

// IWETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWETHSession struct {
	Contract     *IWETH            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWETHCallerSession struct {
	Contract *IWETHCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IWETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWETHTransactorSession struct {
	Contract     *IWETHTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWETHRaw struct {
	Contract *IWETH // Generic contract binding to access the raw methods on
}

// IWETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWETHCallerRaw struct {
	Contract *IWETHCaller // Generic read-only contract binding to access the raw methods on
}

// IWETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWETHTransactorRaw struct {
	Contract *IWETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWETH creates a new instance of IWETH, bound to a specific deployed contract.
func NewIWETH(address common.Address, backend bind.ContractBackend) (*IWETH, error) {
	contract, err := bindIWETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWETH{IWETHCaller: IWETHCaller{contract: contract}, IWETHTransactor: IWETHTransactor{contract: contract}, IWETHFilterer: IWETHFilterer{contract: contract}}, nil
}

// NewIWETHCaller creates a new read-only instance of IWETH, bound to a specific deployed contract.
func NewIWETHCaller(address common.Address, caller bind.ContractCaller) (*IWETHCaller, error) {
	contract, err := bindIWETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWETHCaller{contract: contract}, nil
}

// NewIWETHTransactor creates a new write-only instance of IWETH, bound to a specific deployed contract.
func NewIWETHTransactor(address common.Address, transactor bind.ContractTransactor) (*IWETHTransactor, error) {
	contract, err := bindIWETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWETHTransactor{contract: contract}, nil
}

// NewIWETHFilterer creates a new log filterer instance of IWETH, bound to a specific deployed contract.
func NewIWETHFilterer(address common.Address, filterer bind.ContractFilterer) (*IWETHFilterer, error) {
	contract, err := bindIWETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWETHFilterer{contract: contract}, nil
}

// bindIWETH binds a generic wrapper to an already deployed contract.
func bindIWETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH *IWETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH.Contract.IWETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH *IWETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH.Contract.IWETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH *IWETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH.Contract.IWETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH *IWETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH *IWETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH *IWETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IWETH *IWETHCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IWETH *IWETHSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IWETH.Contract.BalanceOf(&_IWETH.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IWETH *IWETHCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IWETH.Contract.BalanceOf(&_IWETH.CallOpts, owner)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH *IWETHTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH *IWETHSession) Deposit() (*types.Transaction, error) {
	return _IWETH.Contract.Deposit(&_IWETH.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH *IWETHTransactorSession) Deposit() (*types.Transaction, error) {
	return _IWETH.Contract.Deposit(&_IWETH.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IWETH *IWETHTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IWETH *IWETHSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Transfer(&_IWETH.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IWETH *IWETHTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Transfer(&_IWETH.TransactOpts, to, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH *IWETHTransactor) Withdraw(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH.contract.Transact(opts, "withdraw", arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH *IWETHSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Withdraw(&_IWETH.TransactOpts, arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH *IWETHTransactorSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Withdraw(&_IWETH.TransactOpts, arg0)
}
