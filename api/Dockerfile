FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY api/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/main api/main.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Saigon
ENV TZ Asia/Saigon

WORKDIR /app
COPY --from=builder /app/main /app/main
COPY --from=builder /app/etc /app/etc

CMD ["./main", "-f", "etc/main.yaml"]
