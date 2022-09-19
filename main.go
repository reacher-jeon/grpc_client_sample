package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	sampleV1 "grpc_client_sample/pb/sample/v1"
	typeV1 "grpc_client_sample/pb/type/v1"
	"time"
)

const (
	sampleServerHost = "localhost"
	sampleServerPort = 5000
)

func main() {
	ctx := context.Background()
	ctxTimeout, _ := context.WithTimeout(ctx, time.Second*5)

	fmt.Println("Connect gRPC Sample Server : ", sampleServerHost, sampleServerPort)
	sampleServerAddress := fmt.Sprintf("%v:%v", sampleServerHost, sampleServerPort)
	var err error

	sampleClientConn, err := grpc.DialContext(ctxTimeout, sampleServerAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Failed to dial gRPC Sample Server:", err)
	}
	sampleClient := sampleV1.NewSampleServiceClient(sampleClientConn)

	request := &sampleV1.GetInfoInfoRequest{
		Header: &typeV1.Header{
			Version:     "1",
			ToIds:       []string{"11111", "22222"},
			FromId:      1,
			RequesterId: "",
		},
		Info: &sampleV1.GetInfo{
			Id: "hello",
		},
	}

	resMessage, er := sampleClient.GetInfo(ctxTimeout, request)
	if err != nil {
		fmt.Println("GetInfo Error:", er)
	} else {
		fmt.Println("GetInfo Data:", resMessage)
	}
}
