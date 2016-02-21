APP=xmlroom
BIN=$(APP)
lint:
	golint ./...
test:
	go test ./...
build:
	CGO_ENABLED=0 godep go build -v -o $(BIN)
