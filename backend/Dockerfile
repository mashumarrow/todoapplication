FROM golang:1.23 

WORKDIR /go/src/github.com/mashumarrow/todoapplication/backend/

COPY go.mod ./
COPY go.sum ./

RUN go mod download
# ソースコード全体をコンテナにコピー
COPY . .

COPY  /cmd/.env .env
COPY .air.toml .air.toml

# migration用にgooseをインストール
RUN go install github.com/pressly/goose/v3/cmd/goose@v3.21.1

#Airをインストール
RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]