package commands

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/badrootd/sei-tendermint/config"
	"github.com/badrootd/sei-tendermint/internal/dbsync"
)

func MakeSnapshotCommand(confGetter func(*cobra.Command) (*config.Config, error)) *cobra.Command {
	return &cobra.Command{
		Use:   "snapshot [height]",
		Short: "Take DBSync snapshot for given height",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := confGetter(cmd)
			if err != nil {
				return err
			}
			height, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}
			return dbsync.Snapshot(height, *conf.DBSync, conf.BaseConfig)
		},
	}
}
