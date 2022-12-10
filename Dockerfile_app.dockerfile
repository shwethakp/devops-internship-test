FROM golang:1.18-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /mybuild

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v
RUN chown -R root  /mybuild/main.go

# Build the Go app
RUN go build -o /mybuild/main.go .

# Start fresh from a smaller image
FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=build_base /mybuild/main.go /mybuild/main.go

# This container exposes port 8080 to the outside world
EXPOSE 8080

USER root:root
