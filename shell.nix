{ pkgs ? import <nixpkgs> {} }:
with pkgs; mkShell {
    name = "Go";
    buildInputs = [ go_1_12
                  ];
    shellHook = ''
        if [ $(uname -s) = "Darwin" ]; then
            alias ls='ls --color=auto'
            alias ll='ls -al'
        else
            alias open="xdg-open"
        fi
        if [ ! -d ./pngs/ ]; then
            mkdir pngs
        fi
        export GOPATH=`pwd`
        find -maxdepth 4 -type d | grep "gg" >/dev/null
        if (( ! $? == 0 )); then
            go get -v github.com/fogleman/gg
        fi
        gofmts() {
            gofmt -w -s -e $1
            if (( $? == 0 )); then
                awk '{ gsub(/\t/, "    "); print }' < $1 > tmp
                cat tmp > $1
                rm tmp
            fi
        }
        export -f gofmts
    '';
}
