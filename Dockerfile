FROM golang:1.9.2

RUN mkdir -p /app

RUN go get "github.com/gin-gonic/gin"

RUN go get -v "github.com/spf13/viper"

WORKDIR /app

ADD . /app

RUN go build ./server.go

CMD [ "./server"]

