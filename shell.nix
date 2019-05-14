{ pkgs ? import <nixpkgs> {} }:
with pkgs; mkShell {
    name = "Go";
    buildInputs = [
        go
    ];
    shellHook = ''
        . .shellhook
    '';
}
