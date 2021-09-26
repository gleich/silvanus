build:
	go build -v -o bin/ .

test:
	go test -v ./...

lint: lint-golangci lint-gomod

#########
# Linting
#########

lint-golangci:
	golangci-lint run
lint-gomod:
	go mod tidy
	git diff --exit-code go.mod
	git diff --exit-code go.sum
