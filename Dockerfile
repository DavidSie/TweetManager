# syntax=docker/dockerfile:1
FROM golang:1.22 AS build
WORKDIR /app
COPY . .
ARG CGO_ENABLED=1
RUN go mod download
RUN go build -o bin/tweetManager ./cmd/...



FROM ubuntu:22.04
COPY --from=build /app/bin/tweetManager /bin/tweetManager

ENTRYPOINT ["/bin/tweetManager"]
