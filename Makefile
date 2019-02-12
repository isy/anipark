migrate:
	docker-compose run --rm be sh -c "goose up"

migrate/down:
	docker-compose run --rm be sh -c "goose down"

migrate/status:
	docker-compose run --rm be sh -c "goose status"