build:
	@echo "Building binary"
	go build -o tracker .
run:
	@echo "Running binary"
	go run tracker

start: build run

