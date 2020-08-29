package distribution

import (
	"time"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "hschain/types"
	"hschain/x/distribution/keeper"
	"log"
)

// set the proposer for determining distribution during endblock
// and distribute rewards for the previous block
func BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock, k keeper.Keeper) {
	latestBlockTime, err := k.GetLatestBlockTime(ctx)
	now := ctx.BlockTime().Add(8 * time.Hour) //change to shanghai timezone
	if err != nil {
		log.Printf("err: %s", err)
		k.SetLatestBlockTime(ctx, now)
		return
	}

	if now.Year() != latestBlockTime.Year() || now.YearDay() != latestBlockTime.YearDay() {
		log.Printf("latest block time %s, current block time %s", latestBlockTime.String(), now.String())
		log.Printf("transfer coins from feeCollector to feeDistributor")
		if ctx.BlockHeight() > 1 {
			k.DistributeCoins(ctx)
		}
	}

	k.SetLatestBlockTime(ctx, now)
}
