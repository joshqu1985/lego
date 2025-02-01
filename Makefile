NAME = lego

.PHONY: build
build: clean
	@echo "Building..." 
	CGO_ENABLED=0 go build -ldflags "-s -w" -o $(NAME) cmd/main.go

.PHONY: run
run:
	go run -race cmd/main.go

.PHONY: clean
clean:
	@echo "Cleaning"
	@go clean

.PHONY: test
test:
	go test -v -race ./...

.PHONY: tidy
tidy:
	@go mod tidy
