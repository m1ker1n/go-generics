# Go-generics

[![Go Reference](https://pkg.go.dev/badge/github.com/m1ker1n/go-generics.svg)](https://pkg.go.dev/github.com/m1ker1n/go-generics)

Go-generics is a package for processing slices such as map, filter, etc.

# Installation

```shell
go get github.com/m1ker1n/go-generics
```
# Notes for me...

## ...how to test

```shell
# Passing all the tests, generate file 'cover.out' for seeing coverage
go test -coverprofile='cover.out' -v
```

```shell
# See coverage
go tool cover -html '.\cover.out'
```

## ...how to publish further
 
1. [Publishing a module](https://go.dev/doc/modules/publishing)
2. [Module releases and versioning workflow](https://go.dev/doc/modules/release-workflow)
3. [How to add package to pkg.go.dev](https://pkg.go.dev/about)
4. [How to Go Doc Comments](https://go.dev/doc/comment)

