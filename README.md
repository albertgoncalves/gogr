# gogr

![](cover.png)

Needed things
---
  * [Nix](https://nixos.org/nix/)

Quick start
---
```
$ nix-shell
[nix-shell:path/to/gogr]$ ./test spline
[nix-shell:path/to/gogr]$ ./test spline prof
[nix-shell:path/to/gogr]$ go build -o bin/main main.go && ./bin/main && open out/main.png
```
