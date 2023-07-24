# syntax=docker/dockerfile:1

FROM golang:1.20.6-alpine
WORKDIR /app

COPY ./ ./
RUN go mod download

RUN go build -o ./lla-app

EXPOSE 8080

CMD [ "./lla-app" ]
