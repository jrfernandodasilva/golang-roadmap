.PHONY: help build run docker-build docker-run docker-sh docker-logs docker-stop

GREEN := $(shell tput setaf 2)
RESET := $(shell tput sgr0)

help:
	@echo "Available commands:"
	@echo "  $(GREEN)help$(RESET): See this help"
	@echo "  $(GREEN)build$(RESET): Build the application"
	@echo "  $(GREEN)run$(RESET): Run the application"
	@echo "  $(GREEN)docker-build$(RESET): Build the docker image"
	@echo "  $(GREEN)docker-run$(RESET): Run the docker image"
	@echo "  $(GREEN)docker-sh$(RESET): Run a shell in the docker image"
	@echo "  $(GREEN)docker-logs$(RESET): Show the logs of the docker image"
	@echo "  $(GREEN)docker-stop$(RESET): Stop the docker image"

build:
	go build -o golang-roadmap main.go

run:
	./golang-roadmap

docker-build:
	docker build -t golang-roadmap:unstable .

docker-run:
	docker run -d --rm --name golang-roadmap-app golang-roadmap:unstable

docker-sh:
	docker exec -it golang-roadmap-app sh

docker-logs:
	docker logs golang-roadmap-app

docker-stop:
	docker stop golang-roadmap-app
