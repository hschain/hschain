package types

import (
	sdk "github.com/hschain/hschain/types"
	"github.com/hschain/hschain/x/supply/exported"
)

// SupplyKeeper defines the expected supply Keeper (noalias)
type SupplyKeeper interface {
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) sdk.Error
	GetModuleAccount(ctx sdk.Context, moduleName string) exported.ModuleAccountI
	GetModuleAddress(moduleName string) sdk.AccAddress
}
