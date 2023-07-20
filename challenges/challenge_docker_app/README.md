<p align="center">
    <img alt="&quot;a random gopher created by gopherize.me&quot;" src="../../img/gopher-challenge-1.png" width="200px" style="display: block; margin: 0 auto"/>
</p>

<h1 align="center">
  Challenge Dockerize App
</h1>

# index
[hexagonal_architecture.md](_docs%2Fddd%2Fhexagonal_architecture.md)

## Dockerize App

Your task
Write a program that prints `Hello, Docker! <3` message when your do a call:
```bash
curl -v location "http://localhost:8000/health"
```

## main steps

### create module app
```bash
go mod init challenges/challenge_dockerize_app
```

### creates project folders
```bash
mkdir _deploy
mkdir cmd
mkdir internal
```

### create docker files
```bash
touch Dockefile
touch docker-compose.yml
touch docker-build.sh
touch docker-run.sh
```

### build app
```bash
go build -o main main.go
go run main.go
go build ./...
go test ./...
go mod tidy
```

### format code
```bash
go fmt -w main.go
```

## Resources

[1] Non-official project structure standard for Go projects: https://github.com/golang-standards/project-layout#go-directories

[2] Build your Go image https://docs.docker.com/language/golang/build-images/
