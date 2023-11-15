FROM golang:1.21.1-bullseye AS builder

COPY . /parking
WORKDIR /parking
RUN go mod tidy
RUN go build

FROM debian:bullseye-slim
ENV GIN_MODE release

RUN mkdir /app
WORKDIR /app
COPY --from=builder /parking/parking /app

EXPOSE 8003

CMD ["/app/parking"]