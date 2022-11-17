GO111MODULE=on
GOPROXY=https://goproxy.cn,direct


go get -u google.golang.org/protobuf

protoc --go_out=. test.proto

====================
https://blog.csdn.net/weixin_41395435/article/details/114868617
