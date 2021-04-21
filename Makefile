APP_NAME  := go-application-standard-structure
export APP_NAME

generate:
	go generate ./migrations/
	go generate ./pkg/api/
.PHONY: generate
