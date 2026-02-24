ARG APP_NAME=kiwi

FROM golang:bullseye AS builder
ARG APP_NAME

WORKDIR /app
COPY . .

RUN apt-get update && apt-get install -y upx && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o  ./... && \
    upx --best --lzma 

FROM gcr.io/distroless/static
ARG APP_NAME

COPY --from=builder /app/ /usr/local/bin/app

ENTRYPOINT ["/usr/local/bin/app"]