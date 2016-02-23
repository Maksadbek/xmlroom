APP=xmlroom
BIN=$(APP)
lint:
	golint ./...
test:
	go test ./...
build:
	godep go build -v -o $(BIN)
