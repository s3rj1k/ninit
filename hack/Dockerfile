# docker build -t alpine/dnsmasq:latest .
# docker run --interactive --rm --cap-add=NET_ADMIN --name dnsmasq --hostname dnsmasq alpine/dnsmasq:latest
# docker exec -it dnsmasq bash

FROM alpine:3

RUN apk --no-cache add bash coreutils dnsmasq procps psmisc && \
    ln -s /usr/bin/fuser /bin/fuser

EXPOSE 53/tcp 53/udp 67/udp 69/udp
ENTRYPOINT ["/ninit"]

COPY ninit /ninit
COPY zombie /zombie

#ENV INIT_COMMAND_PATH="/zombie"

ENV INIT_COMMAND_PATH="/usr/sbin/dnsmasq"
ENV INIT_COMMAND_ARGS="--no-daemon --user=root"
ENV INIT_WATCH_PATH="/etc/"
