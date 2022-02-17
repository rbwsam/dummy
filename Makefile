build:
	@CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/dummy main.go

build-image: build
	@docker build -t "rbwsam/dummy:latest" .

.PHONY: build build-image