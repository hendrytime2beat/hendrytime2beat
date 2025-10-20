# KTFS

## prerequisite

as installed on dev machine

- git 2.39.5
- go 1.25.3
- docker 28.5.1
- docker compose 2.40.0

## installation

```shell
// clone repository
git clone https://github.com/yogarypr/ktfsfrtzgo.git

// install dependencies
cd ktfsfrtzgo && go clean && go get && go install

// copy environment example into environment
cp .env.example .env
```

## running

### local

directly start http service, run

```shell
go run main.go
```

for database migration, add flag `-migrate=1` or `-migrate=t` or `-migrate=true`

```shell
go run main.go -migrate=true
```

### live reload

install [cosmtrek/air](https://github.com/air-verse/air?tab=readme-ov-file#installation), then run

```shell
$GOPATH/bin/air
```

## docker

```shell
docker-compose build
docker-compose up -d
```