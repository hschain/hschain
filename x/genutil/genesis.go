package genutil

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/hschain/hschain/codec"
	sdk "github.com/hschain/hschain/types"
	"github.com/hschain/hschain/x/genutil/types"
)

// InitGenesis - initialize accounts and deliver genesis transactions
func InitGenesis(
	ctx sdk.Context, cdc *codec.Codec, stakingKeeper types.StakingKeeper,
	deliverTx deliverTxfn, genesisState GenesisState,
) []abci.ValidatorUpdate {

	var validators []abci.ValidatorUpdate
	if len(genesisState.GenTxs) > 0 {
		validators = DeliverGenTxs(ctx, cdc, genesisState.GenTxs, stakingKeeper, deliverTx)
	}

	return validators
}
