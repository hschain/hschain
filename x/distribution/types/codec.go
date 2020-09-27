package types

import (
	"github.com/hschain/hschain/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgWithdrawDelegatorReward{}, "cosmos-sdk/MsgWithdrawDelegationReward", nil)
	cdc.RegisterConcrete(MsgWithdrawValidatorCommission{}, "cosmos-sdk/MsgWithdrawValidatorCommission", nil)
	cdc.RegisterConcrete(MsgSetWithdrawAddress{}, "cosmos-sdk/MsgModifyWithdrawAddress", nil)
	cdc.RegisterConcrete(CommunityPoolSpendProposal{}, "cosmos-sdk/CommunityPoolSpendProposal", nil)

	cdc.RegisterConcrete(MsgSetDistrAddress{}, "cosmos-sdk/MsgSetDistrAddress", nil)
	cdc.RegisterConcrete(MsgDistrCoins{}, "cosmos-sdk/MsgDistrCoins", nil)

}

// generic sealed codec to be used throughout module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
