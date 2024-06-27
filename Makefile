PROJECT=itachi

default: build

build:
	make -C juno rustdeps
	go build -v -o $(PROJECT) ./cmd/node/main.go ./cmd/node/testrequest.go

reset:
	@rm -r yu cairo_db verse_db

clean:
	rm -f $(PROJECT)