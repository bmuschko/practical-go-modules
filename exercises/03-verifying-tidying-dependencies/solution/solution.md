# Solution

Navigate to the directory `start` and have a look at the dependency graph.

```bash
$ cd start
$ go mod graph
github.com/bmuschko/calc github.com/stretchr/testify@v1.4.0
github.com/stretchr/testify@v1.4.0 github.com/davecgh/go-spew@v1.1.0
github.com/stretchr/testify@v1.4.0 github.com/pmezard/go-difflib@v1.0.0
github.com/stretchr/testify@v1.4.0 github.com/stretchr/objx@v0.1.0
github.com/stretchr/testify@v1.4.0 gopkg.in/yaml.v2@v2.2.2
gopkg.in/yaml.v2@v2.2.2 gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
```

You can track down the originating dependency that requires `gopkg.in/yaml.v2`.

```bash
$ go mod why gopkg.in/yaml.v2
# gopkg.in/yaml.v2
github.com/bmuschko/calc
github.com/bmuschko/calc.test
github.com/stretchr/testify/assert
gopkg.in/yaml.v2
```

You can verify the integrity of the code.

```bash
$ go mod verify
all modules verified
```

Modify the file `$GOPATH/pkg/mod/gopkg.in/yaml.v2@v2.2.2/apic.go` by adding a comment. Running the `verify` command should render an error.

```bash
$ go mod verify
gopkg.in/yaml.v2 v2.2.2: dir has been modified (/Users/bmuschko/go/pkg/mod/gopkg.in/yaml.v2@v2.2.2)
```

Be aware that the `build` command does not automatically verify the dependencies and therefore will work just fine.

```bash
$ go build .
```

Remove the comment again and re-run the `verify` command.

```bash
$ go mod verify
all modules verified
```

Add the `errors` package with the `go get` command.

```bash
$ go get github.com/juju/errors
go: finding github.com/juju/errors latest
go: downloading github.com/juju/errors v0.0.0-20190930114154-d42613fe1ab9
go: extracting github.com/juju/errors v0.0.0-20190930114154-d42613fe1ab9
```

The resulting `go.mod` file will look as follows:

```
module github.com/bmuschko/calc

go 1.20

require (
	github.com/juju/errors v0.0.0-20190930114154-d42613fe1ab9 // indirect
	github.com/stretchr/testify v1.4.0
)
```

As we didn't use the `errors` package in the code, running the `tidy` command will simply remove the dependency as its declaration is deemed unnecessary.

```bash
$ go mod tidy
```

The `go.mod` will not contain the dependency declaration anymore.

```
module github.com/bmuschko/calc

go 1.20

require github.com/stretchr/testify v1.4.0
```