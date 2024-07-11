# TweetManager
Tweet caching system that collects tweets with specific phrase with or without emoticon

## Running
create secrets.yaml in current working directory and than:

```bash
go run ./cmd/...

```

## Endpoints 

Get tweets with INTERESTING_SYMBOL for specific time frame. Note INTERESTING_SYMBOL must be first imported into DB.

```bash

curl http://localhost:8080/tweets?symbol=INTERESTING_SYMBOL&start_date=2022-01-02&end_date=2022-01-04

```

## Development

#### Test database 

```
mkdir -p /postgresql-data
docker run -d \
    --rm \
	--name postgres-tweeter \
	-e POSTGRES_PASSWORD=mysecretpassword \
	-e PGDATA=/var/lib/postgresql/data/pgdata \
	-v $PWD/postgresql-data:/var/lib/postgresql/data \
    -p 127.0.0.1:25432:5432 \
	postgres:13.14
```


https://gobuffalo.io/documentation/database/migrations/ 
https://github.com/gobuffalo/fizz/blob/main/README.md
```
soda  generate fizz CreateTweetsTable

soda create
soda migrate
soda migrate down
soda reset
```

## X (twitter) information

Permissions based on product used: https://developer.x.com/en/portal/products/basic as of now Basic is needed to get search recent tweets 

