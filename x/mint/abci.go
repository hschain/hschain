package mint

import (
	"fmt"
	sdk "hschain/types"
	"hschain/x/mint/internal/types"
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k Keeper) {
	// fetch stored minter & params
	minter := k.GetMinter(ctx)
	params := k.GetParams(ctx)

	totalMintedSupply := k.MintedTokenSupply(ctx)

	// mint coins, update supply
	mintedCoin := minter.BlockProvision(params, totalMintedSupply)

	if mintedCoin.IsZero() {
		ctx.Logger().Info("no coins mint")
		return
	}

	mintedCoins := sdk.NewCoins(mintedCoin)

	err := k.MintCoins(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	// send the minted coins to the fee collector account
	err = k.AddMintingCoins(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	totalMintingSupply := k.MintingTokenSupply(ctx)

	ctx.Logger().Info(fmt.Sprintf("mint:TotalSupply:%s, TotalMintingSupply: %s, DistrTokenSupply:%s, CurrentDayProvisions:%s, NextPeroidStartTime:%d, NextPeriodDayProvisions:%s, mintedCoin: %s",
		totalMintedSupply.String(),
		totalMintingSupply.String(),
		k.DistrTokenSupply(ctx).String(),
		minter.CurrentDayProvisions(totalMintedSupply.Sub(totalMintingSupply)).String(),
		minter.NextPeroidStartTime(params, totalMintedSupply),
		minter.NextPeriodDayProvisions(totalMintedSupply).String(),
		mintedCoin.Amount.String(),
	))

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(types.AttributeKeyCurrentDayProvisions, minter.CurrentDayProvisions(totalMintedSupply.Sub(totalMintingSupply)).String()),
			sdk.NewAttribute(types.AttributeKeyNextPeriodDayProvisions, minter.NextPeriodDayProvisions(totalMintedSupply).String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()),
		),
	)
}
