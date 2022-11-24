# BDJuno

BDJuno (shorthand for BigDipper Juno) is the [Juno](https://github.com/forbole/juno) implementation
for [BigDipper](https://github.com/forbole/big-dipper).

All the chains' data that are queried from the RPC and gRPC endpoints are stored inside
a [PostgreSQL](https://www.postgresql.org/) database on top of which [GraphQL](https://graphql.org/) APIs can then be
created using [Hasura](https://hasura.io/).

## Features specific to cheqd

1. Indexing for [cheqd network] DIDs and Resources
2. Changes to workflows/pipelines
3. Optimised Dockerfile

### Configuration files

1. BDJuno configuration is in `deploy` folder
   1. `genesis.json` should be current from the published one for this specific chain (testnet/mainnet) from [cheqd-node](https://github.com/cheqd/cheqd-node).
   2. Edit `deploy/config.yaml` if necessary.
   3. The variables used in config file are populated from DigitalOcean secrets by the entrypoint script `deploy/entrypoint.sh`
2. Hasura configuration is in `hasura` folder
   1. Edit `hasura/config.yaml` if necessary.
   2. The variables used in config file are populated from DigitalOcean secrets by the entrypoint script `deploy/entrypoint.sh`

## Developer guide

This section is reproduced as-is from upstream project.

### Usage

To know how to setup and run BDJuno, please refer to
the [docs website](https://docs.bigdipper.live/cosmos-based/parser/overview/).

## Testing

If you want to test the code, you can do so by running

```shell
make test-unit
```

**Note**: Requires [Docker](https://docker.com).

This will:

1. Create a Docker container running a PostgreSQL database.
2. Run all the tests using that database as support.
