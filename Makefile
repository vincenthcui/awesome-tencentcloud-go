.PHONY: generate

SRC=./cmd/genaction
BIN=bin/genaction
ACTIONS=tencentcloud/actions

generate: $(BIN)
	$(BIN) -workdir $(ACTIONS)
	go fmt ./$(ACTIONS)/...

$(BIN):
	go build -o $(BIN) $(SRC)
