// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"encoding/json"
	"github.com/onrik/ethrpc"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"gorm.io/datatypes"
	"math/big"
	"strconv"
	"strings"
	"time"
)

var (
	_ = decimal.Decimal{}
	_ = big.NewInt
	_ = ethrpc.Transaction{}
	_ = time.Time{}
	_ = strings.Builder{}
)

func getJSONFromInterface(data interface{}) datatypes.JSON {
	var (
		err  error
		temp datatypes.JSON
	)
	temp, err = json.Marshal(data)
	if err != nil {
		log.Error("Error Marshalling Data: " + err.Error())
	}
	return temp
}

func convertToEthWithdrawnEvent(abiEvent *ABIEthWithdrawnEvent, transaction *ethrpc.Transaction, chainID string) EthWithdrawnEvent {
	return EthWithdrawnEvent{
		Target: strings.ToLower(abiEvent.Target.String()),
		Amount: decimal.NewFromBigInt(abiEvent.Amount, 0),

		Gas:             decimal.NewFromInt(int64(transaction.Gas)),
		GasPrice:        decimal.NewFromBigInt(&transaction.GasPrice, 0),
		TxFrom:          strings.ToLower(transaction.From),
		TxTo:            strings.ToLower(transaction.To),
		TxValue:         decimal.NewFromBigInt(&transaction.Value, 0),
		BlockNumber:     uint64(abiEvent.Raw.BlockNumber),
		TxHash:          strings.ToLower(abiEvent.Raw.TransactionHash),
		TxIndex:         uint(abiEvent.Raw.TransactionIndex),
		BlockHash:       strings.ToLower(abiEvent.Raw.BlockHash),
		Index:           uint(abiEvent.Raw.LogIndex),
		BlockTime:       time.Unix(int64(abiEvent.BlockDetail.Timestamp), 0),
		ContractAddress: strings.ToLower(abiEvent.Raw.Address),
		ChainID:         chainID,
	}
}

func convertToSwapTargetAddedEvent(abiEvent *ABISwapTargetAddedEvent, transaction *ethrpc.Transaction, chainID string) SwapTargetAddedEvent {
	return SwapTargetAddedEvent{
		Target: strings.ToLower(abiEvent.Target.String()),

		Gas:             decimal.NewFromInt(int64(transaction.Gas)),
		GasPrice:        decimal.NewFromBigInt(&transaction.GasPrice, 0),
		TxFrom:          strings.ToLower(transaction.From),
		TxTo:            strings.ToLower(transaction.To),
		TxValue:         decimal.NewFromBigInt(&transaction.Value, 0),
		BlockNumber:     uint64(abiEvent.Raw.BlockNumber),
		TxHash:          strings.ToLower(abiEvent.Raw.TransactionHash),
		TxIndex:         uint(abiEvent.Raw.TransactionIndex),
		BlockHash:       strings.ToLower(abiEvent.Raw.BlockHash),
		Index:           uint(abiEvent.Raw.LogIndex),
		BlockTime:       time.Unix(int64(abiEvent.BlockDetail.Timestamp), 0),
		ContractAddress: strings.ToLower(abiEvent.Raw.Address),
		ChainID:         chainID,
	}
}

func convertToSwapTargetRemovedEvent(abiEvent *ABISwapTargetRemovedEvent, transaction *ethrpc.Transaction, chainID string) SwapTargetRemovedEvent {
	return SwapTargetRemovedEvent{
		Target: strings.ToLower(abiEvent.Target.String()),

		Gas:             decimal.NewFromInt(int64(transaction.Gas)),
		GasPrice:        decimal.NewFromBigInt(&transaction.GasPrice, 0),
		TxFrom:          strings.ToLower(transaction.From),
		TxTo:            strings.ToLower(transaction.To),
		TxValue:         decimal.NewFromBigInt(&transaction.Value, 0),
		BlockNumber:     uint64(abiEvent.Raw.BlockNumber),
		TxHash:          strings.ToLower(abiEvent.Raw.TransactionHash),
		TxIndex:         uint(abiEvent.Raw.TransactionIndex),
		BlockHash:       strings.ToLower(abiEvent.Raw.BlockHash),
		Index:           uint(abiEvent.Raw.LogIndex),
		BlockTime:       time.Unix(int64(abiEvent.BlockDetail.Timestamp), 0),
		ContractAddress: strings.ToLower(abiEvent.Raw.Address),
		ChainID:         chainID,
	}
}

