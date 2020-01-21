# Solution

The `glide.yaml` contains a single dependency on `github.com/stretchr/testify`. The file `calc_test.go` uses the API of the dependency.

Execute the `go mod init` command to translate the dependency information into `go.mod`.

```bash
$ go mod init github.com/bmuschko/calc
go: creating new go.mod: module github.com/bmuschko/calc
go: copying requirements from glide.lock
```

The resulting `go.mod` should look as follows:

```
module github.com/bmuschko/calc

go 1.13

require github.com/stretchr/testify v1.4.0
```

After deleting the glide files and the vendor directory, the project structure should look fairly simplistic.

```bash
$ tree
.
├── calc.go
├── calc_test.go
├── go.mod
└── go.sum
```

The `go test` command may need to download dependencies but run through all test cases successfully.

```bash
$ go test ./... -v
=== RUN   TestAddWithTestify
--- PASS: TestAddWithTestify (0.00s)
=== RUN   TestSubtractWithTestify
--- PASS: TestSubtractWithTestify (0.00s)
=== RUN   TestMultiplyWithTestify
--- PASS: TestMultiplyWithTestify (0.00s)
=== RUN   TestDivideWithTestify
--- PASS: TestDivideWithTestify (0.00s)
PASS
ok  	github.com/bmuschko/calc	0.186s
```

For additional dependency cleanup, run the `go mod tidy` command.

```bash
$ go mod tidy
go: downloading gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
go: extracting gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
```