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
            go
            gopls

            # Node
            typescript
            nodejs_20

            # Python
            (python311.withPackages (p: with p; [ pip virtualenv ipython ]))
            pyright

            # Tools
            nodePackages.yaml-language-server
            buf
            protobuf
            protoc-gen-go
          ];

          shellHook = ''
            export PY_COLORS="1"
            export PATH="$PWD/node_modules/.bin:$PATH"

            [ ! -d .venv ] && python -m venv .venv
            source .venv/bin/activate
          '';
        };
      });
}
