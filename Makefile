start-dev:
	docker-compose --env-file=dev.env up -d

stop-dev:
	docker-compose down