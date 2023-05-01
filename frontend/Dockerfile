FROM node:18.12 AS builder

WORKDIR /usr/src/app
ADD . /usr/src/app

RUN yarn install && \
    yarn generate

FROM nginx:latest
COPY --from=builder /usr/src/app/.output/public /usr/share/nginx/html

EXPOSE 80