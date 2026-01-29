FROM golang:1.24.12-alpine AS builder

WORKDIR /app

# ビルドに必要なツールをインストール
RUN apk add --no-cache git

# 依存関係をコピー
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# 静的リンクでビルド
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o todo-api cmd/api/main.go

# ビルド確認
# RUN ls -lh todo-api && file todo-api

# 実行用イメージ
FROM alpine:latest

# 必要なランタイムパッケージをインストール
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# ビルド成果物をコピー
COPY --from=builder /app/todo-api ./todo-api
COPY --from=builder /app/.env .env

# 実行権限を付与
RUN chmod +x ./todo-api

EXPOSE 8080

CMD ["./todo-api"]
