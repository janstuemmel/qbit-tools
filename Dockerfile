FROM alpine

COPY ./qbit-tools /usr/bin
RUN chmod +x /usr/bin/qbit-tools

CMD qbit-tools
