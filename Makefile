ROOTDIR := $(shell pwd)

proto-api:
	@echo "--> Generating gRPC clients for API"
	@rm -rf $(ROODIR)/api-gw/api/goclient/*
	@protoc --proto_path=api-gw/api \
        --go_out=paths=source_relative:api-gw/api/goclient \
        --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:api-gw/api/goclient \
        api-gw/api/v1/*.proto
	@echo "Done"

proto-passengers:
	@echo "--> Generating gRPC clients for Passengers"
	@rm -rf $(ROODIR)/passengers/api/goclient/*
	@protoc --proto_path=passengers/api \
        --go_out=paths=source_relative:passengers/api/goclient \
        --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:passengers/api/goclient \
        passengers/api/v1/*.proto
	@echo "Done"

proto-trains:
	@echo "--> Generating gRPC clients for Trains"
	@rm -rf $(ROODIR)/trains/api/goclient/*
	@protoc --proto_path=trains/api \
        --go_out=paths=source_relative:trains/api/goclient \
        --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:trains/api/goclient \
        trains/api/v1/*.proto
	@echo "Done"


proto: proto-api proto-passengers proto-trains

build:
	mkdir -p ./dist
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/apigw ./api-gw
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/passengers ./passengers
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/trains ./trains
	docker-compose build

clean:
	rm -rf ./dist

run-servers:
	@echo "--> Starting servers"
	@docker-compose up