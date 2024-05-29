package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"connectrpc.com/connect"
	greetv1 "github.com/abitofhelp/connect-go-googleapis/gen/greet/v1"
	"github.com/abitofhelp/connect-go-googleapis/gen/greet/v1/greetv1connect"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
		//connect.WithGRPC(), // To use gRPC rather than Connect
	)
	res, err := client.Greet(
		context.Background(),
		connect.NewRequest(&greetv1.GreetRequest{Name: "Jane", Sent: timestamppb.New(time.Now())}),
	)
	if err != nil {
		panic(err)
	}
	log.Println(res.Msg.Greeting)
}
