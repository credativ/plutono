+++
title = "Install on RPM-based Linux"
description = "Plutono Installation guide for RPM-based Linux, such as Centos, Fedora, OpenSuse, and Red Hat."
keywords = ["plutono", "installation", "documentation", "centos", "fedora", "opensuse", "redhat"]
aliases = ["/docs/plutono/latest/installation/installation/rpm"]
weight = 300
+++

# Install on RPM-based Linux (CentOS, Fedora, OpenSuse, Red Hat)

This page explains how to install Plutono dependencies, download and install Plutono, get the service up and running on your RPM-based Linux system, and the installation package details.

**Note on upgrading:** While the process for upgrading Plutono is very similar to installing Plutono, there are some key backup steps you should perform. Read [Upgrading Plutono]({{< relref "upgrading.md" >}}) for tips and guidance on updating an existing installation.


## 1. Download and install

You can install Plutono from a YUM repository, manually using YUM, manually using RPM, or by downloading a binary `.tar.gz` file.

### Install from YUM repository

If you install from the YUM repository, then Plutono is automatically updated every time you run `sudo yum update`.

| Plutono Version            | Package            | Repository                                         |
|----------------------------|--------------------|----------------------------------------------------|
| Plutono OSS                | plutono            | `https://packages.grafana.com/oss/rpm`             |
| Plutono OSS (Beta)         | plutono            | `https://packages.grafana.com/oss/rpm-beta`        |
| Plutono Enterprise         | plutono-enterprise | `https://packages.grafana.com/enterprise/rpm`      |
| Plutono Enterprise (Beta)  | plutono-enterprise | `https://packages.grafana.com/enterprise/rpm-beta` |


Add a new file to your YUM repo using the method of your choice. The command below uses `nano`.

```bash
sudo nano /etc/yum.repos.d/plutono.repo
```

Choose if you want to install the Open Source or Enterprise edition of Plutono and enter the information from the edition you've chosen into `plutono.repo`. If you want to install the beta version of Plutono you need to replace the URL with a beta URL from the table above.

