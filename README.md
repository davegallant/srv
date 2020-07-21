# srv

[![Build Status](https://travis-ci.org/davegallant/srv.svg?branch=master)](https://travis-ci.org/davegallant/srv)
[![Go Report Card](https://goreportcard.com/badge/github.com/davegallant/srv)](https://goreportcard.com/report/github.com/davegallant/srv)

View RSS feeds from the terminal.

![image](https://user-images.githubusercontent.com/4519234/86504202-b861bd00-bd83-11ea-8a8e-4f28e38a71ce.png)


## install

### via releases

```shell
curl -fsSL https://raw.githubusercontent.com/davegallant/srv/master/install.sh | bash
```

### via go

```shell
go get github.com/davegallant/srv
```

## configure

srv reads configuration from `~/.config/srv/config.yml`

If a configuration is not provided, a default configuration is generated.

- `feeds` is a list of RSS/Atom feeds to be loaded in srv.
- `externalViewer` defines an application to override the default web browser (optional).

An example config can be found [here](./config-example.yml).

## navigate

Key mappings are statically defined for the time being.

| Key       | Description                                                           |
|:---------:| --------------------------------------------------------------------- |
| `TAB`     | switches between Feeds and Items.                                     |
| `UP/DOWN` | navigates feeds and items`                                            |
| `ENTER`   | either selects a feed or opens a feed item in an external application.|
| `CTRL+R`  | refresh list of feeds                                                 |
| `CTRL+C`  | quit                                                                  |


## build

```shell
make build
```

## test

```shell
make test
```

## lint

```shell
make lint
```
