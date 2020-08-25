//nolint
package app

import (
	"io"

	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"hschain/baseapp"
	sdk "hschain/types"
	"hschain/x/staking"
)

var (
	genesisFile        string
	paramsFile         string
	exportParamsPath   string
	exportParamsHeight int
	exportStatePath    string
	exportStatsPath    string
	seed               int64
	initialBlockHeight int
	numBlocks          int
	blockSize          int
	enabled            bool
	verbose            bool
	lean               bool
	commit             bool
	period             int
	onOperation        bool // TODO Remove in favor of binary search for invariant violation
	allInvariants      bool
	genesisTime        int64
)

// DONTCOVER

// NewAppUNSAFE is used for debugging purposes only.
//
// NOTE: to not use this function with non-test code
func NewAppUNSAFE(logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, baseAppOptions ...func(*baseapp.BaseApp),
) (happ *App, keyMain, keyStaking *sdk.KVStoreKey, stakingKeeper staking.Keeper) {

	happ = NewApp(logger, db, traceStore, loadLatest, invCheckPeriod, baseAppOptions...)
	return happ, happ.keys[baseapp.MainStoreKey], happ.keys[staking.StoreKey], happ.stakingKeeper
}
