DOCKER_IMAGE = sosedoff/fargate-sumo-forwarder

docker:
	docker build -t ${DOCKER_IMAGE} .

docker-push:
	docker push ${DOCKER_IMAGE}