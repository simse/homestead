FROM debian:bullseye

# Labels
LABEL org.opencontainers.image.description "The base image that is used as a starting point for all Homestead language-specific images"
LABEL org.opencontainers.image.source="https://github.com/simse/homestead"
LABEL org.opencontainers.image.authors="Simon Sorensen <simse>"

ARG TARGETARCH

# Install C build essentials
RUN apt update
RUN apt install -y build-essential nasm make automake pkg-config gcc g++ wget sudo git

# Create user
RUN useradd -ms /bin/bash homestead
RUN /usr/sbin/usermod -aG sudo homestead

USER homestead

# Install NVM
SHELL ["/bin/bash", "-c"] 
ENV NVM_DIR /home/homestead/nvm
RUN mkdir ${NVM_DIR}
RUN wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash

# Create workspace folder
ENV WORKSPACE /home/homestead/workspace
RUN mkdir ${WORKSPACE}
