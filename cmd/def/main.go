package main

import (
	"log/slog"
	"os"

	"github.com/libdefinite/definite/internal/ctl"
	"github.com/libdefinite/definite/internal/node"
	"github.com/spf13/cobra"
)

func main() {
	var logLevel string

	root := &cobra.Command{
		Use:   "definite",
		Short: "Definite application",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var level slog.Level
			if err := level.UnmarshalText([]byte(logLevel)); err != nil {
				return err
			}
			slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level})))
			return nil
		},
	}

	root.Flags().StringVarP(&logLevel, "log", "l", "info", "log level (debug, info, warn, error)")

	root.AddCommand(nodeCmd(), ctlCmd())

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

func ctlCmd() *cobra.Command {
	c := &cobra.Command{
		Use:   "ctl",
		Short: "Control the server via gRPC",
	}

	c.AddCommand(ctl.StatusCmd(), ctl.ConsoleCmd())
	return c
}

func nodeCmd() *cobra.Command {
	var dataPort int

	cmd := &cobra.Command{
		Use:   "node",
		Short: "Start definite node",
		RunE: func(cmd *cobra.Command, args []string) error {
			return node.Start(dataPort, 10000+dataPort)
		},
	}

	cmd.Flags().IntVarP(&dataPort, "port", "p", 9876, "data plane port")

	return cmd
}
