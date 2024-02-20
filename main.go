package main

import (
	"context"
	"github.com/rima971/first-grpc/invoicer"
	"google.golang.org/grpc"
	"log"
	"net"
)

type myInvoiceServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (m myInvoiceServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	res := &invoicer.CreateResponse{
		Invoice: &invoicer.Invoice{Amount: req.Amount, To: req.To, From: req.From},
		Message: "invoice generated",
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := &myInvoiceServer{}

	invoicer.RegisterInvoicerServer(serverRegistrar, service)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
