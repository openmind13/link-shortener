ENTRY_POINT = ./cmd/apiserver
BIN_NAME = ./apiserver
OS = GOOS=linux


.PHONY: install
install:
	sudo apt update && sudo apt upgrade -y && sudo apt-get install -y mongodb-org && sudo apt install golang-go


.PHONY: install_wsl
install_wsl:
	sudo apt update && sudo apt upgrade -y && sudo apt install mongodb && sudo apt install golang-go


.PHONY: start
start:
	sudo service mongod start && go mod download && go build -v -o apiserver $(ENTRY_POINT) && echo "\n" && $(BIN_NAME)


.PHONY: start_wsl
start_wsl:
	sudo service mongodb start && go mod download && go build -v -o apiserver $(ENTRY_POINT) && echo "\n" && $(BIN_NAME)


.PHONY: test
test:
	sudo service mongod start && go test -v -race -timeout 30s ./app/...


.PHONY: test_wsl
test_wsl:
	sudo service mongodb start && go test -v -race -timeout 30s ./app/...


# .PHONY: build docker image
# image:
# 	docker build -t linkshortener .


.DEFAULT_GOAL := start_wsl