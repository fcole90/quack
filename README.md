# Quack

Simple service in Go to keep you IP updated on Duck DNS

## Instruction

1. Register your domain on https://www.duckdns.org/
2. Install the app with one of the following methods
3. Configure the app providing:
 - token: you can find your token on [Duck DNS](https://www.duckdns.org/) after you log in
 - domain: one of the domains you registered on [Duck DNS](https://www.duckdns.org/)
 - time interval: seconds between each update (defaults to 300)


## Snap

Quack is provided as a daemon snap.

### Installation

```bash
  snap install quack
  sudo quack.config  # You'll be prompted for the configuration details
```
Then you can check its logs with

```bash
  snap logs quack
```

## Brew

```bash
  brew install fcole90/quack/quack
```

You can change where to store your configuration by setting the `QUACK_CONFIG_DIR` environment variable.


## Manual Build

You need:
 - `git`
 - `make`
 - `go`

```bash
  git clone git@github.com:fcole90/quack.git
  cd quack
  make build
```

You can then configure it with
```
  ./bin/quack config
```

You can change where to store your configuration by setting the `QUACK_CONFIG_DIR` environment variable.

## Go Package

```
go get github.com/fcole90/quack
```