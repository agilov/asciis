run:
	@go build -o bin/game .
	@./bin/game
test:
	@go test ./...