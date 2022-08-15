###############################################################
###        	STAGE 1: Build BDJuno pre-requisites        	###
###############################################################

FROM golang:1.17-alpine AS builder

RUN apk update && apk add --no-cache make git bash

WORKDIR /go/src/github.com/forbole/bdjuno
COPY . ./

RUN go mod download && make build


###############################################################
###       STAGE 2: Copy chain-specific BDJuno config        ###
###############################################################

FROM alpine:3.16 AS bdjuno

# Copy BDJuno binary
COPY --from=builder /go/src/github.com/forbole/bdjuno/build/bdjuno /usr/bin/bdjuno

# Set user directory and details
ARG HOME_DIR="/bdjuno"
ARG USER="bdjuno"

# Add cheqd user to use in the container
RUN addgroup --system ${USER} \
    && adduser ${USER} --system --home ${HOME_DIR} --shell /bin/bash

# Set working directory & bash defaults
WORKDIR ${HOME_DIR}
USER ${USER}
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

# Copy chain-specific config file from Git repo
COPY deploy/* .bdjuno/

# Fetch genesis file for network
ARG NETWORK_NAME="testnet"
RUN wget -q https://raw.githubusercontent.com/cheqd/cheqd-node/main/networks/${NETWORK_NAME}/genesis.json \
    -O ${HOME_DIR}/.bdjuno/genesis.json && chown -R bdjuno:bdjuno /bdjuno

USER ${USER}
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

CMD ["bdjuno", "start",  "--home /bdjuno/.bdjuno" ]
