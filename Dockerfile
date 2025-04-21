# Base image with Playwright
FROM mcr.microsoft.com/playwright:v1.52.0-noble

# Install system dependencies
RUN apt-get update && \
  apt-get install -y \
  make \
  git \
  jq \
  curl \
  wget \
  apt-transport-https \
  ca-certificates \
  gnupg-agent \
  software-properties-common \
  && rm -rf /var/lib/apt/lists/*

# Install Docker CLI, Docker Compose v2, and Docker Buildx
RUN mkdir -p /etc/apt/keyrings && \
  curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg && \
  echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu noble stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null && \
  apt-get update && \
  apt-get install -y docker-ce-cli docker-compose-plugin && \
  rm -rf /var/lib/apt/lists/*

# Install kubectl
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && \
  chmod +x kubectl && \
  mv kubectl /usr/local/bin/

# Install kustomize
RUN curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh" | bash && \
  mv kustomize /usr/local/bin/

# Install yq (YAML processor)
RUN wget https://github.com/mikefarah/yq/releases/latest/download/yq_linux_amd64 -O /usr/local/bin/yq && \
  chmod +x /usr/local/bin/yq

# Install Helm
RUN curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

# Install KinD (Kubernetes in Docker)
RUN curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.20.0/kind-linux-amd64 && \
  chmod +x ./kind && \
  mv ./kind /usr/local/bin/kind

# Install Go based on go.mod version (will be set during build)
ARG GO_VERSION=1.24.2
RUN curl -LO https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
  rm -rf /usr/local/go && tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
  rm go${GO_VERSION}.linux-amd64.tar.gz && \
  ln -s /usr/local/go/bin/go /usr/local/bin/go && \
  ln -s /usr/local/go/bin/gofmt /usr/local/bin/gofmt

# Add GitHub Actions user and set permissions
RUN groupadd -g 121 docker && \
  useradd -u 1001 -g docker github && \
  usermod -aG docker github

# Set up workspace
USER github
WORKDIR /home/github
ENV PATH=$PATH:/usr/local/go/bin
