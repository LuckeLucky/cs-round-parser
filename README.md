# Build

Build parser
```sh
    go build
```

# Usage

```sh
     .\cs-round-parser.exe
     .\cs-round-parser.exe .\demos\saw_vs_ftw.dem
```

If the first argument is not a specific file the parser will read all demos stored in ".\demos".

In Windows ".dem" files can be set to open with "cs-round-parser.exe".

# Coverage

```sh
    go test -v -coverprofile cover.out ./
    go tool cover -html cover.out -o cover.html
```

# Acknowledgments
This parser relies on [demoinfocs-golang](https://github.com/markus-wa/demoinfocs-golang).
