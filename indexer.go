// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"context"
	"encoding/json"
	watcher "github.com/HydroProtocol/ethereum-watcher"
	"github.com/HydroProtocol/ethereum-watcher/blockchain"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	sdkTransactionTypes "github.com/eucrypt/unmarshal-go-sdk/pkg/transaction_details/types"
	"github.com/jackc/pgconn"
	"github.com/onrik/ethrpc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	commonLog "log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	_ = sdkTransactionTypes.RawTransaction{}
	_ = pgconn.PgError{}
	_ = big.NewInt
	_ = ethrpc.Transaction{}
	_ = sync.Mutex{}
)

func ToEthLog(txLog blockchain.ReceiptLog) types.Log {
	hashTopics := make([]common.Hash, 0)
	for _, topic := range txLog.GetTopics() {
		hashTopics = append(hashTopics, common.HexToHash(topic))
	}
	return types.Log{
		Address:     common.HexToAddress(txLog.GetAddress()),
		Topics:      hashTopics,
		Data:        common.FromHex(txLog.GetData()),
		BlockNumber: uint64(txLog.GetBlockNum()),
		TxHash:      common.HexToHash(txLog.GetTransactionHash()),
		TxIndex:     uint(txLog.GetTransactionIndex()),
		BlockHash:   common.HexToHash(txLog.GetBlockHash()),
		Index:       uint(txLog.GetLogIndex()),
		Removed:     txLog.GetRemoved(),
	}
}

