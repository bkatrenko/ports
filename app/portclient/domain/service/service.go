package service

import (
	"github.com/bkatrenko/ports/app/portclient/domain/repo"
	"github.com/bkatrenko/ports/proto"
	wrapper "github.com/pkg/errors"
	"github.com/prometheus/common/log"
)

// ClientPortsService contain logic for working with models and repo.
type ClientPortsService struct {
	repo        repo.ClientPortsRepo
	jsonMapPath string
}

// NewPortsService is just simple constructor for ClientPortsService
func NewPortsService(jsonMapPath string, repo repo.ClientPortsRepo) (*ClientPortsService, error) {
	return &ClientPortsService{
		repo:        repo,
		jsonMapPath: jsonMapPath,
	}, nil
}

// PutEvents using jsonMapPath variable from ClientPortsService and with client app repo
// pushing ports to domain ports service.
func (ps *ClientPortsService) PutEvents() error {
	log.Infof("start to put ports from file: %s", ps.jsonMapPath)

	if err := ps.repo.PutPorts(ps.jsonMapPath); err != nil {
		return wrapper.Wrap(err, "error while put ports to ports domain service")
	}

	return nil
}

// RetrievePorts using repo to retrieve events from ports domain app
func (ps *ClientPortsService) RetrievePorts(offset, limit int64) (map[string]*proto.Port, error) {
	log.Infof("start to retrieve ports from DB with offset/limit: %d/%d", offset, limit)

	ports, err := ps.repo.RetrievePorts(offset, limit)
	if err != nil {
		return nil, wrapper.Wrap(err, "error while retrieve ports")
	}

	return ports, nil
}
