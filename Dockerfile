
FROM golang:1.23-alpine3.19 AS builder

ARG VERSION=prod

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk --no-cache add tzdata upx vips

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w -X 'main.Version=${VERSION}'" -tags prod -o /app/chessbet ./cmd/v1/main.go \
    && upx /app/chessbet \
    && wget -q -O /usr/local/bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.5/dumb-init_1.2.5_x86_64 \
    && chmod +x /usr/local/bin/dumb-init \
    && apk del upx


FROM gcr.io/distroless/static:nonroot AS production

LABEL org.opencontainers.image.source=https://github.com/corentings/chessbet
LABEL description="Production stage for Chessbet."

ENV TZ Europe/Paris

WORKDIR /app

COPY --from=builder  /app/chessbet /app/chessbet
COPY --from=builder  /usr/local/bin/dumb-init /usr/bin/dumb-init
COPY --from=busybox:1.37.0-musl /bin/wget /usr/bin/wget

EXPOSE 1815

USER nonroot

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

HEALTHCHECK --interval=5s --timeout=5s --start-period=5s --retries=3 \
    CMD ["/usr/bin/wget", "--no-verbose" ,"--tries=1", "--spider", "http://localhost:1815/health"]

CMD ["/app/chessbet"]
