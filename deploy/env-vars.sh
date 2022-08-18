#!/bin/sh

set -euo

sed -i "s, RPC_ADDRESS, '${hasura-graphql-engine.PRIVATE_URL}',g" .bdjuno/config.yaml
