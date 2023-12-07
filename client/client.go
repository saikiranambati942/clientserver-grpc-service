package main

import (
	"context"
	"fmt"
	"grpcapp/proto"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("starting gRPC client application")
	conn, err := grpc.Dial(":9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println(err)
		return
	}

	client := proto.NewCurrencyConversionServiceClient(conn)

	request := &proto.CurrencyConversionRequest{
		From:  "USD",
		To:    "INR",
		Value: 3,
	}
	response, err := client.Convert(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("converted currency value is: ", response.GetConvertedValue())
}
