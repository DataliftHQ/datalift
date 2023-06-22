module go.datalift.io/datalift

go 1.20

replace (
	go.datalift.io/datalift/api => ./api
	//go.datalift.io/datalift/client => ./client
	go.datalift.io/datalift/migrate => ./migrate
	go.datalift.io/datalift/server => ./server
	go.datalift.io/datalift/worker => ./worker
)

require (
	github.com/pkg/errors v0.9.1
	go.datalift.io/datalift/api v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.56.1
)

require (
	github.com/envoyproxy/protoc-gen-validate v1.0.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230526203410-71b5a4ffd15e // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)
