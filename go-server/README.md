# COSCUP 2017 golang gRPC server demo

To use gRPC you will have to install `grpc`, `protobuf` and `protoc-gen-go`.

* install grpc
    * `go get google.golang.org/grpc`
* install [protocol buffers](https://github.com/google/protobuf/) v3.x
* install the go compiler plugin protoc-gen-go
    * `go get -u github.com/golang/protobuf/protoc-gen-go`
    * `protoc-gen-go` is a compiler to generate Go code from a proto file.


Use this command to generate golang source code of gRPC:

* Windows
```
protoc -I .\proto\ --go_out=plugins=grpc:proto .\proto\*.proto
```

* Linux/Mac
```
protoc -I ./proto/ --go_out=plugins=grpc:proto ./proto/*.proto
```