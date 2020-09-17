package types

import (
	"hschain/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgBurn{}, "cosmos-sdk/MsgBurn", nil)
	cdc.RegisterConcrete(MsgConversionRate{}, "cosmos-sdk/MsgConversionRate", nil)
	cdc.RegisterConcrete(MsgDestory{}, "cosmos-sdk/MsgDestory", nil)
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
