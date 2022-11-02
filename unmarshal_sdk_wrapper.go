// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"errors"
	unmarshal "github.com/eucrypt/unmarshal-go-sdk/pkg"
	sdkConstants "github.com/eucrypt/unmarshal-go-sdk/pkg/constants"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/transaction_details"
	"github.com/eucrypt/unmarshal-go-sdk/pkg/transaction_details/types"
	log "github.com/sirupsen/logrus"
	"time"
)

const MaximumPerPage = 100

var (
	Base10ChainIDMap = map[string]sdkConstants.Chain{
		"1":     sdkConstants.ETH,
		"56":    sdkConstants.BSC,
		"137":   sdkConstants.MATIC,
		"43114": sdkConstants.AVALANCHE,
		"4":     sdkConstants.ETH_RINKEBY,
		"97":    sdkConstants.BSC_TESTNET,
		"80001": sdkConstants.MATIC_TESTNET,
		"10":    sdkConstants.OPTIMISM,
		"42161": sdkConstants.ARBITRUM,
		"42220": sdkConstants.CELO,
		"250":   sdkConstants.FANTOM,
		"8217":  sdkConstants.KLAYTN,
		"122":   sdkConstants.FUSE,
	}
	PriceSupportedChains = map[sdkConstants.Chain]bool{
		sdkConstants.ETH:   true,
		sdkConstants.BSC:   true,
		sdkConstants.MATIC: true,
	}
)

func GetChainFromChainID(chainID string) (sdkConstants.Chain, error) {
	chain := Base10ChainIDMap[chainID]
	if chain == "" {
		return "", errors.New("unsupported chain error")
	}
	return chain, nil
}

func IsPriceSupportedForChain(chain sdkConstants.Chain) bool {
	return PriceSupportedChains[chain]
}

type UnmarshalSDKWrapper struct {
	unmarshalSDK    *unmarshal.Unmarshal
	chain           sdkConstants.Chain
	numberOfRetries int
}

func NewUnmarshalSDKWrapper(unmarshalSDK *unmarshal.Unmarshal, chain sdkConstants.Chain, numberOfRetries int) *UnmarshalSDKWrapper {
	return &UnmarshalSDKWrapper{
		unmarshalSDK:    unmarshalSDK,
		chain:           chain,
		numberOfRetries: numberOfRetries,
	}
}

//GetAllTransactionsBetween fetches all transactions across a block range
func (u UnmarshalSDKWrapper) GetAllTransactionsBetween(from, to int, contractAddress string) ([]types.RawTransaction, error) {
	var resp = make([]types.RawTransaction, 0)
	for pageNumber := 1; true; pageNumber++ {
		fullResp, err := u.getTransactionsBetween(from, to, pageNumber, contractAddress)
		if err != nil {
			return resp, err
		}
		resp = append(resp, fullResp.Transactions...)
		if !fullResp.NextPage || len(fullResp.Transactions) == 0 {
			break
		}
	}

	return resp, nil
}

//getTransactionsBetween fetches transactions across a block range and for a page number.
func (u UnmarshalSDKWrapper) getTransactionsBetween(from, to, pageNumber int, contractAddress string) (fullResp types.RawTransactionsResponseV1, err error) {
	var (
		paginationDetails = transaction_details.TransactionDetailsOpts{
			PaginationOptions: transaction_details.PaginationOptions{
				Page:     pageNumber,
				PageSize: MaximumPerPage,
			},
			BlockLimitsOpts: transaction_details.BlockLimitsOpts{
				FromBlock: uint64(from),
				ToBlock:   uint64(to),
			},
		}
	)

	for true {
		fullResp, err = u.getTransactionsWithRetry(contractAddress, paginationDetails)
		if err != nil {
			return
		}
		if fullResp.LastVerifiedBlock == nil {
			return types.RawTransactionsResponseV1{}, errors.New("received incompatible response from gateway")
		}
		if fullResp.LastVerifiedBlock.Int64() >= int64(to) {
			return
		}
		log.WithField("Last verified block", fullResp.LastVerifiedBlock.String()).Info("Waiting for verifier to catch up.")
		time.Sleep(15 * time.Second)
	}

	return fullResp, err
}

func (u UnmarshalSDKWrapper) getTransactionsWithRetry(contractAddress string, paginationDetails transaction_details.TransactionDetailsOpts) (fullResp types.RawTransactionsResponseV1, err error) {
	for i := 0; i < u.numberOfRetries; i++ {
		fullResp, err = u.unmarshalSDK.GetRawTransactionsForAddress(u.chain, contractAddress, &paginationDetails)
		if err == nil {
			break
		}
	}
	return fullResp, err
}
