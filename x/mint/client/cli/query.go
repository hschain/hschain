package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"hschain/client"
	"hschain/client/context"
	"hschain/codec"
	sdk "hschain/types"
	"hschain/x/mint/internal/types"
)

// GetQueryCmd returns the cli query commands for the minting module.
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	mintingQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the minting module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	mintingQueryCmd.AddCommand(
		client.GetCommands(
			GetCmdQueryParams(cdc),
			GetCmdQueryStatus(cdc),
			GetCmdQueryBonus(cdc),
		)...,
	)

	return mintingQueryCmd
}

// GetCmdQueryParams implements a command to return the current minting
// parameters.
func GetCmdQueryParams(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "Query the current minting parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryParameters)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}

			var params types.Params
			if err := cdc.UnmarshalJSON(res, &params); err != nil {
				return err
			}

			return cliCtx.PrintOutput(params)
		},
	}
}

// GetCmdQueryStatus implements a command to return the current minting
// inflation value.
func GetCmdQueryStatus(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Query minting status",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryStatus)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}

			var minter types.Minter
			if err := cdc.UnmarshalJSON(res, &minter); err != nil {
				return err
			}

			return cliCtx.PrintOutput(minter)
		},
	}
}

// GetCmdQueryBonus implements a command to return the current minting
// inflation value.
func GetCmdQueryBonus(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "bonus [bheight]",
		Short: "Query minting bonus for a block",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryBonus)
			res, _, err := cliCtx.QueryWithData(route, []byte(args[0]))
			if err != nil {
				return err
			}

			var coin sdk.Coin
			if err := cdc.UnmarshalJSON(res, &coin); err != nil {
				return err
			}

			return cliCtx.PrintOutput(coin)
		},
	}
}
