{ pkgs ? import <nixpkgs> {} }:
with pkgs; mkShell {
    name = "Go";
    buildInputs = [ go_1_12
                  ];
    shellHook = ''
        lint() {
            gofmt -w -s -e $1
            if (( $? == 0 )); then
                awk '{ gsub(/\t/, "    "); print }' < $1 > tmp
                cat tmp > $1
                rm tmp
            fi
        }
        export -f lint
        export GOPATH=$(pwd)
    '';
}
