FROM golang:1.18-alpine

WORKDIR /go/src/app

RUN apk update \
    && apk add git mysql-client mysql-dev\
    && apk --no-cache add tzdata \
    && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && apk del tzdata
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
RUN go mod tidy
COPY . ./

EXPOSE 8081

RUN go build -ldflags='-s -w' -o dist/api_server
CMD dist/api_server -e=production
