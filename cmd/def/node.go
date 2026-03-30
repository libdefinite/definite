package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/spf13/cobra"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/libdefinite/definite/gen/control/v1/controlv1connect"
	"github.com/libdefinite/definite/gen/data/v1/datav1connect"
	controlhandler "github.com/libdefinite/definite/internal/node/handler/control"
	datahandler "github.com/libdefinite/definite/internal/node/handler/data"
)

func nodeCmd() *cobra.Command {
	var dataPort int

	cmd := &cobra.Command{
		Use:   "node",
		Short: "Start definite node",
		RunE: func(cmd *cobra.Command, args []string) error {
			errc := make(chan error, 2)
			go serveDataPlane(dataPort, errc)
			go serveControlPlane(dataPort+10000, errc)
			return <-errc
		},
	}

	cmd.Flags().IntVarP(&dataPort, "port", "p", 9876, "data plane port")

	return cmd
}

func serveDataPlane(port int, errc chan<- error) {
	mux := http.NewServeMux()
	path, handler := datav1connect.NewCommonServiceHandler(&datahandler.DataCommonHandler{})
	mux.Handle(path, handler)
	addr := fmt.Sprintf(":%d", port)
	slog.Info("data plane listening", "addr", addr)
	errc <- http.ListenAndServe(addr, h2c.NewHandler(mux, &http2.Server{}))
}

func serveControlPlane(port int, errc chan<- error) {
	mux := http.NewServeMux()
	path, handler := controlv1connect.NewCommonServiceHandler(&controlhandler.ControlCommonHandler{})
	mux.Handle(path, handler)
	addr := fmt.Sprintf(":%d", port)
	slog.Info("control plane listening", "addr", addr)
	errc <- http.ListenAndServe(addr, h2c.NewHandler(mux, &http2.Server{}))
}
