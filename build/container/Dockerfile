# syntax=docker/dockerfile:1

# --- Build Stage ---

  FROM golang:1.23-alpine AS builder

  WORKDIR /app

  COPY go.mod go.sum ./
  RUN go mod download

  COPY . ./

  RUN CGO_ENABLED=0 GOOS=linux \
      go build -o chat-service ./cmd

# --- Packaging Stage ---

  FROM scratch AS packager
  LABEL MAINTAINER="Archisman Mridha <archismanmridha12345@gmail.com>"

  COPY --from=builder /app/build .

  ARG GRPC_SERVER_PORT=4000
  EXPOSE ${GRPC_SERVER_PORT}

  CMD [ "/app/chat-service" ]
