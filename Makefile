DOCKER_TAG := latest
REGISTORY_NAME := go-api-sqlite

build:
	docker build -t ${REGISTORY_NAME}:${DOCKER_TAG} .
run:
	docker run -p 8080:8080 -it \
	-e AWS_REGION=ap-northeast-1 \
	-e AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} \
	-e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} \
	-e AWS_SESSION_TOKEN=${AWS_SESSION_TOKEN} \
	${REGISTORY_NAME}

.PHONY: build run
