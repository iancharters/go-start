build: $(shell find . -name "*.go")
	@mkdir -p build
	go build -o ./build/gostart ./cmd/gostart
