# this image is built and updated weekly
# https://github.com/cse-labs/codespaces-images

FROM ghcr.io/cse-labs/k3d:latest

# some images require specific values
ARG USERNAME=vscode

# log docker build start
RUN echo "$(date +'%Y-%m-%d %H:%M:%S')    docker build start" >> "/home/${USERNAME}/status"

### todo - hot fix due to az cli package
RUN apt-get upgrade -y
#RUN apt-get update && \
#    apt-get upgrade -y
### end hot fix

# change ownership
RUN chown -R ${USERNAME}:${USERNAME} "/home/${USERNAME}"

# log completion - make sure this is last
RUN echo "$(date +'%Y-%m-%d %H:%M:%S')    docker build complete" >> "/home/${USERNAME}/status"
