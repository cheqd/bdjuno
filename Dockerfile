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
COPY --from=builder /go/src/github.com/forbole/bdjuno/build/bdjuno /usr/local/bin/bdjuno

# Set user directory and details
ARG HOME_DIR="/bdjuno"
ARG USER="bdjuno"

# Add non-root user to use in the container
RUN addgroup --system ${USER} \
    && adduser ${USER} --system --home ${HOME_DIR} --shell /bin/sh

# Set working directory & bash defaults
WORKDIR ${HOME_DIR}
USER ${USER}
SHELL ["/bin/sh", "-euo", "pipefail", "-c"]

# Copy chain-specific config file from Git repo
COPY deploy/ .bdjuno/

ENTRYPOINT [ "bdjuno start" ]
CMD [ "--home /bdjuno/.bdjuno" ]