func NewPostgresOrm(postgresConfig PostgresConfig, autoMigrateTypes ...interface{}) (*gorm.DB, error) {
	customLogger := logger.New(
		commonLog.New(os.Stdout, "\r\n", commonLog.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	orm, err := gorm.Open(postgres.Open(postgresConfig.ConnectionString), &gorm.Config{
		Logger:          customLogger,
		CreateBatchSize: postgresConfig.CreateBatchSize,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   postgresConfig.TablePrefix,
			SingularTable: false,
		}})
	if err != nil {
		return nil, err
	}
	err = orm.AutoMigrate(autoMigrateTypes...)
	if err != nil {
		return nil, err
	}
	return orm, nil
}

type EthWithdrawnEventTracker struct {
	Orm          *gorm.DB
	MainFilterer *MainFilterer
	chainID      string
}

func NewEthWithdrawnEventTracker(contractAddress string, orm *gorm.DB, ethClient *ethclient.Client, chainID string) (EthWithdrawnEventTracker, error) {
	err := orm.AutoMigrate(&EthWithdrawnEvent{})
	if err != nil {
		return EthWithdrawnEventTracker{}, err
	}
	filterer, err := NewMainFilterer(common.HexToAddress(contractAddress), ethClient)
	return EthWithdrawnEventTracker{
		Orm:          orm,
		MainFilterer: filterer,
		chainID:      chainID,
	}, nil
}

func (a EthWithdrawnEventTracker) CreateEthWithdrawn(txLog blockchain.ReceiptLog, ethClient *EthBlockChainRPCWithRetry) (EthWithdrawnEvent, error) {
	blockDetail, err := ethClient.GetBlockByNumWithoutDetails(txLog.GetBlockNum())
	if err != nil {
		return EthWithdrawnEvent{}, err
	}
	ethLog := ToEthLog(txLog)
	abiEvent, err := a.MainFilterer.ParseABIEthWithdrawnEvent(ethLog, *txLog.Log, *blockDetail)
	if err != nil {
		return EthWithdrawnEvent{}, err
	}
	transactionData, err := ethClient.GetTxByHash(txLog.TransactionHash)
	if err != nil {
		return EthWithdrawnEvent{}, err
	}
	event := convertToEthWithdrawnEvent(abiEvent, transactionData, a.chainID)
	return event, err
}

type SwapTargetAddedEventTracker struct {
	Orm          *gorm.DB
	MainFilterer *MainFilterer
	chainID      string
}

func NewSwapTargetAddedEventTracker(contractAddress string, orm *gorm.DB, ethClient *ethclient.Client, chainID string) (SwapTargetAddedEventTracker, error) {
	err := orm.AutoMigrate(&SwapTargetAddedEvent{})
	if err != nil {
		return SwapTargetAddedEventTracker{}, err
	}
	filterer, err := NewMainFilterer(common.HexToAddress(contractAddress), ethClient)
	return SwapTargetAddedEventTracker{
		Orm:          orm,
		MainFilterer: filterer,
		chainID:      chainID,
	}, nil
}

func (a SwapTargetAddedEventTracker) CreateSwapTargetAdded(txLog blockchain.ReceiptLog, ethClient *EthBlockChainRPCWithRetry) (SwapTargetAddedEvent, error) {
	blockDetail, err := ethClient.GetBlockByNumWithoutDetails(txLog.GetBlockNum())
	if err != nil {
		return SwapTargetAddedEvent{}, err
	}
	ethLog := ToEthLog(txLog)
	abiEvent, err := a.MainFilterer.ParseABISwapTargetAddedEvent(ethLog, *txLog.Log, *blockDetail)
	if err != nil {
		return SwapTargetAddedEvent{}, err
	}
	transactionData, err := ethClient.GetTxByHash(txLog.TransactionHash)
	if err != nil {
		return SwapTargetAddedEvent{}, err
	}
	event := convertToSwapTargetAddedEvent(abiEvent, transactionData, a.chainID)
	return event, err
}

type SwapTargetRemovedEventTracker struct {
	Orm          *gorm.DB
	MainFilterer *MainFilterer
	chainID      string
}

func NewSwapTargetRemovedEventTracker(contractAddress string, orm *gorm.DB, ethClient *ethclient.Client, chainID string) (SwapTargetRemovedEventTracker, error) {
	err := orm.AutoMigrate(&SwapTargetRemovedEvent{})
	if err != nil {
		return SwapTargetRemovedEventTracker{}, err
	}
	filterer, err := NewMainFilterer(common.HexToAddress(contractAddress), ethClient)
	return SwapTargetRemovedEventTracker{
		Orm:          orm,
		MainFilterer: filterer,
		chainID:      chainID,
	}, nil
}

func (a SwapTargetRemovedEventTracker) CreateSwapTargetRemoved(txLog blockchain.ReceiptLog, ethClient *EthBlockChainRPCWithRetry) (SwapTargetRemovedEvent, error) {
	blockDetail, err := ethClient.GetBlockByNumWithoutDetails(txLog.GetBlockNum())
	if err != nil {
		return SwapTargetRemovedEvent{}, err
	}
	ethLog := ToEthLog(txLog)
	abiEvent, err := a.MainFilterer.ParseABISwapTargetRemovedEvent(ethLog, *txLog.Log, *blockDetail)
	if err != nil {
		return SwapTargetRemovedEvent{}, err
	}
	transactionData, err := ethClient.GetTxByHash(txLog.TransactionHash)
	if err != nil {
		return SwapTargetRemovedEvent{}, err
	}
	event := convertToSwapTargetRemovedEvent(abiEvent, transactionData, a.chainID)
	return event, err
}

type TokenWithdrawnEventTracker struct {
	Orm          *gorm.DB
	MainFilterer *MainFilterer
	chainID      string
}

func NewTokenWithdrawnEventTracker(contractAddress string, orm *gorm.DB, ethClient *ethclient.Client, chainID string) (TokenWithdrawnEventTracker, error) {
	err := orm.AutoMigrate(&TokenWithdrawnEvent{})
	if err != nil {
		return TokenWithdrawnEventTracker{}, err
	}
	filterer, err := NewMainFilterer(common.HexToAddress(contractAddress), ethClient)
	return TokenWithdrawnEventTracker{
		Orm:          orm,
		MainFilterer: filterer,
		chainID:      chainID,
	}, nil
}

func (a TokenWithdrawnEventTracker) CreateTokenWithdrawn(txLog blockchain.ReceiptLog, ethClient *EthBlockChainRPCWithRetry) (TokenWithdrawnEvent, error) {
	blockDetail, err := ethClient.GetBlockByNumWithoutDetails(txLog.GetBlockNum())
	if err != nil {
		return TokenWithdrawnEvent{}, err
	}
	ethLog := ToEthLog(txLog)
	abiEvent, err := a.MainFilterer.ParseABITokenWithdrawnEvent(ethLog, *txLog.Log, *blockDetail)
	if err != nil {
		return TokenWithdrawnEvent{}, err
	}
	transactionData, err := ethClient.GetTxByHash(txLog.TransactionHash)
	if err != nil {
		return TokenWithdrawnEvent{}, err
	}
	event := convertToTokenWithdrawnEvent(abiEvent, transactionData, a.chainID)
	return event, err
}

func EventIndexCallback(orm *gorm.DB, ethWithdrawnEventTracker EthWithdrawnEventTracker, swapTargetAddedEventTracker SwapTargetAddedEventTracker, swapTargetRemovedEventTracker SwapTargetRemovedEventTracker, tokenWithdrawnEventTracker TokenWithdrawnEventTracker, txLogs []ethrpc.Log, ethClient *EthBlockChainRPCWithRetry) error {
	errs, _ := errgroup.WithContext(context.Background())

	ethWithdrawnEventsLock := new(sync.Mutex)
	ethWithdrawnEvents := make([]EthWithdrawnEvent, 0)

	swapTargetAddedEventsLock := new(sync.Mutex)
	swapTargetAddedEvents := make([]SwapTargetAddedEvent, 0)

	swapTargetRemovedEventsLock := new(sync.Mutex)
	swapTargetRemovedEvents := make([]SwapTargetRemovedEvent, 0)

	tokenWithdrawnEventsLock := new(sync.Mutex)
	tokenWithdrawnEvents := make([]TokenWithdrawnEvent, 0)

	for i := range txLogs {
		i := i

		errs.Go(func() error {
			return func(i int) error {
				txLog := blockchain.IReceiptLog(blockchain.ReceiptLog{Log: &txLogs[i]})
				switch txLog.GetTopics()[0] {

				case GetEthWithdrawnEventHash():
					event, err := ethWithdrawnEventTracker.CreateEthWithdrawn(txLog.(blockchain.ReceiptLog), ethClient)
					if err != nil {
						log.WithFields(log.Fields{"txHash": txLog.GetTransactionHash(), "blockNumber": txLog.GetBlockNum(), "err": err}).
							Error("EthWithdrawnEvent: Failed to create event from log")
						return err
					}
					err = event.BeforeCreateHook(orm)
					if err != nil {
						log.WithField("err", err).Error("Error making pre create hook call")
					}
					ethWithdrawnEventsLock.Lock()
					ethWithdrawnEvents = append(ethWithdrawnEvents, event)
					ethWithdrawnEventsLock.Unlock()

				case GetSwapTargetAddedEventHash():
					event, err := swapTargetAddedEventTracker.CreateSwapTargetAdded(txLog.(blockchain.ReceiptLog), ethClient)
					if err != nil {
						log.WithFields(log.Fields{"txHash": txLog.GetTransactionHash(), "blockNumber": txLog.GetBlockNum(), "err": err}).
							Error("SwapTargetAddedEvent: Failed to create event from log")
						return err
					}
					err = event.BeforeCreateHook(orm)
					if err != nil {
						log.WithField("err", err).Error("Error making pre create hook call")
					}
					swapTargetAddedEventsLock.Lock()
					swapTargetAddedEvents = append(swapTargetAddedEvents, event)
					swapTargetAddedEventsLock.Unlock()

				case GetSwapTargetRemovedEventHash():
					event, err := swapTargetRemovedEventTracker.CreateSwapTargetRemoved(txLog.(blockchain.ReceiptLog), ethClient)
					if err != nil {
						log.WithFields(log.Fields{"txHash": txLog.GetTransactionHash(), "blockNumber": txLog.GetBlockNum(), "err": err}).
							Error("SwapTargetRemovedEvent: Failed to create event from log")
						return err
					}
					err = event.BeforeCreateHook(orm)
					if err != nil {
						log.WithField("err", err).Error("Error making pre create hook call")
					}
					swapTargetRemovedEventsLock.Lock()
					swapTargetRemovedEvents = append(swapTargetRemovedEvents, event)
					swapTargetRemovedEventsLock.Unlock()

				case GetTokenWithdrawnEventHash():
					event, err := tokenWithdrawnEventTracker.CreateTokenWithdrawn(txLog.(blockchain.ReceiptLog), ethClient)
					if err != nil {
						log.WithFields(log.Fields{"txHash": txLog.GetTransactionHash(), "blockNumber": txLog.GetBlockNum(), "err": err}).
							Error("TokenWithdrawnEvent: Failed to create event from log")
						return err
					}
					err = event.BeforeCreateHook(orm)
					if err != nil {
						log.WithField("err", err).Error("Error making pre create hook call")
					}
					tokenWithdrawnEventsLock.Lock()
					tokenWithdrawnEvents = append(tokenWithdrawnEvents, event)
					tokenWithdrawnEventsLock.Unlock()

				default:
					return nil
				}
				return nil
			}(i)
		})

	}
	err := errs.Wait()
	if err != nil {
		return err
	}

	wg := new(sync.WaitGroup)

	err = orm.Clauses(clause.OnConflict{DoNothing: true}).Create(&ethWithdrawnEvents).Error
	if err != nil {
		return err
	}
	for _, event := range ethWithdrawnEvents {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			err = event.AfterCreateHook(orm)
			if err != nil {
				log.WithField("err", err).Error("EventIndexCallback: Error making post create hook call")
			}
		}(wg)
	}

	err = orm.Clauses(clause.OnConflict{DoNothing: true}).Create(&swapTargetAddedEvents).Error
	if err != nil {
		return err
	}
	for _, event := range swapTargetAddedEvents {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			err = event.AfterCreateHook(orm)
			if err != nil {
				log.WithField("err", err).Error("EventIndexCallback: Error making post create hook call")
			}
		}(wg)
	}

	err = orm.Clauses(clause.OnConflict{DoNothing: true}).Create(&swapTargetRemovedEvents).Error
	if err != nil {
		return err
	}
	for _, event := range swapTargetRemovedEvents {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			err = event.AfterCreateHook(orm)
			if err != nil {
				log.WithField("err", err).Error("EventIndexCallback: Error making post create hook call")
			}
		}(wg)
	}

	err = orm.Clauses(clause.OnConflict{DoNothing: true}).Create(&tokenWithdrawnEvents).Error
	if err != nil {
		return err
	}
	for _, event := range tokenWithdrawnEvents {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			err = event.AfterCreateHook(orm)
			if err != nil {
				log.WithField("err", err).Error("EventIndexCallback: Error making post create hook call")
			}
		}(wg)
	}

	wg.Wait()
	return nil
}

type FillQuoteEthToTokenMethodTracker struct {
	Orm          *gorm.DB
	MainFilterer *MainFilterer
	chainID      string
}

func NewFillQuoteEthToTokenMethodTracker(contractAddress string, orm *gorm.DB, ethClient *ethclient.Client, chainID string) (FillQuoteEthToTokenMethodTracker, error) {
	err := orm.AutoMigrate(&FillQuoteEthToTokenMethod{})
	if err != nil {
		return FillQuoteEthToTokenMethodTracker{}, err
	}
	filterer, err := NewMainFilterer(common.HexToAddress(contractAddress), ethClient)
	return FillQuoteEthToTokenMethodTracker{Orm: orm, MainFilterer: filterer, chainID: chainID}, nil
}

func (a FillQuoteEthToTokenMethodTracker) CreateFillQuoteEthToToken(tx sdkTransactionTypes.RawTransaction) (FillQuoteEthToTokenMethod, error) {
	abiMethod, err := a.MainFilterer.ParseABIFillQuoteEthToTokenMethod(tx)
	if err != nil {
		return FillQuoteEthToTokenMethod{}, err
	}
	method := convertToFillQuoteEthToTokenMethod(abiMethod, a.chainID)
	return method, err
}

type FillQuoteTokenToTokenMethodTracker struct {
	Orm          *gorm.DB
	MainFilterer *MainFilterer
	chainID      string
}

func NewFillQuoteTokenToTokenMethodTracker(contractAddress string, orm *gorm.DB, ethClient *ethclient.Client, chainID string) (FillQuoteTokenToTokenMethodTracker, error) {
	err := orm.AutoMigrate(&FillQuoteTokenToTokenMethod{})
	if err != nil {
		return FillQuoteTokenToTokenMethodTracker{}, err
	}
	filterer, err := NewMainFilterer(common.HexToAddress(contractAddress), ethClient)
	return FillQuoteTokenToTokenMethodTracker{Orm: orm, MainFilterer: filterer, chainID: chainID}, nil
}

func (a FillQuoteTokenToTokenMethodTracker) CreateFillQuoteTokenToToken(tx sdkTransactionTypes.RawTransaction) (FillQuoteTokenToTokenMethod, error) {
	abiMethod, err := a.MainFilterer.ParseABIFillQuoteTokenToTokenMethod(tx)
	if err != nil {
		return FillQuoteTokenToTokenMethod{}, err
	}
	method := convertToFillQuoteTokenToTokenMethod(abiMethod, a.chainID)
	return method, err
}

type UpdateSwapTargetsMethodTracker struct {
	Orm          *gorm.DB
	MainFilterer *MainFilterer
	chainID      string
}

func NewUpdateSwapTargetsMethodTracker(contractAddress string, orm *gorm.DB, ethClient *ethclient.Client, chainID string) (UpdateSwapTargetsMethodTracker, error) {
	err := orm.AutoMigrate(&UpdateSwapTargetsMethod{})
	if err != nil {
		return UpdateSwapTargetsMethodTracker{}, err
	}
	filterer, err := NewMainFilterer(common.HexToAddress(contractAddress), ethClient)
	return UpdateSwapTargetsMethodTracker{Orm: orm, MainFilterer: filterer, chainID: chainID}, nil
}

func (a UpdateSwapTargetsMethodTracker) CreateUpdateSwapTargets(tx sdkTransactionTypes.RawTransaction) (UpdateSwapTargetsMethod, error) {
	abiMethod, err := a.MainFilterer.ParseABIUpdateSwapTargetsMethod(tx)
	if err != nil {
		return UpdateSwapTargetsMethod{}, err
	}
	method := convertToUpdateSwapTargetsMethod(abiMethod, a.chainID)
	return method, err
}

type WithdrawTokenMethodTracker struct {
	Orm          *gorm.DB
	MainFilterer *MainFilterer
	chainID      string
}

func NewWithdrawTokenMethodTracker(contractAddress string, orm *gorm.DB, ethClient *ethclient.Client, chainID string) (WithdrawTokenMethodTracker, error) {
	err := orm.AutoMigrate(&WithdrawTokenMethod{})
	if err != nil {
		return WithdrawTokenMethodTracker{}, err
	}
	filterer, err := NewMainFilterer(common.HexToAddress(contractAddress), ethClient)
	return WithdrawTokenMethodTracker{Orm: orm, MainFilterer: filterer, chainID: chainID}, nil
}

func (a WithdrawTokenMethodTracker) CreateWithdrawToken(tx sdkTransactionTypes.RawTransaction) (WithdrawTokenMethod, error) {
	abiMethod, err := a.MainFilterer.ParseABIWithdrawTokenMethod(tx)
	if err != nil {
		return WithdrawTokenMethod{}, err
	}
	method := convertToWithdrawTokenMethod(abiMethod, a.chainID)
	return method, err
}

type FillQuoteTokenToEthMethodTracker struct {
	Orm          *gorm.DB
	MainFilterer *MainFilterer
	chainID      string
}

func NewFillQuoteTokenToEthMethodTracker(contractAddress string, orm *gorm.DB, ethClient *ethclient.Client, chainID string) (FillQuoteTokenToEthMethodTracker, error) {
	err := orm.AutoMigrate(&FillQuoteTokenToEthMethod{})
	if err != nil {
		return FillQuoteTokenToEthMethodTracker{}, err
	}
	filterer, err := NewMainFilterer(common.HexToAddress(contractAddress), ethClient)
	return FillQuoteTokenToEthMethodTracker{Orm: orm, MainFilterer: filterer, chainID: chainID}, nil
}

func (a FillQuoteTokenToEthMethodTracker) CreateFillQuoteTokenToEth(tx sdkTransactionTypes.RawTransaction) (FillQuoteTokenToEthMethod, error) {
	abiMethod, err := a.MainFilterer.ParseABIFillQuoteTokenToEthMethod(tx)
	if err != nil {
		return FillQuoteTokenToEthMethod{}, err
	}
	method := convertToFillQuoteTokenToEthMethod(abiMethod, a.chainID)
	return method, err
}

type FillQuoteTokenToEthWithPermitMethodTracker struct {
	Orm          *gorm.DB
	MainFilterer *MainFilterer
	chainID      string
}

func NewFillQuoteTokenToEthWithPermitMethodTracker(contractAddress string, orm *gorm.DB, ethClient *ethclient.Client, chainID string) (FillQuoteTokenToEthWithPermitMethodTracker, error) {
	err := orm.AutoMigrate(&FillQuoteTokenToEthWithPermitMethod{})
	if err != nil {
		return FillQuoteTokenToEthWithPermitMethodTracker{}, err
	}
	filterer, err := NewMainFilterer(common.HexToAddress(contractAddress), ethClient)
	return FillQuoteTokenToEthWithPermitMethodTracker{Orm: orm, MainFilterer: filterer, chainID: chainID}, nil
}

