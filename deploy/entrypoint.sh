#!/bin/bash

set -euo pipefail

# Inject environment variables into config.yaml file

sed -i "s, RPC_ADDRESS, '${RPC_ADDRESS}',g" /bdjuno/.bdjuno/config.yaml
sed -i "s, GRPC_ADDRESS, '${GRPC_ADDRESS}',g" /bdjuno/.bdjuno/config.yaml
sed -i "s, LOG_LEVEL, '${LOG_LEVEL}',g" /bdjuno/.bdjuno/config.yaml
sed -i "s, DATABASE_URL, '${DATABASE_URL}',g" /bdjuno/.bdjuno/config.yaml

bdjuno start --home /bdjuno/.bdjuno
