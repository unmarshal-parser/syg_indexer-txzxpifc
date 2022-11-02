// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"errors"
	"github.com/ethereum/go-ethereum/common/hexutil"
	sdkTransactionTypes "github.com/eucrypt/unmarshal-go-sdk/pkg/transaction_details/types"
	"github.com/onrik/ethrpc"
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
	_ = sdkTransactionTypes.RawTransaction{}
	_ = errors.New("")
	_ = hexutil.Big{}
	_ = ethrpc.Transaction{}
	_ = abi.ABI{}
	_ = abi.ABI{}
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PermitHelperPermit is an auto generated low-level Go binding around an user-defined struct.
type PermitHelperPermit struct {
	Value            *big.Int
	Nonce            *big.Int
	Deadline         *big.Int
	IsDaiStylePermit bool
	V                uint8
	R                [32]byte
	S                [32]byte
}

const MainABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"SwapTargetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"SwapTargetRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyTokenAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"swapCallData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"fillQuoteEthToToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sellTokenAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"swapCallData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercentageBasisPoints\",\"type\":\"uint256\"}],\"name\":\"fillQuoteTokenToEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sellTokenAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"swapCallData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercentageBasisPoints\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isDaiStylePermit\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structPermitHelper.Permit\",\"name\":\"permitData\",\"type\":\"tuple\"}],\"name\":\"fillQuoteTokenToEthWithPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sellTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyTokenAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"swapCallData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"fillQuoteTokenToToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sellTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyTokenAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"swapCallData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isDaiStylePermit\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structPermitHelper.Permit\",\"name\":\"permitData\",\"type\":\"tuple\"}],\"name\":\"fillQuoteTokenToTokenWithPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapTargets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"updateSwapTargets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

type MainFilterer struct {
	abi      *abi.ABI
	contract *bind.BoundContract
}

type ABIEthWithdrawnEvent struct {
	Target      common.Address
	Amount      *big.Int
	Raw         ethrpc.Log
	BlockDetail ethrpc.Block
}

func (_Main *MainFilterer) ParseABIEthWithdrawnEvent(log types.Log, parsedLog ethrpc.Log, blockDetail ethrpc.Block) (*ABIEthWithdrawnEvent, error) {
	e := new(ABIEthWithdrawnEvent)
	if err := _Main.contract.UnpackLog(e, "EthWithdrawn", log); err != nil {
		return nil, err
	}
	e.Raw = parsedLog
	e.BlockDetail = blockDetail
	return e, nil
}

type ABISwapTargetAddedEvent struct {
	Target      common.Address
	Raw         ethrpc.Log
	BlockDetail ethrpc.Block
}

func (_Main *MainFilterer) ParseABISwapTargetAddedEvent(log types.Log, parsedLog ethrpc.Log, blockDetail ethrpc.Block) (*ABISwapTargetAddedEvent, error) {
	e := new(ABISwapTargetAddedEvent)
	if err := _Main.contract.UnpackLog(e, "SwapTargetAdded", log); err != nil {
		return nil, err
	}
	e.Raw = parsedLog
	e.BlockDetail = blockDetail
	return e, nil
}

type ABISwapTargetRemovedEvent struct {
	Target      common.Address
	Raw         ethrpc.Log
	BlockDetail ethrpc.Block
}

func (_Main *MainFilterer) ParseABISwapTargetRemovedEvent(log types.Log, parsedLog ethrpc.Log, blockDetail ethrpc.Block) (*ABISwapTargetRemovedEvent, error) {
	e := new(ABISwapTargetRemovedEvent)
	if err := _Main.contract.UnpackLog(e, "SwapTargetRemoved", log); err != nil {
		return nil, err
	}
	e.Raw = parsedLog
	e.BlockDetail = blockDetail
	return e, nil
}

type ABITokenWithdrawnEvent struct {
	Token       common.Address
	Target      common.Address
	Amount      *big.Int
	Raw         ethrpc.Log
	BlockDetail ethrpc.Block
}

