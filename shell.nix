{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
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

  hardeningDisable = [ "all" ];
}
