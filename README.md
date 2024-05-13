### Linters

- enforced in github actions (by using golangci-lint), and PR can only be merged when linters are passed

- no local linter enforced

An example of local lint setup.

```bash
# edit .git/hooks/pre-commit

go test ./...
go vet ./...
staticcheck ./...
```

