#It's not working
FROM golang:alpine

WORKDIR /simple-go-mysql

ADD . .

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -command="./simple-go-mysql/cmd/main"