func (a FillQuoteTokenToEthWithPermitMethodTracker) CreateFillQuoteTokenToEthWithPermit(tx sdkTransactionTypes.RawTransaction) (FillQuoteTokenToEthWithPermitMethod, error) {
	abiMethod, err := a.MainFilterer.ParseABIFillQuoteTokenToEthWithPermitMethod(tx)
	if err != nil {
		return FillQuoteTokenToEthWithPermitMethod{}, err
	}
	method := convertToFillQuoteTokenToEthWithPermitMethod(abiMethod, a.chainID)
	return method, err
}

type FillQuoteTokenToTokenWithPermitMethodTracker struct {
	Orm          *gorm.DB
	MainFilterer *MainFilterer
	chainID      string
}

func NewFillQuoteTokenToTokenWithPermitMethodTracker(contractAddress string, orm *gorm.DB, ethClient *ethclient.Client, chainID string) (FillQuoteTokenToTokenWithPermitMethodTracker, error) {
	err := orm.AutoMigrate(&FillQuoteTokenToTokenWithPermitMethod{})
	if err != nil {
		return FillQuoteTokenToTokenWithPermitMethodTracker{}, err
	}
	filterer, err := NewMainFilterer(common.HexToAddress(contractAddress), ethClient)
	return FillQuoteTokenToTokenWithPermitMethodTracker{Orm: orm, MainFilterer: filterer, chainID: chainID}, nil
}

func (a FillQuoteTokenToTokenWithPermitMethodTracker) CreateFillQuoteTokenToTokenWithPermit(tx sdkTransactionTypes.RawTransaction) (FillQuoteTokenToTokenWithPermitMethod, error) {
	abiMethod, err := a.MainFilterer.ParseABIFillQuoteTokenToTokenWithPermitMethod(tx)
	if err != nil {
		return FillQuoteTokenToTokenWithPermitMethod{}, err
	}
	method := convertToFillQuoteTokenToTokenWithPermitMethod(abiMethod, a.chainID)
	return method, err
}

type TransferOwnershipMethodTracker struct {
	Orm          *gorm.DB
	MainFilterer *MainFilterer
	chainID      string
}

func NewTransferOwnershipMethodTracker(contractAddress string, orm *gorm.DB, ethClient *ethclient.Client, chainID string) (TransferOwnershipMethodTracker, error) {
	err := orm.AutoMigrate(&TransferOwnershipMethod{})
	if err != nil {
		return TransferOwnershipMethodTracker{}, err
	}
	filterer, err := NewMainFilterer(common.HexToAddress(contractAddress), ethClient)
	return TransferOwnershipMethodTracker{Orm: orm, MainFilterer: filterer, chainID: chainID}, nil
}

func (a TransferOwnershipMethodTracker) CreateTransferOwnership(tx sdkTransactionTypes.RawTransaction) (TransferOwnershipMethod, error) {
	abiMethod, err := a.MainFilterer.ParseABITransferOwnershipMethod(tx)
	if err != nil {
		return TransferOwnershipMethod{}, err
	}
	method := convertToTransferOwnershipMethod(abiMethod, a.chainID)
	return method, err
}

type WithdrawEthMethodTracker struct {
	Orm          *gorm.DB
	MainFilterer *MainFilterer
	chainID      string
}

func NewWithdrawEthMethodTracker(contractAddress string, orm *gorm.DB, ethClient *ethclient.Client, chainID string) (WithdrawEthMethodTracker, error) {
	err := orm.AutoMigrate(&WithdrawEthMethod{})
	if err != nil {
		return WithdrawEthMethodTracker{}, err
	}
	filterer, err := NewMainFilterer(common.HexToAddress(contractAddress), ethClient)
	return WithdrawEthMethodTracker{Orm: orm, MainFilterer: filterer, chainID: chainID}, nil
}

func (a WithdrawEthMethodTracker) CreateWithdrawEth(tx sdkTransactionTypes.RawTransaction) (WithdrawEthMethod, error) {
	abiMethod, err := a.MainFilterer.ParseABIWithdrawEthMethod(tx)
	if err != nil {
		return WithdrawEthMethod{}, err
	}
	method := convertToWithdrawEthMethod(abiMethod, a.chainID)
	return method, err
}

