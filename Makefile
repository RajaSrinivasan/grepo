BIN := bin

all:
	go build -o $(BIN)/ ./...

test:
	go test ./...

setup:
	mkdir $(BIN)

clean:
	rm -rf $(BIN)