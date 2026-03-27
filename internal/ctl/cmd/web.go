package cmd

import (
	"github.com/libdefinite/definite/internal/ctl/web"
	"github.com/spf13/cobra"
)

// WebCtlCmd returns the cobra command for starting the web console.
func WebCtlCmd() *cobra.Command {
	var port int

	cmd := &cobra.Command{
		Use:   "web",
		Short: "Start the web console",
		RunE: func(cmd *cobra.Command, args []string) error {
			return web.NewConsole(port).Start()
		},
	}

	cmd.Flags().IntVarP(&port, "port", "p", 9090, "server port")

	return cmd
}
