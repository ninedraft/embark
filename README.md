# embark

Create new golang project with this smol helper:
- create `pkg` and `cmd` dirs
- generate initial files
- init git repo
- init package manager

## Install

```bash
go get -u -v github.com/ninedraft/embark/cmd/embark
```

## Flags
  | key             | description                                |
  |-----------------|--------------------------------------------|
  |-c, --cli        | generate cli boilerplate                   |
  |-h, --help       | help for embark                            |
  |-l, --lib        | skip main package generation (default true)|
  |-n, --name       | new package name                           |
  |--package-manager| string   package manager to use            |

