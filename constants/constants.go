package constants

import (
	"time"

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
	// the Pod name.
	InstanceID = utils.GetEnv("POD_NAME")
)
