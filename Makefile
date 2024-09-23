buf-generate:
	cd ./api/proto && buf generate

sqlc-generate:
	cd ./pkg/postgres/sql && sqlc generate

build-binary:
	go build -o build/generated/chat-service ./cmd/server/main.go

build-container:
	docker build -f ./build/container/Dockerfile -t archismanmridha:chat-service .

build-wasm:
	GOOS=wasip1 GOARCH=wasm go build -o build/generated/chat-service.wasm cmd/server/main.go
	wasmedge compile build/generated/chat-service.wasm build/generated/chat-service.aot.wasm

compose-up:
	docker-compose -f ./build/compose/compose.yaml --project-directory . up -d

compose-down:
	docker-compose -f ./build/compose/compose.yaml --project-directory . down

query-analyzer-up:
	docker-compose -f ./build/compose/query-analyzer.compose.yaml up -d

query-analyzer-down:
	docker-compose -f ./build/compose/query-analyzer.compose.yaml down

run:
	go run ./cmd --config-file=./config/sample.config.yaml

lint:
	golangci-lint run ./...
