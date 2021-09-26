build:
	go get -v -t -d ./...
	go build -v .

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
