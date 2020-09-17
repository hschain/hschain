package types

import (
	"hschain/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgBurn{}, "cosmos-sdk/MsgBurn", nil)
<<<<<<< HEAD
	cdc.RegisterConcrete(MsgConversionRate{}, "cosmos-sdk/MsgConversionRate", nil)
	cdc.RegisterConcrete(MsgDestory{}, "cosmos-sdk/MsgDestory", nil)
=======
>>>>>>> df41a681ebe3047d8be9520b9858e17a9bf418c1
	cdc.RegisterConcrete(MsgIssue{}, "cosmos-sdk/MsgIssue", nil)
}

// module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
