all: build

build:
	@echo "Building..."
	@go build -o main cmd/libre-izly/main.go

run:
	@go run cmd/libre-izly/main.go

clean:
	@echo "Cleaning..."
	@rm -f main
