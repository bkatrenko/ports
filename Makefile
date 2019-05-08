.PHONY: dependency unit-test integration-test docker-up docker-down clear 

build:
	@go build

integration-test: docker-up dependency
	sleep 5
	@go test -v ./testing
	@go test -v ./app/portclient/infrastructure


docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down

clear: docker-down