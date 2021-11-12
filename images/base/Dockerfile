FROM debian:bullseye

# Labels
LABEL org.opencontainers.image.description "The base image that is used as a starting point for all Homestead language-specific images"
LABEL org.opencontainers.image.source="https://github.com/simse/homestead"
LABEL org.opencontainers.image.authors="Simon Sorensen <simse>"

ARG TARGETARCH

# Install C build essentials
RUN apt update
RUN apt install -y build-essential nasm make automake pkg-config gcc g++ wget sudo git

# Install NVM
RUN wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash