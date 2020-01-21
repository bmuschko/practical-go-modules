# Exercise 7

In this exercise, you will use a different proxy server for resolving dependencies.

1. Create a new directory and initialize Go Modules for this project.
2. Open a browser, open [GoCenter](https://search.gocenter.io/) and search for the package `github.com/goccy/go-yaml`.
3. Navigate the dependencies of the package in the "Dependencies" tab. Find a transitive dependency that are required by other dependencies as indicated by the "Used By" tab.
4. Use GoCenter as proxy for the project. Verify the correct setup by rendering the value of the environment variable `GOPROXY`.
5. Add the dependency `github.com/goccy/go-yaml`.