> We recommend all users to install the Enterprise Edition of Plutono, which can be seamlessly upgraded with a Plutono Enterprise [subscription](https://grafana.com/products/enterprise/).

For Enterprise releases:
```bash
[plutono]
name=plutono
baseurl=https://packages.grafana.com/enterprise/rpm
repo_gpgcheck=1
enabled=1
gpgcheck=1
gpgkey=https://packages.grafana.com/gpg.key
sslverify=1
sslcacert=/etc/pki/tls/certs/ca-bundle.crt
```

For OSS releases:
```bash
[plutono]
name=plutono
baseurl=https://packages.grafana.com/oss/rpm
repo_gpgcheck=1
enabled=1
gpgcheck=1
gpgkey=https://packages.grafana.com/gpg.key
sslverify=1
sslcacert=/etc/pki/tls/certs/ca-bundle.crt
```

Install Plutono with one of the following commands:

```bash
sudo yum install plutono

# or

sudo yum install plutono-enterprise
```

### Install manually with YUM

If you install manually with YUM, then you will need to manually update Plutono for each new version. To enable automatic updates for your Plutono installation please use the instructions below to install via our YUM repository.

1. On the [Plutono download page](https://grafana.com/grafana/download), select the Plutono version you want to install.
   - The most recent Plutono version is selected by default.
   - The **Version** field displays only finished releases. If you want to install a beta version, click **Nightly Builds** and then select a version.
1. Select an **Edition**.
   - **Enterprise** - Recommended download. Functionally identical to the open source version, but includes features you can unlock with a license if you so choose.
   - **Open Source** - Functionally identical to the Enterprise version, but you will need to download the Enterprise version if you want enterprise features.
1. Depending on which system you are running, click **Linux** or **ARM**.
1. Copy and paste the code from the installation page into your command line and run. It follows the pattern shown below.

```bash
wget <rpm package url>
sudo yum localinstall <local rpm package>
```
   You can also install Plutono using YUM directly:

```bash
sudo yum install <rpm package url>
```

### Install with RPM

If you install with RPM, then you will need to manually update Plutono for each new version. This method varies according to which Linux OS you are running. Read the instructions fully before you begin.

**Note:** The .rpm files are signed, you can verify the signature with this [public GPG key](https://packages.grafana.com/gpg.key).

1. On the [Plutono download page](https://grafana.com/grafana/download), select the Plutono version you want to install.
   - The most recent Plutono version is selected by default.
   - The **Version** field displays only finished releases. If you want to install a beta version, click **Nightly Builds** and then select a version.
1. Select an **Edition**.
   - **Enterprise** - Recommended download. Functionally identical to the open source version, but includes features you can unlock with a license if you so choose.
   - **Open Source** - Functionally identical to the Enterprise version, but you will need to download the Enterprise version if you want Enterprise features.
1. Depending on which system you are running, click **Linux** or **ARM**.
1. Copy and paste the .rpm package URL and the local .rpm package information from the installation page into the pattern shown below, then run the commands.

**On CentOS, Fedora, Red Hat, or RHEL:**

```bash
sudo yum install initscripts urw-fonts wget
wget <rpm package url>
sudo rpm -Uvh <local rpm package>
```

**On OpenSUSE or SUSE:**

```bash
wget <rpm package url>
sudo rpm -i --nodeps <local rpm package>
```

### Install from binary .tar.gz file

Download the latest [`.tar.gz` file](https://grafana.com/grafana/download?platform=linux) and extract it. The files are extracted into a folder named after the Plutono version that you downloaded. This folder contains all files required to run Plutono. There are no init scripts or install scripts in this package.

```bash
wget <tar.gz package url>
sudo tar -zxvf <tar.gz package>
```

## 2. Start the server

This starts the `plutono-server` process as the `plutono` user, which was created during the package installation. The systemd commands work in most cases, but some older Linux systems might require init.d. The installer should prompt you with the correct commands.

If you installed with an `.rpm` package, then you can start the server using `systemd` or `init.d`. If you installed a binary `.tar.gz` file, then you need to execute the binary.

### Start the server with systemd

To start the service and verify that the service has started:

```bash
sudo systemctl daemon-reload
sudo systemctl start plutono-server
sudo systemctl status plutono-server
```

Configure the Plutono server to start at boot:

```bash
sudo systemctl enable plutono-server
```

> **SUSE or OpenSUSE users:** You might need to start the server with the systemd method, then use the init.d method to configure Plutono to start at boot.

### Start the server with init.d

To start the service and verify that the service has started:

```bash
sudo service plutono-server start
sudo service plutono-server status
```

Configure the Plutono server to start at boot:

```bash
sudo /sbin/chkconfig --add plutono-server
```

### Execute the binary

The `plutono-server` binary needs the working directory to be the root install directory where the binary and the `public` folder are located.

Start Plutono by running:
```bash
./bin/plutono-server web
```

## Package details

- Installs binary to `/usr/sbin/plutono-server`
- Copies init.d script to `/etc/init.d/plutono-server`
- Installs default file (environment vars) to `/etc/sysconfig/plutono-server`
- Copies configuration file to `/etc/plutono/plutono.ini`
- Installs systemd service (if systemd is available) name `plutono-server.service`
- The default configuration uses a log file at `/var/log/plutono/plutono.log`
- The default configuration specifies an sqlite3 database at `/var/lib/plutono/plutono.db`

## Next steps

Refer to the [Getting Started]({{< relref "../getting-started/getting-started/" >}}) guide for information about logging in, setting up data sources, and so on.

## Configure Plutono

Refer to the [Configuration]({{< relref "../administration/configuration.md" >}}) page for details on options for customizing your environment, logging, database, and so on.
