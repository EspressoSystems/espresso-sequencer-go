# Espresso Sequencer Go SDK

This package provides tools and interfaces for working with the
[Espresso sequencer](https://github.com/EspressoSystems/espresso-sequencer) in Go. It should
(eventually) provide everything needed to integrate a rollup written in Go with the Espresso
sequencer.

## Development

- Obtain code:

        git clone git@github.com:EspressoSystems/espresso-sequencer-go
        git submodule update --init --recursive

- Make sure [nix](https://nixos.org/download.html) is installed.
- Activate the environment with `nix-shell`, or `nix develop`, or `direnv allow` if using [direnv](https://direnv.net/).

## Run the linter and unit tests

    just lint
    just test

## Generating contract bindings

    just bind-light-client
