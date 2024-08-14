package grpc

import (
	"context"
	"go-kit-template/admin"
	pb "go-kit-template/proto"

	"github.com/go-kit/kit/transport/grpc"
)

var _ pb.MathServiceServer = (*serverRepository)(nil)

type serverRepository struct {
    add grpc.Handler // all meothod in Service
    pb.UnimplementedMathServiceServer // dont know what is this
}

func decodeMathRequest(_ context.Context, request interface{}) (interface{}, error) {
    req := request.(*pb.MathRequest)
    return MathReq{NumA: req.NumA, NumB: req.NumB}, nil
}

func encodeMathResponse(_ context.Context, response interface{}) (interface{}, error) {
    resp := response.(MathResp)
    return &pb.MathResponse{Result: resp.Result}, nil
}

func NewServerRepository(svc admin.GRPCService) pb.MathServiceServer {
    return &serverRepository{
        add: grpc.NewServer(
            grpcAddEndpoint(svc),
            decodeMathRequest,
            encodeMathResponse,
        ),
    }
}

func (s *serverRepository) Add(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
    _, resp, err := s.add.ServeGRPC(ctx, req)
    if err != nil {
        return nil, err
    }
    return resp.(*pb.MathResponse), nil
}