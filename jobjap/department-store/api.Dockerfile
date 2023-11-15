FROM golang:1.21.1-bullseye AS builder

COPY . /department-store
WORKDIR /department-store
RUN go mod tidy
RUN go build

FROM debian:bullseye-slim
ENV GIN_MODE release

RUN mkdir /app
WORKDIR /app
COPY --from=builder /department-store/department-store /app

EXPOSE 8091

CMD ["/app/department-store"]