
FROM golang:1.21-bullseye

# Install development tools
RUN apt-get update && apt-get install -y \
    git \
    make \
    pandoc \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Install cosign
RUN go install github.com/sigstore/cosign/cmd/cosign@latest

# Create non-root user
RUN useradd -m -s /bin/bash vscode \
    && mkdir -p /workspace \
    && chown vscode:vscode /workspace

USER vscode
WORKDIR /workspace

# Install dagger CLI
RUN curl -L https://dl.dagger.io/dagger/install.sh | sh

COPY --chown=vscode:vscode . .