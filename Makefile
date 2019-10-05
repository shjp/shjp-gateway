local:
	cd cmd/server && go run main.go

build-functions:
	mkdir -p functions
	go get ./...
	go clean -cache
	go build -o functions/gateway ./cmd/netlify-function
