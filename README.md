## Update proto file after change by running:

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative routeguide/cehcking_stuff.proto
```
$ go get -u github.com/golang/protobuf/protoc-gen-go
$ go get -u github.com/golang/protobuf/protoc-gen-go-grpc
$ go get -u github.com/golang/protobuf/protoc-gen-go-plugin
$ go get -u github.com/golang/protobuf/protoc-gen-go-pure

### default behavior is to use the current working directory as the root of the generated code.
```bash
protoc --go_out=plugins=grpc:. *.proto
```