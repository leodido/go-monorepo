# go-monorepo

A _PoC_ about organizing the code in a **monorepo** fashion using golang **dep**.

This code organization assumes the root is not vendored, only the leafs.

The directories can:

* depend on one (or more) siblings
* have transitive deps
* implement executables
* implement libraries

##### Disclaimer

The name of packages and directories are simply a bias given to me from the last TV series - ie. [lucifer](https://www.rottentomatoes.com/tv/lucifer) - I am watching.

Nothing more.

They are not intended to hurt anyone.

## Installing

1. Clone it
2. Execute `dep ensure -v` in each subfolder