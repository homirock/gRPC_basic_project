package main

import (
	"context"
	"fmt"

	//"log"
	"net"

	info "github.com/homirock/gRPC_basic_project/info"
	"google.golang.org/grpc"
)

type myinfoServer struct {
	info.UnimplementedInfoServer
}

func (m *myinfoServer) Create(context.Context, *info.CreateRequest) (*info.CreateResponse, error) {
	return &info.CreateResponse{
		Pdf:  []byte("test123123"),
		Docx: []byte("test111123"),
	}, nil
}

func main() {
	//go func(){
	fmt.Println("Starting server")
	lis, _ := net.Listen("tcp", ":8089")
	server_register := grpc.NewServer()
	myservice := myinfoServer{}
	info.RegisterInfoServer(server_register, &myservice)
	fmt.Println("registration done")
	err := server_register.Serve(lis)
	if err != nil {
		fmt.Println("errorr")
	}
	// }()
	// conn, err := grpc.Dial("localhost:8089", grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("Failed to dial: %v", err)
	// }
	// defer conn.Close()

	// client := info.NewInfoClient(conn)

	// //Call the Create method on the server
	// resp, err := client.Create(context.Background(), &info.CreateRequest{})
	// if err != nil {
	// 	log.Fatalf("Create request failed: %v", err)
	// }

	// Print the response
	
	// fmt.Println("PDF:", resp.Pdf)
	// fmt.Println("DOCX:", resp.Docx)
	//fmt.Println(resp)

}
