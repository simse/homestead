FROM ghcr.io/simse/homestead/base:latest

# Labels
LABEL org.opencontainers.image.description "A Homestead dev environment for Go"
LABEL org.opencontainers.image.source="https://github.com/simse/homestead/tree/main/images/go"
LABEL org.opencontainers.image.authors="Simon Sorensen <simse>"

ARG TARGETARCH

# Copy scripts
WORKDIR /tmp/homestead
COPY . /tmp/homestead/

# Install Go
RUN bash install_go.sh
ENV PATH="${PATH}:/usr/local/go/bin"

# Create user
RUN useradd -ms /bin/bash homestead
RUN /usr/sbin/usermod -aG sudo homestead

USER homestead

RUN bash install_dev_tools.sh

RUN mkdir /home/homestead/workspace
WORKDIR /home/homestead/workspace