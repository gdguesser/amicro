proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. types/types.proto