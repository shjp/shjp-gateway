build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shjp-gateway .

docker_up: build
	docker-compose build
	docker-compose up -d

docker_update: build
	docker build --no-cache --tag shjp-gateway-server_v$(version) .
	docker run --name shjp-gateway-server -itd shjp-gateway-server_v$(version)
	docker commit shjp-gateway-server sanggonlee/shjp:shjp-gateway-server_v$(version)
	docker stop shjp-gateway-server
	docker rm shjp-gateway-server
	docker push sanggonlee/shjp:shjp-gateway-server_v$(version)
	rm ./shjp-gateway

docker_update_dev: build
	docker build --no-cache --tag shjp-gateway-server_dev .
	docker run --name shjp-gateway-server -itd shjp-gateway-server_dev
	docker commit shjp-gateway-server sanggonlee/shjp:shjp-gateway-server_dev
	docker stop shjp-gateway-server
	docker rm shjp-gateway-server
	docker push sanggonlee/shjp:shjp-gateway-server_dev
	rm ./shjp-gateway