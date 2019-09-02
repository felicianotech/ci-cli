# CI CLI [![CircleCI Build Status](https://circleci.com/gh/felicianotech/ci-cli.svg?style=shield)](https://circleci.com/gh/felicianotech/ci-cli) [![GitHub License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/felicianotech/ci-cli/master/LICENSE)

CI CLI (`ci`) is a command-line tool designed to make working with your CI provider and git a bit easier.


## Requirements

- Ubuntu 18.04, 19.04 - This project supports 64-bit installations of Ubuntu 18.04 "Bionic" and Ubuntu 19.04 "Disco".
For 32-bit support or other distros, please open a Pull Request or Issue.

macOS hasn't been tested yet but support is planned.
Windows 10 support is up in the air.
Once this CLI gets more built-out, I'll determine if Windows support is feasible without a huge amount of work.


## Installation

For now, this can be installed on Ubuntu via Go:

```
go get github.com/felicianotech/ci-cli
ci-cli --help
```

This installation method will install the cli as the `ci-cli` command.
Once proper binaries are released, the command will be `ci`.


## Usage

For now, run:

```bash
ci-cli --help
```


## Development

Instructions coming in the future.


## License

The source code for CI CLI is licensed under the MIT license.
This repo's license can be found [here](./LICENSE).
