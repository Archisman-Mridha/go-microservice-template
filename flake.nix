{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    buildsafe.url = "github:buildsafedev/bsf";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      buildsafe,
      ...
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
          config.allowUnfree = true;
        };
      in
      with pkgs;
      {
        devShells.default = mkShell {
          buildInputs = [
            go_1_24
            golangci-lint
            golines

            protobuf
            buf

            sqlc

            buildsafe.packages.${system}.default
            cue
            docker-compose
            k3d
            kubectl

            go-task
          ];
        };
      }
    );
}
