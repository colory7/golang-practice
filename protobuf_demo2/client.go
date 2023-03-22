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
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	googlerpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"

	dgrpc "github.com/xiexianbin/go-rpc-demo/grpc"
)

var (
	ch             bool
	rootCACrtPath2 string
	clientCrtPath  string
	clientKeyPath  string
)

func init() {
	flag.BoolVar(&ch, "help", false, "show help message")
	flag.StringVar(&rootCACrtPath2, "ca-crt", "", "ca crt file path")
	flag.StringVar(&clientCrtPath, "client-crt", "", "client crt file path")
	flag.StringVar(&clientKeyPath, "client-key", "", "client key file path")

	flag.Usage = func() {
		flag.PrintDefaults()
	}
	flag.Parse()
}

func loadClientTSLCert() (credentials.TransportCredentials, error) {
	caPEMFile, err := ioutil.ReadFile(rootCACrtPath2)
	if err != nil {
		return nil, err
	}

	caPool := x509.NewCertPool()
	if !caPool.AppendCertsFromPEM(caPEMFile) {
		return nil, fmt.Errorf("load %s cert fail", rootCACrtPath2)
	}

	localCert, err := tls.LoadX509KeyPair(clientCrtPath, clientKeyPath)
	if err != nil {
		return nil, fmt.Errorf("load client cert and key file fail: %s", err.Error())
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{localCert},
		ServerName:   "localhost", // client cn name
		RootCAs:      caPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	if ch == true {
		flag.Usage()
		return
	}

	var creds credentials.TransportCredentials
	var err error
	if rootCACrtPath2 != "" && clientCrtPath != "" && clientKeyPath != "" {
		creds, err = loadClientTSLCert()
		if err != nil {
			log.Fatalf("load client cert err: %s", err.Error())
		}
	} else {
		creds = insecure.NewCredentials()
	}
	cc, err := googlerpc.Dial("127.0.0.1:8000", googlerpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("cannot dial server %v", err)
	}
	defer cc.Close()

	rpcClient := dgrpc.NewServiceClient(cc)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// version
	version, err := rpcClient.Version(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("error happen when call gRPC client: %s", err.Error())
	}
	log.Printf("version: %s", version)

	// sum
	nums := &dgrpc.NumRequest{
		Nums: []int64{1, 2},
	}
	sum, err := rpcClient.Sum(ctx, nums)
	if err != nil {
		log.Fatalf("error happen when call gRPC client: %s", err.Error())
	}
	log.Printf("sum: %s", sum)

	// diff
	diff, err := rpcClient.Diff(ctx, nums)
	if err != nil {
		log.Fatalf("error happen when call gRPC client: %s", err.Error())
	}
	log.Printf("diff: %s", diff)

	// read file
	filePath := &dgrpc.FilePath{
		Path: "/etc/hosts",
	}
	fileContent, err := rpcClient.ReadFile(ctx, filePath)
	if err != nil {
		log.Fatalf("error happen when call gRPC client: %s", err.Error())
	}
	log.Printf("fileContent: %s", fileContent)
}
