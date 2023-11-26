up:
	docker compose up -d
	@echo Done!

docker-build:
	cd ./info&& make docker-build
	cd ./support&& make docker-build
	@echo Done!

go-build:
	cd ./info&& make build
	cd ./support&& make build
	@echo Done!