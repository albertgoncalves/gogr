with import <nixpkgs> {};
mkShell {
    buildInputs = [
        go
        shellcheck
    ];
    shellHook = ''
        . .shellhook
    '';
}
