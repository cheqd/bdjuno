#!/bin/bash

set -euo pipefail

# Inject environment variables into config.yaml file

sed -i "s, RPC_ADDRESS, '${RPC_ADDRESS}',g" /callisto/.callisto/config.yaml
sed -i "s, GRPC_ADDRESS, '${GRPC_ADDRESS}',g" /callisto/.callisto/config.yaml
sed -i "s, LOG_LEVEL, '${LOG_LEVEL}',g" /callisto/.callisto/config.yaml
sed -i "s, DATABASE_URL, '${DATABASE_URL}',g" /callisto/.callisto/config.yaml

callisto start --home /callisto/.callisto
