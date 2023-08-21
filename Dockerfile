FROM golang:1.20.6-alpine3.18

WORKDIR /learn-ci

COPY go.mod /learn-ci

COPY . /learn-ci

RUN go mod tidy

RUN go build -o /learn-ci/bin/main /learn-ci/main.go

EXPOSE 80

CMD ["/learn-ci/bin/main"]