func (_Main *MainFilterer) ParseABITokenWithdrawnEvent(log types.Log, parsedLog ethrpc.Log, blockDetail ethrpc.Block) (*ABITokenWithdrawnEvent, error) {
	e := new(ABITokenWithdrawnEvent)
	if err := _Main.contract.UnpackLog(e, "TokenWithdrawn", log); err != nil {
		return nil, err
	}
	e.Raw = parsedLog
	e.BlockDetail = blockDetail
	return e, nil
}

type ABIFillQuoteEthToTokenMethod struct {
	BuyTokenAddress common.Address
	Target          common.Address
	SwapCallData    []byte
	FeeAmount       *big.Int
	RawTransaction  sdkTransactionTypes.RawTransaction
}

func (_Main *MainFilterer) ParseABIFillQuoteEthToTokenMethod(parsedTx sdkTransactionTypes.RawTransaction) (*ABIFillQuoteEthToTokenMethod, error) {
	data, err := hexutil.Decode(parsedTx.AdditionalData.Data)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("invalid tx data")
	}
	m := new(ABIFillQuoteEthToTokenMethod)
	if err := _Main.UnpackMethodIntoInterface(m, "fillQuoteEthToToken", data); err != nil {
		return nil, err
	}
	m.RawTransaction = parsedTx
	return m, nil
}

type ABIFillQuoteTokenToEthMethod struct {
	SellTokenAddress         common.Address
	Target                   common.Address
	SwapCallData             []byte
	SellAmount               *big.Int
	FeePercentageBasisPoints *big.Int
	RawTransaction           sdkTransactionTypes.RawTransaction
}

func (_Main *MainFilterer) ParseABIFillQuoteTokenToEthMethod(parsedTx sdkTransactionTypes.RawTransaction) (*ABIFillQuoteTokenToEthMethod, error) {
	data, err := hexutil.Decode(parsedTx.AdditionalData.Data)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("invalid tx data")
	}
	m := new(ABIFillQuoteTokenToEthMethod)
	if err := _Main.UnpackMethodIntoInterface(m, "fillQuoteTokenToEth", data); err != nil {
		return nil, err
	}
	m.RawTransaction = parsedTx
	return m, nil
}

type ABIFillQuoteTokenToEthWithPermitMethod struct {
	SellTokenAddress         common.Address
	Target                   common.Address
	SwapCallData             []byte
	SellAmount               *big.Int
	FeePercentageBasisPoints *big.Int
	PermitData               PermitHelperPermit
	RawTransaction           sdkTransactionTypes.RawTransaction
}

func (_Main *MainFilterer) ParseABIFillQuoteTokenToEthWithPermitMethod(parsedTx sdkTransactionTypes.RawTransaction) (*ABIFillQuoteTokenToEthWithPermitMethod, error) {
	data, err := hexutil.Decode(parsedTx.AdditionalData.Data)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("invalid tx data")
	}
	m := new(ABIFillQuoteTokenToEthWithPermitMethod)
	if err := _Main.UnpackMethodIntoInterface(m, "fillQuoteTokenToEthWithPermit", data); err != nil {
		return nil, err
	}
	m.RawTransaction = parsedTx
	return m, nil
}

type ABIFillQuoteTokenToTokenMethod struct {
	SellTokenAddress common.Address
	BuyTokenAddress  common.Address
	Target           common.Address
	SwapCallData     []byte
	SellAmount       *big.Int
	FeeAmount        *big.Int
	RawTransaction   sdkTransactionTypes.RawTransaction
}

func (_Main *MainFilterer) ParseABIFillQuoteTokenToTokenMethod(parsedTx sdkTransactionTypes.RawTransaction) (*ABIFillQuoteTokenToTokenMethod, error) {
	data, err := hexutil.Decode(parsedTx.AdditionalData.Data)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("invalid tx data")
	}
	m := new(ABIFillQuoteTokenToTokenMethod)
	if err := _Main.UnpackMethodIntoInterface(m, "fillQuoteTokenToToken", data); err != nil {
		return nil, err
	}
	m.RawTransaction = parsedTx
	return m, nil
}

