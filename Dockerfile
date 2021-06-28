FROM alpine:3.10.1

ARG version=1.0.0

LABEL maintainer="Takanobu Omura" \
      description="Extension of cat command"

RUN    adduser -D ccat \
    && apk --no-cache add --update --virtual .builddeps curl tar \
#    && curl -s -L -O https://github.com/tamada/ccat/realeases/download/v${version}/ccat-${version}_linux_amd64.tar.gz \
    && curl -s -L -o ccat-${version}_linux_amd64.tar.gz https://drive.google.com/file/d/17C2M0EzQyY-8hxNJHrCOmChY1sTu7vLS/view?usp=sharing \
    && tar xfz ccat-${version}_linux_amd64.tar.gz        \
    && mv ccat-${version} /opt                           \
    && ln -s /opt/ccat-${version} /opt/ccat               \
    && ln -s /opt/ccat-${version}/ccat /usr/local/bin/ccat \
    && rm ccat-${version}_linux_amd64.tar.gz             \
    && apk del --purge .builddeps

ENV HOME="/home/ccat"

WORKDIR /home/ccat
USER    ccat

ENTRYPOINT [ "/opt/ccat/ccat" ]