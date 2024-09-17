{
  description = "SDK common properties";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
        corepack = pkgs.stdenv.mkDerivation {
          name = "corepack";
          buildInputs = [ pkgs.nodejs_20 ];
          phases = [ "installPhase" ];
          installPhase = ''
            mkdir -p $out/bin
            corepack enable --install-directory=$out/bin
          '';
        };

      in {
        devShells.default = pkgs.mkShell {
          packages = [ corepack ];

          nativeBuildInputs = with pkgs; [
            # Go
            # go
            # golangci-lint
            # gotests
            # gomodifytags
            # gore
            # gotools
            # goreleaser

            # Node
            typescript
            nodejs_20

            # Tools
            buf
            protoc-gen-go
            protoc-gen-connect-go
          ];
        };
      });
}
