// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractBindings

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

// FlashbotsTraderABI is the input ABI used to generate the binding from.
const FlashbotsTraderABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_WETHAddress\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"convertETHtoWETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAddressWithData\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amounts\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"}],\"name\":\"executeTrade\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// FlashbotsTraderFuncSigs maps the 4-byte function signature to its string representation.
var FlashbotsTraderFuncSigs = map[string]string{
	"2de9800a": "convertETHtoWETH(uint256)",
	"cec43477": "executeTrade(uint256,uint256,address,uint256)",
	"41c0e1b5": "kill()",
	"f14210a6": "withdrawETH(uint256)",
	"89476069": "withdrawToken(address)",
}

// FlashbotsTraderBin is the compiled bytecode used for deploying new contracts.
var FlashbotsTraderBin = "0x60c0604052604051610ffa380380610ffa8339810160408190526100229161003e565b33606090811b6080521b6001600160601b03191660a05261006e565b60006020828403121561005057600080fd5b81516001600160a01b038116811461006757600080fd5b9392505050565b60805160601c60a05160601c610f0d6100ed6000396000818161017e01528181610250015281816109450152610a6901526000818161013f015281816101fe0152818161033c015281816103710152818161042e0152818161053b01528181610a0e01528181610b7601528181610c230152610c970152610f0d6000f3fe60806040526004361061004e5760003560e01c80632de9800a146100ac57806341c0e1b5146100cc57806389476069146100e1578063cec4347714610101578063f14210a61461011457610055565b3661005557005b34801561006157600080fd5b5036156100aa5760405162461bcd60e51b815260206004820152601260248201527111905313109050d2d7d5149251d1d154915160721b60448201526064015b60405180910390fd5b005b3480156100b857600080fd5b506100aa6100c7366004610d1e565b610134565b3480156100d857600080fd5b506100aa6101f3565b3480156100ed57600080fd5b506100aa6100fc366004610cfc565b610366565b6100aa61010f366004610d50565b610530565b34801561012057600080fd5b506100aa61012f366004610d1e565b610a03565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461017c5760405162461bcd60e51b81526004016100a190610e0f565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663d0e30db0826040518263ffffffff1660e01b81526004016000604051808303818588803b1580156101d757600080fd5b505af11580156101eb573d6000803e3d6000fd5b505050505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461023b5760405162461bcd60e51b81526004016100a190610e0f565b6040516370a0823160e01b81523060048201527f0000000000000000000000000000000000000000000000000000000000000000906000906001600160a01b038316906370a082319060240160206040518083038186803b15801561029f57600080fd5b505afa1580156102b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102d79190610d37565b9050801561033a57604051632e1a7d4d60e01b8152600481018290526001600160a01b03831690632e1a7d4d90602401600060405180830381600087803b15801561032157600080fd5b505af1158015610335573d6000803e3d6000fd5b505050505b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316ff5b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146103ae5760405162461bcd60e51b81526004016100a190610e0f565b6040516370a0823160e01b81523060048201526000906001600160a01b038316906370a082319060240160206040518083038186803b1580156103f057600080fd5b505afa158015610404573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104289190610d37565b604080517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b039081166024830152604480830185905283518084039091018152606490920183526020820180516001600160e01b031663a9059cbb60e01b1790529151929350600092918516916104a69190610db9565b6000604051808303816000865af19150503d80600081146104e3576040519150601f19603f3d011682016040523d82523d6000602084013e6104e8565b606091505b505090508061052b5760405162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b60448201526064016100a1565b505050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105785760405162461bcd60e51b81526004016100a190610e0f565b60a884901c43811480610589575080155b6105c65760405162461bcd60e51b815260206004820152600e60248201526d57524f4e475f424c4f434b5f4e4f60901b60448201526064016100a1565b600060a0866001600160a81b0316901c6001146105e45760006105e7565b60015b604080516001600160a01b0389811660248301819052604480840189905284518085039091018152606490930184526020830180516001600160e01b031663a9059cbb60e01b17905292519394509192608089901c926fffffffffffffffffffffffffffffffff8a1692600092918a16916106629190610db9565b6000604051808303816000865af19150503d806000811461069f576040519150601f19603f3d011682016040523d82523d6000602084013e6106a4565b606091505b50509050806106e75760405162461bcd60e51b815260206004820152600f60248201526e1514905394d1915497d19052531151608a1b60448201526064016100a1565b84156107e157604080516000808252602082019092526001600160a01b0386169063022c0d9f906107219084908890309060448101610dd5565b6040516020818303038152906040529060e01b6020820180516001600160e01b03838183161783525050505060405161075a9190610db9565b6000604051808303816000865af19150503d8060008114610797576040519150601f19603f3d011682016040523d82523d6000602084013e61079c565b606091505b50509050806107db5760405162461bcd60e51b815260206004820152600b60248201526a14d5d05417d1905253115160aa1b60448201526064016100a1565b506108d1565b604080516000808252602082019092526001600160a01b0386169063022c0d9f906108159087908590309060448101610e33565b6040516020818303038152906040529060e01b6020820180516001600160e01b03838183161783525050505060405161084e9190610db9565b6000604051808303816000865af19150503d806000811461088b576040519150601f19603f3d011682016040523d82523d6000602084013e610890565b606091505b50509050806108cf5760405162461bcd60e51b815260206004820152600b60248201526a14d5d05417d1905253115160aa1b60448201526064016100a1565b505b34156109295760405141906161a89034906000818181858888f193505050503d806000811461091c576040519150601f19603f3d011682016040523d82523d6000602084013e610921565b606091505b5050506109f7565b81156109f757604051632e1a7d4d60e01b8152600481018390527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690632e1a7d4d90602401600060405180830381600087803b15801561099157600080fd5b505af11580156109a5573d6000803e3d6000fd5b50506040514192506161a8915084906000818181858888f193505050503d80600081146109ee576040519150601f19603f3d011682016040523d82523d6000602084013e6109f3565b606091505b5050505b50505050505050505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610a4b5760405162461bcd60e51b81526004016100a190610e0f565b4780821115610c8a576040516370a0823160e01b81523060048201527f0000000000000000000000000000000000000000000000000000000000000000906000906001600160a01b038316906370a082319060240160206040518083038186803b158015610ab857600080fd5b505afa158015610acc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610af09190610d37565b905083610afd8483610e66565b10610bc7576001600160a01b038216632e1a7d4d610b1b8587610e7e565b6040518263ffffffff1660e01b8152600401610b3991815260200190565b600060405180830381600087803b158015610b5357600080fd5b505af1158015610b67573d6000803e3d6000fd5b50506040516001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016925086156108fc02915086906000818181858888f19350505050158015610bc1573d6000803e3d6000fd5b50610c84565b604051632e1a7d4d60e01b8152600481018290526001600160a01b03831690632e1a7d4d90602401600060405180830381600087803b158015610c0957600080fd5b505af1158015610c1d573d6000803e3d6000fd5b505050507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166108fc8285610c5a9190610e66565b6040518115909202916000818181858888f19350505050158015610c82573d6000803e3d6000fd5b505b50505050565b6040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169083156108fc029084906000818181858888f1935050505015801561052b573d6000803e3d6000fd5b80356001600160a01b0381168114610cf757600080fd5b919050565b600060208284031215610d0e57600080fd5b610d1782610ce0565b9392505050565b600060208284031215610d3057600080fd5b5035919050565b600060208284031215610d4957600080fd5b5051919050565b60008060008060808587031215610d6657600080fd5b8435935060208501359250610d7d60408601610ce0565b9396929550929360600135925050565b60008151808452610da5816020860160208601610e95565b601f01601f19169290920160200192915050565b60008251610dcb818460208701610e95565b9190910192915050565b60ff8516815283602082015260018060a01b0383166040820152608060608201526000610e056080830184610d8d565b9695505050505050565b6020808252600a908201526927a7262cafa7aba722a960b11b604082015260600190565b84815260ff841660208201526001600160a01b0383166040820152608060608201819052600090610e0590830184610d8d565b60008219821115610e7957610e79610ec1565b500190565b600082821015610e9057610e90610ec1565b500390565b60005b83811015610eb0578181015183820152602001610e98565b83811115610c845750506000910152565b634e487b7160e01b600052601160045260246000fdfea264697066735822122069756e669fa76a37a8def0c2b1b8872e883feaf0f70e33407f6ff7484c8c6db764736f6c63430008070033"

