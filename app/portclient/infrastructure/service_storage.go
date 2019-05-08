package infrastructure

import (
	"context"
	"encoding/json"
	"io"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"os"

	"github.com/bkatrenko/ports/proto"
	wrapper "github.com/pkg/errors"
)

// ServiceStorage using to connect to Domain Service to store and retrieve ports info
type ServiceStorage struct {
	portDomainClient      proto.PortsDomainServiceClient
	timeoutSec            time.Duration
	portDomainServiceAddr string
}

// NewServiceStorage is just constructor to ServiceStorage struct
func NewServiceStorage(portsDomainServiceAddr string, timeoutSec time.Duration) (*ServiceStorage, error) {
	log.Infof("start to connect to storage on addr: %s with timeout: %v", portsDomainServiceAddr, timeoutSec)

	conn, err := grpc.Dial(portsDomainServiceAddr, grpc.WithInsecure(), grpc.WithTimeout(timeoutSec), grpc.WithBlock())
	if err != nil {
		return nil, wrapper.Wrap(err, "error while connect to ports domain service")
	}

	log.Infof("connection status: %s", conn.GetState().String())

	portDomainClient := proto.NewPortsDomainServiceClient(conn)

	log.Info("connected to ports domain service")
	return &ServiceStorage{
		portDomainClient:      portDomainClient,
		portDomainServiceAddr: portsDomainServiceAddr,
		timeoutSec:            timeoutSec,
	}, nil
}

// RetrievePorts using to get saved ports from ports domain service
func (ps *ServiceStorage) RetrievePorts(offset, limit int64) (map[string]*proto.Port, error) {
	ctx, cancel := context.WithTimeout(context.Background(), ps.timeoutSec)
	defer cancel()

	ports, err := ps.portDomainClient.GetPorts(ctx, &proto.GetPortsRequest{Offset: offset, Limit: limit})
	if err != nil {
		return nil, wrapper.Wrap(err, "error while get ports from ports domain service")
	}

	return ports.Ports, nil
}

// PutPorts using to save ports in domain service storage
func (ps *ServiceStorage) PutPorts(path string) error {
	log.Info("start to put events")
	reader, err := ps.readFile(path)
	if err != nil {
		return wrapper.Wrap(err, "error while open file")
	}

	ch := ps.decodeFile(reader)

	var counter int
	for v := range ch {
		ctx, cancel := context.WithTimeout(context.Background(), ps.timeoutSec)

		if _, err := ps.portDomainClient.PutPorts(ctx, &proto.PutPortsRequest{Ports: v}); err != nil {
			cancel()
			return wrapper.Wrap(err, "error while put ports")
		}
		counter++
		log.Infof("put events count: %d", counter)
		cancel()
	}

	return nil
}

func (ps *ServiceStorage) decodeFile(r io.Reader) chan map[string]*proto.Port {
	ch := make(chan map[string]*proto.Port)

	go func() {
		defer close(ch)
		dec := json.NewDecoder(r)

		_, err := dec.Token()
		if err != nil {
			log.Fatal(err, "error while read start token")
		}

		for {
			res := make(map[string]*proto.Port)

			token, err := dec.Token()
			switch err {
			case nil:
			case io.EOF:
				return
			default:
				log.Error(err)
			}

			if r, ok := token.(rune); ok && r == '}' {
				return
			}

			port := proto.Port{}

			err = dec.Decode(&port)
			switch err {
			case nil:
			case io.EOF:
				return
			default:
				log.Error(err)
			}

			res[token.(string)] = &port
			ch <- res
		}
	}()

	return ch
}

func (ps *ServiceStorage) readFile(path string) (io.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, wrapper.Wrap(err, "error while open file")
	}

	return file, nil
}
