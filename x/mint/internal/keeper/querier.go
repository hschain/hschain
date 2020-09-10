package keeper

import (
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"

	"hschain/codec"
	sdk "hschain/types"
	"hschain/x/mint/internal/types"
)

// NewQuerier returns a minting Querier handler.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, query abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case types.QueryParameters:
			return queryParams(ctx, k)

		case types.QueryStatus:
			return queryStatus(ctx, k)

		case types.QueryBonus:
			return queryBonus(ctx, k, string(query.Data))

		default:
			return nil, sdk.ErrUnknownRequest(fmt.Sprintf("unknown minting query endpoint: %s", path[0]))
		}
	}
}

func queryParams(ctx sdk.Context, k Keeper) ([]byte, sdk.Error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(k.cdc, params)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to marshal JSON", err.Error()))
	}

	return res, nil
}

func queryStatus(ctx sdk.Context, k Keeper) ([]byte, sdk.Error) {
	minter := k.GetMinter(ctx)

	params := k.GetParams(ctx)

	minter.Status.TotalMintedSupply = k.MintedTokenSupply(ctx)
	minter.Status.TotalMintingSupply = k.MintingTokenSupply(ctx)
	minter.Status.TotalDistrSupply = k.DistrTokenSupply(ctx)
	minter.Status.StatBurnCoins = k.BurnTokenSupply(ctx)
	minter.Status.CurrentDayProvisions = minter.CurrentDayProvisions(minter.Status.TotalMintedSupply.Sub(minter.Status.TotalMintingSupply))
	minter.Status.NextPeriodDayProvisions = minter.NextPeriodDayProvisions(minter.Status.TotalMintedSupply)
	minter.Status.NextPeroidStartTime = minter.NextPeroidStartTime(params, minter.Status.TotalMintedSupply)
	minter.Status.BlockProvision = minter.BlockProvision(params, minter.Status.TotalMintedSupply)

	res, err := codec.MarshalJSONIndent(k.cdc, minter)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to marshal JSON", err.Error()))
	}

	return res, nil
}

func queryBonus(ctx sdk.Context, k Keeper, height string) ([]byte, sdk.Error) {

	//log.Printf("query bonus at height %s", height)
	coin := k.GetBonus(ctx, height)

	res, err := codec.MarshalJSONIndent(k.cdc, coin)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to marshal JSON", err.Error()))
	}

	return res, nil
}
