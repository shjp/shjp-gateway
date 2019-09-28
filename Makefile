local:
	cd cmd/server && go run main.go

build-functions:
	mkdir -p functions
	go get ./...
	go build -o functions/gateway ./cmd/netlify-function
