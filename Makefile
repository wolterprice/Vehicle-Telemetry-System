SHELL := /bin/sh

.PHONY: backend-dev dashboard-dev docker-up

backend-dev:
	cd backend && go run .

dashboard-dev:
	cd dashboard && npm install && npm run dev

docker-up:
	docker compose up --build

