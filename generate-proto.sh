cd ./proto
protoc entity.proto --go-grpc_out=../entity --go_out=../entity --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative
protoc comm.proto --go-grpc_out=../comm --go_out=../comm --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative
