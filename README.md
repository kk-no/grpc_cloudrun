# gRPC tests on Cloud Run

### gRPC generate
```
protoc -I. -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,paths=source_relative:./gen/go proto/v1/cat.proto
```
### gateway generate
```
protoc -I/usr/local/include -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true,paths=source_relative:./gen/go proto/v1/cat.proto
```

### test request
```
go run main.go
curl localhost:8080/cat/get
```