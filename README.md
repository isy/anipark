# go-sandbox

Golang 勉強用の素敵な Repo

## Build Setup

```bash
# start server at localhost: 3001(be)
$ docker-compose up

# db migrate
$ make migrate

# db migrate down
$ make migrate/down

# migrate status
$ make migrate/status

# Create migrate file
$ goose -path="infra/datastore" create filename sql
```
