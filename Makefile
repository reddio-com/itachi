PROJECT=itachi

default: build

build:
	make -C juno rustdeps
	go build -v -o $(PROJECT) ./cmd/node/main.go

clean:
	rm -f $(PROJECT)