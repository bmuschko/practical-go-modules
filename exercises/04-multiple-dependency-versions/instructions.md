# Exercise 4

In this exercise, you will declare, resolve and use a dependency with different major versions. You will learn that the API even for breaking changes by using different import paths.

1. Create a new `main.go` file in a new directory. Initialize Go Modules for this project.
2. Inspect the code and tags of the repository [github.com/bmuschko/calc](https://github.com/bmuschko/calc). Compare the API between the major versions v1 and v2.
2. Add the dependency `github.com/bmuschko/calc` with version `v1.0.0`.
3. Import the library, use the `Add` method in your `main.go` file and print out the return value.
4. Add the dependency `github.com/bmuschko/calc/v2` with version `v2.0.0`.
5. Import the library with version 2, use the `Add` method in your `main.go` file and print out the return value. Function calls from both versions of the library should live alongside.