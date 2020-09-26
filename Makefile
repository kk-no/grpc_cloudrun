.PHONY:run/gw
run/gw:
	go run service/gateway/cmd/gateway/main.go

.PHONY:run/cat
run/cat:
	go run service/cat/cmd/cat/main.go

.PHONY: gen
gen: gen/grpc gen/gw

.PHONY: gen/grpc
gen/grpc:
	protoc -I. -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,paths=source_relative:./gen/go proto/v1/cat.proto

.PHONY: gen/gw
gen/gw:
	protoc -I/usr/local/include -I. -I$(GOPATH)/src -I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true,paths=source_relative:./gen/go proto/v1/cat.proto
