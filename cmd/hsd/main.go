package main

import (
	"encoding/json"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"hschain/app"

	"hschain/baseapp"
	"hschain/client"
	"hschain/server"
	"hschain/store"
	sdk "hschain/types"
	"hschain/x/genaccounts"
	genaccscli "hschain/x/genaccounts/client/cli"
	genutilcli "hschain/x/genutil/client/cli"
	"hschain/x/staking"
)

// hsd custom flags
const flagInvCheckPeriod = "inv-check-period"

var invCheckPeriod uint

func main() {
	cdc := app.MakeCodec()

	/*
		config := sdk.GetConfig()
		app.SetBech32AddressPrefixes(config)
		app.SetBip44CoinType(config)
		config.Seal()
	*/

	ctx := server.NewDefaultContext()
	cobra.EnableCommandSorting = false
	rootCmd := &cobra.Command{
		Use:               "hsd",
		Short:             "hschain Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	rootCmd.AddCommand(genutilcli.InitCmd(ctx, cdc, app.ModuleBasics, app.DefaultNodeHome))
	rootCmd.AddCommand(genutilcli.CollectGenTxsCmd(ctx, cdc, genaccounts.AppModuleBasic{}, app.DefaultNodeHome))
	rootCmd.AddCommand(genutilcli.GenTxCmd(ctx, cdc, app.ModuleBasics, staking.AppModuleBasic{},
		genaccounts.AppModuleBasic{}, app.DefaultNodeHome, app.DefaultCLIHome))
	rootCmd.AddCommand(genutilcli.ValidateGenesisCmd(ctx, cdc, app.ModuleBasics))
	rootCmd.AddCommand(genaccscli.AddGenesisAccountCmd(ctx, cdc, app.DefaultNodeHome, app.DefaultCLIHome))
	rootCmd.AddCommand(client.NewCompletionCmd(rootCmd, true))
	rootCmd.AddCommand(replayCmd())

	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, "HSC", app.DefaultNodeHome)
	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod,
		0, "Assert registered invariants every N blocks")

	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer, prefix string) abci.Application {
	logger.Info("run with subchain " + prefix)
	config := sdk.GetConfig()
	app.SetBech32Prefix(prefix)
	app.SetBech32AddressPrefixes(config)
	app.SetBip44CoinType(config)
	config.Seal()

	return app.NewApp(
		logger, db, traceStore, true, invCheckPeriod,
		baseapp.SetPruning(store.NewPruningOptionsFromString(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
		baseapp.SetHaltHeight(uint64(viper.GetInt(server.FlagHaltHeight))),
	)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {

	if height != -1 {
		hApp := app.NewApp(logger, db, traceStore, false, uint(1))
		err := hApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}

		return hApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}
	hApp := app.NewApp(logger, db, traceStore, true, uint(1))
	return hApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
