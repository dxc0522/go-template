FROM golang:1.21.6-alpine AS builder
LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY . .
RUN cd app

RUN go build -ldflags="-s -w" -o server app/app.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /build/server ./
COPY --from=builder /build/app/migrations/ ./migrations/
COPY --from=builder /build/app/etc/app.yaml app.yaml

EXPOSE 8888
CMD ["./server", "-f", "app.yaml"]
