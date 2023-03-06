FROM alpine

LABEL org.opencontainers.image.source=https://github.com/janstuemmel/qbit-tools

COPY ./qbit-tools /usr/bin/qbit-tools
RUN chmod a+x /usr/bin/qbit-tools
