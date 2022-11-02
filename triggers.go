// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	unmarshal "github.com/eucrypt/unmarshal-go-sdk/pkg"
	conf "github.com/eucrypt/unmarshal-go-sdk/pkg/config"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	sdkTokenDetails "github.com/eucrypt/unmarshal-go-sdk/pkg/token_details"
	tokenPriceTypes "github.com/eucrypt/unmarshal-go-sdk/pkg/token_price/types"
	"github.com/onrik/ethrpc"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"
)

var (
	_   = decimal.Decimal{}
	_   = big.NewInt
	_   = ethrpc.Transaction{}
	sdk = unmarshal.Unmarshal{}
	_   = sdkTokenDetails.TokenDetailsOptions{}
	_   = gorm.Model{}
	_   = time.Time{}
	_   = strconv.NumError{}
	_   = math.NaN()
	_   = log.Error
	_   = tokenPriceTypes.TokenPrice{}
)

func InitPluginModels(db *gorm.DB) error {
	err := db.AutoMigrate(
		&TokenDetails{},

		&TokenDetails{},
	)
	if err != nil {
		return err
	}
	return nil
}

func initUnmarshalSDK(cfg IndexerConfig) {
	key := strings.TrimSpace(cfg.ApiKey)
	if key == "" {
		panic("missing api key")
	}
	sdk = unmarshal.NewWithConfig(conf.Config{
		AuthKey:     key,
		Environment: constants.Prod,
	})
}

func (entity *EthWithdrawnEvent) BeforeCreateHook(tx *gorm.DB) error {

	tokenDetailsAmount, err := getTokenDetails("0x00000000009726632680FB29d3F7A9734E3010E2", tx, entity.ChainID)
	if err == nil {
		amount := formatAmount(entity.Amount, tokenDetailsAmount.Decimal)
		entity.DecimalAdjustedAmount = amount
	}
	tokenPriceAmount := getPriceAtInstant("0x00000000009726632680FB29d3F7A9734E3010E2", tokenDetailsAmount.Symbol, entity.ChainID, entity.BlockTime)
	entity.TokenPriceAmount = mustParseFloat(tokenPriceAmount)

	return nil
}

func (entity *EthWithdrawnEvent) AfterCreateHook(tx *gorm.DB) error {
	return nil
}

func (entity *SwapTargetAddedEvent) BeforeCreateHook(tx *gorm.DB) error {

	return nil
}

func (entity *SwapTargetAddedEvent) AfterCreateHook(tx *gorm.DB) error {
	return nil
}

func (entity *SwapTargetRemovedEvent) BeforeCreateHook(tx *gorm.DB) error {

	return nil
}

func (entity *SwapTargetRemovedEvent) AfterCreateHook(tx *gorm.DB) error {
	return nil
}

func (entity *TokenWithdrawnEvent) BeforeCreateHook(tx *gorm.DB) error {

	tokenDetailsAmount, err := getTokenDetails(entity.Token, tx, entity.ChainID)
	if err == nil {
		amount := formatAmount(entity.Amount, tokenDetailsAmount.Decimal)
		entity.DecimalAdjustedAmount = amount
	}
	tokenPriceAmount := getPriceAtInstant(entity.Token, tokenDetailsAmount.Symbol, entity.ChainID, entity.BlockTime)
	entity.TokenPriceAmount = mustParseFloat(tokenPriceAmount)

	return nil
}

func (entity *TokenWithdrawnEvent) AfterCreateHook(tx *gorm.DB) error {
	return nil
}

func (entity *FillQuoteEthToTokenMethod) BeforeCreateHook(tx *gorm.DB) error {

	tokenDetailsFeeAmount, err := getTokenDetails("0x00000000009726632680FB29d3F7A9734E3010E2", tx, entity.ChainID)
	if err == nil {
		amount := formatAmount(entity.FeeAmount, tokenDetailsFeeAmount.Decimal)
		entity.DecimalAdjustedFeeAmount = amount
	}
	tokenPriceFeeAmount := getPriceAtInstant("0x00000000009726632680FB29d3F7A9734E3010E2", tokenDetailsFeeAmount.Symbol, entity.ChainID, entity.BlockTime)
	entity.TokenPriceFeeAmount = mustParseFloat(tokenPriceFeeAmount)

	return nil
}

func (entity *FillQuoteEthToTokenMethod) AfterCreateHook(tx *gorm.DB) error {
	return nil
}

func (entity *FillQuoteTokenToTokenMethod) BeforeCreateHook(tx *gorm.DB) error {

	tokenDetailsSellAmount, err := getTokenDetails(entity.SellTokenAddress, tx, entity.ChainID)
	if err == nil {
		amount := formatAmount(entity.SellAmount, tokenDetailsSellAmount.Decimal)
		entity.DecimalAdjustedSellAmount = amount
	}
	tokenPriceSellAmount := getPriceAtInstant(entity.SellTokenAddress, tokenDetailsSellAmount.Symbol, entity.ChainID, entity.BlockTime)
	entity.TokenPriceSellAmount = mustParseFloat(tokenPriceSellAmount)

	tokenDetailsFeeAmount, err := getTokenDetails(entity.SellTokenAddress, tx, entity.ChainID)
	if err == nil {
		amount := formatAmount(entity.FeeAmount, tokenDetailsFeeAmount.Decimal)
		entity.DecimalAdjustedFeeAmount = amount
	}
	tokenPriceFeeAmount := getPriceAtInstant(entity.SellTokenAddress, tokenDetailsFeeAmount.Symbol, entity.ChainID, entity.BlockTime)
	entity.TokenPriceFeeAmount = mustParseFloat(tokenPriceFeeAmount)

	return nil
}

func (entity *FillQuoteTokenToTokenMethod) AfterCreateHook(tx *gorm.DB) error {
	return nil
}

func (entity *UpdateSwapTargetsMethod) BeforeCreateHook(tx *gorm.DB) error {

	return nil
}

func (entity *UpdateSwapTargetsMethod) AfterCreateHook(tx *gorm.DB) error {
	return nil
}

func (entity *WithdrawTokenMethod) BeforeCreateHook(tx *gorm.DB) error {

	tokenDetailsAmount, err := getTokenDetails(entity.Token, tx, entity.ChainID)
	if err == nil {
		amount := formatAmount(entity.Amount, tokenDetailsAmount.Decimal)
		entity.DecimalAdjustedAmount = amount
	}
	tokenPriceAmount := getPriceAtInstant(entity.Token, tokenDetailsAmount.Symbol, entity.ChainID, entity.BlockTime)
	entity.TokenPriceAmount = mustParseFloat(tokenPriceAmount)

	return nil
}

func (entity *WithdrawTokenMethod) AfterCreateHook(tx *gorm.DB) error {
	return nil
}

func (entity *FillQuoteTokenToEthMethod) BeforeCreateHook(tx *gorm.DB) error {

	tokenDetailsSellAmount, err := getTokenDetails(entity.SellTokenAddress, tx, entity.ChainID)
	if err == nil {
		amount := formatAmount(entity.SellAmount, tokenDetailsSellAmount.Decimal)
		entity.DecimalAdjustedSellAmount = amount
	}
	tokenPriceSellAmount := getPriceAtInstant(entity.SellTokenAddress, tokenDetailsSellAmount.Symbol, entity.ChainID, entity.BlockTime)
	entity.TokenPriceSellAmount = mustParseFloat(tokenPriceSellAmount)

	tokenDetailsFeePercentageBasisPoints, err := getTokenDetails("0x00000000009726632680FB29d3F7A9734E3010E2", tx, entity.ChainID)
	if err == nil {
		amount := formatAmount(entity.FeePercentageBasisPoints, tokenDetailsFeePercentageBasisPoints.Decimal)
		entity.DecimalAdjustedFeePercentageBasisPoints = amount
	}
	tokenPriceFeePercentageBasisPoints := getPriceAtInstant("0x00000000009726632680FB29d3F7A9734E3010E2", tokenDetailsFeePercentageBasisPoints.Symbol, entity.ChainID, entity.BlockTime)
	entity.TokenPriceFeePercentageBasisPoints = mustParseFloat(tokenPriceFeePercentageBasisPoints)

	return nil
}

