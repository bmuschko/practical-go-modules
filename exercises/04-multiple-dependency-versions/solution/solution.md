# Solution

Create a new directory named `multi-versions` and initialize Go Modules.

```bash
$ mkdir multi-versions
$ cd multi-versions
$ go mod init github.com/bmuschko/multi-versions
```

Create a main.go file with a basic skeleton.

```go
package main

func main() {
}
```

Retrieve version `v1.0.0` with the `go get` command.

```bash
go get github.com/bmuschko/calc@v1.0.0
go: finding github.com/bmuschko/calc v1.0.0
go: downloading github.com/bmuschko/calc v1.0.0
go: extracting github.com/bmuschko/calc v1.0.0
```

Use the API in the `main.go` file.

```go
package main

import (
	"fmt"
	"github.com/bmuschko/calc"
)

func main() {
	fmt.Println(calc.Add(1, 2))
}
```

Running `go run main.go` should return the value 3 on the console.

Now, get version 2 two of the library.

```bash
go get github.com/bmuschko/calc/v2@v2.0.0
go: finding github.com/bmuschko/calc/v2 v2.0.0
go: finding github.com/bmuschko/calc v2.0.0
go: finding github.com/bmuschko v2.0.0
go: finding github.com v2.0.0
go: downloading github.com/bmuschko/calc/v2 v2.0.0
go: extracting github.com/bmuschko/calc/v2 v2.0.0
```

Change the code in `main.go` to use the API. The import for the v2 needs to provide a unique identifier.

```go
package main

import (
	"fmt"
	"github.com/bmuschko/calc"
	calcv2 "github.com/bmuschko/calc/v2"
)

func main() {
	fmt.Println(calc.Add(1, 2))
	fmt.Println(calcv2.Add(1, 2, 5, 8))
}
```

Running `go run main.go` should return the value 3 and 16 on the console.