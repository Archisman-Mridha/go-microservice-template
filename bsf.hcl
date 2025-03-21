packages {
  // Packages that we need in the runtime image i.e the final artifact is shipped to production.
  runtime = []
}

// Parameters to build the Golang module.
gomodule {
  // Name of the binary.
  name = "go-microservice-template"

  // Path to source code. It is relative to root of the project i.e the same place as bsf.hcl
  // resides.
  src = "./."

  // If unit tests should be run on builds
  doCheck = false
}

oci "default" { // "default" represents the environment.
  name          = "go-microservice-template"
  layers        = ["packages.runtime"]
  isBase        = true
  cmd           = []
  entrypoint    = []
  envVars       = []
  exposedPorts  = []
  importConfigs = []
}
