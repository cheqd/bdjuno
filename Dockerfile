FROM golang:1.17-alpine AS builder
RUN apk update && apk add --no-cache make git
WORKDIR /bdjuno
COPY . ./
RUN make install && make build

FROM alpine:latest
WORKDIR /bdjuno
COPY --from=builder build/bdjuno /usr/bin/bdjuno
COPY deploy/config.yaml .bdjuno/config.yaml
CMD [ "bdjuno" ]
