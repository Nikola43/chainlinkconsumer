package datafeedsiface

import "math/big"

type DataFeedsInterface interface {
	LatestRoundData() (error, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int)
}
