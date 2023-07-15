download:
	@echo Download go.mod dependencies
	@go mod download
 
install-tools: download
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

example.basic:
	go run main.go -s ./examples/basic/schema -t ./examples/basic/erd.md -o markdown

build:
	go build -o ./bin/entmaid

test:
	go test -v ./cmd/...