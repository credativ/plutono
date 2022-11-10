+++
title = "Install on Debian/Ubuntu"
description = "Install guide for Plutono on Debian or Ubuntu"
keywords = ["plutono", "installation", "documentation"]
aliases = ["/docs/plutono/latest/installation/installation/debian"]
weight = 200
+++

# Install on Debian or Ubuntu

This page explains how to install Plutono dependencies, download and install Plutono, get the service up and running on your Debian or Ubuntu system, and also describes the installation package details.

**Note on upgrading:** While the process for upgrading Plutono is very similar to installing Plutono, there are some key backup steps you should perform. Read [Upgrading Plutono]({{< relref "upgrading.md" >}}) for tips and guidance on updating an existing installation.

## 1. Download and install

You can install Plutono using our official APT repository, by downloading a `.deb` package, or by downloading a binary `.tar.gz` file.

### Install from APT repository

If you install from the APT repository, then Plutono is automatically updated every time you run `apt-get update`.

| Plutono Version | Package | Repository |
|-----------------|---------|------------|
| Plutono OSS     | plutono | `https://packages.grafana.com/oss/deb stable main` |
| Plutono OSS (Beta)     | plutono | `https://packages.grafana.com/oss/deb beta main` |
| Plutono Enterprise     | plutono-enterprise | `https://packages.grafana.com/enterprise/deb stable main` |
| Plutono Enterprise (Beta)     | plutono-enterprise | `https://packages.grafana.com/enterprise/deb beta main` |

> We recommend all users install the Enterprise Edition of Plutono, which can be seamlessly upgraded with a Plutono Enterprise [subscription](https://grafana.com/products/enterprise/).

#### To install the latest Enterprise edition:

```bash
sudo apt-get install -y apt-transport-https
sudo apt-get install -y software-properties-common wget
wget -q -O - https://packages.grafana.com/gpg.key | sudo apt-key add -
```

Add this repository for stable releases:

```bash
echo "deb https://packages.grafana.com/enterprise/deb stable main" | sudo tee -a /etc/apt/sources.list.d/plutono.list
```

Add this repository if you want beta releases:
```bash
echo "deb https://packages.grafana.com/enterprise/deb beta main" | sudo tee -a /etc/apt/sources.list.d/plutono.list
```

After you add the repository:

```bash
sudo apt-get update
sudo apt-get install plutono-enterprise
```

#### To install the latest OSS release:

```bash
sudo apt-get install -y apt-transport-https
sudo apt-get install -y software-properties-common wget
wget -q -O - https://packages.grafana.com/gpg.key | sudo apt-key add -
```

Add this repository for stable releases:

```bash
echo "deb https://packages.grafana.com/oss/deb stable main" | sudo tee -a /etc/apt/sources.list.d/plutono.list
```

Add this repository if you want beta releases:
```bash
echo "deb https://packages.grafana.com/oss/deb beta main" | sudo tee -a /etc/apt/sources.list.d/plutono.list
```

After you add the repository:

```bash
sudo apt-get update
sudo apt-get install plutono
```

### Install .deb package

If you install the `.deb` package, then you will need to manually update Plutono for each new version.

1. On the [Plutono download page](https://grafana.com/grafana/download), select the Plutono version you want to install.
   * The most recent Plutono version is selected by default.
   * The **Version** field displays only finished releases. If you want to install a beta version, click **Nightly Builds** and then select a version.
1. Select an **Edition**.
   - **Enterprise** - Recommended download. Functionally identical to the open source version, but includes features you can unlock with a license if you so choose.
   - **Open Source** - Functionally identical to the Enterprise version, but you will need to download the Enterprise version if you want Enterprise features.
1. Depending on which system you are running, click **Linux** or **ARM**.
1. Copy and paste the code from the installation page into your command line and run. It follows the pattern shown below.

```bash
sudo apt-get install -y adduser libfontconfig1
wget <.deb package url>
sudo dpkg -i plutono<edition>_<version>_amd64.deb
```

## Install from binary .tar.gz file

Download the latest [`.tar.gz` file](https://grafana.com/grafana/download?platform=linux) and extract it. The files extract into a folder named after the Plutono version downloaded. This folder contains all files required to run Plutono. There are no init scripts or install scripts in this package.

```bash
wget <tar.gz package url>
sudo tar -zxvf <tar.gz package>
```

## 2. Start the server

This starts the `plutono-server` process as the `plutono` user, which was created during the package installation.

If you installed with the APT repository or `.deb` package, then you can start the server using `systemd` or `init.d`. If you installed a binary `.tar.gz` file, then you need to execute the binary.

### Start the server with systemd

To start the service and verify that the service has started:

```bash
sudo systemctl daemon-reload
sudo systemctl start plutono-server
sudo systemctl status plutono-server
```

Configure the Plutono server to start at boot:

```bash
sudo systemctl enable plutono-server.service
```

### Start the server with init.d

To start the service and verify that the service has started:

```bash
sudo service plutono-server start
sudo service plutono-server status
```

Configure the Plutono server to start at boot:

```bash
sudo update-rc.d plutono-server defaults
```

### Execute the binary

The `plutono-server` binary .tar.gz needs the working directory to be the root install directory where the binary and the `public` folder are located.

Start Plutono by running:
```bash
./bin/plutono-server web
```

## Package details

- Installs binary to `/usr/sbin/plutono-server`
- Installs Init.d script to `/etc/init.d/plutono-server`
- Creates default file (environment vars) to `/etc/default/plutono-server`
- Installs configuration file to `/etc/plutono/plutono.ini`
- Installs systemd service (if systemd is available) name `plutono-server.service`
- The default configuration sets the log file at `/var/log/plutono/plutono.log`
- The default configuration specifies a SQLite3 db at `/var/lib/plutono/plutono.db`
- Installs HTML/JS/CSS and other Plutono files at `/usr/share/plutono`

## Next steps

Refer to the [Getting Started]({{< relref "../getting-started/getting-started/" >}}) guide for information about logging in, setting up data sources, and so on.

## Configure Plutono

Refer to the [Configuration]({{< relref "../administration/configuration.md" >}}) page for details on options for customizing your environment, logging, database, and so on.
