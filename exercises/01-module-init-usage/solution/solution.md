# Solution

## Initializing the project

Anywhere in your file system create a new directory and navigate to it.

```bash
$ mkdir json-parsing
$ cd json-parsing
```

Initialize the project for using Go Modules.

```bash
$ go mod init github.com/bmuschko/json-parsing
go: creating new go.mod: module github.com/bmuschko/json-parsing
```

The contents of `go.mod` should look as follows depending on the module path you provided.

```
module github.com/bmuschko/json-parsing

go 1.13
```

Create a new file called `main.go`. The content could look as shown in [this file](./code-without-dependency/main.go).

Run the `go build` command and execute the binary. The binary may use a different file extension depending on your operating system. The parsed JSON data could look as follows:

```bash
$ go build .
$ ./json-parsing
{John 31 New York}
```

## Adding a dependency

Get the latest version of the external dependency.

```bash
$ go get github.com/buger/jsonparser
go: finding github.com/buger/jsonparser latest
```

The contents of `go.mod` should contain an entry for the dependency with an "indirect" comment. This comment is shown because you didn't use the API of the dependency in the code yet. No version of the dependency has been tagged yet on GitHub. Therefore, you will get the latest commit on `master` indicated by a semantic version combination of 0.0.0 + timestamp + commit hash.

```
module github.com/bmuschko/json-parsing

go 1.13

require github.com/buger/jsonparser v0.0.0-20191204142016-1a29609e0929 // indirect
```

The `go list` command will show that no versions have been published yet.

```bash
$ go list -m -versions github.com/buger/jsonparser
github.com/buger/jsonparser
```

Modify the contents of `main.go`. Use the API of the external dependency. The content could look as shown in [this file](./code-with-dependency/main.go). If you are working in an IDE, the "indirect" comment will be removed automatically.

Run the `go build` command and execute the binary. The binary may use a different file extension depending on your operating system. The value of the JSON attribute "city" could look as follows:

```bash
$ go build .
$ ./json-parsing
New York
```