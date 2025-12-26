FROM docker.io/golang:alpine

RUN apk add --no-cache \
    git \
    curl \
    wget \
    unzip \
    bash \
    build-base

RUN addgroup -g 1000 vscode && adduser -D -G vscode -u 1000 -h /home/vscode -s /bin/bash vscode

# Go tools
RUN go install github.com/go-delve/delve/cmd/dlv@latest \
 && go install golang.org/x/tools/cmd/goimports@latest \
 && go install golang.org/x/tools/gopls@latest

# Ensure writable GOPATH
RUN mkdir -p /go && chown -R vscode:vscode /go /home/vscode

USER vscode
WORKDIR /workspace
EXPOSE 2345
