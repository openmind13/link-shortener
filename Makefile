ENTRY_POINT = ./cmd/apiserver
BIN_NAME = ./apiserver
OS = GOOS=linux


.PHONY: start
start:
	sudo service mongodb start && $(OS) go build -v -o apiserver $(ENTRY_POINT) && echo "\n" && $(BIN_NAME)


.PHONY: test
test:
	go test -v -race -timeout 30s ./app/...


.DEFAULT_GOAL := start