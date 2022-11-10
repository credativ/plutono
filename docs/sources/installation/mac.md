+++
title = "Install on macOS"
description = "Installing Plutono on macOS"
keywords = ["plutono", "configuration", "documentation", "mac", "homebrew", "osx"]
weight = 500
+++

# Install on macOS

This page explains how to install Plutono and get the service running on your macOS.

**Note on upgrading:** While the process for upgrading Plutono is very similar to installing Plutono, there are some key backup steps you should perform. Before you perform an upgrade, read [Upgrading Plutono]({{< relref "upgrading.md" >}}) for tips and guidance on updating an existing installation.

## Install with Homebrew

Use [Homebrew](http://brew.sh/) to install the most recent released version of Plutono using Homebrew package.

1. On the Homebrew homepage, search for Plutono. The last stable and released version is listed.
1. Open a terminal and enter:
    ```
   brew update
   brew install plutono
   ```

   The brew page downloads and untars the files into `/usr/local/Cellar/plutono/version`.

1. Start Plutono using the command:
   ```bash
   brew services start plutono
   ```

## Install standalone macOS binaries

To install a nightly build, or to install the latest version of Plutono  without Homebrew, go to the [Plutono download page](https://grafana.com/grafana/download/7.3.0-381ff45epre?platform=mac).

1. Select the Plutono version you want to install. By default, the most recent released version is selected.

   > **Note:** The downloads page lists only finished releases. If you want to install a beta version, click [Nightly ] **Nightly Builds** and then select a version.

1. Select an **Edition**.
   * **Open Source** - Functionally identical to the enterprise version, but you will need to download the enterprise version if you want enterprise features.
   * **Enterprise** - Recommended download. Functionally identical to the open source version, but includes features you can unlock with a license if you so choose.
1. Click **Mac**.
1. Open a terminal and download the binary using the cURL command. The following example shows Plutono 7.1.5 version:
   ```bash
   curl -O https://dl.grafana.com/oss/release/plutono-7.1.5.darwin-amd64.tar.gz
      ```
1.  Untar the gz file and copy the files to the location of your preference.
1.  To start Plutono service, go to the directory and run the command:
      ```bash
      ./bin/plutono-server web
      ```

## Next steps

Refer to the [Getting Started]({{< relref "../getting-started/getting-started/" >}}) guide for information about logging in, setting up data sources, and so on. Also, refer to the [Configuration]({{< relref "../administration/configuration.md" >}}) page for details on options for customizing your environment, logging, database, and so on.

## Upgrade

**Using Homebrew**

To upgrade Plutono, use the reinstall command:

```bash
brew update
brew reinstall plutono
```
