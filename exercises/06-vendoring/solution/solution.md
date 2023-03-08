# Solution

Create a new directory and navigate there in the console. Initialize Go Modules.

```bash
$ mkdir code
$ cd code
$ go mod init github.com/bmuschko/my-app
go: creating new go.mod: module github.com/bmuschko/my-app
```

Get the latest version of the dependency.

```bash
$ go get github.com/sirupsen/logrus
```

The dependency graph contains the dependency `github.com/sirupsen/logrus` and its transitive closure.

```bash
$ go mod graph
github.com/bmuschko/my-app github.com/sirupsen/logrus@v1.9.0
github.com/bmuschko/my-app golang.org/x/sys@v0.0.0-20220715151400-c0bba94af5f8
github.com/sirupsen/logrus@v1.9.0 github.com/davecgh/go-spew@v1.1.1
github.com/sirupsen/logrus@v1.9.0 github.com/stretchr/testify@v1.7.0
github.com/sirupsen/logrus@v1.9.0 golang.org/x/sys@v0.0.0-20220715151400-c0bba94af5f8
github.com/stretchr/testify@v1.7.0 github.com/davecgh/go-spew@v1.1.0
github.com/stretchr/testify@v1.7.0 github.com/pmezard/go-difflib@v1.0.0
github.com/stretchr/testify@v1.7.0 github.com/stretchr/objx@v0.1.0
github.com/stretchr/testify@v1.7.0 gopkg.in/yaml.v3@v3.0.0-20200313102051-9f266ea9e77c
gopkg.in/yaml.v3@v3.0.0-20200313102051-9f266ea9e77c gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
```

Add a `main.go` file with the following contents. As you can see it uses the API of the external dependency.

```go
package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
```

The `go mod vendor` downloads the dependencies and puts them into the `vendor` directory.

```bash
$ go mod vendor
$ ls -l vendor
total 8
drwxr-xr-x  3 bmuschko  staff   96 Mar  8 13:28 github.com
drwxr-xr-x  3 bmuschko  staff   96 Mar  8 13:28 golang.org
-rw-r--r--  1 bmuschko  staff  245 Mar  8 13:28 modules.txt
```

Generate the binary with the `go build` command. Make sure to point to the vendored dependencies.

```bash
$ go build -mod=vendor .
$ ./my-app
INFO[0000] A walrus appears                              animal=walrus
```

# Discussion

> From your perspective, what could be an issue with using vendoring?

- The `vendor` folder is not used by default for the `go` command. If you do not provide `-mod=vendor`, it will not be respected. This can lead to confusion.
- Vendored dependencies can take up a lot of space on disk especially if your project is depending on a lot of external packages. As a result, cloning a repository with vendored dependencies can take a long time.
- Updating a dependency in the `vendor` directory can make reviewing pull requests more difficult.