func MethodIndexerCallback(config IndexerConfig, ethClient *ethclient.Client, chainID string, orm *gorm.DB, from, to int, contractAddress string, unmarshalSDKWrapper *UnmarshalSDKWrapper) error {

	fillQuoteEthToTokenMethodTracker, err := NewFillQuoteEthToTokenMethodTracker(config.ContractAddress, orm, ethClient, chainID)
	if err != nil {
		panic(err)
	}

	fillQuoteTokenToTokenMethodTracker, err := NewFillQuoteTokenToTokenMethodTracker(config.ContractAddress, orm, ethClient, chainID)
	if err != nil {
		panic(err)
	}

	updateSwapTargetsMethodTracker, err := NewUpdateSwapTargetsMethodTracker(config.ContractAddress, orm, ethClient, chainID)
	if err != nil {
		panic(err)
	}

	withdrawTokenMethodTracker, err := NewWithdrawTokenMethodTracker(config.ContractAddress, orm, ethClient, chainID)
	if err != nil {
		panic(err)
	}

	fillQuoteTokenToEthMethodTracker, err := NewFillQuoteTokenToEthMethodTracker(config.ContractAddress, orm, ethClient, chainID)
	if err != nil {
		panic(err)
	}

	fillQuoteTokenToEthWithPermitMethodTracker, err := NewFillQuoteTokenToEthWithPermitMethodTracker(config.ContractAddress, orm, ethClient, chainID)
	if err != nil {
		panic(err)
	}

	fillQuoteTokenToTokenWithPermitMethodTracker, err := NewFillQuoteTokenToTokenWithPermitMethodTracker(config.ContractAddress, orm, ethClient, chainID)
	if err != nil {
		panic(err)
	}

	transferOwnershipMethodTracker, err := NewTransferOwnershipMethodTracker(config.ContractAddress, orm, ethClient, chainID)
	if err != nil {
		panic(err)
	}

	withdrawEthMethodTracker, err := NewWithdrawEthMethodTracker(config.ContractAddress, orm, ethClient, chainID)
	if err != nil {
		panic(err)
	}

	errs, _ := errgroup.WithContext(context.Background())

	fillQuoteEthToTokenMethodsLock := new(sync.Mutex)
	fillQuoteEthToTokenMethods := make([]FillQuoteEthToTokenMethod, 0)

	fillQuoteTokenToTokenMethodsLock := new(sync.Mutex)
	fillQuoteTokenToTokenMethods := make([]FillQuoteTokenToTokenMethod, 0)

	updateSwapTargetsMethodsLock := new(sync.Mutex)
	updateSwapTargetsMethods := make([]UpdateSwapTargetsMethod, 0)

	withdrawTokenMethodsLock := new(sync.Mutex)
	withdrawTokenMethods := make([]WithdrawTokenMethod, 0)

	fillQuoteTokenToEthMethodsLock := new(sync.Mutex)
	fillQuoteTokenToEthMethods := make([]FillQuoteTokenToEthMethod, 0)

	fillQuoteTokenToEthWithPermitMethodsLock := new(sync.Mutex)
	fillQuoteTokenToEthWithPermitMethods := make([]FillQuoteTokenToEthWithPermitMethod, 0)

	fillQuoteTokenToTokenWithPermitMethodsLock := new(sync.Mutex)
	fillQuoteTokenToTokenWithPermitMethods := make([]FillQuoteTokenToTokenWithPermitMethod, 0)

	transferOwnershipMethodsLock := new(sync.Mutex)
	transferOwnershipMethods := make([]TransferOwnershipMethod, 0)

	withdrawEthMethodsLock := new(sync.Mutex)
	withdrawEthMethods := make([]WithdrawEthMethod, 0)

	transactionList, err := unmarshalSDKWrapper.GetAllTransactionsBetween(from, to, contractAddress)
	if err != nil {
		log.WithFields(log.Fields{"start Block": from, "to block": to, "err": err}).
			Error("Failed to fetch transactions for a block")
		return err
	}
	if transactionList == nil {
		return nil
	}
	for j := range transactionList {
		j := j
		errs.Go(func() error {
			return func(j int) error {
				tx := transactionList[j]
				if tx.AdditionalData.Status != 1 {
					return nil
				}
				if tx.To == "" {
					return nil
				}
				if strings.ToLower(tx.To) != strings.ToLower(contractAddress) {
					return nil
				}
				if len(tx.AdditionalData.Data) < 10 {
					return nil
				}
				switch methodHex := tx.AdditionalData.Data[2:10]; methodHex {

				case GetFillQuoteEthToTokenMethodHash():
					method, err := fillQuoteEthToTokenMethodTracker.CreateFillQuoteEthToToken(tx)
					if err != nil {
						log.WithFields(log.Fields{"txHash": tx.TxHash, "blockNumber": tx.BlockNumber, "err": err}).
							Error("FillQuoteEthToTokenMethod: Failed to write the method execution")
						return err
					}
					err = method.BeforeCreateHook(orm)
					if err != nil {
						log.WithField("err", err).Error("Error making pre create hook call")
					}
					fillQuoteEthToTokenMethodsLock.Lock()
					fillQuoteEthToTokenMethods = append(fillQuoteEthToTokenMethods, method)
					fillQuoteEthToTokenMethodsLock.Unlock()

				case GetFillQuoteTokenToTokenMethodHash():
					method, err := fillQuoteTokenToTokenMethodTracker.CreateFillQuoteTokenToToken(tx)
					if err != nil {
						log.WithFields(log.Fields{"txHash": tx.TxHash, "blockNumber": tx.BlockNumber, "err": err}).
							Error("FillQuoteTokenToTokenMethod: Failed to write the method execution")
						return err
					}
					err = method.BeforeCreateHook(orm)
					if err != nil {
						log.WithField("err", err).Error("Error making pre create hook call")
					}
					fillQuoteTokenToTokenMethodsLock.Lock()
					fillQuoteTokenToTokenMethods = append(fillQuoteTokenToTokenMethods, method)
					fillQuoteTokenToTokenMethodsLock.Unlock()

				case GetUpdateSwapTargetsMethodHash():
					method, err := updateSwapTargetsMethodTracker.CreateUpdateSwapTargets(tx)
					if err != nil {
						log.WithFields(log.Fields{"txHash": tx.TxHash, "blockNumber": tx.BlockNumber, "err": err}).
							Error("UpdateSwapTargetsMethod: Failed to write the method execution")
						return err
					}
					err = method.BeforeCreateHook(orm)
					if err != nil {
						log.WithField("err", err).Error("Error making pre create hook call")
					}
					updateSwapTargetsMethodsLock.Lock()
					updateSwapTargetsMethods = append(updateSwapTargetsMethods, method)
					updateSwapTargetsMethodsLock.Unlock()

				case GetWithdrawTokenMethodHash():
					method, err := withdrawTokenMethodTracker.CreateWithdrawToken(tx)
					if err != nil {
						log.WithFields(log.Fields{"txHash": tx.TxHash, "blockNumber": tx.BlockNumber, "err": err}).
							Error("WithdrawTokenMethod: Failed to write the method execution")
						return err
					}
					err = method.BeforeCreateHook(orm)
					if err != nil {
						log.WithField("err", err).Error("Error making pre create hook call")
					}
					withdrawTokenMethodsLock.Lock()
					withdrawTokenMethods = append(withdrawTokenMethods, method)
					withdrawTokenMethodsLock.Unlock()

				case GetFillQuoteTokenToEthMethodHash():
					method, err := fillQuoteTokenToEthMethodTracker.CreateFillQuoteTokenToEth(tx)
					if err != nil {
						log.WithFields(log.Fields{"txHash": tx.TxHash, "blockNumber": tx.BlockNumber, "err": err}).
							Error("FillQuoteTokenToEthMethod: Failed to write the method execution")
						return err
					}
					err = method.BeforeCreateHook(orm)
					if err != nil {
						log.WithField("err", err).Error("Error making pre create hook call")
					}
					fillQuoteTokenToEthMethodsLock.Lock()
					fillQuoteTokenToEthMethods = append(fillQuoteTokenToEthMethods, method)
					fillQuoteTokenToEthMethodsLock.Unlock()

				case GetFillQuoteTokenToEthWithPermitMethodHash():
					method, err := fillQuoteTokenToEthWithPermitMethodTracker.CreateFillQuoteTokenToEthWithPermit(tx)
					if err != nil {
						log.WithFields(log.Fields{"txHash": tx.TxHash, "blockNumber": tx.BlockNumber, "err": err}).
							Error("FillQuoteTokenToEthWithPermitMethod: Failed to write the method execution")
						return err
					}
					err = method.BeforeCreateHook(orm)
					if err != nil {
						log.WithField("err", err).Error("Error making pre create hook call")
					}
					fillQuoteTokenToEthWithPermitMethodsLock.Lock()
					fillQuoteTokenToEthWithPermitMethods = append(fillQuoteTokenToEthWithPermitMethods, method)
					fillQuoteTokenToEthWithPermitMethodsLock.Unlock()

				case GetFillQuoteTokenToTokenWithPermitMethodHash():
					method, err := fillQuoteTokenToTokenWithPermitMethodTracker.CreateFillQuoteTokenToTokenWithPermit(tx)
					if err != nil {
						log.WithFields(log.Fields{"txHash": tx.TxHash, "blockNumber": tx.BlockNumber, "err": err}).
							Error("FillQuoteTokenToTokenWithPermitMethod: Failed to write the method execution")
						return err
					}
					err = method.BeforeCreateHook(orm)
					if err != nil {
						log.WithField("err", err).Error("Error making pre create hook call")
					}
					fillQuoteTokenToTokenWithPermitMethodsLock.Lock()
					fillQuoteTokenToTokenWithPermitMethods = append(fillQuoteTokenToTokenWithPermitMethods, method)
					fillQuoteTokenToTokenWithPermitMethodsLock.Unlock()

				case GetTransferOwnershipMethodHash():
					method, err := transferOwnershipMethodTracker.CreateTransferOwnership(tx)
					if err != nil {
						log.WithFields(log.Fields{"txHash": tx.TxHash, "blockNumber": tx.BlockNumber, "err": err}).
							Error("TransferOwnershipMethod: Failed to write the method execution")
						return err
					}
					err = method.BeforeCreateHook(orm)
					if err != nil {
						log.WithField("err", err).Error("Error making pre create hook call")
					}
					transferOwnershipMethodsLock.Lock()
					transferOwnershipMethods = append(transferOwnershipMethods, method)
					transferOwnershipMethodsLock.Unlock()

				case GetWithdrawEthMethodHash():
					method, err := withdrawEthMethodTracker.CreateWithdrawEth(tx)
					if err != nil {
						log.WithFields(log.Fields{"txHash": tx.TxHash, "blockNumber": tx.BlockNumber, "err": err}).
							Error("WithdrawEthMethod: Failed to write the method execution")
						return err
					}
					err = method.BeforeCreateHook(orm)
					if err != nil {
						log.WithField("err", err).Error("Error making pre create hook call")
					}
					withdrawEthMethodsLock.Lock()
					withdrawEthMethods = append(withdrawEthMethods, method)
					withdrawEthMethodsLock.Unlock()

				default:
					return nil
				}
				return nil
			}(j)
		})
	}

	err = errs.Wait()
	if err != nil {
		return err
	}

	wg := new(sync.WaitGroup)

	err = orm.Clauses(clause.OnConflict{DoNothing: true}).Create(&fillQuoteEthToTokenMethods).Error
	if err != nil {
		return err
	}
	for _, method := range fillQuoteEthToTokenMethods {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			err = method.AfterCreateHook(orm)
			if err != nil {
				log.WithField("err", err).Error("MethodIndexerCallback: Error making post create hook call")
			}
		}(wg)
	}

	err = orm.Clauses(clause.OnConflict{DoNothing: true}).Create(&fillQuoteTokenToTokenMethods).Error
	if err != nil {
		return err
	}
	for _, method := range fillQuoteTokenToTokenMethods {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			err = method.AfterCreateHook(orm)
			if err != nil {
				log.WithField("err", err).Error("MethodIndexerCallback: Error making post create hook call")
			}
		}(wg)
	}

	err = orm.Clauses(clause.OnConflict{DoNothing: true}).Create(&updateSwapTargetsMethods).Error
	if err != nil {
		return err
	}
	for _, method := range updateSwapTargetsMethods {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			err = method.AfterCreateHook(orm)
			if err != nil {
				log.WithField("err", err).Error("MethodIndexerCallback: Error making post create hook call")
			}
		}(wg)
	}

	err = orm.Clauses(clause.OnConflict{DoNothing: true}).Create(&withdrawTokenMethods).Error
	if err != nil {
		return err
	}
	for _, method := range withdrawTokenMethods {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			err = method.AfterCreateHook(orm)
			if err != nil {
				log.WithField("err", err).Error("MethodIndexerCallback: Error making post create hook call")
			}
		}(wg)
	}

	err = orm.Clauses(clause.OnConflict{DoNothing: true}).Create(&fillQuoteTokenToEthMethods).Error
	if err != nil {
		return err
	}
	for _, method := range fillQuoteTokenToEthMethods {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			err = method.AfterCreateHook(orm)
			if err != nil {
				log.WithField("err", err).Error("MethodIndexerCallback: Error making post create hook call")
			}
		}(wg)
	}

	err = orm.Clauses(clause.OnConflict{DoNothing: true}).Create(&fillQuoteTokenToEthWithPermitMethods).Error
	if err != nil {
		return err
	}
	for _, method := range fillQuoteTokenToEthWithPermitMethods {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			err = method.AfterCreateHook(orm)
			if err != nil {
				log.WithField("err", err).Error("MethodIndexerCallback: Error making post create hook call")
			}
		}(wg)
	}

	err = orm.Clauses(clause.OnConflict{DoNothing: true}).Create(&fillQuoteTokenToTokenWithPermitMethods).Error
	if err != nil {
		return err
	}
	for _, method := range fillQuoteTokenToTokenWithPermitMethods {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			err = method.AfterCreateHook(orm)
			if err != nil {
				log.WithField("err", err).Error("MethodIndexerCallback: Error making post create hook call")
			}
		}(wg)
	}

	err = orm.Clauses(clause.OnConflict{DoNothing: true}).Create(&transferOwnershipMethods).Error
	if err != nil {
		return err
	}
	for _, method := range transferOwnershipMethods {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			err = method.AfterCreateHook(orm)
			if err != nil {
				log.WithField("err", err).Error("MethodIndexerCallback: Error making post create hook call")
			}
		}(wg)
	}

	err = orm.Clauses(clause.OnConflict{DoNothing: true}).Create(&withdrawEthMethods).Error
	if err != nil {
		return err
	}
	for _, method := range withdrawEthMethods {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			err = method.AfterCreateHook(orm)
			if err != nil {
				log.WithField("err", err).Error("MethodIndexerCallback: Error making post create hook call")
			}
		}(wg)
	}

	wg.Wait()
	return nil
}

