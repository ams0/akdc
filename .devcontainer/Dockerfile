# this image is built and updated weekly
# https://github.com/cse-labs/codespaces-images

FROM ghcr.io/cse-labs/k3d:latest

# some images require specific values
ARG USERNAME=vscode

ENV PATH $PATH:/usr/local/go/bin:/home/$USERNAME/go/bin

# install flux
RUN echo "$(date)    docker build start" >> /home/$USERNAME/status && \
    mkdir -p /home/$USERNAME/.local/bin && \
    mkdir -p /home/$USERNAME/.oh-my-zsh/completions && \
    mkdir -p /home/$USERNAME/go && \
    echo "compinit" >> /home/$USERNAME/.zshrc && \
    curl -s https://fluxcd.io/install.sh | bash  && \
    chmod +x /usr/local/bin/flux && \
    flux completion zsh > /home/$USERNAME/.oh-my-zsh/completions/_flux && \
    wget https://go.dev/dl/go1.17.6.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.17.6.linux-amd64.tar.gz && \
    rm -f go1.17.6.linux-amd64.tar.gz && \
    echo "$(date)    docker build complete" >> /home/$USERNAME/status && \
    chown -R $USERNAME:$USERNAME /home/$USERNAME
