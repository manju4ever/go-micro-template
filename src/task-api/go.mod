module todo-api

go 1.16

replace taskservice => ../task-service

require (
	github.com/go-micro/plugins/v4/broker/nats v1.1.0
	github.com/go-micro/plugins/v4/registry/nats v1.1.0
	github.com/go-micro/plugins/v4/transport/nats v1.1.0
	github.com/nats-io/nats.go v1.17.0 // indirect
	go-micro.dev/v4 v4.8.1
	golang.org/x/net v0.0.0-20220920203100-d0c6ba3f52d9 // indirect
	taskservice v0.0.0-00010101000000-000000000000
)
