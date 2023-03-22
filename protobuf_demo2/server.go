/*
Copyright [2022] [xiexianbin.cn]

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	googlerpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	dgrpc "github.com/xiexianbin/go-rpc-demo/grpc"
	dgrpcserver "github.com/xiexianbin/go-rpc-demo/grpc/server"
)

var server = &dgrpcserver.DemoServiceServer{}
var (
	h             bool
	rootCACrtPath string
	serverCrtPath string
	serverKeyPath string
)

func init() {
	flag.BoolVar(&h, "help", false, "show help message")
	flag.StringVar(&rootCACrtPath, "ca-crt", "", "ca crt file path")
	flag.StringVar(&serverCrtPath, "server-crt", "", "server crt file path")
	flag.StringVar(&serverKeyPath, "server-key", "", "server key file path")

	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()
}

func loadServerTSLCert() (credentials.TransportCredentials, error) {
	caPEMFile, err := ioutil.ReadFile(rootCACrtPath)
	if err != nil {
		return nil, err
	}

	caPool := x509.NewCertPool()
	if !caPool.AppendCertsFromPEM(caPEMFile) {
		return nil, fmt.Errorf("load %s cert fail", rootCACrtPath)
	}

	localCert, err := tls.LoadX509KeyPair(serverCrtPath, serverKeyPath)
	if err != nil {
		return nil, fmt.Errorf("load server cert and key file fail: %s", err.Error())
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{localCert},
		ClientAuth:   tls.RequireAndVerifyClientCert, // check client's certificate
		ClientCAs:    caPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	if h == true {
		flag.Usage()
		return
	}

	// Listener
	addr := ":8000"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatalf(err.Error())
		}
		log.Println("grpc server closed.")
	}(listener)

	var s *googlerpc.Server
	if rootCACrtPath != "" && serverCrtPath != "" && serverKeyPath != "" {
		tlsCrt, err := loadServerTSLCert()
		if err != nil {
			log.Fatalf("load server cert err: %s", err.Error())
		}
		s = googlerpc.NewServer(googlerpc.Creds(tlsCrt))
	} else {
		s = googlerpc.NewServer()
	}
	dgrpc.RegisterServiceServer(s, server)
	log.Println("grpc server listen on", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to start grpc server: %v", err)
	}
}
