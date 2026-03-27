package main

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/libdefinite/definite/gen/server/v1/serverv1connect"
	"github.com/libdefinite/definite/internal/server/handler"
)

func serverCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Start the gRPC server",
		RunE: func(cmd *cobra.Command, args []string) error {
			mux := http.NewServeMux()

			path, handler := serverv1connect.NewCommonServiceHandler(&handler.CommonServiceHandler{})
			mux.Handle(path, handler)

			log.Println("listening on :8080")
			return http.ListenAndServe(":8080", h2c.NewHandler(mux, &http2.Server{}))
		},
	}
}
