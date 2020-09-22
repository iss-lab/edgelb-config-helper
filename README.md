# EdgeLB Config Helper

## Build

```
go build -o edgelb-config-helper ./cmd/edgelb-config-helper/...
```

## Run

```
export EDGELB_HOST="<EDGELB-API-SERVER>"
export DNS_HOST="example.com"
export PUBLIC_HOST="example-1.com"
export POOL_FILE="/path/to/pool.json"
./edgelb-config-helper
```

## Development Notes

```
swagger generate client -f swagger.json -A edgelb-config-helper
go mod init
go get -u -f ./...
```