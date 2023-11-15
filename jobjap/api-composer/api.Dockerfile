FROM golang:1.21.1-bullseye AS builder

COPY . /api-composer
WORKDIR /api-composer
RUN go mod tidy
RUN go build

FROM debian:bullseye-slim
ENV GIN_MODE release

RUN mkdir /app
WORKDIR /app
COPY --from=builder /api-composer/api-composer /app

EXPOSE 8000

CMD ["/app/api-composer"]