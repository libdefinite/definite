package main

import (
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/libdefinite/definite/gen/server/v1/serverv1connect"
	"github.com/libdefinite/definite/internal/server"
)

func main() {
	mux := http.NewServeMux()

	path, handler := serverv1connect.NewCommonServiceHandler(&server.CommonServiceHandler{})
	mux.Handle(path, handler)

	log.Println("listening on :8080")
	if err := http.ListenAndServe(
		":8080",
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatal(err)
	}
}
