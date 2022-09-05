## File Naming Convention

There are a few other caveats to this that we will cover a bit later that touch on things like
- `export_test.go` to access unexported variables in external tests
- `xxx_internal_test.go` for internal tests
- `example_xxx_test.go` for examples in isolated files 

## Running tests in parallel

Sometimes running many tests in parallel can provide a ton of value

Example use cases
1. Simulating a real-world scenario
2. Verify that a type is truly threadsafe

(1) - A web app with many users
(2) - Verify that your in-memory cache can handle multiple concurrent web requests using it

Prallelism could also mean more work:
- Test can't use as many hard-coded values; eq unique email constraints
- Tests might try to use shared resource incorrectly; eg image manipulation on the same image or sharing a DB that
doesn't support multiple concurrent connections

## Race detection flag

```bash
$ go test -race
$ go run -race thing.go
$ go build -race thing.go
$ go install -race pkg
$ go get -race golang.org/x/blog/support/racy
$ racy
```


## Go test coverage

```shell
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```