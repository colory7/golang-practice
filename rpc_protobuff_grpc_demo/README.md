https://blog.csdn.net/lxw1844912514/article/details/122505293
https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code


protoc --go_out=. message.proto

protoc --go-grpc_out=. message.proto

protoc --go_out=. --go-grpc_out=. message.proto

protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. message.proto

protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. message.proto
