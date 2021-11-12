FROM ghcr.io/simse/homestead/base:latest

# Labels
LABEL org.opencontainers.image.description "A Homestead dev environment for Javascript"
LABEL org.opencontainers.image.source="https://github.com/simse/homestead"
LABEL org.opencontainers.image.authors="Simon Sorensen <simse>"

ARG TARGETARCH

# Copy scripts
# WORKDIR /tmp/homestead
# COPY . /tmp/homestead/

# nvm environment variables
ENV NODE_VERSION 16.13.0

# install node and npm
RUN source $NVM_DIR/nvm.sh \
    && nvm install $NODE_VERSION \
    && nvm alias default $NODE_VERSION \
    && nvm use default

# add node and npm to path so the commands are available
ENV NODE_PATH $NVM_DIR/v$NODE_VERSION/lib/node_modules
ENV PATH $NVM_DIR/versions/node/v$NODE_VERSION/bin:$PATH

# Install yarn
RUN npm install -g yarn

WORKDIR ${WORKSPACE}