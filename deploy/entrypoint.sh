#!/bin/bash

set -euo pipefail

# Change configuration file values
echo "${RPC_ADDRESS}"
echo "${DATABASE_NAME}"

sed -i "s, DATABASE_NAME, '${DATABASE_NAME}',g" /bdjuno/.bdjuno/config.yaml
sed -i "s, DATABASE_HOST, '${DATABASE_HOST}',g" /bdjuno/.bdjuno/config.yaml
sed -i "s, DATABASE_PORT, '${DATABASE_PORT}',g" /bdjuno/.bdjuno/config.yaml
sed -i "s, DATABASE_USER, '${DATABASE_USER}',g" /bdjuno/.bdjuno/config.yaml
sed -i "s, DATABASE_PASSWORD, '${DATABASE_PASSWORD}',g" /bdjuno/.bdjuno/config.yaml
