migrate:
	docker-compose run --rm be sh -c "goose -path='infra/datastore' up"

migrate/down:
	docker-compose run --rm be sh -c "goose -path='infra/datastore' down"

migrate/status:
	docker-compose run --rm be sh -c "goose -path='infra/datastore' status"