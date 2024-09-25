protoc --go_out=../webapi/internal/proto --go-grpc_out=../webapi/internal/proto orderservice/order_service.proto

protoc --go_out=../order-service/internal/proto --go-grpc_out=../order-service/internal/proto orderservice/order_service.proto