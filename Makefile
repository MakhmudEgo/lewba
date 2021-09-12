all:
	go build -o matcha cmd/web/main.go

run:
	@go run cmd/web/main.go