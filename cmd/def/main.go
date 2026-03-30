package main

import (
	"log/slog"
	"os"

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
