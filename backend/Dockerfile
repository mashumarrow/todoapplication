FROM golang:1.23 
WORKDIR /app

# ソースコード全体をコンテナにコピー
COPY . .

# アプリケーションをビルドしてmainバイナリを作成
RUN go build -o main ./cmd

CMD ["./main"]