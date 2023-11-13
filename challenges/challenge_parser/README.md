# Parser

## create a empty module

```bash
# create a folder
mkdir /challenges/challenges_parser
# init the go module. It creates a go.mod file inside
go mod init github.com/learning-go/challenges/challenges_parser
#
go mod tidy
```
Once it's done, you have a go.mod file inside the folder like this: 

```bash
module github.com/antonio-martin/learning-go/challenges/challenges_parser

go 1.21.1
```

## execute application

Over the main root of the project /challenges/challenge_parser execute

```bash
go run cmd/cli/main.go
```

## compile and build the artifact

```bash
# build
go build ./...
# or
go build -o ./bin ./...
# or to generate a binary executable
go build -o ./bin/main cmd/cli/main.go
# setup grants
chmod +x ./bin/main
# execute
./bin/main 
```

## adding/install dependencies 

```bash
go get  github.com/stretchr/testify/mock

```

## testing
```bash
go test ./...
#or
go test github.com/yourusername/yourproject/...
```

