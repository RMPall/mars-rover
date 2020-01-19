.PHONY:  tools lint


# Tools required
TOOLS = curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.22.2

# Lint
lint: tools
	golangci-lint run

# build
build:
	PATH=$(PATH) go build ./...

# test
test:
	PATH=$(PATH) go test ./...

# run
run:
	PATH=$(PATH) go run cmd/mars-rover/main.go -fPath="cmd/mars-rover/sample.txt"
