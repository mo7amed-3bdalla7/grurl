## Grurl

A command line tool to extract a git repository remote url for a specific reference without using git.


## How It Works?

> Just type in terminal `grurl --remote=origin --path=repo/path` or `grurl` for default mostly used remote origin in the current path.

[![asciicast](https://asciinema.org/a/oJvGhVk3LEuaPd9rprILyBPNC.svg)](https://asciinema.org/a/oJvGhVk3LEuaPd9rprILyBPNC)

## Installation

1. `git clone https://github.com/mo7amed-3bdalla7/grurl.git`.
1. `cd grurl`.
2. `go install` or `go build -o /usr/bin/grurl` for global installation.
3. Just cd to any git repo add run `grurl`.

## Usage

```
Usage of grurl:
  -list
        list all remote names with URLs and exit
  -path string
        to set the git repository path (default ".")
  -remote string
        to set the remote name you for required url (default "origin")
```

## TODO
- [x] Handle remote http urls.
- [x] Handle ssh urls.
- [ ] Handle git urls.
- [ ] Handle different fetch and push URLS.