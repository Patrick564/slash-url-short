# syntax=docker/dockerfile:1
FROM 1.19.5-bullseye as development

ENV GOOS linux
ENV CGO_ENABLED 0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app

EXPOSE 8080

CMD ./app
