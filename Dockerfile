FROM golang:1.16.6-alpine3.14 AS builder
RUN mkdir /build
COPY . /build
WORKDIR /build
RUN go build

FROM alpine:3.14
RUN adduser -S -D -H -h /app appuser
RUN mkdir /app
RUN mkdir /app/credential
RUN chown appuser /app
RUN chown appuser /app/credential
USER appuser
COPY --from=builder /build/almercadito-api-restful /app/
WORKDIR /app
CMD ["./almercadito-api-restful","-workdir","./credential"]
