.PHONY: build run docker-build docker-run docker-sh docker-logs docker-stop

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
