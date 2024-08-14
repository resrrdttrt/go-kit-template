package grpc

import (
	"context"
	pb "go-kit-template/proto"

	"github.com/go-kit/kit/endpoint"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"google.golang.org/grpc"
)

var _ pb.MathServiceClient = (*clientRepository)(nil)

type clientRepository struct {
	add endpoint.Endpoint
}

func NewClientRepository(conn *grpc.ClientConn, logger log.Logger) pb.MathServiceClient {
	return &clientRepository{
		add: kitgrpc.NewClient(
			conn,
			pb.MathService_Add_FullMethodName,
			"MathService",
			encodeMathRequest,
			decodeMathResponse,
			pb.MathResponse{},
		).Endpoint(),
	}
}

func (c *clientRepository) Add(ctx context.Context, in *pb.MathRequest, opts ...grpc.CallOption) (*pb.MathResponse, error) {
	resp, err := c.add(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.MathResponse), nil
}

func encodeMathRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(MathReq)
	return &pb.MathRequest{NumA: req.NumA, NumB: req.NumB}, nil
}

func decodeMathResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.MathResponse)
	return MathResp{Result: resp.Result}, nil
}