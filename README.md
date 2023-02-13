# Possible error in go-flags

The help message is being printed twice.

Steps to reproduce the error:

```bash
go build -o flagstest cmd/flagtest/main.go
./flagstest --help
Usage:
  flagstest [OPTIONS] Job

Application Options:
  -v, --verbose  Show verbose debug information
  -s, --server=  Local archive server URL or blank if using AWS
  -d, --docker   Running in docker

Help Options:
  -h, --help     Show this help message

Arguments:
  Job:           JSON job object string

2023/02/13 14:57:45 Usage:
  flagstest [OPTIONS] Job

Application Options:
  -v, --verbose  Show verbose debug information
  -s, --server=  Local archive server URL or blank if using AWS
  -d, --docker   Running in docker

Help Options:
  -h, --help     Show this help message

Arguments:
  Job:           JSON job object string
```