func convertToTokenWithdrawnEvent(abiEvent *ABITokenWithdrawnEvent, transaction *ethrpc.Transaction, chainID string) TokenWithdrawnEvent {
	return TokenWithdrawnEvent{
		Token:  strings.ToLower(abiEvent.Token.String()),
		Target: strings.ToLower(abiEvent.Target.String()),
		Amount: decimal.NewFromBigInt(abiEvent.Amount, 0),

		Gas:             decimal.NewFromInt(int64(transaction.Gas)),
		GasPrice:        decimal.NewFromBigInt(&transaction.GasPrice, 0),
		TxFrom:          strings.ToLower(transaction.From),
		TxTo:            strings.ToLower(transaction.To),
		TxValue:         decimal.NewFromBigInt(&transaction.Value, 0),
		BlockNumber:     uint64(abiEvent.Raw.BlockNumber),
		TxHash:          strings.ToLower(abiEvent.Raw.TransactionHash),
		TxIndex:         uint(abiEvent.Raw.TransactionIndex),
		BlockHash:       strings.ToLower(abiEvent.Raw.BlockHash),
		Index:           uint(abiEvent.Raw.LogIndex),
		BlockTime:       time.Unix(int64(abiEvent.BlockDetail.Timestamp), 0),
		ContractAddress: strings.ToLower(abiEvent.Raw.Address),
		ChainID:         chainID,
	}
}

func convertToFillQuoteEthToTokenMethod(abiMethod *ABIFillQuoteEthToTokenMethod, chainID string) FillQuoteEthToTokenMethod {
	return FillQuoteEthToTokenMethod{
		BuyTokenAddress: strings.ToLower(abiMethod.BuyTokenAddress.String()),
		Target:          strings.ToLower(abiMethod.Target.String()),
		SwapCallData:    abiMethod.SwapCallData,
		FeeAmount:       decimal.NewFromBigInt(abiMethod.FeeAmount, 0),

		Gas:             decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasUsed, 0),
		GasPrice:        decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasPrice, 0),
		TxFrom:          strings.ToLower(abiMethod.RawTransaction.From),
		TxTo:            strings.ToLower(abiMethod.RawTransaction.To),
		TxValue:         decimal.NewFromBigInt(abiMethod.RawTransaction.Value, 0),
		BlockNumber:     getUint64FromString(abiMethod.RawTransaction.BlockNumber),
		TxHash:          strings.ToLower(abiMethod.RawTransaction.TxHash),
		TxIndex:         abiMethod.RawTransaction.TxIndex,
		BlockHash:       strings.ToLower(abiMethod.RawTransaction.BlockHash),
		BlockTime:       time.Unix(abiMethod.RawTransaction.BlockTime.Int64(), 0),
		ContractAddress: strings.ToLower(abiMethod.RawTransaction.To),
		ChainID:         chainID,
	}
}

func convertToFillQuoteTokenToTokenMethod(abiMethod *ABIFillQuoteTokenToTokenMethod, chainID string) FillQuoteTokenToTokenMethod {
	return FillQuoteTokenToTokenMethod{
		SellTokenAddress: strings.ToLower(abiMethod.SellTokenAddress.String()),
		BuyTokenAddress:  strings.ToLower(abiMethod.BuyTokenAddress.String()),
		Target:           strings.ToLower(abiMethod.Target.String()),
		SwapCallData:     abiMethod.SwapCallData,
		SellAmount:       decimal.NewFromBigInt(abiMethod.SellAmount, 0),
		FeeAmount:        decimal.NewFromBigInt(abiMethod.FeeAmount, 0),

		Gas:             decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasUsed, 0),
		GasPrice:        decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasPrice, 0),
		TxFrom:          strings.ToLower(abiMethod.RawTransaction.From),
		TxTo:            strings.ToLower(abiMethod.RawTransaction.To),
		TxValue:         decimal.NewFromBigInt(abiMethod.RawTransaction.Value, 0),
		BlockNumber:     getUint64FromString(abiMethod.RawTransaction.BlockNumber),
		TxHash:          strings.ToLower(abiMethod.RawTransaction.TxHash),
		TxIndex:         abiMethod.RawTransaction.TxIndex,
		BlockHash:       strings.ToLower(abiMethod.RawTransaction.BlockHash),
		BlockTime:       time.Unix(abiMethod.RawTransaction.BlockTime.Int64(), 0),
		ContractAddress: strings.ToLower(abiMethod.RawTransaction.To),
		ChainID:         chainID,
	}
}

func convertToUpdateSwapTargetsMethod(abiMethod *ABIUpdateSwapTargetsMethod, chainID string) UpdateSwapTargetsMethod {
	return UpdateSwapTargetsMethod{
		Target: strings.ToLower(abiMethod.Target.String()),
		Add:    abiMethod.Add,

		Gas:             decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasUsed, 0),
		GasPrice:        decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasPrice, 0),
		TxFrom:          strings.ToLower(abiMethod.RawTransaction.From),
		TxTo:            strings.ToLower(abiMethod.RawTransaction.To),
		TxValue:         decimal.NewFromBigInt(abiMethod.RawTransaction.Value, 0),
		BlockNumber:     getUint64FromString(abiMethod.RawTransaction.BlockNumber),
		TxHash:          strings.ToLower(abiMethod.RawTransaction.TxHash),
		TxIndex:         abiMethod.RawTransaction.TxIndex,
		BlockHash:       strings.ToLower(abiMethod.RawTransaction.BlockHash),
		BlockTime:       time.Unix(abiMethod.RawTransaction.BlockTime.Int64(), 0),
		ContractAddress: strings.ToLower(abiMethod.RawTransaction.To),
		ChainID:         chainID,
	}
}

