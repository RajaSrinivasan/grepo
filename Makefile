BIN := bin

all:
	go build -o $(BIN)/ ./...

test:
	go test ./...

setup:
	mkdir $(BIN)

release:
	srctrace -L go -m 0 -n 1 -b 0 -o version
	mv version.go impl/version

clean:
	rm -rf $(BIN)