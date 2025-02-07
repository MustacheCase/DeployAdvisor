package commands

import (
	"advisor/cmd/internal/ui"
	"fmt"

	"github.com/anchore/clio"
	"github.com/spf13/cobra"
)

func Root(app clio.Application, packagesCmd *cobra.Command) *cobra.Command {
	//id := app.ID()

	//opts := defaultScanOptions()

	return app.SetupRootCommand(&cobra.Command{
		Use:     fmt.Sprintf("%s [SOURCE]", app.ID().Name),
		Short:   packagesCmd.Short,
		Long:    packagesCmd.Long,
		Args:    packagesCmd.Args,
		Example: packagesCmd.Example,
		RunE: func(cmd *cobra.Command, args []string) error {
			restoreStdout := ui.CaptureStdoutToTraceLog()
			defer restoreStdout()

			return nil
		},
	}, nil)
}
