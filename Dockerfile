FROM --platform=linux/amd64 golang:1.24-alpine

WORKDIR /service

COPY go.mod go.sum /service/

RUN go mod download

COPY . /service/

RUN go build -a -o ./bin/api ./cmd/api/

CMD [ "/service/bin/api" ]

