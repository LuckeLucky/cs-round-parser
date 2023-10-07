# Build
Build parser to parse .dem inserted in "\demos" folder:
```sh
    go build
```

Build parser to parse a single .dem file:
```sh
    go build -tags single
```

# Command-line arguments

| Command | Value type | Default |
|---|---|---|
| -rM | int | 800 is the default regular phase start money |
| -otM | int | 16000 is the default overtime start money |

Example usage:
```sh
     .\cs-round-parser.exe -ot 13000
```

If we are parsing the overtime all rounds in overtime should have "mp_overtime_startmoney" equal to 13000


# Coverage

```sh
    go test -v -coverprofile cover.out ./
    go tool cover -html cover.out -o cover.html
```

