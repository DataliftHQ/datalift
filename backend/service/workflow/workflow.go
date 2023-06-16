package workflow

import (
	"context"
	"fmt"
	"github.com/DataliftHQ/datalift/backend/service"
	"github.com/uber-go/tally/v4"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/anypb"
	"log"
	"time"

	"go.temporal.io/sdk/client"
)

const Name = "datalift.service.workflow"

type svc struct {
	logger *zap.Logger
	scope  tally.Scope
	client client.Client
}

type Service interface {
	Test(ctx context.Context)
}

func New(_ *anypb.Any, logger *zap.Logger, scope tally.Scope) (service.Service, error) {

	c, err := client.Dial(client.Options{
		HostPort: "127.0.0.1:7233",
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}

	return &svc{logger: logger, scope: scope, client: c}, nil
}

func (s *svc) Test(ctx context.Context) {
	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: "GREETING_TASK_QUEUE",
	}

	name := "World"
	we, err := s.client.ExecuteWorkflow(context.Background(), options, GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}
	// Get the results
	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}

	printResults(greeting, we.GetID(), we.GetRunID())
}

func printResults(greeting string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", greeting)
}

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	err := workflow.ExecuteActivity(ctx, ComposeGreeting, name).Get(ctx, &result)

	return result, err
}

func ComposeGreeting(ctx context.Context, name string) (string, error) {
	greeting := fmt.Sprintf("Hello %s!", name)
	return greeting, nil
}
