WEB_BIN=webapi
SRC_BIN=xmldumper
lint:
	golint ./...
test:
	go test -cover ./...
build:
	godep go build -o $(WEB_BIN) web/main.go
	godep go build -o $(SRC_BIN) src/main.go
