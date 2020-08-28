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
	totalMintingSupply := k.MintingTokenSupply(ctx)

	totalSupply := totalMintedSupply.Sub(totalMintingSupply)

	// mint coins, update supply
	mintedCoin := minter.BlockProvision(params, totalSupply)
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

	ctx.Logger().Info(fmt.Sprintf("mint:totalSupply:%s, totalMintingSupply: %s, DistrTokenSupply:%s, DayProvisions:%s, PeriodProvisions:%s, mintedCoin: %s",
		totalSupply.String(),
		totalMintingSupply.String(),
		k.DistrTokenSupply(ctx).String(),
		minter.CurrentDayProvisions(totalSupply).String(),
		minter.NextPeriodProvisions(totalSupply).String(),
		mintedCoin.Amount.String(),
	))

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(types.AttributeKeyDayProvisions, minter.CurrentDayProvisions(totalSupply).String()),
			sdk.NewAttribute(types.AttributeKeyPeriodProvisions, minter.NextPeriodProvisions(totalSupply).String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()),
		),
	)
}
