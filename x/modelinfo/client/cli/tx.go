package cli

import (
	"fmt"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/cli"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/utils/conversions"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/modelinfo/internal/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	modelinfoTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Modelinfo transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	modelinfoTxCmd.AddCommand(cli.SignedCommands(client.PostCommands(
		GetCmdAddModel(cdc),
		GetCmdUpdateModel(cdc),
		//GetCmdDeleteModel(cdc), Disable deletion
	)...)...)

	return modelinfoTxCmd
}

//nolint dupl
func GetCmdAddModel(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-model",
		Short: "Add new Model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid, err := conversions.ParseVID(viper.GetString(FlagVID))
			if err != nil {
				return err
			}

			pid, err := conversions.ParsePID(viper.GetString(FlagPID))
			if err != nil {
				return err
			}

			name := viper.GetString(FlagName)

			description, err_ := cliCtx.ReadFromFile(viper.GetString(FlagDescription))
			if err_ != nil {
				return err_
			}

			sku := viper.GetString(FlagSKU)
			firmwareVersion := viper.GetString(FlagFirmwareVersion)
			hardwareVersion := viper.GetString(FlagHardwareVersion)

			tisOrTrpTestingCompleted, err_ := strconv.ParseBool(viper.GetString(FlagTisOrTrpTestingCompleted))
			if err_ != nil {
				return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid Tis-or-trp-testing-completed: Parsing Error: \"%v\" must be boolean", viper.GetString(FlagTisOrTrpTestingCompleted)))
			}

			custom, err_ := cliCtx.ReadFromFile(viper.GetString(FlagCustom))
			if err_ != nil {
				return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid custom:\"%v\"", err_))
			}

			var cid uint16
			if cidStr := viper.GetString(FlagCID); len(cidStr) != 0 {
				cid, err = conversions.ParseCID(cidStr)
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgAddModelInfo(vid, pid, cid, name, description, sku, firmwareVersion, hardwareVersion,
				custom, tisOrTrpTestingCompleted, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().String(FlagCID, "", "Model category ID")
	cmd.Flags().StringP(FlagName, FlagNameShortcut, "", "Model name")
	cmd.Flags().StringP(FlagDescription, FlagDescriptionShortcut, "", "Model description (string or path to file containing data)")
	cmd.Flags().String(FlagSKU, "", "Model stock keeping unit")
	cmd.Flags().StringP(FlagFirmwareVersion, FlagFirmwareVersionShortcut, "", "Version of model firmware")
	cmd.Flags().StringP(FlagHardwareVersion, FlagHardwareVersionShortcut,"", "Version of model hardware")
	cmd.Flags().StringP(FlagCustom, FlagCustomShortcut, "", "Custom information (string or path to file containing data)")
	cmd.Flags().StringP(FlagTisOrTrpTestingCompleted, FlagTisOrTrpTestingCompletedShortcut, "", "Whether model has successfully completed TIS/TRP testing")

	cmd.MarkFlagRequired(FlagVID)
	cmd.MarkFlagRequired(FlagPID)
	cmd.MarkFlagRequired(FlagName)
	cmd.MarkFlagRequired(FlagSKU)
	cmd.MarkFlagRequired(FlagFirmwareVersion)
	cmd.MarkFlagRequired(FlagHardwareVersion)
	cmd.MarkFlagRequired(FlagTisOrTrpTestingCompleted)

	return cmd
}

//nolint dupl
func GetCmdUpdateModel(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-model",
		Short: "Update existing Model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid, err := conversions.ParseVID(viper.GetString(FlagVID))
			if err != nil {
				return err
			}

			pid, err := conversions.ParsePID(viper.GetString(FlagPID))
			if err != nil {
				return err
			}

			tisOrTrpTestingCompleted, err_ := strconv.ParseBool(viper.GetString(FlagTisOrTrpTestingCompleted))
			if err_ != nil {
				return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid tis-or-trp-testing-completed: Parsing Error: \"%v\" must be boolean", viper.GetString(FlagTisOrTrpTestingCompleted)))
			}

			description, err_ := cliCtx.ReadFromFile(viper.GetString(FlagDescription))
			if err_ != nil {
				return err_
			}

			var cid uint16
			if cidStr := viper.GetString(FlagCID); len(cidStr) != 0 {
				cid, err = conversions.ParseCID(cidStr)
				if err != nil {
					return err
				}
			}

			custom, err_ := cliCtx.ReadFromFile(viper.GetString(FlagCustom))
			if err_ != nil {
				return sdk.ErrUnknownRequest(fmt.Sprintf("Invalid custom:\"%v\"", err_))
			}

			msg := types.NewMsgUpdateModelInfo(vid, pid, cid, description, custom, tisOrTrpTestingCompleted, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().String(FlagCID, "", "Model category ID")
	cmd.Flags().StringP(FlagDescription, FlagDescriptionShortcut, "", "Model description (string or path to file containing data)")
	cmd.Flags().StringP(FlagCustom, FlagCustomShortcut, "", "Custom information (string or path to file containing data)")
	cmd.Flags().StringP(FlagTisOrTrpTestingCompleted, FlagTisOrTrpTestingCompletedShortcut, "", "Whether model has successfully completed TIS/TRP testing")


	cmd.MarkFlagRequired(FlagVID)
	cmd.MarkFlagRequired(FlagPID)
	cmd.MarkFlagRequired(FlagTisOrTrpTestingCompleted)

	return cmd
}

func GetCmdDeleteModel(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-model",
		Short: "Delete existing Model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := cli.NewCLIContext().WithCodec(cdc)

			vid, err := conversions.ParseVID(viper.GetString(FlagVID))
			if err != nil {
				return err
			}

			pid, err := conversions.ParsePID(viper.GetString(FlagPID))
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteModelInfo(vid, pid, cliCtx.FromAddress())

			return cliCtx.HandleWriteMessage(msg)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")

	cmd.MarkFlagRequired(FlagVID)
	cmd.MarkFlagRequired(FlagPID)

	return cmd
}