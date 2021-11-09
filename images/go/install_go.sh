echo $TARGETARCH
echo $ARCH
wget "https://golang.org/dl/go1.17.3.linux-${TARGETARCH}.tar.gz"
tar -C /usr/local -xzf "go1.17.3.linux-${TARGETARCH}.tar.gz"