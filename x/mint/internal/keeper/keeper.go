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
	sk                   types.StakingKeeper
	supplyKeeper         types.SupplyKeeper
	coinsCollectorName   string
	coinsDistributorName string
	coinsBurnerName      string
	coinsDestoryerName   string
}

// NewKeeper creates a new mint Keeper instance
func NewKeeper(
	cdc *codec.Codec, key sdk.StoreKey, paramSpace params.Subspace,
	sk types.StakingKeeper, supplyKeeper types.SupplyKeeper, coinsCollectorName, coinsDistributorName, coinsBurnerName, coinsDestoryerName string) Keeper {

	// ensure mint module account is set
	if addr := supplyKeeper.GetModuleAddress(types.ModuleName); addr == nil {
		panic("the mint module account has not been set")
	}

	return Keeper{
		cdc:                  cdc,
		storeKey:             key,
		paramSpace:           paramSpace.WithKeyTable(types.ParamKeyTable()),
		sk:                   sk,
		supplyKeeper:         supplyKeeper,
		coinsCollectorName:   coinsCollectorName,
		coinsDistributorName: coinsDistributorName,
		coinsBurnerName:      coinsBurnerName,
		coinsDestoryerName:   coinsDestoryerName,
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

func (k Keeper) GetBonus(ctx sdk.Context, height string) (coin sdk.Coin) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get([]byte(fmt.Sprintf("%s_bns", height)))
	if b == nil {
		params := k.GetParams(ctx)
		return sdk.NewInt64Coin(params.MintDenom, 0)
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &coin)
	return
}

func (k Keeper) SetBonus(ctx sdk.Context, height int64, coin sdk.Coin) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(coin)
	store.Set([]byte(fmt.Sprintf("%d_bns", height)), b)
}

func (k Keeper) GetConversionRates(ctx sdk.Context, denom string) (rates sdk.Coin) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get([]byte(fmt.Sprintf("%s_rates", denom)))
	if b == nil {
		params := k.GetParams(ctx)
		return sdk.NewInt64Coin(params.MintDenom, 70000000)
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &rates)
	return
}

func (k Keeper) SetConversionRates(ctx sdk.Context, denom string, rates sdk.Coin) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(rates)
	store.Set([]byte(fmt.Sprintf("%s_rates", denom)), b)
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
//GetDenomSupply get supply of spec denom
func (k Keeper) GetDenomSupply(ctx sdk.Context, denom string) sdk.Int {
	return k.supplyKeeper.GetSupply(ctx).GetTotal().AmountOf(denom)
}

// MintedTokenSupply implements an alias call to the underlying supply keeper's
// MintedTokenSupply to be used in BeginBlocker.
func (k Keeper) MintedTokenSupply(ctx sdk.Context) sdk.Int {
	return k.GetDenomSupply(ctx, k.GetParams(ctx).MintDenom)
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

//已燃烧
func (k Keeper) BurnTokenSupply(ctx sdk.Context) sdk.Coins {
	coinsBurnerAcc := k.supplyKeeper.GetModuleAccount(ctx, k.coinsBurnerName)
	return coinsBurnerAcc.GetCoins()
}

//已销毁
func (k Keeper) DestoryTokenSupply(ctx sdk.Context) sdk.Coins {
	coinsdestoriserAcc := k.supplyKeeper.GetModuleAccount(ctx, k.coinsDestoryerName)
	return coinsdestoriserAcc.GetCoins()
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

func (k Keeper) BurnCoins(ctx sdk.Context, sender sdk.AccAddress, amt sdk.Coins) sdk.Error {

	params := k.GetParams(ctx)
	if amt.AmountOf(params.MintDenom).IsZero() {
		errMsg := fmt.Sprintf("sender must hold %s", params.MintDenom)
		return sdk.ErrUnknownRequest(errMsg)
	}

	if err := k.supplyKeeper.SendCoinsFromAccountToModule(ctx, sender, k.coinsBurnerName, amt); err != nil {
		return err
	}

	rates := k.GetConversionRates(ctx, params.MintDenom)

	powerAmount := amt.AmountOf(params.MintDenom)
	powerAmount = powerAmount.Quo(rates.Amount).MulRaw(1000000)

	power := sdk.NewCoin(types.PowerDenom, powerAmount)

	mintedCoins := sdk.NewCoins(power)
	if err := k.MintCoins(ctx, mintedCoins); err != nil {
		return err
	}

	return k.supplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, mintedCoins)
}

func (k Keeper) DestoryCoins(ctx sdk.Context, fromAddr sdk.AccAddress, amt sdk.Coins) sdk.Error {

	return k.supplyKeeper.SendCoinsFromAccountToModule(ctx, fromAddr, k.coinsDestoryerName, amt)
}

// AddMintingCoins implements an alias call to the underlying supply keeper's
// AddMintingCoins to be used in BeginBlocker.
func (k Keeper) AddMintingCoins(ctx sdk.Context, amt sdk.Coins) sdk.Error {
	return k.supplyKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.coinsCollectorName, amt)
}

//IssueCoins
func (k Keeper) IssueCoins(ctx sdk.Context, toAddress sdk.AccAddress, amt sdk.Coins) sdk.Error {

	for _, coin := range amt {
		if coin.Denom != types.PowerDenom && len(coin.Denom) <= 2 {
			errMsg := fmt.Sprintf("Denominations can be 3 ~ 16 characters long")
			return sdk.ErrUnknownRequest(errMsg)
		}

		if coin.Denom[0:1] != types.DenomBegin {
			errMsg := fmt.Sprintf("Denominations begin must be u")
			return sdk.ErrUnknownRequest(errMsg)

		}
	}

	if err := k.supplyKeeper.MintCoins(ctx, types.ModuleName, amt); err != nil {
		return err
	}

	return k.supplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, toAddress, amt)
}

func (k Keeper) GetBalance(ctx sdk.Context, acc sdk.AccAddress) sdk.Coins {
	return k.supplyKeeper.GetBalance(ctx, acc)
}

func (k Keeper) BondDenom(ctx sdk.Context) string {
	return k.sk.BondDenom(ctx)
}
