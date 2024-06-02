install:
	go mod tidy
	go mod download
	go mod vendor

start:
	go run main.go start

add-balance:
	go run main.go add-balance --email $(email)