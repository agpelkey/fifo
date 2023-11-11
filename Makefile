build:
	@echo 'Building fifo app'
	go build -ldflags='-s' -o=./bin/fifo ./cmd/api

run: build
	./bin/fifo
