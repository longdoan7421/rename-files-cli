# Rename Files

## Build

```bash
go build -ldflags "-s -w" -o rename-files main.go
```

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
