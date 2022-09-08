#!/bin/bash

set -euo pipefail

sed -i "s, HASURA_GRAPHQL_ENDPOINT, '${HASURA_GRAPHQL_ENDPOINT}',g" config.yaml
sed -i "s, ACTION_BASE_URL, '${ACTION_BASE_URL}',g" config.yaml

hasura metadata apply --endpoint "$HASURA_GRAPHQL_ENDPOINT" --admin-secret "$HASURA_GRAPHQL_ADMIN_SECRET"
