package main

import (
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/bkatrenko/ports/app/portclient/domain/service"
	"github.com/bkatrenko/ports/app/portclient/infrastructure"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const (
	dotFileName = "app.env"

	timeoutEnv    = "PORTS_DOMAIN_SERVICE_TIMEOUT_SEC"
	serverAddrEnv = "PORTS_DOMAIN_SERVICE_ADDR"
	mapPathEnv    = "PORTS_MAP_JSON_PATH"
	endpointEnv   = "REST_ENDPOINT"
)

func main() {
	log.Info("client app started")
	if err := godotenv.Load(dotFileName); err != nil {
		log.Fatal(err)
	}

	domainServiceTimeout, err := strconv.ParseInt(os.Getenv(timeoutEnv), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	repo, err := infrastructure.NewServiceStorage(
		os.Getenv(serverAddrEnv),
		time.Duration(domainServiceTimeout)*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	service, err := service.NewPortsService(os.Getenv(mapPathEnv), repo)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			err := service.PutEvents()
			switch err {
			case nil:
				return
			default:
				log.Errorf("can't put events: can't connect to domain service: %+v", err)
			}

			<-time.After(time.Duration(domainServiceTimeout) * time.Second)
		}
	}()

	handlers, err := infrastructure.NewAPIHandler(service)
	if err != nil {
		log.Fatal(err)
	}

	server := infrastructure.NewServer(os.Getenv(endpointEnv), handlers)

	interruptSignal := make(chan os.Signal)
	done := make(chan bool)
	signal.Notify(interruptSignal, os.Interrupt, syscall.SIGTERM)

	go server.Start()

	go func() {
		<-interruptSignal
		server.Shutdown()
		done <- true
	}()

	<-done
}
