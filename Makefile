.PHONY: build
build:
	@go build -gcflags=all="-N -l" -o ./bin/debug/pocone ./cmd/main.go
