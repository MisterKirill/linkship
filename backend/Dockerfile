FROM golang:1.23-alpine

WORKDIR /usr/src/app

COPY . .
RUN go mod download

RUN go build -o /linkship

CMD ["/linkship"]
