BUILD=go build -ldflags="-w -s"

default: build

build:
	@echo "Building cry..."
	$(BUILD) -o bin/cry.exe
	@echo "Building server..."
	cd web/ && $(BUILD) -o ../bin/web.exe

clean:
	@rm -rf bin/
	@rm -f debug debug.test web/debug web/debug.test

all:
	GOOS=windows GOARCH=amd64 $(BUILD) -o bin/windows_amd64.exe
	GOOS=windows GOARCH=386 $(BUILD) -o bin/windows_x86.exe
	GOOS=linux GOARCH=amd64 $(BUILD) -o bin/linux_amd64
	GOOS=linux GOARCH=386 $(BUILD) -o bin/linux_x86
	GOOS=darwin GOARCH=amd64 $(BUILD) -o bin/macos
