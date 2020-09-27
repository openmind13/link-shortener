ENTRY_POINT = ./cmd/apiserver
BIN_NAME = ./apiserver
OS = GOOS=linux

.PHONY: start
start:
	sudo service mongodb start && go mod download && go build -v -o apiserver $(ENTRY_POINT) && echo "\n" && $(BIN_NAME)

.PHONY: test
test:
	sudo service mongodb start && go test -v -race -timeout 30s ./app/...

.PHONY: build docker image
image:
	docker build -t linkshortener .

.PHONY: compile
compile:
	go build -v -o apiserver $(ENTRY_POINT)

.DEFAULT_GOAL := start