type ABIFillQuoteTokenToTokenWithPermitMethod struct {
	SellTokenAddress common.Address
	BuyTokenAddress  common.Address
	Target           common.Address
	SwapCallData     []byte
	SellAmount       *big.Int
	FeeAmount        *big.Int
	PermitData       PermitHelperPermit
	RawTransaction   sdkTransactionTypes.RawTransaction
}

func (_Main *MainFilterer) ParseABIFillQuoteTokenToTokenWithPermitMethod(parsedTx sdkTransactionTypes.RawTransaction) (*ABIFillQuoteTokenToTokenWithPermitMethod, error) {
	data, err := hexutil.Decode(parsedTx.AdditionalData.Data)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("invalid tx data")
	}
	m := new(ABIFillQuoteTokenToTokenWithPermitMethod)
	if err := _Main.UnpackMethodIntoInterface(m, "fillQuoteTokenToTokenWithPermit", data); err != nil {
		return nil, err
	}
	m.RawTransaction = parsedTx
	return m, nil
}

type ABITransferOwnershipMethod struct {
	NewOwner       common.Address
	RawTransaction sdkTransactionTypes.RawTransaction
}

func (_Main *MainFilterer) ParseABITransferOwnershipMethod(parsedTx sdkTransactionTypes.RawTransaction) (*ABITransferOwnershipMethod, error) {
	data, err := hexutil.Decode(parsedTx.AdditionalData.Data)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("invalid tx data")
	}
	m := new(ABITransferOwnershipMethod)
	if err := _Main.UnpackMethodIntoInterface(m, "transferOwnership", data); err != nil {
		return nil, err
	}
	m.RawTransaction = parsedTx
	return m, nil
}

type ABIUpdateSwapTargetsMethod struct {
	Target         common.Address
	Add            bool
	RawTransaction sdkTransactionTypes.RawTransaction
}

func (_Main *MainFilterer) ParseABIUpdateSwapTargetsMethod(parsedTx sdkTransactionTypes.RawTransaction) (*ABIUpdateSwapTargetsMethod, error) {
	data, err := hexutil.Decode(parsedTx.AdditionalData.Data)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("invalid tx data")
	}
	m := new(ABIUpdateSwapTargetsMethod)
	if err := _Main.UnpackMethodIntoInterface(m, "updateSwapTargets", data); err != nil {
		return nil, err
	}
	m.RawTransaction = parsedTx
	return m, nil
}

type ABIWithdrawEthMethod struct {
	To             common.Address
	Amount         *big.Int
	RawTransaction sdkTransactionTypes.RawTransaction
}

func (_Main *MainFilterer) ParseABIWithdrawEthMethod(parsedTx sdkTransactionTypes.RawTransaction) (*ABIWithdrawEthMethod, error) {
	data, err := hexutil.Decode(parsedTx.AdditionalData.Data)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("invalid tx data")
	}
	m := new(ABIWithdrawEthMethod)
	if err := _Main.UnpackMethodIntoInterface(m, "withdrawEth", data); err != nil {
		return nil, err
	}
	m.RawTransaction = parsedTx
	return m, nil
}

type ABIWithdrawTokenMethod struct {
	Token          common.Address
	To             common.Address
	Amount         *big.Int
	RawTransaction sdkTransactionTypes.RawTransaction
}

func (_Main *MainFilterer) ParseABIWithdrawTokenMethod(parsedTx sdkTransactionTypes.RawTransaction) (*ABIWithdrawTokenMethod, error) {
	data, err := hexutil.Decode(parsedTx.AdditionalData.Data)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("invalid tx data")
	}
	m := new(ABIWithdrawTokenMethod)
	if err := _Main.UnpackMethodIntoInterface(m, "withdrawToken", data); err != nil {
		return nil, err
	}
	m.RawTransaction = parsedTx
	return m, nil
}

func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, contractAbi, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract, abi: contractAbi}, nil
}

func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, *abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return nil, nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), &parsed, nil
}

func (_Main MainFilterer) UnpackMethodIntoInterface(v interface{}, name string, data []byte) error {
	args := _Main.abi.Methods[name].Inputs
	unpacked, err := args.Unpack(data[4:])
	if err != nil {
		return err
	}
	return args.Copy(v, unpacked)
}
