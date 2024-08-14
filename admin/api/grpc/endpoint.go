package grpc

import (
	"context"

	"go-kit-template/admin"

	"github.com/go-kit/kit/endpoint"
)

func grpcAddEndpoint(svc admin.GRPCService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(MathReq)
		result, err := svc.Add(ctx, req.NumA, req.NumB)
		if err != nil {
			return nil, err
		}
		return MathResp{Result: result}, nil
	}
}