// DeployFlashbotsTrader deploys a new Ethereum contract, binding an instance of FlashbotsTrader to it.
func DeployFlashbotsTrader(auth *bind.TransactOpts, backend bind.ContractBackend, _WETHAddress common.Address) (common.Address, *types.Transaction, *FlashbotsTrader, error) {
	parsed, err := abi.JSON(strings.NewReader(FlashbotsTraderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FlashbotsTraderBin), backend, _WETHAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FlashbotsTrader{FlashbotsTraderCaller: FlashbotsTraderCaller{contract: contract}, FlashbotsTraderTransactor: FlashbotsTraderTransactor{contract: contract}, FlashbotsTraderFilterer: FlashbotsTraderFilterer{contract: contract}}, nil
}

// FlashbotsTrader is an auto generated Go binding around an Ethereum contract.
type FlashbotsTrader struct {
	FlashbotsTraderCaller     // Read-only binding to the contract
	FlashbotsTraderTransactor // Write-only binding to the contract
	FlashbotsTraderFilterer   // Log filterer for contract events
}

// FlashbotsTraderCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlashbotsTraderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashbotsTraderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlashbotsTraderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashbotsTraderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlashbotsTraderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashbotsTraderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlashbotsTraderSession struct {
	Contract     *FlashbotsTrader  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FlashbotsTraderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlashbotsTraderCallerSession struct {
	Contract *FlashbotsTraderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// FlashbotsTraderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlashbotsTraderTransactorSession struct {
	Contract     *FlashbotsTraderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// FlashbotsTraderRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlashbotsTraderRaw struct {
	Contract *FlashbotsTrader // Generic contract binding to access the raw methods on
}

// FlashbotsTraderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlashbotsTraderCallerRaw struct {
	Contract *FlashbotsTraderCaller // Generic read-only contract binding to access the raw methods on
}

// FlashbotsTraderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlashbotsTraderTransactorRaw struct {
	Contract *FlashbotsTraderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlashbotsTrader creates a new instance of FlashbotsTrader, bound to a specific deployed contract.
func NewFlashbotsTrader(address common.Address, backend bind.ContractBackend) (*FlashbotsTrader, error) {
	contract, err := bindFlashbotsTrader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FlashbotsTrader{FlashbotsTraderCaller: FlashbotsTraderCaller{contract: contract}, FlashbotsTraderTransactor: FlashbotsTraderTransactor{contract: contract}, FlashbotsTraderFilterer: FlashbotsTraderFilterer{contract: contract}}, nil
}

// NewFlashbotsTraderCaller creates a new read-only instance of FlashbotsTrader, bound to a specific deployed contract.
func NewFlashbotsTraderCaller(address common.Address, caller bind.ContractCaller) (*FlashbotsTraderCaller, error) {
	contract, err := bindFlashbotsTrader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlashbotsTraderCaller{contract: contract}, nil
}

// NewFlashbotsTraderTransactor creates a new write-only instance of FlashbotsTrader, bound to a specific deployed contract.
func NewFlashbotsTraderTransactor(address common.Address, transactor bind.ContractTransactor) (*FlashbotsTraderTransactor, error) {
	contract, err := bindFlashbotsTrader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlashbotsTraderTransactor{contract: contract}, nil
}

// NewFlashbotsTraderFilterer creates a new log filterer instance of FlashbotsTrader, bound to a specific deployed contract.
func NewFlashbotsTraderFilterer(address common.Address, filterer bind.ContractFilterer) (*FlashbotsTraderFilterer, error) {
	contract, err := bindFlashbotsTrader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlashbotsTraderFilterer{contract: contract}, nil
}

// bindFlashbotsTrader binds a generic wrapper to an already deployed contract.
func bindFlashbotsTrader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FlashbotsTraderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashbotsTrader *FlashbotsTraderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashbotsTrader.Contract.FlashbotsTraderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashbotsTrader *FlashbotsTraderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashbotsTrader.Contract.FlashbotsTraderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashbotsTrader *FlashbotsTraderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashbotsTrader.Contract.FlashbotsTraderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FlashbotsTrader *FlashbotsTraderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FlashbotsTrader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FlashbotsTrader *FlashbotsTraderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashbotsTrader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FlashbotsTrader *FlashbotsTraderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FlashbotsTrader.Contract.contract.Transact(opts, method, params...)
}

// ConvertETHtoWETH is a paid mutator transaction binding the contract method 0x2de9800a.
//
// Solidity: function convertETHtoWETH(uint256 amount) returns()
func (_FlashbotsTrader *FlashbotsTraderTransactor) ConvertETHtoWETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _FlashbotsTrader.contract.Transact(opts, "convertETHtoWETH", amount)
}

// ConvertETHtoWETH is a paid mutator transaction binding the contract method 0x2de9800a.
//
// Solidity: function convertETHtoWETH(uint256 amount) returns()
func (_FlashbotsTrader *FlashbotsTraderSession) ConvertETHtoWETH(amount *big.Int) (*types.Transaction, error) {
	return _FlashbotsTrader.Contract.ConvertETHtoWETH(&_FlashbotsTrader.TransactOpts, amount)
}

