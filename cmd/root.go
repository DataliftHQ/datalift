package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"go.datalift.io/datalift/cmd/migrate"
	"go.datalift.io/datalift/cmd/server"
	"go.datalift.io/datalift/cmd/worker"
	"go.datalift.io/datalift/internal/version"
)

type rootCmd struct {
	cmd     *cobra.Command
	log     *zap.Logger
	verbose bool
	exit    func(int)
}

func Execute(version version.Info, exit func(int), args []string) {
	newRootCmd(version, exit).Execute(args)
}

func newRootCmd(version version.Info, exit func(int)) *rootCmd {
	//version version.Info,
	atom := zap.NewAtomicLevel()
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.Lock(os.Stdout),
		atom,
	))
	defer logger.Sync()

	root := &rootCmd{
		exit: exit,
		log:  logger,
	}

	cmd := &cobra.Command{
		Use:               "datalift",
		Short:             "Platform Orchestrator that helps developers build, deploy, and manage their applications",
		Version:           version.String(),
		SilenceUsage:      true,
		SilenceErrors:     true,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			if root.verbose {
				atom.SetLevel(zap.DebugLevel)
				root.log.Debug("verbose output enabled")
			}
		},
	}
	cmd.SetVersionTemplate("{{.Version}}")
	cmd.CompletionOptions.DisableDefaultCmd = true

	cmd.PersistentFlags().BoolVar(&root.verbose, "verbose", false, "Enable verbose mode")
	cmd.AddCommand(
		migrate.NewCmd(logger).Cmd,
		server.NewServerCmd().Cmd,
		worker.NewWorkerCmd().Cmd,
	)
	root.cmd = cmd

	return root
}

func (cmd *rootCmd) Execute(args []string) {
	cmd.cmd.SetArgs(args)

	if err := cmd.cmd.Execute(); err != nil {
		code := 1
		msg := "command failed"
		eerr := &exitError{}
		if errors.As(err, &eerr) {
			code = eerr.code
			if eerr.details != "" {
				msg = eerr.details
			}
		}
		cmd.log.Error(msg, zap.Error(err))
		cmd.exit(code)
	}
}
