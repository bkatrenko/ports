package service

import (
	"context"

	"github.com/bkatrenko/ports/app/portdomain/domain/repo"
	"github.com/bkatrenko/ports/proto"
	wrapper "github.com/pkg/errors"
	"github.com/prometheus/common/log"
)

// DomainPortsService provide some custom logic on models and repo structures
type DomainPortsService struct {
	repo repo.PortsStorage
}

// NewPortsService ...
func NewPortsService(repo repo.PortsStorage) (*DomainPortsService, error) {
	return &DomainPortsService{
		repo: repo,
	}, nil
}

// GetPorts using app repo to retrieve ports from storage
func (ps *DomainPortsService) GetPorts(ctx context.Context, r *proto.GetPortsRequest) (*proto.GetPortsResponse, error) {
	log.Infof("get ports request with offset/limit: %d/%d", r.Offset, r.Limit)

	ports, err := ps.repo.GetPorts(r.Offset, r.Limit)
	if err != nil {
		return nil, wrapper.Wrap(err, "error while get ports from storage")
	}

	return &proto.GetPortsResponse{
		Ports: ports,
	}, nil
}

// PutPorts using repo to put ports to storage
func (ps *DomainPortsService) PutPorts(ctx context.Context, r *proto.PutPortsRequest) (*proto.PutPortsResponse, error) {
	log.Infof("put ports request with map len: %d", len(r.Ports))

	if err := ps.repo.PutPorts(r.Ports); err != nil {
		return nil, wrapper.Wrap(err, "error while put ports")
	}

	return &proto.PutPortsResponse{}, nil
}
