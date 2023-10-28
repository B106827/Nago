ARG BASE=nginx:latest
FROM ${BASE}

ENV NGINX_PORT=80
ENV NGINX_PROXY_API=nago-api
ENV TZ=Asia/Tokyo

RUN echo "${TZ}" > /etc/timezone \
    && dpkg-reconfigure -f noninteractive tzdata

COPY etc/nginx/conf/dev/nago.conf /etc/nginx/templates/default.conf.template
COPY etc/nginx/conf/dev/http.conf /etc/nginx/conf.d/http.conf
# デフォルトログをLTSV形式に
COPY etc/nginx/conf/common/ltsv.conf /etc/nginx/conf.d/00_ltsv.conf
# デフォルトの設定で不要なものをコメントアウト
RUN sed -i 's/keepalive_timeout/#keepalive_timeout/' /etc/nginx/nginx.conf
RUN sed -i 's/access_log/#access_log/' /etc/nginx/nginx.conf
