FROM golang:1.21 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

ARG DB_DSN
ARG JWT_SECRET
ARG JWT_TTL
ARG SENTRY_DSN

RUN echo "$DB_DSN" > "./.env"
RUN echo "$JWT_SECRET" >> "./.env"
RUN echo "$JWT_TTL" >> "./.env"
RUN echo "$SENTRY_DSN" >> "./.env"

RUN go build ./cmd/app/main.go


FROM alpine:latest
RUN apk add --no-cache ca-certificates libc6-compat

WORKDIR /app
COPY --from=builder /app .

EXPOSE 9002
CMD ["./main"]