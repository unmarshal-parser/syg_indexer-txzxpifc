// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"github.com/shopspring/decimal"
	"gorm.io/datatypes"
	"math/big"
	"time"
)

var (
	_ = decimal.Decimal{}
	_ = big.NewInt
	_ = datatypes.JSON{}
	_ = time.Time{}
)

func GetEthWithdrawnEventHash() string {
	return "0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b"
}

type EthWithdrawnEvent struct {
	Target string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:0d0e7c3e-2437-4c47-b447-86136b50c109,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:0d0e7c3e-2437-4c47-b447-86136b50c109,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:0d0e7c3e-2437-4c47-b447-86136b50c109,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetSwapTargetAddedEventHash() string {
	return "0xb907822409611d127ab6a64611591b98e03a6a85ade4f258bae26b7c1efdfeaf"
}

type SwapTargetAddedEvent struct {
	Target string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:65e75c5c-e84d-400e-9506-39465a17f578,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:65e75c5c-e84d-400e-9506-39465a17f578,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:65e75c5c-e84d-400e-9506-39465a17f578,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetSwapTargetRemovedEventHash() string {
	return "0x393b8be3e26787f19285ecd039dfd80bc6507828750f4d50367e6efe2524695c"
}

type SwapTargetRemovedEvent struct {
	Target string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:4e0cbd9f-1f80-45e3-a806-311ae3b1ae93,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:4e0cbd9f-1f80-45e3-a806-311ae3b1ae93,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:4e0cbd9f-1f80-45e3-a806-311ae3b1ae93,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetTokenWithdrawnEventHash() string {
	return "0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620"
}

type TokenWithdrawnEvent struct {
	Token  string
	Target string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:2eed1e92-a48e-4eb5-b4f9-706c0b5e5b2a,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:2eed1e92-a48e-4eb5-b4f9-706c0b5e5b2a,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	Index           uint            `gorm:"uniqueIndex:2eed1e92-a48e-4eb5-b4f9-706c0b5e5b2a,unique"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteEthToTokenMethodHash() string {
	return "3c2b9a7d"
}

type FillQuoteEthToTokenMethod struct {
	BuyTokenAddress string
	Target          string
	SwapCallData    []byte
	FeeAmount       decimal.Decimal `gorm:"type:numeric"`

	TokenPriceFeeAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeeAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:ee0e5dc2-0b5d-4e90-ae3d-e4c8dde93e2b,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:ee0e5dc2-0b5d-4e90-ae3d-e4c8dde93e2b,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToTokenMethodHash() string {
	return "55e4b7be"
}

type FillQuoteTokenToTokenMethod struct {
	SellTokenAddress string
	BuyTokenAddress  string
	Target           string
	SwapCallData     []byte
	SellAmount       decimal.Decimal `gorm:"type:numeric"`
	FeeAmount        decimal.Decimal `gorm:"type:numeric"`

	TokenPriceSellAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`

	DecimalAdjustedFeeAmount float64 `gorm:"type:numeric"`
	TokenPriceFeeAmount      float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:2309bca9-b771-4165-88b8-26c0d5701062,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:2309bca9-b771-4165-88b8-26c0d5701062,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetUpdateSwapTargetsMethodHash() string {
	return "97bbda0e"
}

type UpdateSwapTargetsMethod struct {
	Target string
	Add    bool

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:9e7cf5e8-1e1c-45e1-842f-70d33db32e35,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:9e7cf5e8-1e1c-45e1-842f-70d33db32e35,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetWithdrawTokenMethodHash() string {
	return "01e33667"
}

type WithdrawTokenMethod struct {
	Token  string
	To     string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:3febc50f-d3dc-4928-ae39-9985f35a13a2,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:3febc50f-d3dc-4928-ae39-9985f35a13a2,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToEthMethodHash() string {
	return "999b6464"
}

type FillQuoteTokenToEthMethod struct {
	SellTokenAddress         string
	Target                   string
	SwapCallData             []byte
	SellAmount               decimal.Decimal `gorm:"type:numeric"`
	FeePercentageBasisPoints decimal.Decimal `gorm:"type:numeric"`

	TokenPriceSellAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`

	TokenPriceFeePercentageBasisPoints      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeePercentageBasisPoints float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:256bb934-bcda-47be-9f6a-15fd73e8c378,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:256bb934-bcda-47be-9f6a-15fd73e8c378,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToEthWithPermitMethodHash() string {
	return "b3093838"
}

type FillQuoteTokenToEthWithPermitMethod struct {
	SellTokenAddress         string
	Target                   string
	SwapCallData             []byte
	SellAmount               decimal.Decimal `gorm:"type:numeric"`
	FeePercentageBasisPoints decimal.Decimal `gorm:"type:numeric"`
	PermitData               datatypes.JSON

	TokenPriceSellAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`

	TokenPriceFeePercentageBasisPoints      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeePercentageBasisPoints float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:d3ab252f-3c33-46d4-be83-c9d16751de84,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:d3ab252f-3c33-46d4-be83-c9d16751de84,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetFillQuoteTokenToTokenWithPermitMethodHash() string {
	return "b0480bbd"
}

type FillQuoteTokenToTokenWithPermitMethod struct {
	SellTokenAddress string
	BuyTokenAddress  string
	Target           string
	SwapCallData     []byte
	SellAmount       decimal.Decimal `gorm:"type:numeric"`
	FeeAmount        decimal.Decimal `gorm:"type:numeric"`
	PermitData       datatypes.JSON

	TokenPriceSellAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedSellAmount float64 `gorm:"type:numeric"`

	TokenPriceFeeAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedFeeAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:52c8f70a-ba5a-4963-830d-e3e86dc25279,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:52c8f70a-ba5a-4963-830d-e3e86dc25279,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetTransferOwnershipMethodHash() string {
	return "f2fde38b"
}

type TransferOwnershipMethod struct {
	NewOwner string

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:0e8d0902-06f7-455a-b947-dbc7a640fc0f,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:0e8d0902-06f7-455a-b947-dbc7a640fc0f,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

func GetWithdrawEthMethodHash() string {
	return "1b9a91a4"
}

type WithdrawEthMethod struct {
	To     string
	Amount decimal.Decimal `gorm:"type:numeric"`

	TokenPriceAmount      float64 `gorm:"type:numeric"`
	DecimalAdjustedAmount float64 `gorm:"type:numeric"`

	ID              uint   `gorm:"primaryKey"`
	BlockNumber     uint64 `gorm:"uniqueIndex:310a32c6-005e-4f10-a917-d878a68d8836,unique;index"`
	TxHash          string
	TxIndex         uint `gorm:"uniqueIndex:310a32c6-005e-4f10-a917-d878a68d8836,unique"`
	BlockHash       string
	Gas             decimal.Decimal `gorm:"type:numeric"`
	GasPrice        decimal.Decimal `gorm:"type:numeric"`
	TxFrom          string          `gorm:"index"`
	TxTo            string          `gorm:"index"`
	TxValue         decimal.Decimal `gorm:"type:numeric"`
	BlockTime       time.Time       `gorm:"index"`
	ContractAddress string
	ChainID         string
}

type LastSyncedBlock struct {
	Contract    string `gorm:"primaryKey"`
	ChainID     string `gorm:"primaryKey"`
	SyncType    string `gorm:"primaryKey"`
	BlockNumber uint64
}

// Plugin Models

type TokenDetails struct {
	ID      int
	Address string `gorm:"uniqueIndex:address_and_chain"`
	Symbol  string
	ChainID string `gorm:"uniqueIndex:address_and_chain"`
	Decimal int
}

// Config

type PostgresConfig struct {
	ConnectionString string `mapstructure:"connection_string"`
	TablePrefix      string `mapstructure:"table_prefix"`
	CreateBatchSize  int    `mapstructure:"create_batch_size"`
}

type IndexerConfig struct {
	EthEndpoint       string `mapstructure:"eth_endpoint"`
	ContractAddress   string `mapstructure:"contract_address"`
	StartBlock        int    `mapstructure:"start_block"`
	ApiKey            string `mapstructure:"api_key"`
	PostgresConfig    `mapstructure:"postgres_config"`
	LagToHighestBlock int `mapstructure:"lag_to_highest_block"`
	StepSize          int `mapstructure:"step_size"`
	ParallelCalls     int `mapstructure:"parallel_calls_for_logs"`
}

func (i *IndexerConfig) AssignDefaults() {
	if i.PostgresConfig.CreateBatchSize == 0 {
		i.PostgresConfig.CreateBatchSize = 100
	}
	if i.StepSize == 0 {
		i.StepSize = 50
	}
	if i.LagToHighestBlock == 0 {
		i.LagToHighestBlock = 10
	}
}
