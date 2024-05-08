# cs-round-parser
Output round by round score of a counterstrike match, along with the details of the participants and spectators.

# Example Output
![image](https://github.com/LuckeLucky/cs-round-parser/assets/43279191/f7de0c7e-07f1-46dc-b915-69bf85eb3818)

# Build

Build parser
```sh
    go build
    go build -ldflags "-X main.readerType=simple" -o cs-round-simple-parser
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
