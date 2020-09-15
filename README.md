# EdgeLB Config Helper

## Build

```
go build -o edgelb-config-helper ./cmd/edgelb-config-helper/...
```

## Run

```
export EDGELB_HOST="<EDGELB-API-SERVER>"
./edgelb-config-helper
```

## Development Notes

```
swagger generate client -f swagger.json -A edgelb-config-helper
go mod init
go get -u -f ./...
```