# How to use dep in your golang project

* Repo: [https://github.com/golang/dep](https://github.com/golang/dep)
* Doc: [https://golang.github.io/dep/](https://golang.github.io/dep/)

## dep init
```bash
$ dep init
```
To initialize dep in your project, all you have to do is launch the next command, this will add two files, Gopkg.toml and Gopkg.lock. 

## add dependencies
```bash
$ dep ensure -add github.com/travelgateX/go-jwt-tools
```
To add a dependency, it must be added in the command line, not editing the Gopkg.toml file. This command will add the necessary dependencies to the Gopkg.toml file and update de Gopkg.lock file.

## add subpackages
```bash
$ dep ensure -add github.com/travelgateX/go-jwt-tools/jwt
```
We will use the example repository github.com/travelgateX/go-jwt-tools that have one "authorization" package, and two sub-packages that have a dependency on the first one, the "jwt" and "cache" packages.

If we only want the jwt package, we will use the previous command, dep will resolve the dependency and will download the files included in the jwt folder and the parent project (authorization).

But instead of that if we want only the authorization package, the dep ensure command of the repo (not of the jwt package) ONLY the authorization package will be downloaded but not the two sub-packages because there is no dependency between them.

Finally, the version of the sub-packages can not be fixed in the Gopkg.toml, the branch or fixed version will be that of the original repository.

## download vendoring for your project
```bash
$ dep ensure
```
To download all project dependencies

## update Gopkg.lock
```bash
$ dep ensure -update
```
update dep depencies

