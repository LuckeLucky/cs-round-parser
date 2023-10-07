# Build
Build parser
```sh
    go build
```

# Usage

| Command | Value type | Default |
|---|---|---|
| -rM | int | 800 is the default regular phase start money |
| -otM | int | 16000 is the default overtime start money |

Example usage:
```sh
     .\cs-round-parser.exe -ot 13000
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
This analyser relies on [demoinfocs-golang](https://github.com/markus-wa/demoinfocs-golang).
