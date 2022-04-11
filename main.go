package main

import (
	"fmt"
	web3manager "github.com/nikola43/web3manager/datafeeds"
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

	dataFeedConsumer := web3manager.NewDataFeedConsumer(web3manager.BSC, web3manager.BNB_USD)

	err, roundId, answer, startedAt, updatedAt, answeredInRound := dataFeedConsumer.LatestRoundData()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("RoundId: ", roundId)
	fmt.Println("Answer: ", answer)
	fmt.Println("Started At:", startedAt)
	fmt.Println("Updated At: ", updatedAt)
	fmt.Println("Answered In Round: ", answeredInRound)

}
