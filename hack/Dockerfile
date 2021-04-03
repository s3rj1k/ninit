# docker build -t alpine/dnsmasq:latest .
# docker run --interactive --rm --cap-add=NET_ADMIN --name dnsmasq --hostname dnsmasq alpine/dnsmasq:latest
# docker exec -it dnsmasq bash

FROM alpine:3

RUN apk --no-cache add bash coreutils dnsmasq procps psmisc && \
    ln -s /usr/bin/fuser /bin/fuser && \
    mkdir -p /etc/k8s.d/

# RUN rm -vf /etc/dnsmasq.conf

EXPOSE 53/tcp 53/udp 67/udp 69/udp
ENTRYPOINT ["/ninit"]

COPY ninit /ninit
COPY ninit-k8s-cm /ninit-k8s-cm

COPY zombie /zombie
# ENV INIT_COMMAND_PATH="/zombie"

ENV INIT_COMMAND_PATH="/usr/sbin/dnsmasq"
ENV INIT_COMMAND_ARGS="--no-daemon --user=root"
# ENV INIT_WORK_DIRECTORY_PATH="/etc/"

# ENV INIT_RELOAD_SIGNAL="SIGHUP"
# ENV INIT_RELOAD_SIGNAL_TO_PGID="true"
ENV INIT_SIGNAL_TO_DIRECT_CHILD_ONLY="true"

ENV INIT_WATCH_INTERVAL="5s"
ENV INIT_WATCH_PATH="/etc/"

# ENV INIT_K8S_BASE_DIRECTORY_PATH="/etc/k8s.d/"
# ENV INIT_K8S_NAMESPACE="default"
# ENV INIT_K8S_CONFIG_MAP_NAME="dnsmasq-config"
