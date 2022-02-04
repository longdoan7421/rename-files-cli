# Rename Files

A personal command line tool for renaming files to conventional formats, e.g. title case, snake case, kebab case, ...

## Build

```bash
go build -ldflags "-s -w" -o rename-files main.go
```
*Note: have been tested with go@1.16 and above only.*

## Usage

```bash
rename-files -path "path to file or directory" -case "case type"
```

If you want to check result before renaming, use `--dry-run` flag.
```bash
rename-files -dry-run -path "path to file or directory" -case "case type"
```

Check helps to see supporting flags and more detail.
```bash
rename-files -h
```

## License
GNU GPLv3. See the [LICENSE](./LICENSE) file for details.
