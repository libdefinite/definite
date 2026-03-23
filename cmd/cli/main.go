package main

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"

	definitev1 "github.com/libdefinite/definite/gen/server/v1"
	"github.com/libdefinite/definite/gen/server/v1/serverv1connect"
)

func main() {
	// h2c transport: HTTP/2 cleartext for gRPC without TLS
	transport := &http2.Transport{
		AllowHTTP: true,
		DialTLSContext: func(ctx context.Context, network, addr string, _ *tls.Config) (net.Conn, error) {
			return (&net.Dialer{}).DialContext(ctx, network, addr)
		},
	}

	client := serverv1connect.NewCommonServiceClient(
		&http.Client{Transport: transport},
		"http://localhost:8080",
		connect.WithGRPC(),
	)

	res, err := client.GetStatus(context.Background(), connect.NewRequest(&definitev1.GetStatusRequest{}))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("status:", res.Msg.Status)
}
