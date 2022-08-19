#!/bin/bash

set -euo pipefail

sed -i "s, RPC_ADDRESS, '${RPC_ADDRESS}',g" /bdjuno/.bdjuno/config.yaml
sed -i "s, GRPC_ADDRESS, '${GRPC_ADDRESS}',g" /bdjuno/.bdjuno/config.yaml
sed -i "s, LOG_LEVEL, '${LOG_LEVEL}',g" /bdjuno/.bdjuno/config.yaml
sed -i "s, DATABASE_NAME, '${DATABASE_NAME}',g" /bdjuno/.bdjuno/config.yaml
sed -i "s, DATABASE_HOST, '${DATABASE_HOST}',g" /bdjuno/.bdjuno/config.yaml
sed -i "s, DATABASE_PORT, '${DATABASE_PORT}',g" /bdjuno/.bdjuno/config.yaml
sed -i "s, DATABASE_USER, '${DATABASE_USER}',g" /bdjuno/.bdjuno/config.yaml
sed -i "s, DATABASE_PASSWORD, '${DATABASE_PASSWORD}',g" /bdjuno/.bdjuno/config.yaml

bdjuno start --home /bdjuno/.bdjuno
