{
  description = "Go 1.24 development shell with treefmt-nix";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/21808d22b1cda1898b71cf1a1beb524a97add2c4";
    flake-parts.url = "github:hercules-ci/flake-parts";
    treefmt-nix.url = "github:numtide/treefmt-nix";
    flake-root.url = "github:srid/flake-root";
  };

  outputs = inputs@{ self, nixpkgs, flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];

      imports = [
        inputs.treefmt-nix.flakeModule
        inputs.flake-root.flakeModule
      ];

      perSystem = { config, self', pkgs, ... }: {
        treefmt = {
          projectRootFile = "flake.nix";
          programs.gofumpt.enable = true;
          programs.gofmt.enable = true;
        };

        devShells.default = pkgs.mkShell {
          buildInputs = [
            pkgs.go_1_24
            config.treefmt.build.wrapper
            pkgs.just
          ];
        };
      };
    };
}
