ARG BASE=node:16.14.0-alpine
FROM ${BASE} as dev

ARG ROOT=/usr/src/app
WORKDIR ${ROOT}

COPY ./package.json ./
COPY ./yarn.lock ./
RUN yarn install
COPY . ./

EXPOSE 3001
CMD yarn start
