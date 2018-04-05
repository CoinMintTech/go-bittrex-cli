# agent-bittrex-firehose

[![Build Status](https://travis-ci.org/steenzout/agent-bittrex-firehose.svg?branch=master)](https://travis-ci.org/steenzout/agent-bittrex-firehose/)
[![Coverage Status](https://coveralls.io/repos/steenzout/agent-bittrex-firehose/badge.svg?branch=master&service=github)](https://coveralls.io/github/steenzout/agent-bittrex-firehose?branch=master)

This repository contains the Go agent that connects to the [Bittrex][bitrex] [API][bitrex-api].

## Build

```bash
$ go build main.go
```

## Run

```bash
$ go run main.go
```

## Usage

```bash
NAME:
   main - bittrex-cli is a command-line interface to the Bittrex API.

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     get      get data
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

[bittrex]:  https://bittrex.com "Bittrex"
[bittrex-api]:  https://bittrex.com/home/api "Bittrex API"
[license]:  https://raw.githubusercontent.com/steenzout/agent-bittrex-firehose/master/LICENSE   "License"
