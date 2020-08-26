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
	return func(ctx sdk.Context, path []string, _ abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case types.QueryParameters:
			return queryParams(ctx, k)

		case types.QueryDayProvisions:
			return queryDayProvisions(ctx, k)

		case types.QueryPeriodProvisions:
			return queryPeriodProvisions(ctx, k)

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

func queryDayProvisions(ctx sdk.Context, k Keeper) ([]byte, sdk.Error) {
	minter := k.GetMinter(ctx)
	supply := k.StakingTokenSupply(ctx)

	res, err := codec.MarshalJSONIndent(k.cdc, minter.CurrentDayProvisions(supply))
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to marshal JSON", err.Error()))
	}

	return res, nil
}

func queryPeriodProvisions(ctx sdk.Context, k Keeper) ([]byte, sdk.Error) {
	minter := k.GetMinter(ctx)
	supply := k.StakingTokenSupply(ctx)

	res, err := codec.MarshalJSONIndent(k.cdc, minter.NextPeriodProvisions(supply))
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to marshal JSON", err.Error()))
	}

	return res, nil
}
