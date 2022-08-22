# go-tady

A project template generator for Go

Eliminate manual re-creation of similar codes for Go  
When starting a new project, building a learning codebase, etc...

By registering your favorite codebase with `go-tady`  
You can use it as a new project with a single command at any time.

For simple projects, you can also use any of the default projects under `/assets/presets`.

# Install

`go install github.com/TadayoshiOtsuka/go-tady@latest`

# Usage

0. Run `go-tady init` for the first time only. (Create `$HOME/.go-tady.toml`)

1. Go to the Root of the project you wish to register as a preset.
2. Run `go-tady register <your preset name>`
3. Go to the directory where you want to create the new project.
4. Run `go-tady create`
5. Follow the interactive CLI. (It's so simple!)