func convertToWithdrawTokenMethod(abiMethod *ABIWithdrawTokenMethod, chainID string) WithdrawTokenMethod {
	return WithdrawTokenMethod{
		Token:  strings.ToLower(abiMethod.Token.String()),
		To:     strings.ToLower(abiMethod.To.String()),
		Amount: decimal.NewFromBigInt(abiMethod.Amount, 0),

		Gas:             decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasUsed, 0),
		GasPrice:        decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasPrice, 0),
		TxFrom:          strings.ToLower(abiMethod.RawTransaction.From),
		TxTo:            strings.ToLower(abiMethod.RawTransaction.To),
		TxValue:         decimal.NewFromBigInt(abiMethod.RawTransaction.Value, 0),
		BlockNumber:     getUint64FromString(abiMethod.RawTransaction.BlockNumber),
		TxHash:          strings.ToLower(abiMethod.RawTransaction.TxHash),
		TxIndex:         abiMethod.RawTransaction.TxIndex,
		BlockHash:       strings.ToLower(abiMethod.RawTransaction.BlockHash),
		BlockTime:       time.Unix(abiMethod.RawTransaction.BlockTime.Int64(), 0),
		ContractAddress: strings.ToLower(abiMethod.RawTransaction.To),
		ChainID:         chainID,
	}
}

func convertToFillQuoteTokenToEthMethod(abiMethod *ABIFillQuoteTokenToEthMethod, chainID string) FillQuoteTokenToEthMethod {
	return FillQuoteTokenToEthMethod{
		SellTokenAddress:         strings.ToLower(abiMethod.SellTokenAddress.String()),
		Target:                   strings.ToLower(abiMethod.Target.String()),
		SwapCallData:             abiMethod.SwapCallData,
		SellAmount:               decimal.NewFromBigInt(abiMethod.SellAmount, 0),
		FeePercentageBasisPoints: decimal.NewFromBigInt(abiMethod.FeePercentageBasisPoints, 0),

		Gas:             decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasUsed, 0),
		GasPrice:        decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasPrice, 0),
		TxFrom:          strings.ToLower(abiMethod.RawTransaction.From),
		TxTo:            strings.ToLower(abiMethod.RawTransaction.To),
		TxValue:         decimal.NewFromBigInt(abiMethod.RawTransaction.Value, 0),
		BlockNumber:     getUint64FromString(abiMethod.RawTransaction.BlockNumber),
		TxHash:          strings.ToLower(abiMethod.RawTransaction.TxHash),
		TxIndex:         abiMethod.RawTransaction.TxIndex,
		BlockHash:       strings.ToLower(abiMethod.RawTransaction.BlockHash),
		BlockTime:       time.Unix(abiMethod.RawTransaction.BlockTime.Int64(), 0),
		ContractAddress: strings.ToLower(abiMethod.RawTransaction.To),
		ChainID:         chainID,
	}
}

func convertToFillQuoteTokenToEthWithPermitMethod(abiMethod *ABIFillQuoteTokenToEthWithPermitMethod, chainID string) FillQuoteTokenToEthWithPermitMethod {
	return FillQuoteTokenToEthWithPermitMethod{
		SellTokenAddress:         strings.ToLower(abiMethod.SellTokenAddress.String()),
		Target:                   strings.ToLower(abiMethod.Target.String()),
		SwapCallData:             abiMethod.SwapCallData,
		SellAmount:               decimal.NewFromBigInt(abiMethod.SellAmount, 0),
		FeePercentageBasisPoints: decimal.NewFromBigInt(abiMethod.FeePercentageBasisPoints, 0),
		PermitData:               getJSONFromInterface(abiMethod.PermitData),

		Gas:             decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasUsed, 0),
		GasPrice:        decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasPrice, 0),
		TxFrom:          strings.ToLower(abiMethod.RawTransaction.From),
		TxTo:            strings.ToLower(abiMethod.RawTransaction.To),
		TxValue:         decimal.NewFromBigInt(abiMethod.RawTransaction.Value, 0),
		BlockNumber:     getUint64FromString(abiMethod.RawTransaction.BlockNumber),
		TxHash:          strings.ToLower(abiMethod.RawTransaction.TxHash),
		TxIndex:         abiMethod.RawTransaction.TxIndex,
		BlockHash:       strings.ToLower(abiMethod.RawTransaction.BlockHash),
		BlockTime:       time.Unix(abiMethod.RawTransaction.BlockTime.Int64(), 0),
		ContractAddress: strings.ToLower(abiMethod.RawTransaction.To),
		ChainID:         chainID,
	}
}

