ARG BASE=nginx:latest
FROM ${BASE}

ENV TZ=Asia/Tokyo

RUN echo "${TZ}" > /etc/timezone \
    && dpkg-reconfigure -f noninteractive tzdata

# デフォルトログをLTSV形式に
COPY etc/nginx/conf/ltsv.conf /etc/nginx/conf.d/00_ltsv.conf
# デフォルトログの access_log 設定をコメントアウトする
RUN sed -i 's/access_log/#access_log/' /etc/nginx/nginx.conf
