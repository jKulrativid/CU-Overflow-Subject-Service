FROM golang:1.21.1-alpine3.18

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./src/*.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /subject-service

EXPOSE 8081

CMD ["/subject-service"]
