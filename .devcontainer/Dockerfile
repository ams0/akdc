# this image is built and updated weekly
# https://github.com/cse-labs/codespaces-images

FROM ghcr.io/cse-labs/k3d:latest

# some images require specific values
ARG USERNAME=vscode

# install flux
RUN mkdir -p $HOME/.local/bin && \
    curl -s https://fluxcd.io/install.sh | bash  && \
    chmod +x /usr/local/bin/flux && \
    mkdir -p /home/$USERNAME/.oh-my-zsh/completions && \
    flux completion zsh > /home/$USERNAME/.oh-my-zsh/completions/_flux && \
    chown -R $USERNAME:$USERNAME /home/$USERNAME
