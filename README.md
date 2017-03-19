# Motivation
I find myself always creating small projects to experiment with, and I wanted
a tool which makes this easier. I didn't want the tool to make assumptions
about the type of project I am creating, and I wanted it to be intelligent
enough to be usable from anywhere.

Thus `go-new` was born.

# Usage
`go-new` provides two distinct usage models. Inside the `GOPATH` and outside.
Both of the use cases only require the name of the project.

## Inside the `GOPATH`
```bash
$ pwd
  /User/achernyak/go

# run the command
$ go-new sampleProject
```

This will create a new directory `/User/achernyak/go/sampleProject` and a 
blank Go file, `sampleProject.go`, inside.

If you run `go-new` from inside the `GOPATH`, it is assumed that you are in the
location you with the directory to be created.

## Outside the `GOPATH`
```bash
$ echo $GOPATH
  /User/achernyak/go
$ pwd
  /User/achernyak

# run the command
$ go-new sampleProject
```

This will create a new directory
`/User/achernyak/go/src/github.com/achernyak/sampleProject` and a blank Go
file, `sampleProject.go`, inside.

If you run the command outside the `GOPATH`, two assumptions are made. 

1. You are creating a github project.
2. You want the project to be under your github user name. `~/.gitconfig` is
   read to determine this.

# Todo
* Make code repository configurable via flags.
* Make code repository configurable via config file.
* Make `.gitconfig` location configurable via flags.
* Make `.gitconfig` location  configurable via config file.
