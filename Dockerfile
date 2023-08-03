FROM golang:alpine

WORKDIR /rmsubekti

COPY . .

RUN go mod tidy

RUN go build -o gotodo

ENTRYPOINT [ "/rmsubekti/gotodo" ]