{ pkgs ? import <nixpkgs> {} }:
with pkgs; mkShell {
    name = "gogr";
    buildInputs = [
        go
    ];
    shellHook = ''
        . .shellhook
    '';
}
