PROJECT=itachi

default: build

build:
	make -C juno rustdeps
	go build -v -o $(PROJECT) ./cmd/node/main.go

reset:
	@rm yu.log
	@rm -r yu cairo_db

clean:
	rm -f $(PROJECT)