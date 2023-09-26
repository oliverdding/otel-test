FROM --platform=$TARGETPLATFORM golang:1.21-bookworm as builder

ARG TARGETOS TARGETARCH

COPY ./ /go/src/github.com/oliverdding/otel-test

WORKDIR /go/src/github.com/oliverdding/otel-test

RUN <<EOF
set -eux

CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /go/bin/demo-server ./cmd/server/main.go
CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /go/bin/demo-client ./cmd/client/main.go
EOF

FROM debian:bookworm

ENV LANG='en_US.UTF-8' LANGUAGE='en_US:en' LC_ALL='en_US.UTF-8'

RUN <<EOF
set -ex
ln -s /lib /lib64

apt-get update
DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends tzdata curl ca-certificates fontconfig locales

sed -i 's/http:\/\/deb.\(.*\)/https:\/\/deb.\1/g' /etc/apt/sources.list.d/debian.sources
apt-get update
DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends bash neovim iputils-ping less unzip procps
rm /bin/sh
ln -sv /bin/bash /bin/sh
echo "auth required pam_wheel.so use_uid" >> /etc/pam.d/su
chgrp root /etc/passwd && chmod ug+rw /etc/passwd
rm -rf /var/lib/apt/lists/*

echo "en_US.UTF-8 UTF-8" >> /etc/locale.gen
locale-gen en_US.UTF-8
EOF

ENV EDITOR=nvim PAGER=less

ARG USER_UID=10001
USER ${USER_UID}

COPY --from=builder /go/bin/demo-server /usr/local/bin/demo-server
COPY --from=builder /go/bin/demo-client /usr/local/bin/demo-client