// ConvertETHtoWETH is a paid mutator transaction binding the contract method 0x2de9800a.
//
// Solidity: function convertETHtoWETH(uint256 amount) returns()
func (_FlashbotsTrader *FlashbotsTraderTransactorSession) ConvertETHtoWETH(amount *big.Int) (*types.Transaction, error) {
	return _FlashbotsTrader.Contract.ConvertETHtoWETH(&_FlashbotsTrader.TransactOpts, amount)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0xcec43477.
//
// Solidity: function executeTrade(uint256 poolAddressWithData, uint256 amounts, address from, uint256 inputAmount) payable returns()
func (_FlashbotsTrader *FlashbotsTraderTransactor) ExecuteTrade(opts *bind.TransactOpts, poolAddressWithData *big.Int, amounts *big.Int, from common.Address, inputAmount *big.Int) (*types.Transaction, error) {
	return _FlashbotsTrader.contract.Transact(opts, "executeTrade", poolAddressWithData, amounts, from, inputAmount)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0xcec43477.
//
// Solidity: function executeTrade(uint256 poolAddressWithData, uint256 amounts, address from, uint256 inputAmount) payable returns()
func (_FlashbotsTrader *FlashbotsTraderSession) ExecuteTrade(poolAddressWithData *big.Int, amounts *big.Int, from common.Address, inputAmount *big.Int) (*types.Transaction, error) {
	return _FlashbotsTrader.Contract.ExecuteTrade(&_FlashbotsTrader.TransactOpts, poolAddressWithData, amounts, from, inputAmount)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0xcec43477.
//
// Solidity: function executeTrade(uint256 poolAddressWithData, uint256 amounts, address from, uint256 inputAmount) payable returns()
func (_FlashbotsTrader *FlashbotsTraderTransactorSession) ExecuteTrade(poolAddressWithData *big.Int, amounts *big.Int, from common.Address, inputAmount *big.Int) (*types.Transaction, error) {
	return _FlashbotsTrader.Contract.ExecuteTrade(&_FlashbotsTrader.TransactOpts, poolAddressWithData, amounts, from, inputAmount)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_FlashbotsTrader *FlashbotsTraderTransactor) Kill(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashbotsTrader.contract.Transact(opts, "kill")
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_FlashbotsTrader *FlashbotsTraderSession) Kill() (*types.Transaction, error) {
	return _FlashbotsTrader.Contract.Kill(&_FlashbotsTrader.TransactOpts)
}

// Kill is a paid mutator transaction binding the contract method 0x41c0e1b5.
//
// Solidity: function kill() returns()
func (_FlashbotsTrader *FlashbotsTraderTransactorSession) Kill() (*types.Transaction, error) {
	return _FlashbotsTrader.Contract.Kill(&_FlashbotsTrader.TransactOpts)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_FlashbotsTrader *FlashbotsTraderTransactor) WithdrawETH(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _FlashbotsTrader.contract.Transact(opts, "withdrawETH", amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_FlashbotsTrader *FlashbotsTraderSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _FlashbotsTrader.Contract.WithdrawETH(&_FlashbotsTrader.TransactOpts, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 amount) returns()
func (_FlashbotsTrader *FlashbotsTraderTransactorSession) WithdrawETH(amount *big.Int) (*types.Transaction, error) {
	return _FlashbotsTrader.Contract.WithdrawETH(&_FlashbotsTrader.TransactOpts, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address token) returns()
func (_FlashbotsTrader *FlashbotsTraderTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _FlashbotsTrader.contract.Transact(opts, "withdrawToken", token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address token) returns()
func (_FlashbotsTrader *FlashbotsTraderSession) WithdrawToken(token common.Address) (*types.Transaction, error) {
	return _FlashbotsTrader.Contract.WithdrawToken(&_FlashbotsTrader.TransactOpts, token)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address token) returns()
func (_FlashbotsTrader *FlashbotsTraderTransactorSession) WithdrawToken(token common.Address) (*types.Transaction, error) {
	return _FlashbotsTrader.Contract.WithdrawToken(&_FlashbotsTrader.TransactOpts, token)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_FlashbotsTrader *FlashbotsTraderTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _FlashbotsTrader.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_FlashbotsTrader *FlashbotsTraderSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FlashbotsTrader.Contract.Fallback(&_FlashbotsTrader.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_FlashbotsTrader *FlashbotsTraderTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FlashbotsTrader.Contract.Fallback(&_FlashbotsTrader.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FlashbotsTrader *FlashbotsTraderTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FlashbotsTrader.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FlashbotsTrader *FlashbotsTraderSession) Receive() (*types.Transaction, error) {
	return _FlashbotsTrader.Contract.Receive(&_FlashbotsTrader.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FlashbotsTrader *FlashbotsTraderTransactorSession) Receive() (*types.Transaction, error) {
	return _FlashbotsTrader.Contract.Receive(&_FlashbotsTrader.TransactOpts)
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
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

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
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

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20Session) Decimals() (uint8, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20 *IERC20CallerSession) Decimals() (uint8, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20Session) Name() (string, error) {
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20 *IERC20CallerSession) Name() (string, error) {
	return _IERC20.Contract.Name(&_IERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20Session) Symbol() (string, error) {
	return _IERC20.Contract.Symbol(&_IERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20 *IERC20CallerSession) Symbol() (string, error) {
	return _IERC20.Contract.Symbol(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
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

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWETHABI is the input ABI used to generate the binding from.
const IWETHABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IWETHFuncSigs maps the 4-byte function signature to its string representation.
var IWETHFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"d0e30db0": "deposit()",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
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

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH *IWETHCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH *IWETHSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IWETH.Contract.Allowance(&_IWETH.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH *IWETHCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IWETH.Contract.Allowance(&_IWETH.CallOpts, owner, spender)
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

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IWETH *IWETHCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IWETH.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IWETH *IWETHSession) Decimals() (uint8, error) {
	return _IWETH.Contract.Decimals(&_IWETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IWETH *IWETHCallerSession) Decimals() (uint8, error) {
	return _IWETH.Contract.Decimals(&_IWETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IWETH *IWETHCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IWETH.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IWETH *IWETHSession) Name() (string, error) {
	return _IWETH.Contract.Name(&_IWETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IWETH *IWETHCallerSession) Name() (string, error) {
	return _IWETH.Contract.Name(&_IWETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IWETH *IWETHCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IWETH.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IWETH *IWETHSession) Symbol() (string, error) {
	return _IWETH.Contract.Symbol(&_IWETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IWETH *IWETHCallerSession) Symbol() (string, error) {
	return _IWETH.Contract.Symbol(&_IWETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH *IWETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWETH.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH *IWETHSession) TotalSupply() (*big.Int, error) {
	return _IWETH.Contract.TotalSupply(&_IWETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH *IWETHCallerSession) TotalSupply() (*big.Int, error) {
	return _IWETH.Contract.TotalSupply(&_IWETH.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IWETH *IWETHTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IWETH *IWETHSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Approve(&_IWETH.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IWETH *IWETHTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.Approve(&_IWETH.TransactOpts, spender, value)
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

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IWETH *IWETHTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IWETH *IWETHSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.TransferFrom(&_IWETH.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IWETH *IWETHTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IWETH.Contract.TransferFrom(&_IWETH.TransactOpts, from, to, value)
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

// IWETHApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IWETH contract.
type IWETHApprovalIterator struct {
	Event *IWETHApproval // Event containing the contract specifics and raw log

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
func (it *IWETHApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWETHApproval)
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
		it.Event = new(IWETHApproval)
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
func (it *IWETHApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWETHApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWETHApproval represents a Approval event raised by the IWETH contract.
type IWETHApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWETH *IWETHFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IWETHApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IWETH.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IWETHApprovalIterator{contract: _IWETH.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWETH *IWETHFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IWETHApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IWETH.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWETHApproval)
				if err := _IWETH.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IWETH *IWETHFilterer) ParseApproval(log types.Log) (*IWETHApproval, error) {
	event := new(IWETHApproval)
	if err := _IWETH.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWETHTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IWETH contract.
type IWETHTransferIterator struct {
	Event *IWETHTransfer // Event containing the contract specifics and raw log

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
func (it *IWETHTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWETHTransfer)
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
		it.Event = new(IWETHTransfer)
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
func (it *IWETHTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWETHTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWETHTransfer represents a Transfer event raised by the IWETH contract.
type IWETHTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWETH *IWETHFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IWETHTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IWETH.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IWETHTransferIterator{contract: _IWETH.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWETH *IWETHFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IWETHTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IWETH.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWETHTransfer)
				if err := _IWETH.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IWETH *IWETHFilterer) ParseTransfer(log types.Log) (*IWETHTransfer, error) {
	event := new(IWETHTransfer)
	if err := _IWETH.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
