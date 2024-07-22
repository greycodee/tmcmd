.PHONY: all windows linux-arm linux-amd64 mac-arm mac-intel

BINARY_NAME=tmcmd

build: bin/darwin/arm64/$(BINARY_NAME) bin/darwin/intel/$(BINARY_NAME) bin/linux/arm64/$(BINARY_NAME) bin/linux/amd64/$(BINARY_NAME) bin/windows/amd64/$(BINARY_NAME).exe
	echo "Build complete"

tar.gz: bin/darwin/arm64/$(BINARY_NAME) bin/darwin/intel/$(BINARY_NAME) bin/linux/arm64/$(BINARY_NAME) bin/linux/amd64/$(BINARY_NAME) bin/windows/amd64/$(BINARY_NAME).exe
	tar -czvf bin/darwin-arm64-$(BINARY_NAME).tar.gz bin/darwin/arm64/$(BINARY_NAME)
	tar -czvf bin/darwin-intel-$(BINARY_NAME).tar.gz bin/darwin/intel/$(BINARY_NAME)
	tar -czvf bin/linux-arm64-$(BINARY_NAME).tar.gz bin/linux/arm64/$(BINARY_NAME)
	tar -czvf bin/linux-amd64-$(BINARY_NAME).tar.gz bin/linux/amd64/$(BINARY_NAME)
	zip -r bin/windows-amd64-$(BINARY_NAME).zip bin/windows/amd64/$(BINARY_NAME).exe

bin/darwin/arm64/$(BINARY_NAME):
	GOOS=darwin GOARCH=arm64 go build -o bin/darwin/arm64/$(BINARY_NAME) .
bin/darwin/intel/$(BINARY_NAME):
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/intel/$(BINARY_NAME) .
bin/linux/arm64/$(BINARY_NAME):
	GOOS=linux GOARCH=arm64 go build -o bin/linux/arm64/$(BINARY_NAME) .
bin/linux/amd64/$(BINARY_NAME):
	GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/$(BINARY_NAME) .
bin/windows/amd64/$(BINARY_NAME).exe:
	GOOS=windows GOARCH=amd64 go build -o bin/windows/amd64/$(BINARY_NAME).exe .

clean:
	rm -rf bin