package repo

import (
	"github.com/bkatrenko/ports/proto"
)

// PortsStorage represents repo for ports info
type PortsStorage interface {
	GetPorts(offset, limit int64) (map[string]*proto.Port, error)
	PutPorts(map[string]*proto.Port) error
}
