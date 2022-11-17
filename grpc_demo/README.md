https://www.jianshu.com/p/6cb5293aed03

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/hello.proto

protoc --go_out=.  --go-grpc_out=. hello.proto

