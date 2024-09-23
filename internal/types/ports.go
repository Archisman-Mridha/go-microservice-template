package types

import "github.com/Archisman-Mridha/chat-service/pkg/healthcheck"

type (
	KVStore interface {
		healthcheck.Healthcheckable
	}

	Database interface {
		healthcheck.Healthcheckable
	}
)
