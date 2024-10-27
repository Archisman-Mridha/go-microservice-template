.PHONY: buf-generate
buf-generate:
	@cd ./api/proto && buf generate

.PHONY: sqlc-generate
sqlc-generate:
	@cd ./pkg/postgres/sql && sqlc generate

.PHONY: build-binary
build-binary:
	@go build -o build/generated/chat-service ./cmd/server/main.go

.PHONY: build-container
build-container:
	@docker build -f ./build/container/Dockerfile -t archismanmridha:chat-service .

.PHONY: build-wasm
build-wasm:
	GOOS=wasip1 GOARCH=wasm go build -o build/generated/chat-service.wasm cmd/server/main.go
	@wasmedge compile build/generated/chat-service.wasm build/generated/chat-service.aot.wasm

.PHONY: compose-up
compose-up:
	@docker-compose -f ./build/compose/compose.yaml --project-directory . up -d

.PHONY: compose-down
compose-down:
	@docker-compose -f ./build/compose/compose.yaml --project-directory . down

.PHONY: query-analyzer-up
query-analyzer-up:
	@docker-compose -f ./build/compose/query-analyzer.compose.yaml up -d

.PHONY: query-analyzer-down
query-analyzer-down:
	@docker-compose -f ./build/compose/query-analyzer.compose.yaml down

.PHONY: run
run:
	@go run ./cmd --config-file=./config/sample.config.yaml

.PHONY: lint
lint:
	@golangci-lint run ./...
