build:
	@go build -o bin/caltrack

run: build
	@./bin/caltrack
