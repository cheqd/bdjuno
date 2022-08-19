#!/bin/bash

set -euo pipefail

sed -i "s, DATABASE_NAME, '${testnet-explorer-database.DATABASE}',g" config.yaml
sed -i "s, DATABASE_HOST, '${testnet-explorer-database.HOSTNAME}',g" config.yaml
sed -i "s, DATABASE_PORT, '${testnet-explorer-database.PORT}',g" config.yaml
sed -i "s, DATABASE_USER, '${testnet-explorer-database.USERNAME}',g" config.yaml
sed -i "s, DATABASE_PASSWORD, '${testnet-explorer-database.PASSWORD}',g" config.yaml
