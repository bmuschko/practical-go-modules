# Exercise 1

In this exercise, you will initialize a project to use Go Modules. Later, you will add a new dependency and use its API in the application code. Feel free to use the IDE of your choice.

## Initializing the project

1. Create a new directory that serves as project root directory anywhere on your local disk.
2. Navigate to the directory in the console.
3. Initialize the project by running the `go mod init` command. Provide a module path of your choice. Inspect the generated `go.mod` file.
4. Create a `main.go` file that defines the `main` function. The function is supposed to read a JSON file and parse the data into a struct representation by using the `encoding/json` package. Render the struct in the console. You can find the JSON file in the [`data` folder](../../data).
5. Execute the `go build` command to generate a binary file you can use to execute the program.
6. Run the program with the help of the generated binary file.

## Adding a dependency

1. Instead of using the `encoding/json` package for parsing the JSON data, use the external package [`buger/jsonparser`](https://github.com/buger/jsonparser). Add the latest version of this dependency to the project.
2. Inspect the dependency declaration in the `go.mod` file. What version do you see? Why is the dependency marked as "indirect"?
3. Verify if other versions are available for the dependency with `go list` or by looking at the GitHub repository.
4. Locate the dependency files in the cache.
5. Import the package of the external dependency in your `main.go` file. Use the API of the package e.g. get the value of the attribute "city". Execute the `go build` command to generate a binary file you can use to execute the program. Does the dependency declaration change in `go.mod`?