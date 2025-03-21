package main

import (
	"github.com/archisman-mridha/go-microservice-template/deploy/cue.mod/gen/compose:compose"
)

compose.#Project & {
	name:    "go-microservice-template"

	networks: "go-microservice-template": {
		name: "go-microservice-template"
	}

	services: {}
}

