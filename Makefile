compile_proto:
	protoc -I ./src/proto --go_out=./src/grpc --go-grpc_out=./src/grpc ./src/proto/*.proto

compose_all:
	docker compose up -d --build

down_compose_all:
	docker compose down

compose_subject_service:
	docker compose up --build -d subject-service
