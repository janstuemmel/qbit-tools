FROM alpine

COPY ./qbit-tools /usr/local/bin
RUN chmod +x /usr/local/bin/qbit-tools

CMD qbit-tools