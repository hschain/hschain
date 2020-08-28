package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"hschain/codec"
	sdk "hschain/types"
	"hschain/x/mint/internal/types"
	"hschain/x/params"
)

// Keeper of the mint store
type Keeper struct {
	cdc                  *codec.Codec
	storeKey             sdk.StoreKey
	paramSpace           params.Subspace
	supplyKeeper         types.SupplyKeeper
	coinsCollectorName   string
	coinsDistributorName string
	coinsBurnerName      string
}

// NewKeeper creates a new mint Keeper instance
func NewKeeper(
	cdc *codec.Codec, key sdk.StoreKey, paramSpace params.Subspace,
	supplyKeeper types.SupplyKeeper, coinsCollectorName, coinsDistributorName, coinsBurnerName string) Keeper {

	// ensure mint module account is set
	if addr := supplyKeeper.GetModuleAddress(types.ModuleName); addr == nil {
		panic("the mint module account has not been set")
	}

	return Keeper{
		cdc:                  cdc,
		storeKey:             key,
		paramSpace:           paramSpace.WithKeyTable(types.ParamKeyTable()),
		supplyKeeper:         supplyKeeper,
		coinsCollectorName:   coinsCollectorName,
		coinsDistributorName: coinsDistributorName,
		coinsBurnerName:      coinsBurnerName,
	}
}

//______________________________________________________________________

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// get the minter
func (k Keeper) GetMinter(ctx sdk.Context) (minter types.Minter) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.MinterKey)
	if b == nil {
		panic("stored minter should not have been nil")
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &minter)
	return
}

// set the minter
func (k Keeper) SetMinter(ctx sdk.Context, minter types.Minter) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(minter)
	store.Set(types.MinterKey, b)
}

//______________________________________________________________________

// GetParams returns the total set of minting parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of minting parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

//______________________________________________________________________

// MintedTokenSupply implements an alias call to the underlying supply keeper's
// MintedTokenSupply to be used in BeginBlocker.
func (k Keeper) MintedTokenSupply(ctx sdk.Context) sdk.Int {
	return k.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(k.GetParams(ctx).MintDenom)
}

//已挖不可分配
func (k Keeper) MintingTokenSupply(ctx sdk.Context) sdk.Int {
	coinsCollectorAcc := k.supplyKeeper.GetModuleAccount(ctx, k.coinsCollectorName)
	return coinsCollectorAcc.GetCoins().AmountOf(k.GetParams(ctx).MintDenom)
}

//已挖等待分配
func (k Keeper) DistrTokenSupply(ctx sdk.Context) sdk.Int {
	coinsDistributorAcc := k.supplyKeeper.GetModuleAccount(ctx, k.coinsDistributorName)
	return coinsDistributorAcc.GetCoins().AmountOf(k.GetParams(ctx).MintDenom)
}

// MintCoins implements an alias call to the underlying supply keeper's
// MintCoins to be used in BeginBlocker.
func (k Keeper) MintCoins(ctx sdk.Context, newCoins sdk.Coins) sdk.Error {
	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}
	return k.supplyKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}

func (k Keeper) BurnCoins(ctx sdk.Context, fromAddr sdk.AccAddress, amt sdk.Coins) sdk.Error {
	return k.supplyKeeper.SendCoinsFromAccountToModule(ctx, fromAddr, k.coinsBurnerName, amt)
}

// AddMintingCoins implements an alias call to the underlying supply keeper's
// AddMintingCoins to be used in BeginBlocker.
func (k Keeper) AddMintingCoins(ctx sdk.Context, amt sdk.Coins) sdk.Error {
	return k.supplyKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.coinsCollectorName, amt)
}
