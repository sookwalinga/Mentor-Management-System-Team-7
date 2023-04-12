test:
	go test -v -cover ./...

server:
	go run main.go

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: test redis server
