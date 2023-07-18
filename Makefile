install-tools:
	go install entgo.io/ent/cmd/ent@latest
	curl -sSf https://atlasgo.sh | sh

example.readme:
	go run main.go -s ./examples/start/schema -t ./README.md -o markdown --startPattern "<!-- #start:entmaidReadme1 -->" --endPattern "<!-- #end:entmaidReadme1 -->"

example.start:
	go run main.go -s ./examples/start/schema -t ./examples/start/readme.md -o markdown

example.m2m2types:
	go run main.go -s ./examples/m2m2types/schema -t ./examples/m2m2types/readme.md -o markdown

example.edgeschema:
	go run main.go -s ./examples/edgeschema/schema -t ./examples/edgeschema/readme.md -o markdown

example.all: example.readme example.start example.m2m2types example.edgeschema

build:
	go build -o ./bin/entmaid

test:
	go test -v ./cmd/... -race -covermode=atomic -coverprofile=coverage.out