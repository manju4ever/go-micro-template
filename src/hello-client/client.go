package main

import (
	"context"
	"fmt"
	natsbroker "github.com/go-micro/plugins/v4/broker/nats"
	natsregistry "github.com/go-micro/plugins/v4/registry/nats"
	natsport "github.com/go-micro/plugins/v4/transport/nats"
	micro "go-micro.dev/v4"
	schema "helloworld/proto"
)

func main() {
	service := micro.NewService(
		micro.Registry(natsregistry.NewRegistry()),
		micro.Broker(natsbroker.NewBroker()),
		micro.Transport(natsport.NewTransport()),
	)
	service.Init()
	client := service.Client()
	req := client.NewRequest("hello-world", "HelloWorld.Call", &schema.CallRequest{Name: "Manju"})
	res := &schema.CallResponse{}
	ctx := context.Background()
	if err := client.Call(ctx, req, res); err != nil {
		fmt.Println("Something is wrong:", err, res)
		return
	}
	fmt.Println(res)
}
