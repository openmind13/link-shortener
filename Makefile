ENTRY_POINT = ./cmd/apiserver
BIN_NAME = ./apiserver


.PHONY: start
start:
	sudo service mongodb start && go build -v -o apiserver $(ENTRY_POINT) && echo "\n" && $(BIN_NAME)


.PHONY: test
test:
	sudo service mongodb start && go test -v -race -timeout 30s ./app/...


.DEFAULT_GOAL := start