APP_NAME  := go-application-standard-structure
export APP_NAME

generate:
	go generate ./migrations/
	go generate ./pkg/api/
.PHONY: generate

test:
	go test -race github.com/fahdjamy/standard-structure-layout/..

run-dev:
	#go build cmd/myapp/main.go -o bin/app && ./bin/app
