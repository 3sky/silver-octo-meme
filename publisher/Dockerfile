# Dockerfile
FROM golang:alpine as builder
RUN apk add --no-cache git gcc libc-dev libzmq pkgconf zeromq-dev
WORKDIR /build/app
# Get depedences
COPY go.mod ./
RUN go mod download
COPY . .
# Build app
RUN go build -o pub
FROM alpine
RUN apk add --no-cache libzmq
COPY --from=builder /build/app/pub ./pub
ENTRYPOINT ["./pub"]

EXPOSE 3000
EXPOSE 5553
