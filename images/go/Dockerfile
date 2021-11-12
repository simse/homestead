FROM ghcr.io/simse/homestead/base:latest

# Labels
LABEL org.opencontainers.image.description "A Homestead dev environment for Go"
LABEL org.opencontainers.image.source="https://github.com/simse/homestead"
LABEL org.opencontainers.image.authors="Simon Sorensen <simse>"

ARG TARGETARCH

USER root

# Copy scripts
WORKDIR /tmp/homestead
COPY . /tmp/homestead/

# Install Go
RUN bash install_go.sh
ENV PATH="${PATH}:/usr/local/go/bin"

# Install Go dev tools
USER homestead
RUN bash install_dev_tools.sh


WORKDIR ${WORKSPACE}