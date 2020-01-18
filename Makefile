.PHONY:  tools lint


# Tools required
TOOLS = github.com/golangci/golangci-lint/cmd/golangci-lint

# Lint
lint: tools
	PATH=$(PATH) golangci-lint run

# build
build:
	PATH=$(PATH) go build ./...

# test
test:
	PATH=$(PATH) go test ./...

# run
run:
	PATH=$(PATH) go run cmd/mars-rover/main.go -fPath="cmd/mars-rover/sample.txt"
