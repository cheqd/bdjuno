#!/bin/bash

set -euo pipefail

sed -i "s, RPC_ADDRESS, '${RPC_ADDRESS}',g" config.yaml
sed -i "s, GRPC_ADDRESS, '${GRPC_ADDRESS}',g" config.yaml
sed -i "s, LOG_LEVEL, '${LOG_LEVEL}',g" config.yaml
sed -i "s, DATABASE_NAME, '${testnet-explorer-database.DATABASE}',g" config.yaml
sed -i "s, DATABASE_HOST, '${testnet-explorer-database.HOSTNAME}',g" config.yaml
sed -i "s, DATABASE_PORT, '${testnet-explorer-database.PORT}',g" config.yaml
sed -i "s, DATABASE_USER, '${testnet-explorer-database.USERNAME}',g" config.yaml
sed -i "s, DATABASE_PASSWORD, '${testnet-explorer-database.PASSWORD}',g" config.yaml
