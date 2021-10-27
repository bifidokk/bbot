FROM golang:1.12-alpine as builder

RUN apk update && apk upgrade && \
    apk add build-base && \
    apk add --no-cache bash git openssh

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . ./
RUN go build -o main .


FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app .

EXPOSE 9000
CMD ["./main"]