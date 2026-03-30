package ctl

import (
	"github.com/libdefinite/definite/internal/ctl/console"
	"github.com/spf13/cobra"
)

// WebCtlCmd returns the cobra command for starting the web console.
func WebCtlCmd() *cobra.Command {
	var port int

	cmd := &cobra.Command{
		Use:   "console",
		Short: "Start the web console",
		RunE: func(cmd *cobra.Command, args []string) error {
			return console.NewConsole(port).Start()
		},
	}

	cmd.Flags().IntVarP(&port, "port", "p", 3000, "server port")

	return cmd
}

// FetchCtlCmd returns the cobra command for fetching server status.
func FetchCtlCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Get server status",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
