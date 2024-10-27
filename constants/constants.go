package constants

import (
	"errors"
	"time"

	"github.com/Archisman-Mridha/chat-service/api/proto/generated"
	"github.com/Archisman-Mridha/chat-service/pkg/utils"
)

const (
	SERVICE_NAME    = "chats-microservice"
	SERVICE_VERSION = "0.0.1"

	HEALTHCHECK_FREQUENCY           = 5 * time.Second
	RESOURCE_CLEANUP_TIMEOUT        = time.Second
	TRACE_EXPORTER_SHUTDOWN_TIMEOUT = time.Second

	UNIMPLEMENTED = "unimplemented"
)

// Global Variables.
var (
	// The instanceID is provided by the orchestrator. For example, in case of Kubernetes, it'll be
	// the Pod IP.
	InstanceID = utils.GetEnv("POD_IP")

	// userID -> gRPC Chat RPC bi-directional stream.
	UserIDToStream map[int32]generated.ChatService_ChatServer
)

// Errors.
var (
	ErrServer = errors.New("internal server error")

	ErrHandshakeFailed  = errors.New("handshake failed")
	ErrHandshakeNotDone = errors.New("handshake not done")

	ErrUserNotConnectedToServer = errors.New("user isn't connected to this server")
)
