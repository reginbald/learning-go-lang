# Notes
## Go initialization
package initialization is executed in dependency graph order which means that we can be guaranteed that when our main function runs all of our dependencies have been initialized[*](https://vimeo.com/53221558#t=1217s)

## Go tool is a Zero Configuration tool
No Makefiles or scripts, only Go code. The build schema and code are one and the same. This is because import paths mirror code location on the file system. Import paths can be URLs.[*](https://vimeo.com/53221558#t=1311s)

## Godoc
Extracts documentation from Go code and presents it in a variety of forms.[*](https://vimeo.com/53221558#t=1480s)

## Git pre commit hooks for go
* [gofmt](https://pkg.go.dev/cmd/gofmt) is the official tool for formatting go code.
* [errcheck](https://github.com/kisielk/errcheck) is a program for checking for unchecked errors in go programs.
* [vet](https://pkg.go.dev/cmd/vet) examines Go source code and reports suspicious constructs

[Guide to setup pre-commit hook for Git](https://levelup.gitconnected.com/committing-better-go-code-with-static-analysis-and-git-hooks-23216b8a1674)