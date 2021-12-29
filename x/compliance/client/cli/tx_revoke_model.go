package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

var _ = strconv.Itoa(0)

func CmdRevokeModel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revoke-model",
		Short: "Revoke an existing model",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVid, err := cast.ToInt32E(viper.GetString(FlagVID))
			if err != nil {
				return err
			}
			argPid, err := cast.ToInt32E(viper.GetString(FlagPID))
			if err != nil {
				return err
			}
			argSoftwareVersion, err := cast.ToUint32E(viper.GetString(FlagSoftwareVersion))
			if err != nil {
				return err
			}
			argSoftwareVersionString := viper.GetString(FlagSoftwareVersionString)
			argRevocationDate := viper.GetString(FlagRevocationDate)
			argCertificationType := viper.GetString(FlagCertificationType)
			argReason := viper.GetString(FlagReason)
			argCDVersionNumber, err := cast.ToUint32E(viper.GetString(FlagCDVersionNumber))
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRevokeModel(
				clientCtx.GetFromAddress().String(),
				argVid,
				argPid,
				argSoftwareVersion,
				argSoftwareVersionString,
				argCDVersionNumber,
				argRevocationDate,
				argCertificationType,
				argReason,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().String(FlagSoftwareVersion, "", "Model software version")
	cmd.Flags().String(FlagSoftwareVersionString, "", "Model software version string")
	cmd.Flags().StringP(FlagCertificationType, FlagCertificationTypeShortcut, "", TextCertificationType)
	cmd.Flags().StringP(FlagRevocationDate, FlagDateShortcut, "",
		"The date of model revocation (rfc3339 encoded)")
	cmd.Flags().StringP(FlagReason, FlagReasonShortcut, "",
		"Optional comment describing the reason of revocation")
	cmd.Flags().String(FlagCDVersionNumber, "", "CD Version Number of the certification")

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersionString)
	_ = cmd.MarkFlagRequired(FlagCertificationType)
	_ = cmd.MarkFlagRequired(FlagRevocationDate)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
