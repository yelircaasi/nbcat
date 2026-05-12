{
  description = "Nix flake providing the nbcat CLI, along with a compatible development shell.";
  inputs = {
    nixpkgs = {
      url = "github:nixos/nixpkgs/a77b87719ff11d3d7ea2439c406361b6b0c4c56a"; # 2025-10-01
    };

    flake-utils = {
      url = "github:numtide/flake-utils/11707dc2f618dd54ca8739b309ec4fc024de578b"; # 2025-10-01
    };
  };
  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {
        inherit system;
        # config.permittedInsecurePackages = ["olm-3.2.16"];
      };

      packageName = "nbcat";

      package = pkgs.buildGoModule {
        pname = packageName;
        version = "0.1.0"; # Update this as you version your tool

        src = ./.;
        # vendorHash is required for reproducible builds.
        # If you don't know it, leave it as pkgs.lib.fakeHash,
        # run the build, and Nix will tell you the correct hash.
        # vendorHash = pkgs.lib.fakeHash;
        vendorHash = "sha256-wxfkEg8SfVHIj5xaEJnqaxXJg6l1/mIqniSLw4pPpIw=";

        meta = with pkgs.lib; {
          description = "CLI tool to display the contents of a Jupyter notebook.";
          homepage = "https://github.com/yelircaasi/nbcat";
          license = licenses.gpl3;
          maintainers = [];
        };
      };

      devEnv = pkgs.mkShell {
        # nativeBuildInputs is usually what you want -- tools you need to run
        nativeBuildInputs = with pkgs; [
          # nixpkgs-fmt
          # rnix-lsp
          # docker-client
          # gnumake
          bash

          # go development
          go
          # go-outline
          # gopls
          # gopkgs
          # go-tools
          # delve
          cobra-cli
        ];

        hardeningDisable = ["all"];
      };
    in {
      packages = {
        default = package;
        "${packageName}" = package;
      };

      apps = {
        default = {
          type = "app";
          program = "${package}/bin/${packageName}";
        };
        ${packageName} = package;
      };

      devShells = rec {
        default = devEnv;
        "${packageName}" = devEnv;
      };

      legacyPackages = pkgs;
    });
}
