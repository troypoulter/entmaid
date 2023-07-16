download:
	@echo Download go.mod dependencies
	@go mod download
 
install-tools: download
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

example.readme:
	go run main.go -s ./examples/start/schema -t ./README.md -o markdown --startPattern "<!-- #start:entmaidReadme1 -->" --endPattern "<!-- #end:entmaidReadme1 -->"

example.start:
	go run main.go -s ./examples/start/schema -t ./examples/start/readme.md -o markdown

example.m2m2types:
	go run main.go -s ./examples/m2m2types/schema -t ./examples/m2m2types/readme.md -o markdown

example.all: example.readme example.start example.m2m2types

build:
	go build -o ./bin/entmaid

test:
	go test -v ./cmd/... -race -covermode=atomic -coverprofile=coverage.out