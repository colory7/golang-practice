https://www.jianshu.com/p/6cb5293aed03

go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

mkdir gen
protoc --go_out=./gen --go_opt=paths=source_relative --go-grpc_out=./gen --go-grpc_opt=paths=source_relative hello.proto

或者 protoc --go_out=. --go-grpc_out=. hello.proto