func (entity *FillQuoteTokenToEthMethod) AfterCreateHook(tx *gorm.DB) error {
	return nil
}

func (entity *FillQuoteTokenToEthWithPermitMethod) BeforeCreateHook(tx *gorm.DB) error {

	tokenDetailsSellAmount, err := getTokenDetails(entity.SellTokenAddress, tx, entity.ChainID)
	if err == nil {
		amount := formatAmount(entity.SellAmount, tokenDetailsSellAmount.Decimal)
		entity.DecimalAdjustedSellAmount = amount
	}
	tokenPriceSellAmount := getPriceAtInstant(entity.SellTokenAddress, tokenDetailsSellAmount.Symbol, entity.ChainID, entity.BlockTime)
	entity.TokenPriceSellAmount = mustParseFloat(tokenPriceSellAmount)

	tokenDetailsFeePercentageBasisPoints, err := getTokenDetails("0x00000000009726632680FB29d3F7A9734E3010E2", tx, entity.ChainID)
	if err == nil {
		amount := formatAmount(entity.FeePercentageBasisPoints, tokenDetailsFeePercentageBasisPoints.Decimal)
		entity.DecimalAdjustedFeePercentageBasisPoints = amount
	}
	tokenPriceFeePercentageBasisPoints := getPriceAtInstant("0x00000000009726632680FB29d3F7A9734E3010E2", tokenDetailsFeePercentageBasisPoints.Symbol, entity.ChainID, entity.BlockTime)
	entity.TokenPriceFeePercentageBasisPoints = mustParseFloat(tokenPriceFeePercentageBasisPoints)

	return nil
}

func (entity *FillQuoteTokenToEthWithPermitMethod) AfterCreateHook(tx *gorm.DB) error {
	return nil
}

func (entity *FillQuoteTokenToTokenWithPermitMethod) BeforeCreateHook(tx *gorm.DB) error {

	tokenDetailsSellAmount, err := getTokenDetails(entity.SellTokenAddress, tx, entity.ChainID)
	if err == nil {
		amount := formatAmount(entity.SellAmount, tokenDetailsSellAmount.Decimal)
		entity.DecimalAdjustedSellAmount = amount
	}
	tokenPriceSellAmount := getPriceAtInstant(entity.SellTokenAddress, tokenDetailsSellAmount.Symbol, entity.ChainID, entity.BlockTime)
	entity.TokenPriceSellAmount = mustParseFloat(tokenPriceSellAmount)

	tokenDetailsFeeAmount, err := getTokenDetails(entity.SellTokenAddress, tx, entity.ChainID)
	if err == nil {
		amount := formatAmount(entity.FeeAmount, tokenDetailsFeeAmount.Decimal)
		entity.DecimalAdjustedFeeAmount = amount
	}
	tokenPriceFeeAmount := getPriceAtInstant(entity.SellTokenAddress, tokenDetailsFeeAmount.Symbol, entity.ChainID, entity.BlockTime)
	entity.TokenPriceFeeAmount = mustParseFloat(tokenPriceFeeAmount)

	return nil
}

func (entity *FillQuoteTokenToTokenWithPermitMethod) AfterCreateHook(tx *gorm.DB) error {
	return nil
}

func (entity *TransferOwnershipMethod) BeforeCreateHook(tx *gorm.DB) error {

	return nil
}

func (entity *TransferOwnershipMethod) AfterCreateHook(tx *gorm.DB) error {
	return nil
}

