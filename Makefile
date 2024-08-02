PROJECT=itachi

default: build

build:
	make -C juno rustdeps
	go build -v -o $(PROJECT) ./cmd/node/main.go

reset:
	@rm -r yu cairo_db
	@rm -r yu cairo_state_diff

clean:
	rm -f $(PROJECT)