###############################################################
###        	STAGE 1: Build BDJuno pre-requisites        	###
###############################################################

FROM golang:1.18-alpine AS builder

RUN apk update && apk add --no-cache make git bash

WORKDIR /go/src/github.com/forbole/bdjuno
COPY . ./

######################################################
## Enabe the lines below if chain supports cosmwasm ##
## module to properly build docker image            ##
######################################################
#RUN apk update && apk add --no-cache ca-certificates build-base git
#ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.1.1/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
#ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.1.1/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
#RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep 9ecb037336bd56076573dc18c26631a9d2099a7f2b40dc04b6cae31ffb4c8f9a
#RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep 6e4de7ba9bad4ae9679c7f9ecf7e283dd0160e71567c6a7be6ae47c81ebe7f32
## Copy the library you want to the final location that will be found by the linker flag `-lwasmvm_muslc`
#RUN cp /lib/libwasmvm_muslc.$(uname -m).a /lib/libwasmvm_muslc.a

RUN go mod download && make build

##################################################
## Enabe line below if chain supports cosmwasm  ##
## module to properly build docker image        ##
##################################################
#RUN LINK_STATICALLY=true BUILD_TAGS="muslc" make build


###############################################################
###       STAGE 2: Copy chain-specific BDJuno config        ###
###############################################################

FROM alpine:3 AS bdjuno

##################################################
## Enabe line below if chain supports cosmwasm  ##
## module to properly build docker image        ##
##################################################
#RUN apk update && apk add --no-cache ca-certificates build-base
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