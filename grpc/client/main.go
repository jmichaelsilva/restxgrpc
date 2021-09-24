package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	pb "perf_test/proto"

	"google.golang.org/grpc"
)

func main() {

	// Read the entire file into a byte slice
	imgBytes, err := ioutil.ReadFile("./img/IMG_2416.png")
	if err != nil {
		log.Printf("Load image error")
		log.Fatal(err)
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewImageServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SendImage(ctx, &pb.SendImageRequest{Image: imgBytes})
	log.Println(r.Message)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Success")
}
