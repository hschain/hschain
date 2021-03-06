// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: hschain/x/mint/internal/keeper
// ALIASGEN: hschain/x/mint/internal/types
package mint

import (
	"github.com/hschain/hschain/x/mint/internal/keeper"
	"github.com/hschain/hschain/x/mint/internal/types"
)

const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	DefaultParamspace = types.DefaultParamspace
	StoreKey          = types.StoreKey
	QuerierRoute      = types.QuerierRoute
	QueryParameters   = types.QueryParameters
	QueryStatus       = types.QueryStatus
)

var (
	// functions aliases
	NewKeeper            = keeper.NewKeeper
	NewQuerier           = keeper.NewQuerier
	NewMinter            = types.NewMinter
	InitialMinter        = types.InitialMinter
	DefaultInitialMinter = types.DefaultInitialMinter
	ValidateMinter       = types.ValidateMinter
	ParamKeyTable        = types.ParamKeyTable
	NewParams            = types.NewParams
	DefaultParams        = types.DefaultParams
	ValidateParams       = types.ValidateParams
	DefaultMintPlans     = types.DefaultMintPlans

	// variable aliases
	ModuleCdc       = types.ModuleCdc
	RegisterCodec   = types.RegisterCodec
	MinterKey       = types.MinterKey
	KeyMintDenom    = types.KeyMintDenom
	KeyBlocksPerDay = types.KeyBlocksPerDay
)

type (
	Keeper            = keeper.Keeper
	Minter            = types.Minter
	Params            = types.Params
	MsgBurn           = types.MsgBurn
	MsgDestory        = types.MsgDestory
	MsgDestoryUser    = types.MsgDestoryUser
	MsgConversionRate = types.MsgConversionRate
)
