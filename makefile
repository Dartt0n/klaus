all: build run

build:
	docker build -t klaus-image .

run:
	docker run --rm --name klaus klaus-image