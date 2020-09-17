package mint

import (
	"fmt"
<<<<<<< HEAD

=======
>>>>>>> df41a681ebe3047d8be9520b9858e17a9bf418c1
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

<<<<<<< HEAD
	BurnAmount := k.BurnTokenSupply(ctx)
=======
>>>>>>> df41a681ebe3047d8be9520b9858e17a9bf418c1
	// send the minted coins to the fee collector account
	err = k.AddMintingCoins(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	k.SetBonus(ctx, ctx.BlockHeight(), mintedCoin)

	totalMintingSupply := k.MintingTokenSupply(ctx)

<<<<<<< HEAD
	ctx.Logger().Info(fmt.Sprintf("mint:TotalSupply:%s, TotalMintingSupply: %s, DistrTokenSupply:%s, CurrentDayProvisions:%s, NextPeroidStartTime:%d, NextPeriodDayProvisions:%s, mintedCoin: %s, BurnAmount: %s",
=======
	ctx.Logger().Info(fmt.Sprintf("mint:TotalSupply:%s, TotalMintingSupply: %s, DistrTokenSupply:%s, CurrentDayProvisions:%s, NextPeroidStartTime:%d, NextPeriodDayProvisions:%s, mintedCoin: %s",
>>>>>>> df41a681ebe3047d8be9520b9858e17a9bf418c1
		totalMintedSupply.String(),
		totalMintingSupply.String(),
		k.DistrTokenSupply(ctx).String(),
		minter.CurrentDayProvisions(totalMintedSupply.Sub(totalMintingSupply)).String(),
		minter.NextPeroidStartTime(params, totalMintedSupply),
		minter.NextPeriodDayProvisions(totalMintedSupply).String(),
		mintedCoin.Amount.String(),
<<<<<<< HEAD
		BurnAmount.String(),
=======
>>>>>>> df41a681ebe3047d8be9520b9858e17a9bf418c1
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
