ARG BASE=node:16.14.0-alpine
FROM ${BASE}

ARG ROOT=/usr/src/app
WORKDIR ${ROOT}

COPY ./package.json ./
COPY ./yarn.lock ./
RUN yarn install --network-timeout 240000
COPY . ./

EXPOSE 3001
