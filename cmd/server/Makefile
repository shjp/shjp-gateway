build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shjp-gateway .

docker_up: build
	docker-compose build
	docker-compose up -d

docker_update: build
	docker build --no-cache --tag shjp_gateway_server .
	docker run --name shjp_gateway_server -itd shjp_gateway_server
	docker commit shjp_gateway_server sanggonlee/shjp:shjp_gateway_server
	docker stop shjp_gateway_server
	docker rm shjp_gateway_server
	docker push sanggonlee/shjp:shjp_gateway_server
