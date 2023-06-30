package main

import (
	_ "embed"
	"os"

	"go.datalift.io/datalift/cmd"
	"go.datalift.io/datalift/internal/version"
)

//go:embed art.txt
var asciiArt string

func main() {
	cmd.Execute(
		buildInfo(),
		os.Exit,
		os.Args[1:],
	)
}

func buildInfo() version.Info {
	return version.GetVersionInfo(
		version.WithAppDetails("Datalift", "Platform Orchestrator that helps developers build, deploy, and manage their applications more quickly and easily", "https://datalift.io"),
		version.WithASCIIName(asciiArt),
	)
}
