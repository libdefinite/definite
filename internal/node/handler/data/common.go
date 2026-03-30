package grpc

import (
	"context"

	"connectrpc.com/connect"

	datav1 "github.com/libdefinite/definite/gen/data/v1"
	"github.com/libdefinite/definite/gen/data/v1/datav1connect"
)

// DataCommonHandler implements the CommonService gRPC handlers.
type DataCommonHandler struct {
	datav1connect.UnimplementedCommonServiceHandler
}

// GetStatus returns the current service status.
func (cs *DataCommonHandler) GetStatus(
	ctx context.Context,
	req *connect.Request[datav1.GetStatusRequest],
) (*connect.Response[datav1.GetStatusResponse], error) {
	return connect.NewResponse(&datav1.GetStatusResponse{
		Status: "ok",
	}), nil
}
