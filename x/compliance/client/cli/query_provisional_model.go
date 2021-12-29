package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zigbee-alliance/distributed-compliance-ledger/utils/cli"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/compliance/types"
)

func CmdListProvisionalModel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all-provisional-models",
		Short: "Query the list of all provisional models",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllProvisionalModelRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ProvisionalModelAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowProvisionalModel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "provisional-model",
		Short: "Gets a boolean if the given Model (identified by the `vid`, `pid`, `softwareVersion` and `certification_type`) is in provisional state",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

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
			argCertificationType := viper.GetString(FlagCertificationType)

			params := &types.QueryGetProvisionalModelRequest{
				Vid:               argVid,
				Pid:               argPid,
				SoftwareVersion:   argSoftwareVersion,
				CertificationType: argCertificationType,
			}

			res, err := queryClient.ProvisionalModel(context.Background(), params)
			if cli.HandleError(err) != nil {
				return err
			}
			if err != nil {
				// show default (empty) value in CLI
				res = &types.QueryGetProvisionalModelResponse{ProvisionalModel: nil}
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String(FlagVID, "", "Model vendor ID")
	cmd.Flags().String(FlagPID, "", "Model product ID")
	cmd.Flags().String(FlagSoftwareVersion, "", "Model software version")
	cmd.Flags().StringP(FlagCertificationType, FlagCertificationTypeShortcut, "", TextCertificationType)

	_ = cmd.MarkFlagRequired(FlagVID)
	_ = cmd.MarkFlagRequired(FlagPID)
	_ = cmd.MarkFlagRequired(FlagSoftwareVersion)
	_ = cmd.MarkFlagRequired(FlagCertificationType)

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
