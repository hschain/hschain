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

	totalMintingSupply := k.MintingTokenSupply(ctx)

	// mint coins, update supply
	mintedCoin := minter.BlockProvision(params, totalMintingSupply)
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

	log.Printf("mint:totalMintingSupply:%s, undistSupply: %s, DayProvisions:%s, PeriodProvisions:%s, mintedCoin: %s",
		totalMintingSupply.String(),
		k.UndistMintedTokenSupply(ctx).String(),
		minter.CurrentDayProvisions(totalMintingSupply).String(),
		minter.NextPeriodProvisions(totalMintingSupply).String(),
		mintedCoin.Amount.String(),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(types.AttributeKeyDayProvisions, minter.CurrentDayProvisions(totalMintingSupply).String()),
			sdk.NewAttribute(types.AttributeKeyPeriodProvisions, minter.NextPeriodProvisions(totalMintingSupply).String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()),
		),
	)
}
