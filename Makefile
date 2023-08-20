# Makefile

build:
	docker-compose build

rebuild:
	docker-compose build --no-cache

up:
	docker-compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f

.PHONY: build rebuild up down logs