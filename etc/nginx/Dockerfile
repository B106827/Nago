ARG BASE=nginx:latest
FROM ${BASE}

ENV TZ=Asia/Tokyo

RUN echo "${TZ}" > /etc/timezone \
    && dpkg-reconfigure -f noninteractive tzdata

COPY etc/nginx/conf/http.conf /etc/nginx/conf.d/http.conf
# デフォルトログをLTSV形式に
COPY etc/nginx/conf/ltsv.conf /etc/nginx/conf.d/00_ltsv.conf
# デフォルトの設定で不要なものをコメントアウト
RUN sed -i 's/keepalive_timeout/#keepalive_timeout/' /etc/nginx/nginx.conf
RUN sed -i 's/access_log/#access_log/' /etc/nginx/nginx.conf
