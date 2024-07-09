# ctf01d-training-platform

Service used go && psql

Roadmap [here](TODO.md)

## Project struct

```sh
.
├── api # http openapi schema
│   └── openapi.yaml
├── cmd # main app
│   └── ctf01d
│       └── main.go
├── config
│   ├── config.go
│   └── config.yml
├── html
└── internal
    └── app
        ├── database
        ├── db
        │   ├── game.go
        │   ├── result.go
        │   ├── service.go
        │   ├── team.go
        │   ├── university.go
        │   └── user.go
        ├── handlers # backend handlers
        │   ├── games.go
        │   ├── interface.go
        │   ├── results.go
        │   ├── services.go
        │   ├── sessions.go
        │   ├── teams.go
        │   ├── universities.go
        │   └── users.go
        ├── logger
        │   └── logger.go
        ├── repository # work with db, query ...
        │   ├── game.go
        │   ├── result.go
        │   ├── service.go
        │   ├── session.go
        │   ├── team.go
        │   ├── university.go
        │   └── user.go
        ├── server # autogenerated code from oapi-codegen
        │   └── server.gen.go
        ├── utils # shared code
        │   └── helpers.go
        └── view # api json template
            ├── game.go
            ├── result.go
            ├── service.go
            ├── team.go
            ├── university.go
            └── user.go
```

## Development

### Install requirements

```shell
$ go mod download && go mod tidy
```

### Build server

```shell
$ go build cmd/ctf01d/main.go
```

### Run server

```shell
$ go run cmd/ctf01d/main.go
```

will be available on - [http://localhost:4102](http://localhost:4102)


### Generate code from openapi schema

```shell
oapi-codegen -generate models,chi -o internal/app/server/server.gen.go --package routers api/openapi.yaml
```

## DataBase

### psql

connect to db configure in `src/config/config.yaml`

```yaml
...
db:
  driver: postgres
  data_source: postgres://postgres@localhost:5432/ctf01d
...
```

### run db local container

```shell
$ docker run -d --name ctf01d-postgres -e POSTGRES_DB=ctf01d -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 postgres
```

### attach to db container

```shell
$ docker exec -it ctf01d-postgres psql -U postgres
```

## Experimental

### fuzz api

```shell
docker run --net=host --volume /home/user/ctf01d-training-platform/api:/api/ ghcr.io/matusf/openapi-fuzzer run -s '/api/openapi.yaml' --url http://localhost:4102

docker run --net=host --volume /home/user/ctf01d-training-platform/api:/api/ kisspeter/apifuzzer --src_file '/api/openapi.yaml' --url http://localhost:4102 -r /api/
```


## Generate Go server boilerplate from OpenAPI 3 - oapi-codegen

install:

```shell
$ go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
$ export PATH="$PATH:$HOME/bin:$HOME/go/bin"
```


## Database local

```shell
$ make run-db
$ psql postgresql://postgres:postgres@localhost:4112/ctf01d_training_platform
ctf01d_training_platform=#
```
