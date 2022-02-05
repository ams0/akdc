# this image is built and updated weekly
# https://github.com/cse-labs/codespaces-images

FROM ghcr.io/cse-labs/k3d:latest

# some images require specific values
ARG USERNAME=vscode

# install flux
RUN curl -s https://fluxcd.io/install.sh | bash  && \
    chmod +x /usr/local/bin/flux && \
    mkdir -p /home/$USERNAME/.oh-my-zsh/completions && \
    flux completion zsh > /home/$USERNAME/.oh-my-zsh/completions/_flux && \
    chown -R $USERNAME:$USERNAME /home/$USERNAME && \
    echo "" >> /home/$USERNAME/.zshrc && \
    echo "# helper function to ssh by store name" >> /home/$USERNAME/.zshrc && \
    echo "ss() {" >> /home/$USERNAME/.zshrc && \
    echo "  # run ssh using the ips file" >> /home/$USERNAME/.zshrc && \
    echo '  ssh -p 2222 -o "StrictHostKeyChecking=no" akdc@$(cat ips | grep $1 | tail -1 | cut -f2)' >> /home/$USERNAME/.zshrc && \
    echo "}" >> /home/$USERNAME/.zshrc
