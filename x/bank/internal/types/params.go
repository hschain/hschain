package types

import (
	"hschain/x/params"
)

const (
	// DefaultParamspace for params keeper
	DefaultParamspace = ModuleName
	// DefaultSendEnabled enabled
	DefaultSendEnabled = true
)

// ParamStoreKeySendEnabled is store's key for SendEnabled
var ParamStoreKeySendEnabled = []byte("sendenabled")

// ParamKeyTable type declaration for parameters
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		ParamStoreKeySendEnabled, false,
	)
}
