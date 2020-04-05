# srv

View RSS feeds from the terminal.

![image](https://user-images.githubusercontent.com/4519234/78465683-bc1f6e00-76c6-11ea-96e7-1cdd4a5c294f.png)

## configure

srv reads configuration from `~/.config/srv/config.yaml`

If a configuration is not provided, a default configuration is generated.

- `feeds` is a list of RSS/Atom feeds to be loaded in srv.
- `externalViewer` defines an application to override the default web browser (optional).

An example config can be copied:

```shell
cp ./config-example.yaml ~/.config/srv/config.yaml
```

## control

Key mappings are statically defined for the time being.

- `TAB` switches between Feeds and Items.
- `UP/DOWN` navigates feeds and items`
- `ENTER` either selects a feed or opens a feed item in an external application.
- `F5` refresh list of feeds

## build

```shell
make build
```

## test

```shell
make test
```
