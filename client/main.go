package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "gotokens.com/proto"
)

//command line flags
var (
	host   = flag.String("host", "localhost", "server host")
	port   = flag.String("port", "50051", "server port")
	create = flag.Bool("create", false, "command")
	write  = flag.Bool("write", false, "command")
	read   = flag.Bool("read", false, "command")
	drop   = flag.Bool("drop", false, "command")

	id   = flag.String("id", "", "id")
	name = flag.String("name", "", "name")
	low  = flag.Uint64("low", 0, "low")
	mid  = flag.Uint64("mid", 0, "mid")
	high = flag.Uint64("high", 0, "high")
)

func main() {
	flag.Parse()
	addr := *host + ":" + *port

	// Set up a connection to the server.
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	// we can then call to the generated proto library
	c := pb.NewTokenManagerClient(conn)
	// Connect to the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if *create {
		r, err := c.Create(ctx, &pb.CreateRequest{Id: *id})
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		log.Printf("message: %s", r.GetMessage())
	} else if *write {
		r, err := c.Write(ctx, &pb.WriteRequest{Id: *id, Name: *name, Low: *low, Mid: *mid, High: *high})
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		log.Printf("message: %s, partial: %d", r.GetMessage(), r.GetPartial())
	} else if *read {
		r, err := c.Read(ctx, &pb.ReadRequest{Id: *id})
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		log.Printf("message: %s, final: %d", r.GetMessage(), r.GetFinal())
	} else if *drop {
		r, err := c.Drop(ctx, &pb.DropRequest{Id: *id})
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		log.Printf("message: %s", r.GetMessage())
	} else {
		log.Fatalf("unknown command")
	}

}
