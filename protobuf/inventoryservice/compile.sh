protoc --go_out=../webapi/internal/proto --go-grpc_out=../webapi/internal/proto inventoryservice/inventory_service.proto

protoc --go_out=../inventory-service/internal/proto --go-grpc_out=../inventory-service/internal/proto inventoryservice/inventory_service.proto

protoc --go_out=../order-service/internal/proto --go-grpc_out=../order-service/internal/proto inventoryservice/inventory_service.proto