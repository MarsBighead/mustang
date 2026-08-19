BINARY := xm
CMD    := ./cmd/xm
GO     := go

.PHONY: build clean test run

build:
	$(GO) build -o bin/$(BINARY) $(CMD)

clean:
	rm -rf bin/

test:
	$(GO) test ./...

run: build
	./bin/$(BINARY)

# Cross-compilation targets
.PHONY: build-darwin-arm64 build-darwin-amd64 build-linux-amd64

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 $(GO) build -o bin/$(BINARY)-darwin-arm64 $(CMD)

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 $(GO) build -o bin/$(BINARY)-darwin-amd64 $(CMD)

build-linux-amd64:
	GOOS=linux GOARCH=amd64 $(GO) build -o bin/$(BINARY)-linux-amd64 $(CMD)
