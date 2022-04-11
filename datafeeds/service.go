package web3manager

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	aggregatorV3Interface "github.com/nikola43/web3manager/contracts/AggregatorV3Interface"
	"log"
	"math/big"
	_ "net/url"
	//web3utils "github.com/nikola43/goweb3manager/goweb3manager/util"
)

type Network int
type Pair string

const (
	BSC       Network = 0
	AVALANCHE Network = 1

	BNB_USD  Pair = "BNB/USD"
	AVAX_USD Pair = "AVAX/USD"
)

type DataFeeds struct {
	httpClient         *ethclient.Client
	aggregatorInstance *aggregatorV3Interface.AggregatorV3Interface
}

func NewDataFeedConsumer(network Network, pair Pair) *DataFeeds {
	net := "https://data-seed-prebsc-2-s3.binance.org:8545"
	switch network {
	case BSC:
		net = "https://data-seed-prebsc-2-s3.binance.org:8545"
		break
	case AVALANCHE:
		net = "https://data-seed-prebsc-2-s3.binance.org:8545"
		break
	}

	client, err := ethclient.Dial(net)
	if err != nil {
		log.Fatal(err)
	}

	pairAddress := "0x2514895c72f50D8bd4B4F9b1110F0D6bD2c97526"
	switch pair {
	case AVAX_USD:
		pairAddress = "0x2514895c72f50D8bd4B4F9b1110F0D6bD2c97526"
		break
	case BNB_USD:
		pairAddress = "0x2514895c72f50D8bd4B4F9b1110F0D6bD2c97526"
		break
	}

	address := common.HexToAddress(pairAddress) // bnb usdt
	instance, err := aggregatorV3Interface.NewAggregatorV3Interface(address, client)
	if err != nil {
		log.Fatal(err)
	}
	return &DataFeeds{
		httpClient:         client,
		aggregatorInstance: instance,
	}
}

func (df *DataFeeds) LatestRoundData() (error, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int) {
	latestRoundData, err := df.aggregatorInstance.LatestRoundData(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     context.Background(),
	})
	if err != nil {
		return err, nil, nil, nil, nil, nil
	}
	return nil, latestRoundData.RoundId, latestRoundData.Answer, latestRoundData.StartedAt, latestRoundData.UpdatedAt, latestRoundData.AnsweredInRound
}
