package worker

import (
	"github.com/spf13/cobra"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
	"os"
)

type WorkerCmd struct {
	Cmd  *cobra.Command
	Opts workerOpts
}

type workerOpts struct {
}

func NewWorkerCmd() *WorkerCmd {
	root := &WorkerCmd{}
	cmd := &cobra.Command{
		Use:               "worker",
		Aliases:           []string{"w"},
		Short:             "Generates GoReleaser's command line docs",
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("this will be the future home of the datalift worker 😀")

			hostPort, isSet := os.LookupEnv("TEMPORAL_ADDRESS")
			if !isSet {
				hostPort = "127.0.0.1:7233"
			}

			log.Println("client connecting to " + hostPort)
			c, err := client.Dial(client.Options{HostPort: hostPort})
			if err != nil {
				log.Fatalln("Unable to create client", err)
			}
			defer c.Close()

			w := worker.New(c, "hello-world", worker.Options{})
			err = w.Run(worker.InterruptCh())
			if err != nil {
				log.Fatalln("Unable to start worker", err)
			}

			return nil
		},
	}
	root.Cmd = cmd
	return root
}
