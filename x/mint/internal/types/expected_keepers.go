package types // noalias

import (
	sdk "github.com/hschain/hschain/types"
	"github.com/hschain/hschain/x/supply/exported"
)

// SupplyKeeper defines the expected supply keeper
type SupplyKeeper interface {
	GetModuleAddress(name string) sdk.AccAddress

	GetSupply(sdk.Context) exported.SupplyI

	GetBalance(sdk.Context, sdk.AccAddress) sdk.Coins

	SetModuleAccount(sdk.Context, exported.ModuleAccountI)

	GetModuleAccount(sdk.Context, string) exported.ModuleAccountI

	SendCoinsFromAccountToModule(sdk.Context, sdk.AccAddress, string, sdk.Coins) sdk.Error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) sdk.Error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) sdk.Error
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) sdk.Error
	BurnCoins(ctx sdk.Context, name string, amt sdk.Coins) sdk.Error
}

// StakingKeeper expected staking keeper (Validator and Delegator sets)
type StakingKeeper interface {
	BondDenom(sdk.Context) string
}
