# use this Dockerfile to build an image for playing with steve locally.
# Build it with something like `docker build -t steve-testing testing/`.
# Run it with `docker run -it --rm -v $(pwd)/build/steve:/bin/steve steve-testing`.
FROM ubuntu:18.04

RUN apt-get update && apt-get install -y git tree curl
RUN curl -L -o /bin/caddy https://github.com/caddyserver/caddy/releases/download/v2.0.0-beta.13/caddy2_beta13_linux_amd64 && chmod 744 /bin/caddy

# This will be the directory where the test repository is. To prevent issues with Git on the host
# it's being added instead of mounted as a volume.
ADD test-blog /workspace
WORKDIR /workspace

RUN git init && \
	git config --global user.name "test" && \
	git config --global user.email "test@test" && \
	git add . && \
	git commit -m "Initial commit"

