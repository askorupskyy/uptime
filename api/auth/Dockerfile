FROM golang:latest

LABEL maintainer = "rcbxd"

WORKDIR /usr/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build

EXPOSE 9000

CMD ["./auth"]
