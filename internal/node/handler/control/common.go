package grpc

import (
	"context"

	"connectrpc.com/connect"

	controlv1 "github.com/libdefinite/definite/gen/control/v1"
	"github.com/libdefinite/definite/gen/control/v1/controlv1connect"
)

// ControlCommonHandler implements the ControlService gRPC handlers.
type ControlCommonHandler struct {
	controlv1connect.UnimplementedCommonServiceHandler
}

// HealthCheck returns the current service health.
func (h *ControlCommonHandler) HealthCheck(
	ctx context.Context,
	req *connect.Request[controlv1.HealthCheckRequest],
) (*connect.Response[controlv1.HealthCheckResponse], error) {
	return connect.NewResponse(&controlv1.HealthCheckResponse{
		Status: "ok",
	}), nil
}
