# syntax=docker/dockerfile:1

FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/gotodo
COPY . .
RUN go mod tidy
RUN mkdir /build
RUN cp .env /build/
RUN go build -o /build/gotodo main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /rmsubekti
COPY --from=builder /build .
ENTRYPOINT [ "/rmsubekti/gotodo" ]
