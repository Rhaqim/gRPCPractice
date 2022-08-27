## Update proto file after change by running:

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative checks/checks.proto
```

### Default behavior is to use the current working directory as the root of the generated code.
```bash
protoc --go_out=plugins=grpc:. *.proto
```
