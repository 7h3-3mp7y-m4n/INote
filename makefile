run:
	go run ./cmd/

build:
	go build -o inote ./cmd/

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down
