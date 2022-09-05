# Chapter 1 Notes

## Installing the Go Tools
* Go programs compile to a single binary and do not require any additional software to be installed in order to run them.

## The Go Workspace
* It's a good idea to explicitly define GOPATH and to put the $GOPATH/bin directory in your executable path. Makes it clear where your Go workspace is and makes it easy to run third party tools.
* Add the following to .profile/.zshrc

```
# Go Paths
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

## The go Command
* ```go run```
    * builds the binaries in a temporary directory and then deletes it when the program has completed.
    * Makes it simple to test small programs an use the language like a scripting language.
* ```go build```
    * creates an executable binary
    * -o flag is used to pass in a file name or directory path where the binary will be created
* ```go install```
    * will build and install into your workspace from source code
    * Takes location argument and version separated by @
    * E.g. ```go install github.com/rakyll/hey@latest```
    * Pulls code and dependencies
    * convenient way to distribute Go programs to other Go developers
* ```go fmt```
    * Reformats code to match the standard format
* ```go vet```
    * Detects errors that are syntactically correct but are mistakes, e.g. passing wrong number of parameters or assigning values to variables that are never used.
    * Should require this tool as part of your automated build


## Getting Third-Party Go Tools
* Go developers don't rely on a centrally hosted service, e.g. NPM or Maven
* Projects are shared via source code repositories

## Formatting Your Code
* Go enforces a standard format 
* Makes it easy to write tools that manipulate source code.
* Simplifies the compiler
* Allows the creation of clever tools to generate code.
* Always run ```go fmt``` or ```goimports``` before you build and run your code.

## Semicolon insertion rule
* If the last token before a new line is any of the following, the lexer inserts a semicolon after the token:
    * Identifier, int, float64
    * Basic literal, number or string constant
    * One of these tokens: break, continue, fallthrough, return, ++, --, ) or }
* One of the things that makes the Go compiler simpler and faster

## Linting and vetting
* ```golint``` helps enforce style guidelines.
    * Some of the changes it will suggest is:
        * properly naming variables
        * formatting error messages
        * Adding comments to public methods and types
    * Should require this as part of your code review process
    * Use this command to install:  ```go install golang.org/x/lint/golint@latest```
    * Use this command to run it: ```golint ./...```
* ```golangci-lint``` runs multiple tools together
    * combines golint, go vet, and ever increasing set of code quality tools.
    * running multiple tools slows down the build process, ```golangci-lint``` fixes that

## Makefiles
* Go developers have adopted make as the tool for repeatable and automatable builds
* Each operation is called a target
* `.DEFAULT_GOAL` defines what target should run if none is specified
* `.PHONY` keeps make from getting confused if something is named the same as a target in the directory
