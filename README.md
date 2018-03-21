# go-monorepo

A _PoC_ about organizing the code in a **monorepo** fashion using golang **dep**.

This code organization assumes the root is not vendored, only the leafs.

The directories can:

* depend on one (or more) siblings
* have transitive deps
* implement executables
* implement libraries
* use different tags of same package as dependency

    * see [uuid@1.2.0](https://github.com/leodido/go-monorepo/blob/develop/amenadiel/Gopkg.lock#L18)
    * see [uuid@master](https://github.com/leodido/go-monorepo/blob/develop/mazekeen/Gopkg.lock#L17)

#### Disclaimer

The name of packages and directories are simply a bias given to me from the last TV series - ie. [lucifer](https://www.rottentomatoes.com/tv/lucifer) - I am watching.

_Nothing more_.

_They are not intended to hurt anyone_.

## Installing

1. Clone it

2. Download deps of each subrepo

    ```bash
    for dir in */; do pushd $dir; dep ensure -v; popd; done
    ```

