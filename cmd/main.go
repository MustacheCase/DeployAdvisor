package main

import (
	"advisor/cmd/internal"

	_ "modernc.org/sqlite"

	"advisor/cmd/cli"

	"github.com/anchore/clio"
)

// applicationName is the non-capitalized name of the application (do not change this)
const applicationName = "advisor"

// all variables here are provided as build-time arguments, with clear default values
var (
	version        = internal.NotProvided
	buildDate      = internal.NotProvided
	gitCommit      = internal.NotProvided
	gitDescription = internal.NotProvided
)

func main() {
	app := cli.Application(
		clio.Identification{
			Name:           applicationName,
			Version:        version,
			BuildDate:      buildDate,
			GitCommit:      gitCommit,
			GitDescription: gitDescription,
		},
	)

	app.Run()
}
