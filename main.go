package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	aggregatorV3Interface "github.com/nikola43/web3manager/contracts/AggregatorV3Interface"
	"log"
)

/*
   // NETWORKS
   bsc: {
     url: 'https://bsc-dataseed.binance.org',
     chainId: 56,
   },
   bsctestnet: {
     url: 'https://data-seed-prebsc-2-s3.binance.org:8545',
     chainId: 97,
   },
*/

func main() {
	client, err := ethclient.Dial("https://data-seed-prebsc-2-s3.binance.org:8545")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x2514895c72f50D8bd4B4F9b1110F0D6bD2c97526") // bnb usdt
	instance, err := aggregatorV3Interface.NewAggregatorV3Interface(address, client)
	if err != nil {
		log.Fatal(err)
	}

	c := &bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}
	latestRoundData, err := instance.LatestRoundData(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(latestRoundData)
}
