package main

import (
	"GolangGRPC/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type myGRPCServer struct {
}

// actual implementation of the Convert method
func (s *myGRPCServer) Convert(ctx context.Context, request *proto.CurrencyConversionRequest) (*proto.CurrencyConversionResponse, error) {
	fmt.Printf("request received to convert currency %+v \n", request)
	return &proto.CurrencyConversionResponse{
		ConvertedValue: 74.02*request.Value,
	}, nil
}

func main() {
	fmt.Println("Starting gRPC Server application")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println(err)
		return
	}
	server := grpc.NewServer()
	proto.RegisterCurrencyConversionServiceServer(server, &myGRPCServer{})
	err = server.Serve(lis)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
