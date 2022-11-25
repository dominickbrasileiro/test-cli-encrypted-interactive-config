default: build

install_deps:
	go get -v ./...

build: install_deps
	go build -o ./bin/test-cli ./cmd

clean:
	rm -rf ./bin