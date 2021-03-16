package mint

import (
	"fmt"
	"time"

	sdk "github.com/hschain/hschain/types"
	"github.com/hschain/hschain/x/mint/internal/types"
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k Keeper) {
	// fetch stored minter & params
	minter := k.GetMinter(ctx)
	params := k.GetParams(ctx)

	//height := 2186143
	height := 2331716
	if ctx.BlockHeight() < (int64)(height) {
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

		k.SetBonus(ctx, ctx.BlockHeight(), mintedCoin)

		totalMintingSupply := k.MintingTokenSupply(ctx)

		BurnAmount := k.BurnTokenSupply(ctx)

		now := ctx.BlockTime().Add(8 * time.Hour) //change to shanghai timezone

		LastDistributeTime := k.GetLastDistributeTime(ctx)

		if now.Year() != LastDistributeTime.Year() || now.YearDay() != LastDistributeTime.YearDay() && now.Hour() >= 10 {

			amt := k.DistrTokenSupply(ctx)
			if !amt.IsZero() {
				coin := sdk.NewCoins(sdk.NewCoin(params.MintDenom, amt))
				if err := k.MintingCoinsIssueAddress(ctx, coin); err == nil {
					k.SetLastDistributeTime(ctx, now)
				}
			}
		}

		ctx.Logger().Info(fmt.Sprintf("mint:TotalSupply:%s, TotalMintingSupply: %s, DistrTokenSupply:%s, CurrentDayProvisions:%s, NextPeroidStartTime:%d, NextPeriodDayProvisions:%s, mintedCoin: %s, BurnAmount: %s",
			totalMintedSupply.String(),
			totalMintingSupply.String(),
			k.DistrTokenSupply(ctx).String(),
			minter.CurrentDayProvisions(totalMintedSupply.Sub(totalMintingSupply)).String(),
			minter.NextPeroidStartTime(params, totalMintedSupply),
			minter.NextPeriodDayProvisions(totalMintedSupply).String(),
			mintedCoin.Amount.String(),
			BurnAmount.String(),
		))

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeMint,
				sdk.NewAttribute(types.AttributeKeyCurrentDayProvisions, minter.CurrentDayProvisions(totalMintedSupply.Sub(totalMintingSupply)).String()),
				sdk.NewAttribute(types.AttributeKeyNextPeriodDayProvisions, minter.NextPeriodDayProvisions(totalMintedSupply).String()),
				sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()),
			),
		)
	} else {
		totalMintedSupply := k.MintedTokenSupply(ctx)
		BurnAmount := k.BurnTokenSupply(ctx)
		ctx.Logger().Info(fmt.Sprintf("mint:TotalSupply:%s, TotalMintingSupply: %s, DistrTokenSupply:%s, CurrentDayProvisions:%s, NextPeroidStartTime:%d, NextPeriodDayProvisions:%s, mintedCoin: %s, BurnAmount: %s",
			totalMintedSupply.String(),
			"0",
			k.DistrTokenSupply(ctx).String(),
			"0",
			0,
			"0",
			"0",
			BurnAmount.String(),
		))

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeMint,
				sdk.NewAttribute(types.AttributeKeyCurrentDayProvisions, "0"),
				sdk.NewAttribute(types.AttributeKeyNextPeriodDayProvisions, minter.NextPeriodDayProvisions(totalMintedSupply).String()),
				sdk.NewAttribute(sdk.AttributeKeyAmount, "0"),
			),
		)
	}
}
