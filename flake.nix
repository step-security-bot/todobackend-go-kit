{
  description = "A simple Todo-Backend application written using Go kit";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    goflake.url = "github:sagikazarmark/go-flake";
    goflake.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs = { self, nixpkgs, flake-utils, goflake, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;

          overlays = [
            goflake.overlay
          ];
        };
      in
      rec
      {
        devShells = {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [
              git
              gnumake

              go_1_20
              golangci-lint
              gotestsum

              buf
              protobuf
              protoc-gen-go
              protoc-gen-go-grpc
              protoc-gen-go-kit
              openapi-generator-cli

              regctl
              skopeo
              syft
            ];

            shellHook = ''
              go version
              golangci-lint --version
              gotestsum --version
              protoc --version
              openapi-generator-cli --version | head -1
            '';
          };

          ci = devShells.default.overrideAttrs (final: prev: {
            buildInputs = prev.buildInputs ++ (with pkgs; [
              flyctl

              cosign
              syft
            ]);

            shellHook = ''
              ${prev.shellHook}
              flyctl version
            '';
          });
        };
      }
    );
}
