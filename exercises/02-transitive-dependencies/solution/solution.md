# Solution

The initial [`go.mod`](../start/go.mod) file declare a dependency on `github.com/stretchr/testify` with version `v1.4.0`. Testify pull in various transitive dependencies.

The dependency graph looks as follows:

```bash
$ go mod graph
github.com/bmuschko/calc github.com/stretchr/testify@v1.4.0
github.com/stretchr/testify@v1.4.0 github.com/davecgh/go-spew@v1.1.0
github.com/stretchr/testify@v1.4.0 github.com/pmezard/go-difflib@v1.0.0
github.com/stretchr/testify@v1.4.0 github.com/stretchr/objx@v0.1.0
github.com/stretchr/testify@v1.4.0 gopkg.in/yaml.v2@v2.2.2
gopkg.in/yaml.v2@v2.2.2 gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
```

The selected versions can be listed with the `go list` command. Currently, there are no conflicting versions for transitive dependencies.

```
$ go list -m all
github.com/bmuschko/calc
github.com/davecgh/go-spew v1.1.0
github.com/pmezard/go-difflib v1.0.0
github.com/stretchr/objx v0.1.0
github.com/stretchr/testify v1.4.0
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
gopkg.in/yaml.v2 v2.2.2
```

Add the latest version of the dependency `github.com/stretchr/objx`.

```bash
$ go get github.com/stretchr/objx
```

As a result, the `go.mod` file contain the `objx` library as top-level dependency. We are not using the dependency on our code. For that reason, it is marked as "indirect".

```
module github.com/bmuschko/calc

go 1.13

require (
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.4.0
)
```

You will find that [`objx`](https://github.com/stretchr/objx/blob/v0.2.0/go.mod) has a dependency on `github.com/davecgh/go-spew` and `github.com/stretchr/testify` as well but with a different version than already available in the dependency graph.

```
module github.com/stretchr/objx

go 1.12

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/stretchr/testify v1.3.0
)
```

You will see this fact reflected in the dependency graph of your project.

```bash
$ go mod graph
github.com/bmuschko/calc github.com/stretchr/objx@v0.2.0
github.com/bmuschko/calc github.com/stretchr/testify@v1.4.0
github.com/stretchr/testify@v1.4.0 github.com/davecgh/go-spew@v1.1.0
github.com/stretchr/testify@v1.4.0 github.com/pmezard/go-difflib@v1.0.0
github.com/stretchr/testify@v1.4.0 github.com/stretchr/objx@v0.1.0
github.com/stretchr/testify@v1.4.0 gopkg.in/yaml.v2@v2.2.2
github.com/stretchr/objx@v0.2.0 github.com/davecgh/go-spew@v1.1.1
github.com/stretchr/objx@v0.2.0 github.com/stretchr/testify@v1.3.0
gopkg.in/yaml.v2@v2.2.2 gopkg.in/check.v1@v0.0.0-20161208181325-20d25e280405
github.com/stretchr/testify@v1.3.0 github.com/davecgh/go-spew@v1.1.0
github.com/stretchr/testify@v1.3.0 github.com/pmezard/go-difflib@v1.0.0
github.com/stretchr/testify@v1.3.0 github.com/stretchr/objx@v0.1.0
```

The select versions us the "minimal version selection" strategy. It uses `go-spew` with version v1.1.1, `objx` with version v0.2.0 and `testify` with version v1.4.0.

```bash
go list -m all
github.com/bmuschko/calc
github.com/davecgh/go-spew v1.1.1
github.com/pmezard/go-difflib v1.0.0
github.com/stretchr/objx v0.2.0
github.com/stretchr/testify v1.4.0
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
gopkg.in/yaml.v2 v2.2.2
```

You can enforce the use of the version v1.3.0 for `testify` with a replace statement in the `go.mod` file.

```
module github.com/bmuschko/calc

go 1.13

require (
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.4.0
)

replace github.com/stretchr/testify => github.com/stretchr/testify v1.3.0
```

The selected version is now `v1.3.0`.

```bash
$ go list -m all
github.com/bmuschko/calc
github.com/davecgh/go-spew v1.1.1
github.com/pmezard/go-difflib v1.0.0
github.com/stretchr/objx v0.2.0
github.com/stretchr/testify v1.4.0 => github.com/stretchr/testify v1.3.0
```