package cli

import (
	"github.com/spf13/cobra"

	"github.com/hschain/hschain/client"
	"github.com/hschain/hschain/client/context"
	"github.com/hschain/hschain/codec"
	sdk "github.com/hschain/hschain/types"
	"github.com/hschain/hschain/x/auth"
	"github.com/hschain/hschain/x/auth/client/utils"
	"github.com/hschain/hschain/x/mint/internal/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Mint transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	burnCmd := BurnTxCmd(cdc)
	burnCmd.AddCommand(
		ConversionRateTxCmd(cdc),
	)

	txCmd.AddCommand(
		burnCmd,
		IssueTxCmd(cdc),
		DestroyTxCmd(cdc),
		DestroyUserTxCmd(cdc),
		PermissionsTxCmd(cdc),
		AddSysAddressTxCmd(cdc),
		SupplementTxCmd(cdc),
		VanishTxCmd(cdc),
		VanishUserTxCmd(cdc),
	)
	return txCmd
}

// BurnTxCmd will create a send tx and sign it with the given key.
func BurnTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn [from_key_or_address] [amount]",
		Short: "Create and sign a burn tx",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithFrom(args[0]).WithCodec(cdc)

			// parse coins trying to be sent
			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgBurn(cliCtx.GetFromAddress(), coins)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = client.PostCommands(cmd)[0]

	return cmd
}

func ConversionRateTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "conversion-rate [rate] --from=[name]",
		Short: "Create and sign a burn conversion-rate tx,the rate is *uhst = 1t",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// parse coins trying to be sent
			coins, err := sdk.ParseCoins(args[0])
			if err != nil {
				return err
			}

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgConversionRate(cliCtx.GetFromAddress(), coins)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = client.PostCommands(cmd)[0]

	return cmd
}

func PermissionsTxCmd(cdc *codec.Codec) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "permissions [command] [to_address] [bool] --from=[name]",
		Short: "Add or remove permissions for the operation [command] for the [to_address]",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			to, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			// parse coins trying to be sent
			command := args[0]

			status := 0
			if args[2] == "true" {
				status = 1
			}
			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgPermissions(cliCtx.GetFromAddress(), to, command, status)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = client.PostCommands(cmd)[0]

	return cmd
}

func AddSysAddressTxCmd(cdc *codec.Codec) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "system-address [command] [to_address] --from=[name]",
		Short: "Add a system address [to_address] for [command]]",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			to, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			// parse coins trying to be sent
			command := args[0]

			// build and sign the transaction, then broadcast to Tendermint

			msg := types.NewMsgAddSysAddress(cliCtx.GetFromAddress(), to, command)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = client.PostCommands(cmd)[0]

	return cmd
}

// DestroyTxCmd will create a send tx and sign it with the given key.
func DestroyTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "destroy [from_key_or_address] [amount]",
		Short: "Create and sign a destroy tx",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithFrom(args[0]).WithCodec(cdc)

			// parse coins trying to be sent
			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgDestory(cliCtx.GetFromAddress(), coins)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = client.PostCommands(cmd)[0]

	return cmd
}

func DestroyUserTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "destroy-user [to_address] [amount] --from=[name] ",
		Short: "Create and sign a destroy user tx",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			to, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			// parse coins trying to be sent
			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgDestoryUser(cliCtx.GetFromAddress(), to, coins)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = client.PostCommands(cmd)[0]

	return cmd
}

//IssueTxCmd will ipo new coins if no exist
func IssueTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue [to_address] [amount] --from=[name]",
		Short: "Create and sign a issue tx",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			to, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			// parse coins trying to be sent
			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgIssue(cliCtx.GetFromAddress(), to, coins)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = client.PostCommands(cmd)[0]

	return cmd
}

//SupplementTxCmd will ipo new coins if no exist
func SupplementTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "supplement  [amount] --from=[name]",
		Short: "Create and sign a supplement tx",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// parse coins trying to be sent
			coins, err := sdk.ParseCoins(args[0])
			if err != nil {
				return err
			}

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgSupplement(cliCtx.GetFromAddress(), coins)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = client.PostCommands(cmd)[0]

	return cmd
}

//VanishTxCmd will ipo new coins if no exist
func VanishTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vanish  [amount] --from=[name]",
		Short: "Create and sign a vanish tx",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// parse coins trying to be sent
			coins, err := sdk.ParseCoins(args[0])
			if err != nil {
				return err
			}

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgVanish(cliCtx.GetFromAddress(), coins)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = client.PostCommands(cmd)[0]

	return cmd
}

//VanishTxCmd will ipo new coins if no exist
func VanishUserTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vanishuser [from] [amount] --from=[name]",
		Short: "Create and sign a vanishuser tx",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			from, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			// parse coins trying to be sent
			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			// build and sign the transaction, then broadcast to Tendermint
			msg := types.NewMsgVanishUser(cliCtx.GetFromAddress(), from, coins)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd = client.PostCommands(cmd)[0]

	return cmd
}
