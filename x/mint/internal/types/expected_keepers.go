package types // noalias

import (
	sdk "hschain/types"
	"hschain/x/supply/exported"
)

// SupplyKeeper defines the expected supply keeper
type SupplyKeeper interface {
	GetModuleAddress(name string) sdk.AccAddress

	GetSupply(sdk.Context) exported.SupplyI

	SetModuleAccount(sdk.Context, exported.ModuleAccountI)

	GetModuleAccount(sdk.Context, string) exported.ModuleAccountI

	SendCoinsFromAccountToModule(sdk.Context, sdk.AccAddress, string, sdk.Coins) sdk.Error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) sdk.Error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) sdk.Error
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) sdk.Error
}
