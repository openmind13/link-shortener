ENTRY_POINT = ./cmd/apiserver
BIN_NAME = ./apiserver


start:
	sudo service mongodb start && go build -v -o apiserver $(ENTRY_POINT) && echo "\n" && $(BIN_NAME)


.DEFAULT_GOAL := start