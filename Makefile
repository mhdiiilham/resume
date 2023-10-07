.PHONY: .

generate:
	go generate ./...

test:
	go test ./... -cover
