package main

import (
	"github.com/spf13/cobra"
	"github.com/suborbital/atmo/atmo"
	"github.com/suborbital/atmo/atmo/release"
)

func main() {
	cmd := rootCommand()
	cmd.Execute()
}

func rootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "atmo [bundle-path]",
		Short: "Atmo function-based web service runner",
		Long: `
Atmo is an all-in-one function-based web service platform that enables 
building backend systems using composable WebAssembly modules in a declarative manner.

Atmo automatically scales using a meshed message bus, job scheduler, and 
flexible API gateway to handle any workload. 

Handling API and event-based traffic is made simple using the declarative 
Directive format and the powerful Runnable API using a variety of languages.`,
		Version: release.AtmoDotVersion,
		RunE: func(cmd *cobra.Command, args []string) error {
			path := "./runnables.wasm.zip"
			if len(args) > 0 {
				path = args[0]
			}

			server := atmo.New()

			return server.Start(path)
		},
	}

	cmd.SetVersionTemplate("{{.Version}}\n")

	return cmd
}
