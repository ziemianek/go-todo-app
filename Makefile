build:
	@echo "Building in $(PWD)/bin/..."
	go build -o bin/$(shell basename $(PWD)) main.go
