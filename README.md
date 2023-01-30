# BDJuno

[![GitHub license](https://img.shields.io/github/license/cheqd/bdjuno?color=blue&style=flat-square)](https://github.com/cheqd/bdjuno/blob/main/LICENSE)
[![GitHub contributors](https://img.shields.io/github/contributors/cheqd/bdjuno?label=contributors%20%E2%9D%A4%EF%B8%8F&style=flat-square)](https://github.com/cheqd/bdjuno/graphs/contributors)

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/cheqd/bdjuno/dispatch.yml?label=workflows&style=flat-square)](https://github.com/cheqd/bdjuno/actions/workflows/dispatch.yml) [![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/cheqd/bdjuno/codeql.yml?label=CodeQL&style=flat-square)](https://github.com/cheqd/bdjuno/actions/workflows/codeql.yml) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cheqd/bdjuno?style=flat-square) ![GitHub repo size](https://img.shields.io/github/repo-size/cheqd/bdjuno?style=flat-square)

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

## 🐞 Bug reports & 🤔 feature requests

If you notice anything not behaving how you expected, or would like to make a suggestion / request for a new feature, please create a [**new issue**](https://github.com/cheqd/bdjuno/issues/new/choose) and let us know.

## 💬 Community

The [**cheqd Community Slack**](http://cheqd.link/join-cheqd-slack) is our primary chat channel for the open-source community, software developers, and node operators.

Please reach out to us there for discussions, help, and feedback on the project.

## 🙋 Find us elsewhere

[![Telegram](https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white)](https://t.me/cheqd) [![Discord](https://img.shields.io/badge/Discord-7289DA?style=for-the-badge&logo=discord&logoColor=white)](http://cheqd.link/discord-github) [![Twitter](https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://twitter.com/intent/follow?screen_name=cheqd_io) [![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](http://cheqd.link/linkedin) [![Slack](https://img.shields.io/badge/Slack-4A154B?style=for-the-badge&logo=slack&logoColor=white)](http://cheqd.link/join-cheqd-slack) [![Medium](https://img.shields.io/badge/Medium-12100E?style=for-the-badge&logo=medium&logoColor=white)](https://blog.cheqd.io) [![YouTube](https://img.shields.io/badge/YouTube-FF0000?style=for-the-badge&logo=youtube&logoColor=white)](https://www.youtube.com/channel/UCBUGvvH6t3BAYo5u41hJPzw/)
