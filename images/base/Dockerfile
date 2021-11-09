FROM debian:bullseye

ARG TARGETARCH

# Install C build essentials
RUN apt update
RUN apt install -y build-essential nasm make automake pkg-config gcc g++ wget sudo git