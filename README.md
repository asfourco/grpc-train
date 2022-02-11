# gRPC Demo: Train Booking Service

## Description

This is a train booking app where a passenger can book a ticket on a train; however, they cannot book a train that is travelling, leaving, or arriving at the same time as the train they are trying to book. The architecture of this app is such that there is one api gateway to ties the rest of the services together.

## Usage

To compile and run the docker services:

```console
make clean
make proto
make build
docker-compose
```

## Contributing

Development Tools:

- [pre-commit](https://pre-commit.com/index.html#install)
- [golangci-lint](https://golangci-lint.run/usage/install/)

### compiling protocol buffer files

```
cd ${service}
protoc -I api --go_out=paths=source_relative:api/goclient --go-grpc_out=paths=source_relative:api/goclient api/v1/${service_name}.proto
```

## Requirements


## Credits

The code was modelled after https://github.com/kostyay/grpc-api-gateway-example