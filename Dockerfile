FROM golang:1.18.4-alpine as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR $GOPATH/src/github.com/alirezazeynali75/simple-to-do-app

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD [ "simple-to-do-app" ]