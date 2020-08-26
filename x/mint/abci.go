package mint

import (
	sdk "hschain/types"
	"hschain/x/mint/internal/types"
	"log"
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k Keeper) {
	// fetch stored minter & params
	minter := k.GetMinter(ctx)
	params := k.GetParams(ctx)

	totalStakingSupply := k.StakingTokenSupply(ctx)

	// mint coins, update supply
	mintedCoin := minter.BlockProvision(params, totalStakingSupply)
	mintedCoins := sdk.NewCoins(mintedCoin)

	err := k.MintCoins(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	// send the minted coins to the fee collector account
	err = k.AddCollectedFees(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	log.Printf("mint:totalStakingSupply:%s, undistSupply: %s, DayProvisions:%s, PeriodProvisions:%s, mintedCoin: %s",
		totalStakingSupply.String(),
		k.UndistStakingTokenSupply(ctx).String(),
		minter.CurrentDayProvisions(totalStakingSupply).String(),
		minter.NextPeriodProvisions(totalStakingSupply).String(),
		mintedCoin.Amount.String(),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(types.AttributeKeyDayProvisions, minter.CurrentDayProvisions(totalStakingSupply).String()),
			sdk.NewAttribute(types.AttributeKeyPeriodProvisions, minter.NextPeriodProvisions(totalStakingSupply).String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()),
		),
	)
}
