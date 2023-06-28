package main

import (
	_ "embed"
	"os"

	"go.datalift.io/datalift/cmd"
	ver "go.datalift.io/datalift/internal/version"
)

// nolint: gochecknoglobals
var (
	version   = ""
	commit    = ""
	treeState = ""
	date      = ""
	builtBy   = ""
)

func main() {
	cmd.Execute(
		buildVersion(version, commit, date, builtBy),
		os.Exit,
		os.Args[1:],
	)
}

const website = "https://datalift.io"

//go:embed art.txt
var asciiArt string

func buildVersion(version, commit, date, builtBy string) ver.Info {
	return ver.GetVersionInfo(
		ver.WithAppDetails("Datalift", "Platform Orchestrator that helps developers build, deploy, and manage their applications more quickly and easily", website),
		ver.WithASCIIName(asciiArt),
		func(i *ver.Info) {
			if commit != "" {
				i.GitCommit = commit
			}
			if treeState != "" {
				i.GitTreeState = treeState
			}
			if date != "" {
				i.BuildDate = date
			}
			if version != "" {
				i.GitVersion = version
			}
			if builtBy != "" {
				i.BuiltBy = builtBy
			}
		},
	)
}
