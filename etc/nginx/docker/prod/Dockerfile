ARG BASE=nginx:latest
FROM ${BASE}

ENV NGINX_PORT=80
# ECSの同一タスク定義内のコンテナ間はlocalhostで通信可能である
ENV NGINX_PROXY_API=localhost
ENV TZ=Asia/Tokyo

RUN echo "${TZ}" > /etc/timezone \
    && dpkg-reconfigure -f noninteractive tzdata

COPY etc/nginx/conf/prod/nago.conf /etc/nginx/templates/default.conf.template
COPY etc/nginx/conf/prod/http.conf /etc/nginx/conf.d/http.conf
# デフォルトログをLTSV形式に
COPY etc/nginx/conf/common/ltsv.conf /etc/nginx/conf.d/00_ltsv.conf
# デフォルトの設定で不要なものをコメントアウト
RUN sed -i 's/keepalive_timeout/#keepalive_timeout/' /etc/nginx/nginx.conf
RUN sed -i 's/access_log/#access_log/' /etc/nginx/nginx.conf
# （本番用）フロント側コードを展開
COPY ./front/build/ /var/www/html/public/
