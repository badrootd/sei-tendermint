package main

import (
	"context"

	"github.com/badrootd/sei-tendermint/cmd/tendermint/commands"
	"github.com/badrootd/sei-tendermint/cmd/tendermint/commands/debug"
	"github.com/badrootd/sei-tendermint/config"
	"github.com/badrootd/sei-tendermint/libs/cli"
	"github.com/badrootd/sei-tendermint/libs/log"
	"github.com/badrootd/sei-tendermint/node"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf, err := commands.ParseConfig(config.DefaultConfig())
	if err != nil {
		panic(err)
	}

	logger, err := log.NewDefaultLogger(conf.LogFormat, conf.LogLevel)
	if err != nil {
		panic(err)
	}

	// A stop gap solution to restart the node for certain failures (such as network partition)
	restartCh := make(chan struct{})

	rcmd := commands.RootCommand(conf, logger)
	rcmd.AddCommand(
		commands.MakeGenValidatorCommand(),
		commands.MakeReindexEventCommand(conf, logger),
		commands.MakeInitFilesCommand(conf, logger),
		commands.MakeLightCommand(conf, logger),
		commands.MakeReplayCommand(conf, logger),
		commands.MakeReplayConsoleCommand(conf, logger),
		commands.MakeResetCommand(conf, logger),
		commands.MakeUnsafeResetAllCommand(conf, logger),
		commands.MakeShowValidatorCommand(conf, logger),
		commands.MakeTestnetFilesCommand(conf, logger),
		commands.MakeShowNodeIDCommand(conf),
		commands.GenNodeKeyCmd,
		commands.VersionCmd,
		commands.MakeInspectCommand(conf, logger),
		commands.MakeRollbackStateCommand(conf),
		commands.MakeKeyMigrateCommand(conf, logger),
		debug.GetDebugCommand(logger),
		commands.NewCompletionCmd(rcmd, true),
		commands.MakeCompactDBCommand(conf, logger),
	)

	// NOTE:
	// Users wishing to:
	//	* Use an external signer for their validators
	//	* Supply an in-proc abci app
	//	* Supply a genesis doc file from another source
	//	* Provide their own DB implementation
	// can copy this file and use something other than the
	// node.NewDefault function
	nodeFunc := node.NewDefault

	// Create & start node
	rcmd.AddCommand(commands.NewRunNodeCmd(nodeFunc, conf, logger, restartCh))

	if err := cli.RunWithTrace(ctx, rcmd); err != nil {
		panic(err)
	}
}
