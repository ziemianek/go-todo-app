build:
	@echo "Building in $(PWD)/bin/..."
	go build -o bin/todolist main.go

test:
	@echo "Running tests..."
	go test ./...