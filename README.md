# finance-protobuf

## inside

```
PATH=$PATH:~/go/bin protoc --go_out=paths=source_relative:. inside/inside.proto
```

## porter

```
PATH=$PATH:~/go/bin protoc -I. -I~/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. porter/porter.proto

PATH=$PATH:~/go/bin protoc -I. -I~/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. porter/porter.proto
```
