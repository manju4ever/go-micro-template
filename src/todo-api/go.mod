module todo-api

go 1.16

require (
	github.com/go-micro/plugins/v4/broker/nats v1.1.0
	github.com/go-micro/plugins/v4/registry/nats v1.1.0
	github.com/go-micro/plugins/v4/transport/nats v1.1.0
	go-micro.dev/v4 v4.8.1
	helloworld v0.0.0-00010101000000-000000000000
)

replace helloworld => ../hello-world
