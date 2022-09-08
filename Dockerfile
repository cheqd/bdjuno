###############################################################
###        	STAGE 1: Build BDJuno pre-requisites        	###
###############################################################

FROM golang:1.19.1-alpine AS builder

RUN apk update && apk add --no-cache make git bash

WORKDIR /go/src/github.com/forbole/bdjuno
COPY . ./

RUN go mod download && make build


###############################################################
###       STAGE 2: Copy chain-specific BDJuno config        ###
###############################################################

FROM alpine:3.16 AS bdjuno

RUN apk update && apk add --no-cache bash ca-certificates curl

# Copy BDJuno binary
COPY --from=builder /go/src/github.com/forbole/bdjuno/build/bdjuno /usr/local/bin/bdjuno

# Set user directory and details
ARG HOME_DIR="/bdjuno"
ARG USER="bdjuno"
SHELL ["/bin/sh", "-euo", "pipefail", "-c"]

# Add non-root user to use in the container
RUN addgroup --system $USER \
    && adduser $USER --system --home $HOME_DIR --shell /bin/bash

# Set working directory & bash defaults
WORKDIR $HOME_DIR
USER $USER

# Copy chain-specific config file from Git repo
COPY --chown=$USER:$USER deploy/ .bdjuno/
RUN mv .bdjuno/entrypoint.sh . && chmod +x entrypoint.sh

ENTRYPOINT [ "/bdjuno/entrypoint.sh" ]
