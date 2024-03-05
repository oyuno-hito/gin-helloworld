FROM golang:1.22.0-alpine
# NOTE: 外部依存ライブラリのダンロードにgitコマンドが必要になる
RUN apk update && apk add git
WORKDIR /go/
COPY go.mod /go/
COPY go.sum /go/
RUN go mod download

ENTRYPOINT ["go", "build", "-o", "/go/build/artifact", "/go/src/main.go"]
