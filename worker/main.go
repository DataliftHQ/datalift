package main

import (
	"log"
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
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
}
