# go-tutorial

## Go commands
```bash
# Run and build commands
go run .
go build -o [executable_name]
go list -f '{{.Target}}' # Discover the install path
go install # Add install path executable in path before installing

# Enable dependency tracking in a Go project by creating a go.mod file in the module the project source code is
go mod init [module_name]
go get [package_name] # Downloads package
go mod tidy
go clean -modcache
go mod download

# Workspace
go work init [module_or_folder_name]
go work use [-r] [dir] # adds a use directive to the go.work file for dir, if it exists, and removes the use directory if the argument directory doesn’t exist. The -r flag examines subdirectories of dir recursively.
go work sync # syncs dependencies from the workspace’s build list into each of the workspace modules.

# Testing
go test -v # Execute all tests
go test # Output only failing tests
```

## Go documentation
 - [Go documentation](https://go.dev/doc/)
 - [Effective Go](https://go.dev/doc/effective_go)