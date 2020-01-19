.PHONY:  lint

# Lint
lint:
	 golangci-lint run

# build
build:
	PATH=$(PATH) go build ./...

# test
test:
	PATH=$(PATH) go test ./...

# godog tests
godog-tests:
	(cd integration-tests && godog)

# run
run:
	PATH=$(PATH) go run cmd/mars-rover/main.go -fPath="cmd/mars-rover/sample.txt"
