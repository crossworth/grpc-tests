package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	"grpc-tests/myservice"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := myservice.NewMyServiceClient(conn)

	req := new(myservice.DoSomethingRequest)
	req.Name = "GoTest"

	resp, err := client.DoSomething(context.TODO(), req)
	if err != nil {
		log.Fatalln("client.DoSomething error", err)
	}

	log.Println(resp.GetMessage())
}
