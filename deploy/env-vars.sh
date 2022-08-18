#!/bin/bash

set -euo pipefail

sed -i "s, RPC_ADDRESS, '${RPC_ADDRESS}',g" .bdjuno/config.yaml
sed -i "s, GRPC_ADDRESS, '${GRPC_ADDRESS}',g" .bdjuno/config.yaml
sed -i "s, LOG_LEVEL, '${LOG_LEVEL}',g" .bdjuno/config.yaml
sed -i "s, DATABASE_NAME, '${DATABASE}',g" .bdjuno/config.yaml
sed -i "s, DATABASE_HOST, '${HOSTNAME}',g" .bdjuno/config.yaml
sed -i "s, DATABASE_PORT, '${PORT}',g" .bdjuno/config.yaml
sed -i "s, DATABASE_USER, '${USERNAME}',g" .bdjuno/config.yaml
sed -i "s, DATABASE_PASSWORD, '${PASSWORD}',g" .bdjuno/config.yaml
