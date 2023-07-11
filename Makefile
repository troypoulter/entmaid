download:
	@echo Download go.mod dependencies
	@go mod download
 
install-tools: download
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

generate-examples:
	@go generate ./examples/...

example.basic:
	go run main.go -s ./examples/basic/schema -t ./examples/basic/erd.md