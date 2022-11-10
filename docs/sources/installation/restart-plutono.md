+++
title = "Restart Plutono"
description = "Instructions for restarting Plutono"
keywords = ["plutono", "restart", "documentation"]
weight = 750
+++

# Restart Plutono

Users often need to restart Plutono after they have made configuration changes. This topic provides detailed instructions on how to restart Plutono supported operating systems.

- [Windows](#windows)
- [MacOS](#macos)
- [Linux](#linux)
- [Docker](#docker)

## Windows

To restart Plutono:

1. Open the Services app.
1. Right-click on the **Plutono** service.
1. In the context menu, click **Restart**.

## macOS

Restart methods differ depending on whether you installed Plutono using Homebrew or as standalone macOS binaries.

### Restart Plutono using Homebrew

Use the [Homebrew](http://brew.sh/) restart command:

```bash
brew services restart plutono
```
### Restart standalone macOS binaries

To restart Plutono:

1. Open a terminal and go to the directory where you copied the install setup files.
1. Run the command:

```bash
./bin/plutono-server web
```
## Linux

Restart methods differ depending on whether your Linux system uses `systemd` or `init.d`.

### Restart the server with systemd

To restart the service and verify that the service has started, run the following commands:

```bash
sudo systemctl restart plutono-server
sudo systemctl status plutono-server
```

Alternately, you can configure the Plutono server to restart at boot:

```bash
sudo systemctl enable plutono-server.service
```

> **Note:** SUSE or OpenSUSE users may need to start the server with the systemd method, then use the init.d method to configure Plutono to start at boot.

### Restart the server with init.d

To restart the service, run the following command:

```bash
sudo service plutono-server restart
```

or

```bash
sudo /etc/init.d/plutono-server restart
```

Verify the status:

```bash
sudo service plutono-server status
```

or

```bash
sudo /etc/init.d/plutono-server status
```

Alternately, you can configure the Plutono server to restart at boot:

```bash
sudo update-rc.d plutono-server defaults
```
## Docker

To restart the Plutono service, use the `docker restart` command.

`docker restart plutono`

Alternately, you can use the `docker compose restart` command to restart Plutono. For more information, refer to [docker compose documentation](https://docs.docker.com/compose/).

### Docker compose example

Configure your `docker-compose.yml` file. For example:

```bash
plutono:
  image: plutono/plutono:latest
  ports:
    - "3000:3000"
  environment:
    - TERM=linux
    - GF_INSTALL_PLUGINS=plutono-clock-panel,plutono-piechart-panel,plutono-polystat-panel
```

Start the Plutono server:

`docker-compose up`

This starts the Plutono server along with the three plugins specified in the YAML file.

To restart the running container, use this command:

`docker-compose restart plutono`
