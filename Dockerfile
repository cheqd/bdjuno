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

WORKDIR /bdjuno
SHELL ["/bin/bash", "-euo", "pipefail", "-c"]

# Copy BDJuno binary
COPY --from=builder /go/src/github.com/forbole/bdjuno/build/bdjuno /usr/bin/bdjuno

# Copy chain-specific config file from Git repo
COPY deploy/* .bdjuno/

# Fetch genesis file for network
ARG NETWORK_NAME="testnet"
RUN wget -q https://raw.githubusercontent.com/cheqd/cheqd-node/main/networks/${NETWORK_NAME}/genesis.json \
	-O .bdjuno/genesis.json

ENTRYPOINT [ "bdjuno start" ]
CMD [ "--home /bdjuno/.bdjuno" ]
