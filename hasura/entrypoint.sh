#!/bin/bash

set -euo pipefail

sed -i "s, HASURA_GRAPHQL_ENDPOINT, '${HASURA_GRAPHQL_ENDPOINT}',g" config.yaml
sed -i "s, ACTIONS_BASE_URL, '${ACTIONS_BASE_URL}',g" config.yaml

hasura metadata apply --config /hasura/config.yaml
