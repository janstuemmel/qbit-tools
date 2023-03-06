FROM alpine

LABEL org.opencontainers.image.source=https://github.com/janstuemmel/qbit-tools

COPY ./qbit-tools /usr/bin
RUN chmod +x /usr/bin/qbit-tools

CMD qbit-tools
