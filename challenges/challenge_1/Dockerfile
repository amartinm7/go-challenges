# FROM adv_go_alpine:0.17.1 AS buildMyads

FROM golang:alpine AS buildMyAds
# setup a working folder on the alpine instance
WORKDIR /app
# copy from the outside intellij current folder, all the files to the /app folder
COPY . .

RUN go mod download
# compile the files on the /app to the output folder
RUN CGO_ENABLED=0 go build -o /app/bin/ms-ma-myads cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=buildMyAds /app/bin/ms-ma-myads /go/bin/ms-ma-myads
ENTRYPOINT ["/go/bin/ms-ma-myads"]