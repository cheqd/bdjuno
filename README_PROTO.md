# Proto File Generation for Gogoproto Support

## Generating Proto Files

1. Install required tools:
```bash
go install github.com/gogo/protobuf/protoc-gen-gogofaster@latest
```

2. Generate the Go code:
```bash
make -f Makefile.proto proto-gen
```

Or manually:
```bash
protoc \
  --proto_path=proto \
  --proto_path=$(go env GOMODCACHE)/github.com/cosmos/gogoproto@v1.7.0 \
  --proto_path=$(go env GOMODCACHE)/github.com/cosmos/cosmos-proto@v1.0.0-beta.5 \
  --gogoproto_out=plugins=grpc,Mgogoproto/gogo.proto=github.com/cosmos/gogoproto/gogoproto:proto/gen \
  proto/cosmos/mint/v1beta1/query.proto \
  proto/cosmos/mint/v1beta1/mint.proto
```

3. Update the gogoproto_client.go to use the generated types from `proto/gen/cosmos/mint/v1beta1/` instead of `github.com/cosmos/cosmos-sdk/x/mint/types`.


## Important: Preventing Duplicate Proto Type Registration

After generating the proto files, you **must** comment out the `init()` functions in the generated `.pb.go` files to prevent duplicate proto type registration errors. The cosmos-sdk proto types are already registered, and our generated types use the same package name.

## Testing

After generating and editing the proto files, rebuild and test:
```bash
go build ./...
```


