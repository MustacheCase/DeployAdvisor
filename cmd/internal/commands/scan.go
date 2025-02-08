package commands

import (
	"advisor/cmd/internal/ui"

	"advisor/internal"

	"github.com/anchore/clio"
	"github.com/spf13/cobra"
)

const (
	dirHelp = `    {{.appName}} {{.command}} dir:path/to/yourproject read directly from a path on disk (any directory)
    {{.appName}} {{.command}} file:path/to/yourproject/file            read directly from a path on disk (any single file)
`
	scanSchemeHelp = "\n  " + dirHelp

	scanHelp = dirHelp + scanSchemeHelp
)

func Scan(app clio.Application) *cobra.Command {
	id := app.ID()

	return app.SetupCommand(&cobra.Command{
		Use:   "scan [SOURCE]",
		Short: "Generate an SBOM",
		Long:  "Generate a packaged-based Software Bill Of Materials (SBOM) from container images and filesystems",
		Example: internal.Tprintf(scanHelp, map[string]interface{}{
			"appName": id.Name,
			"command": "scan",
		}),
		//Args:    validateScanArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			restoreStdout := ui.CaptureStdoutToTraceLog()
			defer restoreStdout()

			return nil
		},
	})
}
