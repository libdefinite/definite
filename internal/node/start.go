package node

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/libdefinite/definite/gen/control/v1/controlv1connect"
	"github.com/libdefinite/definite/gen/data/v1/datav1connect"
	controlhandler "github.com/libdefinite/definite/internal/node/handler/control"
	datahandler "github.com/libdefinite/definite/internal/node/handler/data"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// Start entry point for node
func Start(dataPort, controlPort int) error {
	errc := make(chan error, 2)
	go serveDataPlane(dataPort, errc)
	go serveControlPlane(controlPort, errc)
	return <-errc

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
