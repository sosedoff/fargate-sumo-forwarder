DOCKER_IMAGE = sosedoff/fargate-sumo-forwarder

test:
	go test -cover -race ./...

docker:
	docker build -t ${DOCKER_IMAGE} .

docker-push:
	docker push ${DOCKER_IMAGE}