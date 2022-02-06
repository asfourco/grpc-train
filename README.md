# gRPC Demo: Train Booking Service

## Description

## Usage

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