{ pkgs ? import <nixpkgs> {} }:
with pkgs; mkShell {
    name = "gogr";
    buildInputs = [
        go
        shellcheck
    ];
    shellHook = ''
        . .shellhook
    '';
}