type EthBlockChainRPCWithRetry struct {
	client        *ethrpc.EthRPC
	maxRetryTimes int
}

func NewEthRPCWithRetry(api string, maxRetryCount int) *EthBlockChainRPCWithRetry {
	rpc := ethrpc.New(api)
	return &EthBlockChainRPCWithRetry{rpc, maxRetryCount}
}

func (rpc EthBlockChainRPCWithRetry) GetBlockByNumWithoutDetails(num int) (resp *ethrpc.Block, err error) {
	for i := 0; i <= rpc.maxRetryTimes; i++ {
		resp, err = rpc.client.EthGetBlockByNumber(num, false)
		if err == nil && isNotEmptyBlockDetails(resp) {
			break
		}
		time.Sleep(time.Duration(500*(i+1)) * time.Millisecond)
	}
	return
}

func isNotEmptyBlockDetails(block *ethrpc.Block) bool {
	if block == nil {
		return false
	}
	if block.Number == 0 && block.Timestamp == 0 && block.Hash == "" {
		return false
	}
	return true
}

func (rpc EthBlockChainRPCWithRetry) GetTxByHash(txHash string) (resp *ethrpc.Transaction, err error) {
	for i := 0; i <= rpc.maxRetryTimes; i++ {
		resp, err = rpc.client.EthGetTransactionByHash(txHash)
		if err == nil && isNotEmptyTxDetails(resp) {
			break
		}
		time.Sleep(time.Duration(500*(i+1)) * time.Millisecond)
	}
	return
}

func isNotEmptyTxDetails(tx *ethrpc.Transaction) bool {
	if tx == nil {
		return false
	}
	if tx.BlockNumber == nil && tx.TransactionIndex == nil && tx.Hash == "" {
		return false
	}
	return true
}

