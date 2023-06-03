FROM golang:alpine3.13 as builder

RUN apk update && apk upgrade && \
  apk --no-cache --update add git make

WORKDIR /go/src/github.com/Aldiwildan77/rust-updater

# author
LABEL maintainer="Aldiwildan77"

COPY . .

RUN go mod download && \
  go build -v -o engine && \
  chmod +x engine

## Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
  apk --no-cache --update add ca-certificates tzdata && \
  mkdir /app && mkdir rust-updater

WORKDIR /rust-updater

EXPOSE 8000

COPY --from=builder /go/src/github.com/Aldiwildan77/rust-updater/engine /app

CMD /app/engine
