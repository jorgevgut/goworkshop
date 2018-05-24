# Go-lang worshop
---

## Index

* [Introduction](https://github.com/jorgevgut/goworkshop#introduction)
* [Requirements](https://github.com/jorgevgut/goworkshop#requirements)
* [Language overview](https://github.com/jorgevgut/goworkshop#language-overview)
* Designing modular apps, packaging, and vendoring
* [Best practices and linters (Gometalinter)](https://github.com/jorgevgut/goworkshop#best-practices)
* [Tips related to 'documentation'](https://github.com/jorgevgut/goworkshop#documentation-tips)
* [Writing unit tests](https://github.com/jorgevgut/goworkshop#unit-testing)
---
## Introduction

This repository is intended to showcase bits of Go lang's essential concepts, including (but not limited to grow over time with more stuff) working with interfaces, go routines, channels, unit tests, and of course,  references to already well known sources.

##### Who is this workshop for?

Anyone with programming experience that is interested in learning Go essential concepts without worrying to much about trivial stuff such as the syntax. We will 'go' straight to the cool stuff. For more trivial concepts, references will be included.

##### How is this workshop organized?

This workshop is divided in sections, some sections are more theoretical than others, and since plenty of content already exists for more "general" topics such as language syntax, references will be included. For more advanced and practical topics, instructions and code will be included. Don't worry, just follow the yellow brick road.


This repository may also be used as a reference, to do so each section that requires a practical example include a reference to the source code. To run such examples, follow the instructions and locally edit the main.go file if required to run a particular function provided in this repository.

---
## Requirements

What you need...

* Go >= 1.10.1 -> https://golang.org/doc/install
* Some comfortable text editor
* This project, retrieve it using `go get`
```bash
go get github.com/jorgevgut/goworkshop
```
* Install `go dep` for dependency related content in this workshop
```bash
go get -u github.com/golang/dep/cmd/dep
```

Notes:
- For top Editors/Plugins for Go check [this resource from ardanlabs](https://github.com/ardanlabs/gotraining#editors)
- `go get` is the primary tool used to retrieve packages written in go, [check this reference ](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) to learn more

---
## Language overview
Go is pretty much a compiled programming language that happens to be statically typed. It is a language that intends to be simple and yet powerful enough to take advantage of multicore systems. Here are some concepts and references that will let you learn more about this language.

* Programming Paradigms: Mostly Imperative, kind of NOT so much object oriented([read this about Go and OOP](https://golang.org/doc/faq#Is_Go_an_object-oriented_language)), concurrent, functional-ish.

* Some concepts to remember(Work In Progress):
 * arrays(which have fixed size) and slices(variable size) read [this introduction](https://blog.golang.org/go-slices-usage-and-internals) and [tricks/operations with slices](https://github.com/golang/go/wiki/SliceTricks)
 * maps([example](https://gobyexample.com/maps))
 * value and pointer receivers([important info](https://stackoverflow.com/questions/27775376/value-receiver-vs-pointer-receiver-in-golang))
 * channels
 * goroutines

* [Great reference on language mechanics (by ardanlabs)](https://github.com/ardanlabs/gotraining/blob/master/topics/courses/go/language/README.md)

---
## Documentation tips

Documenting code in go is quite easy, there are quite simple rules to follow to have basic decent documentation.

1.  Always document Exported functions, types, constants. In order to do this follow this example.

```go
/* To document a function, type or else, write a comment whose
 first word is the name of the function/type/other and follow
 with the description of what it does */

// ExportedFunction does something and then does something more
func ExportedFunction() {
  // do something
}

/* DocumentedType can be documented to by using 'big' paragraphs as comments
*/
type DocumentedType struct {
    id      string
    Value   int
}
```

2. Use **godoc**, it is your friend
```sh
godoc -http=:6060
```
This command will access all packages accessible from `$GOPATH` and `$GOROOT` within your system. Then head over to http://localhost:3000

#### Other tips and tricks

#### Visualizing dependencies
Visualizing dependencies is highly important since in Go circular dependencies are not allowed, godoc is a great tool for generating documentation, but using graphviz is way better to visualize our project's dependency tree.

Follow [this guide](https://golang.github.io/dep/docs/daily-dep.html#visualizing-dependencies) to visualize your code.

---
## Unit testing

This section covers the basics of how to do unit tests when using the Go programming language.

The Go programming language is shipped with a set of tools, one of the most important is 'test' which is use for testing.

##### Writing unit tests

* Tests must be written by 'package' in files with the suffix     **_test.go** (This suffix is what tells go 'test' tool that this is in fact a test).
* Inside the file, functions that start with 'Test' are the ones that will be executed, these may be called test functions.
* Test functions must expect a pointer to testing.T as parameter  ```(t *testing.T)```.
 * [Basic test file example](https://github.com/jorgevgut/goworkshop/blob/master/example_1_test.go).

##### Test coverage

In addition to the Go 'test' tool, 'cover' can be used in conjunction with 'test' to analyze the code coverage from the available unit tests, to generate a report run the following commands.

```bash
# run go test on a given package, visualize coverage and generate report
go test -coverprofile=coverage.out
go tool cover -func=coverage.out # view each function's coverage
go tool cover -html=coverage.out # generate html report

# same as before, but generate report on all sub-packages
go test ./... -coverprofile=coverage.out -v # added -v flag for verbose
go tool cover -func=coverage.out
go tool cover -html=coverage.out
```

###### Good reads
* Go cover story -> https://blog.golang.org/cover