func (rpc *EthBlockChainRPCWithRetry) makeEthCallAndSave(methodName string, result interface{}) (err error) {
	var resp json.RawMessage
	for i := 0; i <= rpc.maxRetryTimes; i++ {
		resp, err = rpc.client.Call(methodName)
		if err != nil || resp == nil {
			log.WithField("Method Name", methodName).Error("RPC Call failed")
			time.Sleep(time.Duration(500*(i+1)) * time.Millisecond)
			continue
		}
		err = json.Unmarshal(resp, result)
		if err != nil {
			log.WithField("Method Name", methodName).Error("Failed to unmarshal response")
			time.Sleep(time.Duration(500*(i+1)) * time.Millisecond)
			continue
		}
		break
	}
	return
}

type SyncDB struct {
	orm             *gorm.DB
	contractAddress string
	chainID         string
	syncType        string
}

func NewSyncDB(orm *gorm.DB, contract, chainID string, syncType string) (SyncDB, error) {
	err := orm.AutoMigrate(&LastSyncedBlock{})
	if err != nil {
		return SyncDB{}, err
	}
	return SyncDB{orm: orm, contractAddress: contract, chainID: chainID, syncType: syncType}, nil
}

func (sync SyncDB) UpdateLastSynced(blockNum int) error {
	res := sync.orm.Save(&LastSyncedBlock{
		Contract:    sync.contractAddress,
		ChainID:     sync.chainID,
		BlockNumber: uint64(blockNum),
		SyncType:    sync.syncType,
	})
	if res.Error == nil {
		log.Infof("LastSyncedBlock: success, SyncType: %s , BlockNumber: %d", sync.syncType, blockNum)
	}
	return res.Error
}

func (sync SyncDB) GetLastSyncedBlock() (int, error) {
	lastSync := &LastSyncedBlock{Contract: sync.contractAddress}
	res := sync.orm.First(lastSync, "contract = ? and chain_id = ? and sync_type = ?", sync.contractAddress, sync.chainID, sync.syncType)
	return int(lastSync.BlockNumber), res.Error
}

func RunIndexer(config IndexerConfig) {
	initUnmarshalSDK(config)
	topicsInterestedIn := []string{
		GetEthWithdrawnEventHash(),
		GetSwapTargetAddedEventHash(),
		GetSwapTargetRemovedEventHash(),
		GetTokenWithdrawnEventHash(),
	}
	methodsInterestedIn := []string{
		GetFillQuoteEthToTokenMethodHash(),
		GetFillQuoteTokenToTokenMethodHash(),
		GetUpdateSwapTargetsMethodHash(),
		GetWithdrawTokenMethodHash(),
		GetFillQuoteTokenToEthMethodHash(),
		GetFillQuoteTokenToEthWithPermitMethodHash(),
		GetFillQuoteTokenToTokenWithPermitMethodHash(),
		GetTransferOwnershipMethodHash(),
		GetWithdrawEthMethodHash(),
	}
	orm, err := NewPostgresOrm(config.PostgresConfig)
	if err != nil {
		panic(err)
	}

	ethClient := new(ethclient.Client)
	ethRPCClient := NewEthRPCWithRetry(config.EthEndpoint, 3)
	chainID, err := getChainID(ethRPCClient)
	if err != nil || chainID == "" {
		log.Error("Failed to get chain id, Shutting down.")
		panic(err)
	}
	chain, err := GetChainFromChainID(chainID)
	if err != nil || chainID == "" {
		log.Error("Chain not supported by SDK")
		panic(err)
	}

	err = InitPluginModels(orm)
	if err != nil {
		panic(err)
	}

	var unmarshalSDKWrapper = NewUnmarshalSDKWrapper(&sdk, chain, 4)
	receiptsChannel := make(chan Receipts, 10)
	var nearHighestBlock bool = false

	wg := new(sync.WaitGroup)

	//syncing events
	if len(topicsInterestedIn) > 0 {
		wg.Add(2)
		fetchAndPushLogsToChannel(config, chainID, topicsInterestedIn, receiptsChannel, wg, &nearHighestBlock)
		processLogsFromChannel(chainID, config, ethClient, receiptsChannel, topicsInterestedIn, ethRPCClient, unmarshalSDKWrapper, wg, &nearHighestBlock)
	}

	//syncing methods
	if len(methodsInterestedIn) > 0 {
		wg.Add(1)
		syncMethods(config, ethClient, chainID, orm, chainID, methodsInterestedIn, unmarshalSDKWrapper, wg)
	}
	wg.Wait()
}

func fetchAndPushLogsToChannel(config IndexerConfig, chainID string, topicsInterestedIn []string, receiptsChannel chan Receipts, wg *sync.WaitGroup, nearHighestBlock *bool) {
	orm, err := NewPostgresOrm(config.PostgresConfig)
	if err != nil {
		panic(err)
	}

	syncDB, err := NewSyncDB(orm, config.ContractAddress, chainID, "events")
	if err != nil {
		panic(err)
	}

	go func() {
		log.Infof("Started go routine to fetch logs")
		startBlock := config.StartBlock
		lastUpdatedBlock, err := syncDB.GetLastSyncedBlock()
		if err == nil {
			if lastUpdatedBlock > startBlock {
				startBlock = lastUpdatedBlock + 1
			}
		}
		rpcClient := CreateEthClientConnectionWithRetry(config)
		log.Infof("StartBlock ----> %d\n", startBlock)
		PARALLEL_CALLS_FOR_LOGS := config.ParallelCalls
		STEP_SIZE := config.StepSize
		mutex := sync.Mutex{}
		mutex1 := sync.Mutex{}
		for true {
			errs, _ := errgroup.WithContext(context.Background())
			receiptsMap := map[int]Receipts{}
			highestBlockNumber, e := rpcClient.EthBlockNumber()
			if e != nil {
				log.Error("fetchAndPushLogsToChannel --> could not get highest block number")
				panic(err)
			}
			highestBlockNumber -= config.LagToHighestBlock
			if highestBlockNumber-(startBlock+(PARALLEL_CALLS_FOR_LOGS*STEP_SIZE)) < 0 {
				log.Infof("Not enough blocks, waiting for new blocks , PARALLEL_CALLS_FOR_LOGS=%d, STEP_SIZE=%d", PARALLEL_CALLS_FOR_LOGS, STEP_SIZE)
				mutex1.Lock()
				*nearHighestBlock = true
				mutex1.Unlock()
				if PARALLEL_CALLS_FOR_LOGS == 1 {
					time.Sleep(3 * time.Second)
					continue
				}
				PARALLEL_CALLS_FOR_LOGS = 1
				STEP_SIZE = 1
				log.Infof("Updated PARALLEL_CALLS_FOR_LOGS and STEP_SIZE , PARALLEL_CALLS_FOR_LOGS=%d, STEP_SIZE=%d", PARALLEL_CALLS_FOR_LOGS, STEP_SIZE)
				continue
			}
			for i := 0; i < PARALLEL_CALLS_FOR_LOGS; i++ {
				i := i
				start := startBlock
				errs.Go(func() error {
					return func(i int, start int) error {
						fetchedReceipts := Receipts{
							receipts: FetchLogs(config.ContractAddress, rpcClient, start+(i*STEP_SIZE), start+(i*STEP_SIZE)+STEP_SIZE-1, topicsInterestedIn),
							i:        i,
							from:     start + (i * STEP_SIZE),
							to:       start + (i * STEP_SIZE) + STEP_SIZE - 1,
						}
						mutex.Lock()
						receiptsMap[i] = fetchedReceipts
						mutex.Unlock()
						return nil
					}(i, start)
					return nil
				})
			}
			startBlock = startBlock + (PARALLEL_CALLS_FOR_LOGS * STEP_SIZE)
			err := errs.Wait()
			if err != nil {
				log.Error("fetchAndPushLogsToChannel, " + err.Error())
				panic(err)
			}
			for i := 0; i < PARALLEL_CALLS_FOR_LOGS; i++ {
				receiptsChannel <- receiptsMap[i]
			}
		}
		wg.Done()
	}()
}

