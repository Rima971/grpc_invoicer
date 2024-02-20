package main

import (
	"context"
	"github.com/rima971/first-grpc/invoicer"
	"google.golang.org/grpc"
	"log"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	i := invoicer.NewInvoicerClient(conn)

	createReq := &invoicer.CreateRequest{
		Amount: &invoicer.Amount{
			Amount:   20,
			Currency: "INR",
		},
		From: "rima",
		To:   "poddar",
	}

	res, err := i.Create(context.Background(), createReq)

	if err != nil {
		log.Fatalf("some error occurred at the server while calling create: %s", err)
	}
	log.Printf("response from the body: %s", res)
}
