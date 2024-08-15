package grpcutils

import (
	"fmt"
	"go-kit-template/pkg/errors"
	"go-kit-template/pkg/logger"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

var (
	errTLS      = errors.New("create TLS credentials failed")
	errDialGRPC = errors.New("dial GRPC failed")
)

const defaultScheme = "dns"

func NewConnGrpc(scheme, url string, clientCACerts string, clientTLS bool, logger logger.Logger) *grpc.ClientConn {
	conn, target, err := CreateClientConn(scheme, url, clientCACerts, clientTLS, grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:    time.Second * 10,
		Timeout: time.Second * 10,
	}))
	if err != nil {
		logger.Error(fmt.Sprintf("gRPC connect to [%s] failed: %s", target, err))
		os.Exit(1)
	}
	logger.Infof("gRPC connect to [%s] success", target)
	return conn
}

func CreateClientConn(scheme, url, clientCACerts string, clientTLS bool, opts ...grpc.DialOption) (*grpc.ClientConn, string, error) {
	if scheme == "" {
		scheme = defaultScheme
	}
	target := fmt.Sprintf("%s:///%s", scheme, url)
	opts = append(opts, []grpc.DialOption{
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`), // This sets the initial balancing policy.
	}...)
	if clientTLS && clientCACerts != "" {
		//TODO check logic TLS
		tpc, err := credentials.NewClientTLSFromFile(clientCACerts, "")
		if err != nil {
			return nil, target, errors.Wrap(errTLS, err)
		}
		opts = append(opts, grpc.WithTransportCredentials(tpc))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		return nil, target, errors.Wrap(errDialGRPC, err)
	}
	return conn, target, nil
}
