package types

// the one key to use for the keeper store
var MinterKey = []byte{0x00}

// nolint
const (
	// module name
	ModuleName = "mint"

<<<<<<< HEAD
	// Power Denom
	PowerDenom = "ut"

	// Denom Begin
	DenomBegin = "u"
=======
>>>>>>> df41a681ebe3047d8be9520b9858e17a9bf418c1
	// default paramspace for params keeper
	DefaultParamspace = ModuleName

	// StoreKey is the default store key for mint
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the minting store.
	QuerierRoute = StoreKey

	// Query endpoints supported by the minting querier
	QueryParameters = "parameters"
	QueryStatus     = "status"
	QueryBonus      = "bonus"
)
