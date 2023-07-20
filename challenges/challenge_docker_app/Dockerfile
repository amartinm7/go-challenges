FROM golang:1.19 AS buildMyApp

WORKDIR /app

COPY . .
RUN go mod download

# compile the files on the /app to the output folder
RUN CGO_ENABLED=0 go build -o /app/bin/my-docker-app cmd/api/main.go

FROM scratch
COPY --from=buildMyApp /app/bin/my-docker-app /go/bin/my-docker-app
ENTRYPOINT ["/go/bin/my-docker-app"]