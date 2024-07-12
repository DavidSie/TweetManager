# syntax=docker/dockerfile:1
FROM golang:1.22 AS build
WORKDIR /app
COPY . .
ARG CGO_ENABLED=1
RUN go mod download
RUN go build -o bin/tweetManager ./cmd/...
# install database migration tool
RUN go install github.com/gobuffalo/pop/v6/soda@latest



FROM ubuntu:22.04
COPY --from=build /app/bin/tweetManager /bin/tweetManager
COPY --from=build /go/bin/soda /bin/soda
COPY ./migrations /migrations

ENTRYPOINT ["/bin/tweetManager"]
