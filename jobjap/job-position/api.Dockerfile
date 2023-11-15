FROM golang:1.21.1-bullseye AS builder

COPY . /job-position
WORKDIR /job-position
RUN go mod tidy
RUN go build

FROM debian:bullseye-slim
ENV GIN_MODE release

RUN mkdir /app
WORKDIR /app
COPY --from=builder /job-position/job-position /app

EXPOSE 8002

CMD ["/app/job-position"]