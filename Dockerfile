FROM golang:1.20.0-buster

ENV TZ=Asia/Tokyo
WORKDIR /go/src/app
COPY . ./

RUN go mod tidy && \
    go install github.com/cosmtrek/air@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest

CMD ["air", "-c", ".air.toml"]