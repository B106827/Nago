ARG BASE=golang:1.18-alpine
FROM ${BASE}

ARG ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update \
    && apk add git \
    && apk --no-cache add tzdata \
    && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && apk del tzdata
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
RUN go mod tidy
COPY . ./

# ホットリロード
RUN go install github.com/cosmtrek/air@latest
EXPOSE 8081
CMD air -c .air.toml
