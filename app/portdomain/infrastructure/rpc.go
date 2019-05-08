package infrastructure

import (
	"context"
	"net"

	"github.com/bkatrenko/ports/app/portdomain/domain/service"
	"github.com/bkatrenko/ports/proto"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
)

// GRPCServer contains domain app service to handle incoming calls
type GRPCServer struct {
	service    *service.DomainPortsService
	serverAddr string
}

// NewServer is just simple constructor for server instance
func NewServer(serverAddr string, service *service.DomainPortsService) (*GRPCServer, error) {
	return &GRPCServer{
		service:    service,
		serverAddr: serverAddr,
	}, nil
}

// PutPorts implements PortsDomainServiceServer interface to handle gRPC calls with app service
func (s *GRPCServer) PutPorts(ctx context.Context, r *proto.PutPortsRequest) (*proto.PutPortsResponse, error) {
	return s.service.PutPorts(ctx, r)
}

// GetPorts implements PortsDomainServiceServer interface to handle gRPC calls with app service
func (s *GRPCServer) GetPorts(ctx context.Context, r *proto.GetPortsRequest) (*proto.GetPortsResponse, error) {
	return s.service.GetPorts(ctx, r)
}

// GRPC starting new server on [port]
func GRPC(portsService proto.PortsDomainServiceServer, serverAddr string) {
	server := grpc.NewServer()
	proto.RegisterPortsDomainServiceServer(server, portsService)

	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Infof("start to listen grpc conn on addr %s", serverAddr)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
