protoc --go_out=../webapi/internal/proto --go-grpc_out=../webapi/internal/proto authservice/auth_service.proto

protoc --go_out=../auth-service/internal/proto --go-grpc_out=../auth-service/internal/proto authservice/auth_service.proto