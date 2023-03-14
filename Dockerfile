FROM golang:1.19-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o out

FROM alpine:3.17

WORKDIR /bin

COPY --from=builder /app/out /bin/out

ENV GIN_MODE=${GIN_MODE}
ENV REDIS_HOST=${REDIS_HOST}
ENV REDIS_PASSWORD=${REDIS_PASSWORD}
ENV REDIS_USER=${REDIS_USER}
ENV PORT=${PORT}

EXPOSE ${PORT}

CMD [ "/bin/out" ]
