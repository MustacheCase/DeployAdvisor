package cli

import (
	"advisor/cmd/internal"

	"github.com/anchore/clio"
	"github.com/spf13/cobra"

	"advisor/cmd/internal/commands"
	"io"
	"os"
)

// Application constructs the root command to `scan ci`.
// It is also responsible for organizing flag usage and injecting the application config for each command.
func Application(id clio.Identification) clio.Application {
	app, _ := create(id, os.Stdout)
	return app
}

// Command returns the root command for the DeployAdvisor CLI application. This is useful for embedding the entire DeployAdvisor CLI
// into an existing application.
func Command(id clio.Identification) *cobra.Command {
	_, cmd := create(id, os.Stdout)
	return cmd
}

func create(id clio.Identification, out io.Writer) (clio.Application, *cobra.Command) {
	clioCfg := internal.AppClioSetupConfig(id, out)

	app := clio.New(*clioCfg)

	// since root is aliased as the packages cmd we need to construct this command first
	// we also need the command to have information about the `root` options because of this alias
	scanCmd := commands.Scan(app)

	// root is currently an alias for the scan command
	rootCmd := commands.Root(app, scanCmd)

	// add sub-commands
	rootCmd.AddCommand(
		scanCmd,
		clio.VersionCommand(id),
		clio.ConfigCommand(app, nil),
	)

	return app, rootCmd
}
