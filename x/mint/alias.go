// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: hschain/x/mint/internal/keeper
// ALIASGEN: hschain/x/mint/internal/types
package mint

import (
	"hschain/x/mint/internal/keeper"
	"hschain/x/mint/internal/types"
)

const (
	ModuleName            = types.ModuleName
	DefaultParamspace     = types.DefaultParamspace
	StoreKey              = types.StoreKey
	QuerierRoute          = types.QuerierRoute
	QueryParameters       = types.QueryParameters
	QueryInflation        = types.QueryInflation
	QueryAnnualProvisions = types.QueryAnnualProvisions
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

	// variable aliases
	ModuleCdc              = types.ModuleCdc
	MinterKey              = types.MinterKey
	KeyMintDenom           = types.KeyMintDenom
	KeyInflationRateChange = types.KeyInflationRateChange
	KeyInflationMax        = types.KeyInflationMax
	KeyInflationMin        = types.KeyInflationMin
	KeyGoalBonded          = types.KeyGoalBonded
	KeyBlocksPerYear       = types.KeyBlocksPerYear
)

type (
	Keeper = keeper.Keeper
	Minter = types.Minter
	Params = types.Params

	MintPlan = types.MintPlan
)
