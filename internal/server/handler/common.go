package handler

import (
	"context"

	"connectrpc.com/connect"

	definitev1 "github.com/libdefinite/definite/gen/server/v1"
	"github.com/libdefinite/definite/gen/server/v1/serverv1connect"
)

// CommonServiceHandler implements the CommonService gRPC handlers.
type CommonServiceHandler struct {
	serverv1connect.UnimplementedCommonServiceHandler
}

// GetStatus returns the current service status.
func (s *CommonServiceHandler) GetStatus(
	ctx context.Context,
	req *connect.Request[definitev1.GetStatusRequest],
) (*connect.Response[definitev1.GetStatusResponse], error) {
	return connect.NewResponse(&definitev1.GetStatusResponse{
		Status: "ok",
	}), nil
}
