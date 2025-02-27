FROM golang:1.21.6-alpine AS builder
LABEL stage=gobuilder

ARG DEV_MODE

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY . .
RUN cd "$DEV_MODE"
#COPY "app/etc/$DEV_MODE-api.yaml" "./app/etc/$DEV_MODE-api.yaml"

RUN go build -ldflags="-s -w" -o ./app/app app/app.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

ENV ENV_MODE=config
WORKDIR /app
COPY --from=builder /build/app/app ./app
COPY --from=builder /build/app/migrations/ ./migrations/
COPY --from=builder /build/app/etc/app-api.yaml app-api.yaml

EXPOSE 8888
CMD ["./app", "-f", "app-api.yaml"]
