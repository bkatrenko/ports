package repo

import (
	"github.com/bkatrenko/ports/proto"
)

// ClientPortsRepo using like repo for client application and should have a functions to retrieve and
// put ports info.
type ClientPortsRepo interface {
	// RetrievePorts using to get saved ports from domain app with some offset and limit -
	// expecter that limit and offset will be more then 0
	RetrievePorts(offset, limit int64) (map[string]*proto.Port, error)
	// PutPorts should read JSON file with ports info and put it to ports domain service
	PutPorts(jsonPortsFilePath string) error
}
