FROM golang:1.21.5-alpine

ENV ROOT=/usr/local/go/src/app

RUN apk update && apk add git

RUN mkdir -p ${ROOT}

WORKDIR ${ROOT}

COPY ./app ./

RUN go mod tidy

ENV GO_ENV=dev

EXPOSE 8080

CMD ["go", "run", "main.go"]