ARG BASE=golang:1.16-alpine
FROM ${BASE} as dev

ARG ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update \
    && apk add git
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY . ./

# ホットリロード
RUN go get -u github.com/cosmtrek/air
EXPOSE 8081
CMD air -c .air.toml
