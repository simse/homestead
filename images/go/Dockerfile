FROM simsemand/homestead/base:latest

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
RUN mkdir /home/homestead/workspace
WORKDIR /home/homestead/workspace