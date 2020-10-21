## Getting Started

1. `go build` - Compiles bunch of files
2. `go run` - Compiles and executed one or two files
3. `go fmt` - Formats all code in each file in the current directory
4. `go install` - Compiles and installs the package
5. `go get` - Downloads the raw source
6. `go test` - Runs any tests associated with the current project

A `package` is the collection of commom source code files
An app will have a single package name
_Package == Project == Workspace_

The very first line should in the file should name the package it belongs to

There are two types of packages in go:
1. **Excecutable** - Generates a file that we can run
2. **Resuable** - Code dependencies/helpers

The type of package is determined by the package name. `package main` specifies that the file is an executable package

The executable file should contain the function _main_

### Skeleton of a go file
1. Package declarations
2. Import statements
3. Function declarations 