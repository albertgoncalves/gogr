{ pkgs ? import <nixpkgs> {} }:
with pkgs; mkShell {
    name = "Go";
    buildInputs = [
        go_1_12
    ];
    shellHook = ''
        . .shellhook
    '';
}
