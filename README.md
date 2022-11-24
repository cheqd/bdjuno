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
