FROM mcr.microsoft.com/devcontainers/go:1.22

# Install additional fonts for multi-language support
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
    fonts-noto-cjk \
    fonts-noto-core \
    fonts-noto-extra \
    fonts-noto-ui-core \
    fonts-dejavu \
    fonts-liberation \
    && rm -rf /var/lib/apt/lists/*

# Install Dagger CLI
RUN go install dagger.io/dagger@v0.9.3

# Install Node.js dependencies
RUN npm install -g @marp-team/marp-cli
