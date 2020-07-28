# Achievement Engine

The service is API layer for achievement management for users of client applications, like web wallets.

It provides:
- API over HTTP protocol for achievement management
- Encrypted key verification - Users authenticate with the Merit market using the same Public/Private key encryption that they use with the Merit Protocol.

## Setting up the app

Achievement Engine is built in the wonderful Go language, and utilizes a baseline set of libraries to keep it very lightweight.

1. To get set up, first install Golang: [download](http://golang.org/downloads)
1. Then install dependencies: `make bootstrap`
1. Install MongoDB if you don't already have it.
1. Optionally, install Air for hot code reloading: [Air Repo](https://github.com/cosmtrek/air)
1. Compile `meritd` and configure it like it is done for MWS.

## Building the app

Achievement Engine and migrations binaries can be built with make:

```
make # builds both achievement engine and migrations binary
make build-achievement-engine # builds market binary
make build-achievement-engine-migrations # builds migrations binary
```

To remove binaries and dependencies use the clean command:

```
make clean
```

## Running the app

Merit Market is easy to run.

1. Allow RPC connections and start `meritd` 
1. Copy `config.sample.yml` to `config.yml` and update configuration with your dev environment
1. Migrate the database: `go run cmd/migrations/main.go`
1. Run the main server: `go run cmd/api/main.go`

## Contributing

Please, check out our [Contribution guide](./CONTRIBUTING.md) and [Code of Conduct](./CODE_OF_CONDUCT.md).

## License

**Code released under [the MIT license](./LICENSE).**

Copyright (C) 2017 - 2020 The Merit Foundation.