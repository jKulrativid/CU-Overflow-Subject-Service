compile_proto:
	protoc -I ./src/proto --go_out=./src/grpc --go-grpc_out=./src/grpc ./src/proto/*.proto

compose_all:
	docker compose up --build -d
