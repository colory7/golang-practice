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

package server

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/types/known/emptypb"

	dgrpc "github.com/xiexianbin/go-rpc-demo/grpc"
)

type ServerError struct {
	message string
}

func (e *ServerError) Error() string {
	return e.message
}

type DemoServiceServer struct {
	dgrpc.UnimplementedServiceServer
}

func (s *DemoServiceServer) Sum(ctx context.Context, numRequest *dgrpc.NumRequest) (*dgrpc.NumResponse, error) {
	numResponse := &dgrpc.NumResponse{
		Result: numRequest.Nums[0] + numRequest.Nums[1],
	}
	return numResponse, nil
}

func (s *DemoServiceServer) Diff(ctx context.Context, numRequest *dgrpc.NumRequest) (*dgrpc.NumResponse, error) {
	numResponse := &dgrpc.NumResponse{
		Result: numRequest.Nums[0] - numRequest.Nums[1],
	}
	return numResponse, nil
}

func (s *DemoServiceServer) Version(ctx context.Context, empty *emptypb.Empty) (*dgrpc.VersionResponse, error) {
	version := &dgrpc.VersionResponse{
		Version: "v0.1.0",
	}
	return version, nil
}

func (s *DemoServiceServer) ReadFile(ctx context.Context, filePath *dgrpc.FilePath) (*dgrpc.FileResponse, error) {
	_, err := os.Stat(filePath.GetPath())
	if err != nil {
		// os.IsNotExist(err)
		message := fmt.Sprintf("file path %s not exist", filePath.GetPath())
		log.Printf(message)
		return nil, &ServerError{message: message}
	}

	fileContent, err := os.ReadFile(filePath.GetPath())
	if err != nil {
		message := fmt.Sprintf("read file error %v", err)
		log.Printf(message)
		return nil, &ServerError{message: message}
	}

	return &dgrpc.FileResponse{Content: fileContent}, nil
}
