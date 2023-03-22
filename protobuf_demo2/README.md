# go-rpc-demo

golang grpc demo

## creat TSL cert(option)

install and use [xca](https://github.com/x-ca/go-ca) to create tsl cert.

```
# 生成根证书
xca -create-ca true \
  -root-cert x-ca/ca/root-ca.crt \
  -root-key x-ca/ca/root-ca/private/root-ca.key \
  -tls-cert x-ca/ca/tls-ca.crt \
  -tls-key x-ca/ca/tls-ca/private/tls-ca.key

# 生成 server 证书
xca -cn server \
  --domains "localhost" \
  --ips 127.0.0.1 \
  -tls-cert x-ca/ca/tls-ca.crt \
  -tls-key x-ca/ca/tls-ca/private/tls-ca.key

# 生成 client 证书
xca -cn client \
  --domains "localhost" \
  --ips 127.0.0.1 \
  -tls-cert x-ca/ca/tls-ca.crt \
  -tls-key x-ca/ca/tls-ca/private/tls-ca.key
```

## start server

```
# no tsl
go run server.go

# tsl
go run server.go -ca-crt ./x-ca/ca/root-ca.crt -server-crt ./x-ca/certs/server/server.bundle.crt -server-key ./x-ca/certs/server/server.key
```

## run client

```
# no tsl
go run client.go

# tsl
go run client.go -ca-crt ./x-ca/ca/root-ca.crt -client-crt ./x-ca/certs/client/client.bundle.crt -client-key ./x-ca/certs/client/client.key
```
