# Exercise 3

In this exercise, you will verify the integrity of a dependency located in the cache. You will also learn how to ask why a specific dependency was pulled in. Finally, you will automatically clean up the `go.mod` file with the help of the `tidy` command.

1. Navigate to the subdirectory named [start](./start). It already contains go code and an existing Go Modules setup.
2. Render the dependency graph again to see the follow list of dependencies.
3. Use the `go mod why` command to identify where the dependency `gopkg.in/yaml.v2` comes from.
4. Run the `go mod verify` command. You should see that all modules can be verified.
5. Modify the source code of the dependency `gopkg.in/yaml.v2` in your local cache e.g. add a comment or a new function. You will likely have to overwrite the contents of the file to add write permissions.
6. Run the `go mod verify` command again. What do you see? What do you think would happen if you'd run `go build .`.
7. Fix the underlying issue of the transitive dependency. Ensure the correctness by running `go mod verify` again.
8. Add the dependency `github.com/juju/errors` with the latest version. Do not use the API in the code. Check `go.mod` to ensure that the dependency has been added.
9. Clean up the dependency definition by running `go mod tidy`. What do you see?