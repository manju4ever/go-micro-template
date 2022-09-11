module hello-world

go 1.18

require (
	github.com/go-micro/plugins/v4/broker/nats v1.1.0
	github.com/go-micro/plugins/v4/registry/nats v1.1.0
	github.com/go-micro/plugins/v4/transport/nats v1.1.0
	go-micro.dev/v4 v4.8.0
	google.golang.org/protobuf v1.26.0
)

// Uncomment if you use etcd
// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
// replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace hello-world => ./
