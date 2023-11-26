BINARY=infoApp
SERVICE_NAME=info-service
DOCKER_FILE_NAME=info_service.dockerfile
VERSION=1.0

build:
	set GOOS=linux&& set GOARCH=amd64&& go build -o ${BINARY}
	@echo Done!


docker-build:
	docker build -f ./${DOCKER_FILE_NAME} -t ${SERVICE_NAME}:${VERSION} .
	@echo Done!