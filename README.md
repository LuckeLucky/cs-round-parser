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
     .\demo-analyser-csgo.exe -ot 13000
```

If we are parsing the overtime all rounds should have "mp_overtime_startmoney" equal to 13000

