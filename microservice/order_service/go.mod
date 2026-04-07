module order_service

go 1.25.0

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/lib/pq v1.12.3
	google.golang.org/grpc v1.80.0
	shared-proto v0.0.0
   shared v0.0.0
)

require (
	golang.org/x/net v0.51.0 // indirect
	golang.org/x/sys v0.41.0 // indirect
	golang.org/x/text v0.34.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260120221211-b8f7ae30c516 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

replace shared => ../shared