func (entity *WithdrawEthMethod) BeforeCreateHook(tx *gorm.DB) error {

	tokenDetailsAmount, err := getTokenDetails("0x00000000009726632680FB29d3F7A9734E3010E2", tx, entity.ChainID)
	if err == nil {
		amount := formatAmount(entity.Amount, tokenDetailsAmount.Decimal)
		entity.DecimalAdjustedAmount = amount
	}
	tokenPriceAmount := getPriceAtInstant("0x00000000009726632680FB29d3F7A9734E3010E2", tokenDetailsAmount.Symbol, entity.ChainID, entity.BlockTime)
	entity.TokenPriceAmount = mustParseFloat(tokenPriceAmount)

	return nil
}

func (entity *WithdrawEthMethod) AfterCreateHook(tx *gorm.DB) error {
	return nil
}

func formatAmount(amountWithoutDecimal decimal.Decimal, decimals int) float64 {
	value := mustParseFloat(amountWithoutDecimal.String())
	divider := math.Pow(10, float64(decimals))
	value = value / divider
	return value
}

func mustParseFloat(floatVal string) float64 {
	value, err := strconv.ParseFloat(floatVal, 64)
	if err != nil {
		return 0.0
	}
	return value
}

func getPriceAtInstant(contractAddress, tokenSymbol string, indexerChainId string, timestamp time.Time) string {
	chainName, _ := GetChainFromChainID(indexerChainId)
	if IsPriceSupportedForChain(chainName) {
		log.WithFields(log.Fields{
			"chain":   chainName,
			"address": contractAddress,
		}).Debug("Attempting fetch price by contract address")
		price, err := sdk.PriceStore.GetTokenPrice(chainName, contractAddress, &tokenPriceTypes.GetPriceOptions{
			Timestamp: uint64(timestamp.Unix()),
		})
		if err == nil {
			return price.Price
		}
	}

	log.WithFields(log.Fields{
		"chain":        chainName,
		"token_symbol": tokenSymbol,
	}).Debug("Attempting fetch price by symbol")
	priceDetails, err := sdk.PriceStore.GetTokenPriceBySymbol(tokenSymbol, &tokenPriceTypes.GetPriceWithSymbolOptions{
		Timestamp: uint64(timestamp.Unix()),
	})
	if err != nil || len(priceDetails) == 0 {
		log.WithFields(log.Fields{"token": tokenSymbol, "time": time.Now().Unix()}).Error("Price not found")
		return "0"
	}
	if err == nil {
		for _, priceDetail := range priceDetails {
			if strings.ToLower(priceDetail.Blockchain) == strings.ToLower(chainName.String()) {
				return priceDetail.Price
			}
		}
	}
	return priceDetails[0].Price
}

func getTokenDetails(tokenAddress string, db *gorm.DB, chainID string) (TokenDetails, error) {
	tokenDetails := TokenDetails{}
	err := db.Where("address = ? and chain_id = ?", strings.ToLower(tokenAddress), chainID).First(&tokenDetails).Error
	if err == nil {
		return tokenDetails, nil
	}
	chain, _ := GetChainFromChainID(chainID)
	fetchedTokenDetails, err := sdk.TokenDetails.GetTokenDetailsByContract(tokenAddress, &sdkTokenDetails.TokenDetailsOptions{chain})
	if err != nil {
		log.WithFields(log.Fields{"token": tokenAddress}).Error("Details not found")
		return TokenDetails{}, err
	}
	tokenDetails = TokenDetails{
		Address: strings.ToLower(tokenAddress),
		Symbol:  fetchedTokenDetails.Symbol,
		Decimal: fetchedTokenDetails.Decimal,
		ChainID: chainID,
	}
	err = db.Create(&tokenDetails).Error
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return tokenDetails, nil
		}
		log.WithFields(log.Fields{"token": tokenAddress}).Error("Token details population error")
		return tokenDetails, err
	}
	return tokenDetails, nil
}