func processLogsFromChannel(chainID string, config IndexerConfig, ethClient *ethclient.Client, receiptsChannel chan Receipts, topicsInterestedIn []string, ethRPCClient *EthBlockChainRPCWithRetry, unmarshalSDKWrapper *UnmarshalSDKWrapper, wg *sync.WaitGroup, nearHighestBlock *bool) {
	orm, err := NewPostgresOrm(config.PostgresConfig)
	if err != nil {
		panic(err)
	}

	syncDB, err := NewSyncDB(orm, config.ContractAddress, chainID, "events")
	if err != nil {
		panic(err)
	}

	go func() {

		ethWithdrawnEventTracker, err := NewEthWithdrawnEventTracker(config.ContractAddress, orm, ethClient, chainID)
		if err != nil {
			panic(err)
		}

		swapTargetAddedEventTracker, err := NewSwapTargetAddedEventTracker(config.ContractAddress, orm, ethClient, chainID)
		if err != nil {
			panic(err)
		}

		swapTargetRemovedEventTracker, err := NewSwapTargetRemovedEventTracker(config.ContractAddress, orm, ethClient, chainID)
		if err != nil {
			panic(err)
		}

		tokenWithdrawnEventTracker, err := NewTokenWithdrawnEventTracker(config.ContractAddress, orm, ethClient, chainID)
		if err != nil {
			panic(err)
		}

		log.Info("Starting go routine for processing logs")
		for true {
			receiptLogs := <-receiptsChannel
			from := receiptLogs.from
			to := receiptLogs.to

			log.Infof("Processing logs from %d to %d, total logs %d", from, to, len(receiptLogs.receipts))
			if len(topicsInterestedIn) > 0 {

				err := EventIndexCallback(orm, ethWithdrawnEventTracker, swapTargetAddedEventTracker, swapTargetRemovedEventTracker, tokenWithdrawnEventTracker, receiptLogs.receipts, ethRPCClient)
				if err != nil {
					log.Fatalf(err.Error())
				}
			}
			err := syncDB.UpdateLastSynced(to)
			if err != nil {
				log.Error("Error while updating last synced block")
				panic(err)
			}
			log.Infof("Finished Processing logs from %d to %d", from, to)
		}
		wg.Done()
	}()
}

func syncMethods(config IndexerConfig, ethClient *ethclient.Client, id string, orm *gorm.DB, chainID string, methodsInterestedIn []string, unmarshalSDKWrapper *UnmarshalSDKWrapper, wg *sync.WaitGroup) {
	orm, err := NewPostgresOrm(config.PostgresConfig)
	if err != nil {
		panic(err)
	}

	syncDB, err := NewSyncDB(orm, config.ContractAddress, chainID, "methods")
	if err != nil {
		panic(err)
	}

	STEP_SIZE := config.StepSize
	startBlock := config.StartBlock
	lastUpdatedBlock, err := syncDB.GetLastSyncedBlock()
	if err == nil {
		if lastUpdatedBlock > startBlock {
			startBlock = lastUpdatedBlock + 1
		}
	}
	rpcClient := CreateEthClientConnectionWithRetry(config)

	go func() {
		log.Infof("syncMethods, StartBlock ----> %d\n", startBlock)
		for true {
			highestBlockNumber, e := rpcClient.EthBlockNumber()
			if e != nil {
				log.Error("syncMethods --> could not get highest block number")
				panic(err)
			}
			highestBlockNumber -= config.LagToHighestBlock
			if highestBlockNumber-(startBlock+STEP_SIZE) < 0 {
				log.Infof("syncMethods ---> Not enough blocks, waiting for new blocks , STEP_SIZE=%d", STEP_SIZE)
				time.Sleep(3 * time.Second)
				STEP_SIZE = 1
				log.Infof("syncMethods ---> Updated STEP_SIZE , STEP_SIZE=%d", STEP_SIZE)
				continue
			}

			for true {
				from := startBlock
				to := startBlock + STEP_SIZE
				err = MethodIndexerCallback(config, ethClient, chainID, orm, from, to, config.ContractAddress, unmarshalSDKWrapper)
				if err != nil {
					log.Panic(err)
				}
				err = syncDB.UpdateLastSynced(to)
				if err != nil {
					panic(err)
				}
				startBlock = to
			}
		}
	}()
}

type Receipts struct {
	from     int
	to       int
	i        int
	receipts []ethrpc.Log
	pairMap  map[string]bool
}

func CreateEthClientConnectionWithRetry(config IndexerConfig) *ethrpc.EthRPC {
	return ethrpc.New(config.EthEndpoint)
}

func FetchLogs(contractAddress string, rpcClient *ethrpc.EthRPC, from int, to int, topicsInterested []string) []ethrpc.Log {
	filterParam := ethrpc.FilterParams{
		FromBlock: "0x" + strconv.FormatUint(uint64(from), 16),
		ToBlock:   "0x" + strconv.FormatUint(uint64(to), 16),
		Address:   []string{contractAddress},
		Topics:    [][]string{topicsInterested},
	}
	var data []ethrpc.Log
	data, err := rpcClient.EthGetLogs(filterParam)
	if err != nil {
		log.Error("FetchLogs, error while fetching logs", err)
		panic(err)
		return nil
	}
	return data
}

func assignDefaultWatcherConfig(config *watcher.ReceiptLogWatcherConfig) {
	if config.StepSizeForBigLag == 0 {
		config.StepSizeForBigLag = 50
	}
	if config.LagToHighestBlock == 0 {
		config.LagToHighestBlock = 10
	}
}

func LoadConfig(file string, path string, cfg *IndexerConfig) (err error) {
	viper.SetConfigName(file)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Errorf("error reading config file %s", err)
		return
	}
	log.Println("Config loaded successfully...")
	log.Println("Getting environment variables...")
	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, getEnvOrPanic(strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")))
		}
	}
	err = viper.Unmarshal(cfg)
	if err != nil {
		log.Errorf("unable to decode config to struct, %v", err)
		return
	}
	return
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(res) == 0 {
		panic("Mandatory env variable not found:" + env)
	}
	return res
}

func getChainID(ethRPCClient *EthBlockChainRPCWithRetry) (chainID string, err error) {
	var hexChainID string
	err = ethRPCClient.makeEthCallAndSave("eth_chainId", &hexChainID)
	if err != nil {
		return
	}
	chainIDInt, err := strconv.ParseUint(hexChainID, 0, 0)
	if err != nil {
		return
	}
	chainID = strconv.FormatUint(chainIDInt, 10)
	log.Info("Chain ID: ", chainID)
	return
}
