FROM golang:1.13.1 as builder

MAINTAINER MDH <dh.duyhai@gmail.com>
RUN mkdir /app
WORKDIR /app

COPY . /app

RUN go mod download

RUN GOOS=linux

RUN go build -o main ./main.go

FROM ubuntu:16.04

RUN apt-get update && apt-get install -y locales && rm -rf /var/lib/apt/lists/*

WORKDIR /app
RUN mkdir storage
RUN echo 123 > storage/test_save.txt

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 80

CMD ["/app/main"]
