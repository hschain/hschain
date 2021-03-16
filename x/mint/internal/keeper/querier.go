package keeper

import (
	"encoding/json"
	"fmt"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/hschain/hschain/codec"
	sdk "github.com/hschain/hschain/types"
	"github.com/hschain/hschain/x/mint/internal/types"
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

		case types.QueryPermissions:
			return queryPermissions(ctx, k, query.Data)
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

	//height := 2186143
	height := 2331716
	if ctx.BlockHeight() < (int64)(height) {
		minter.Status.TotalMintedSupply = k.MintedTokenSupply(ctx)
		minter.Status.TotalMintingSupply = k.MintingTokenSupply(ctx)
		minter.Status.TotalDistrSupply = k.DistrTokenSupply(ctx)

		minter.Status.TotalCirculationSupply = k.MintedTokenSupply(ctx).Sub(k.MintingTokenSupply(ctx))
		minter.Status.TotalCirculationSupply = minter.Status.TotalCirculationSupply.Sub(k.BurnTokenSupply(ctx).AmountOf(params.MintDenom))
		minter.Status.TotalCirculationSupply = minter.Status.TotalCirculationSupply.Sub(k.DestoryTokenSupply(ctx).AmountOf(params.MintDenom))

		minter.Status.CurrentDayProvisions = minter.CurrentDayProvisions(minter.Status.TotalMintedSupply.Sub(minter.Status.TotalMintingSupply))
		minter.Status.NextPeriodDayProvisions = minter.NextPeriodDayProvisions(minter.Status.TotalMintedSupply)
		minter.Status.NextPeroidStartTime = minter.NextPeroidStartTime(params, minter.Status.TotalMintedSupply)
		minter.Status.BlockProvision = minter.BlockProvision(params, minter.Status.TotalMintedSupply)
		minter.Status.BurnAmount = k.BurnTokenSupply(ctx)
		minter.Status.DestoryAmount = k.DestoryTokenSupply(ctx)
		minter.Status.ConversionRates = sdk.NewCoins(k.GetConversionRates(ctx, params.MintDenom))
	} else {
		minter.Status.TotalMintedSupply = k.MintedTokenSupply(ctx)
		minter.Status.TotalMintingSupply = sdk.NewInt(0)
		minter.Status.TotalCirculationSupply = sdk.NewInt(0)
		minter.Status.CurrentDayProvisions = sdk.ZeroDec()
		minter.Status.NextPeriodDayProvisions = sdk.ZeroDec()
		minter.Status.NextPeroidStartTime = 0
		minter.Status.BlockProvision = sdk.NewInt64Coin(params.MintDenom, 0)
		minter.Status.BurnAmount = k.BurnTokenSupply(ctx)
		minter.Status.DestoryAmount = k.DestoryTokenSupply(ctx)
		minter.Status.ConversionRates = sdk.NewCoins(k.GetConversionRates(ctx, params.MintDenom))

	}
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

func queryPermissions(ctx sdk.Context, k Keeper, key []byte) ([]byte, sdk.Error) {

	//log.Printf("query bonus at height %s", height)
	var addressPermissions types.MsgAddressPermissions
	if err := json.Unmarshal(key, &addressPermissions); err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to marshal JSON", err.Error()))
	}

	status := k.GetPermissions(ctx, addressPermissions.Address, addressPermissions.Command)
	if status != 1 {
		addressPermissions.Status = status
	}

	res, err := codec.MarshalJSONIndent(k.cdc, addressPermissions)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to marshal JSON", err.Error()))
	}

	return res, nil
}

func querySysAddress(ctx sdk.Context, k Keeper, cmd string) ([]byte, sdk.Error) {

	//log.Printf("query bonus at height %s", height)
	address := k.GetSysAddress(ctx, cmd)
	return address, nil
}
