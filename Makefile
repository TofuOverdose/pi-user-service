start-dev:
	docker-compose --env-file=dev.env up -d

stop-dev:
	docker-compose --env-file=dev.env down

run-tests:
	./scripts/run-tests.sh

proto-gen:
	protoc -I=./api/proto \
		--go_out=internal/users/ports/grpc/protogen --go_opt=paths=source_relative \
		--go-grpc_out=internal/users/ports/grpc/protogen --go-grpc_opt=paths=source_relative \
		api/proto/users.proto
