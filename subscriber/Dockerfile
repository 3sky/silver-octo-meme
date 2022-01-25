FROM golang:alpine as builder
RUN apk add --no-cache git gcc libc-dev libzmq pkgconf zeromq-dev
WORKDIR /build/app
# Get depedences
COPY go.mod ./
RUN go mod download
COPY . .
# Build app
RUN go build -o sub
FROM alpine
RUN apk add --no-cache libzmq
COPY --from=builder /build/app/sub ./sub
ENTRYPOINT ["./sub"]
