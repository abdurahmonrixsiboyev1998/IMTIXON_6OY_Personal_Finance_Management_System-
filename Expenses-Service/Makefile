.PHONY: build run migrate

build:
	docker build -t personal-finance .

run:
	docker run -p 8080:8080 -p 50051:50051 personal-finance

migrate:
	go run db/migrate.go