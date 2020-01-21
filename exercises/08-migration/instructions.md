# Exercise 8

In this exercise, you will migrate an existing project with dependencies to Go Modules. The pre-existing package manager used is [glide](https://github.com/Masterminds/glide). You will not need to install glide to complete the exercise.

1. Navigate to the subdirectory named [start](./start). It already contains Go code and the glide package information.
2. Inspect the file `glide.yaml`.
3. Run the `go mod init` command to translate the dependency information to a `go.mod` file.
4. Inspect the generated `go.mod` file. Delete the `vendor` directory as well as the `glide.yaml` and `glide.lock` files.
5. Execute the `go test ./... -v` command to ensure that the code compiles and tests can be executed.