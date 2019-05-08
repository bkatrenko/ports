package main

import (
	"os"
	"strconv"
	"time"

	"github.com/bkatrenko/ports/app/portdomain/domain/service"
	"github.com/bkatrenko/ports/app/portdomain/infrastructure"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

const (
	dotFileName = "app.env"

	mongoTimeoutEnv = "MONGO_TIMEOUT_SEC"
	mongoAddrEnv    = "MONGO_ADDR"
	serverAddrEnv   = "SERVER_ADDR"
)

func main() {
	log.Info("domain app started")

	if err := godotenv.Load(dotFileName); err != nil {
		log.Fatal(err)
	}

	mongoTimeout, err := strconv.ParseInt(os.Getenv(mongoTimeoutEnv), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	storage, err := infrastructure.NewMongoStorage(os.Getenv(mongoAddrEnv), time.Duration(mongoTimeout)*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	service, err := service.NewPortsService(storage)
	if err != nil {
		log.Fatal(err)
	}

	infrastructure.GRPC(service, os.Getenv(serverAddrEnv))
}
