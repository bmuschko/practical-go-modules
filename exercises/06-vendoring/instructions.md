# Exercise 6

In this exercise, you will vendor the dependencies declared in a `go.mod` file and use them going forward.

1. Create a new directory and initialize Go Modules for this project.
2. Add the dependency `github.com/sirupsen/logrus` with the latest version.
3. Inspect the `go.mod` file and render the dependency graph. Use the API of these dependency in the application code. You can find an [example in the project's README file](https://github.com/sirupsen/logrus/blob/master/README.md#example).
4. Vendor the dependencies. Inspect the contents of the `vendor` directory.
5. Execute the `go build` command. Ensure that the vendored dependencies are used.
6. Open the IDE of your choice and enable the vendoring option for that project.