func convertToFillQuoteTokenToTokenWithPermitMethod(abiMethod *ABIFillQuoteTokenToTokenWithPermitMethod, chainID string) FillQuoteTokenToTokenWithPermitMethod {
	return FillQuoteTokenToTokenWithPermitMethod{
		SellTokenAddress: strings.ToLower(abiMethod.SellTokenAddress.String()),
		BuyTokenAddress:  strings.ToLower(abiMethod.BuyTokenAddress.String()),
		Target:           strings.ToLower(abiMethod.Target.String()),
		SwapCallData:     abiMethod.SwapCallData,
		SellAmount:       decimal.NewFromBigInt(abiMethod.SellAmount, 0),
		FeeAmount:        decimal.NewFromBigInt(abiMethod.FeeAmount, 0),
		PermitData:       getJSONFromInterface(abiMethod.PermitData),

		Gas:             decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasUsed, 0),
		GasPrice:        decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasPrice, 0),
		TxFrom:          strings.ToLower(abiMethod.RawTransaction.From),
		TxTo:            strings.ToLower(abiMethod.RawTransaction.To),
		TxValue:         decimal.NewFromBigInt(abiMethod.RawTransaction.Value, 0),
		BlockNumber:     getUint64FromString(abiMethod.RawTransaction.BlockNumber),
		TxHash:          strings.ToLower(abiMethod.RawTransaction.TxHash),
		TxIndex:         abiMethod.RawTransaction.TxIndex,
		BlockHash:       strings.ToLower(abiMethod.RawTransaction.BlockHash),
		BlockTime:       time.Unix(abiMethod.RawTransaction.BlockTime.Int64(), 0),
		ContractAddress: strings.ToLower(abiMethod.RawTransaction.To),
		ChainID:         chainID,
	}
}

func convertToTransferOwnershipMethod(abiMethod *ABITransferOwnershipMethod, chainID string) TransferOwnershipMethod {
	return TransferOwnershipMethod{
		NewOwner: strings.ToLower(abiMethod.NewOwner.String()),

		Gas:             decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasUsed, 0),
		GasPrice:        decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasPrice, 0),
		TxFrom:          strings.ToLower(abiMethod.RawTransaction.From),
		TxTo:            strings.ToLower(abiMethod.RawTransaction.To),
		TxValue:         decimal.NewFromBigInt(abiMethod.RawTransaction.Value, 0),
		BlockNumber:     getUint64FromString(abiMethod.RawTransaction.BlockNumber),
		TxHash:          strings.ToLower(abiMethod.RawTransaction.TxHash),
		TxIndex:         abiMethod.RawTransaction.TxIndex,
		BlockHash:       strings.ToLower(abiMethod.RawTransaction.BlockHash),
		BlockTime:       time.Unix(abiMethod.RawTransaction.BlockTime.Int64(), 0),
		ContractAddress: strings.ToLower(abiMethod.RawTransaction.To),
		ChainID:         chainID,
	}
}

func convertToWithdrawEthMethod(abiMethod *ABIWithdrawEthMethod, chainID string) WithdrawEthMethod {
	return WithdrawEthMethod{
		To:     strings.ToLower(abiMethod.To.String()),
		Amount: decimal.NewFromBigInt(abiMethod.Amount, 0),

		Gas:             decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasUsed, 0),
		GasPrice:        decimal.NewFromBigInt(abiMethod.RawTransaction.AdditionalData.GasPrice, 0),
		TxFrom:          strings.ToLower(abiMethod.RawTransaction.From),
		TxTo:            strings.ToLower(abiMethod.RawTransaction.To),
		TxValue:         decimal.NewFromBigInt(abiMethod.RawTransaction.Value, 0),
		BlockNumber:     getUint64FromString(abiMethod.RawTransaction.BlockNumber),
		TxHash:          strings.ToLower(abiMethod.RawTransaction.TxHash),
		TxIndex:         abiMethod.RawTransaction.TxIndex,
		BlockHash:       strings.ToLower(abiMethod.RawTransaction.BlockHash),
		BlockTime:       time.Unix(abiMethod.RawTransaction.BlockTime.Int64(), 0),
		ContractAddress: strings.ToLower(abiMethod.RawTransaction.To),
		ChainID:         chainID,
	}
}

func getUint64FromString(numberString string) uint64 {
	number, err := strconv.ParseUint(numberString, 10, 64)
	if err != nil {
		return 0
	}
	return number
}
