.PHONY:  lint

# Lint
lint:
	 golangci-lint run

# build
build:
	PATH=$(PATH) go build ./...

# test (unit test + godog behaviour tests)
test:
	PATH=$(PATH) go test ./...

# run
run:
	PATH=$(PATH) go run cmd/mars-rover/main.go -fPath="cmd/mars-rover/sample.txt"
