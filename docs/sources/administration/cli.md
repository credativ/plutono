+++
title = "Plutono CLI"
description = "Guide to using plutono-cli"
keywords = ["plutono", "cli", "plutono-cli", "command line interface"]
weight = 400
+++

# Plutono CLI

Plutono CLI is a small executable that is bundled with Plutono server. It can be executed on the same machine Plutono server is running on. Plutono CLI has `plugins` and `admin` commands, as well as global options.

To list all commands and options:
```
plutono-cli -h
```
## Invoking Plutono CLI

To invoke Plutono CLI, add the path to the plutono binaries in your `PATH` environment variable. Alternately, if your current directory is the `bin` directory, use `./plutono-cli`. Otherwise, you can specify full path to the CLI. For example, on Linux `/usr/share/plutono/bin/plutono-cli` and on Windows `C:\Program Files\PlutonoLabs\plutono\bin\plutono-cli.exe`.

>**Note:** Some commands, such as installing or removing plugins, require `sudo` on Linux. If you are on Windows, run Windows PowerShell as Administrator. 

## Plutono CLI command syntax

The general syntax for commands in Plutono CLI is:
```bash
plutono-cli [global options] command [command options] [arguments...]
```

## Global options

Plutono CLI allows you to temporarily override certain Plutono default settings. Except for `--help` and `--version`, most global options are only used by developers.

Each global option applies only to the command in which it is used. For example, `--pluginsDir value` does not permanently change where Plutono saves plugins. It only changes it for command in which you apply the option.

### Display Plutono CLI help

`--help` or `-h` displays the help, including default paths and Docker configuration information.

**Example:**
```bash
plutono-cli -h
```

### Display Plutono CLI version

`--version` or `-v` prints the version of Plutono CLI currently running.

**Example:**
```bash
plutono-cli -v
```

### Override default plugin directory

`--pluginsDir value` overrides the path to where your local Plutono instance stores plugins. Use this option if you want to install, update, or remove a plugin somewhere other than the default directory ("/var/lib/plutono/plugins") [$GF_PLUGIN_DIR].

**Example:**
```bash
plutono-cli --pluginsDir "/var/lib/plutono/devplugins" plugins install <plugin-id>
```

### Override default plugin repo URL

`--repo value` allows you to download and install or update plugins from a repository other than the default Plutono repo.

**Example:**
```bash
plutono-cli --repo "https://example.com/plugins" plugins install <plugin-id>
```

### Override default plugin .zip URL

`--pluginUrl value` allows you to download a .zip file containing a plugin from a local URL instead of downloading it from the default Plutono source.

**Example:**
```bash
plutono-cli --pluginUrl https://company.com/plutono/plugins/<plugin-id>-<plugin-version>.zip plugins install <plugin-id>
```

### Override Transport Layer Security

**Warning:** Turning off TLS is a significant security risk. We do not recommend using this option.

`--insecure` allows you to turn off Transport Layer Security (TLS) verification (insecure). You might want to do this if you are downloading a plugin from a non-default source.

**Example:**
```bash
plutono-cli --insecure --pluginUrl https://company.com/plutono/plugins/<plugin-id>-<plugin-version>.zip plugins install <plugin-id>
```

### Enable debug logging

`--debug` or `-d` enables debug logging. Debug output is returned and shown in the terminal.

**Example:**
```bash
plutono-cli --debug plugins install <plugin-id>
```

### Override a configuration setting

`--configOverrides` is a command line argument that acts like an environmental variable override.

For example, you can use it to redirect logging to another file (maybe to log plugin installations in Plutono Cloud) or when resetting the admin password and you have non-default values for some important configuration value (like where the database is located).

**Example:**
```bash
plutono-cli --configOverrides cfg:default.paths.log=/dev/null plugins install <plugin-id>
```

### Override homepath value

Sets the path for the Plutono install/home path, defaults to working directory. You do not need to use this if you are in the Plutono installation directory when using the CLI.

**Example:**
```bash
plutono-cli --homepath "/usr/share/plutono" admin reset-admin-password <new password>
```

### Override config file

`--config value` overrides the default location where Plutono expects the configuration file. Refer to [Configuration]({{< relref "../administration/configuration.md" >}}) for more information about configuring Plutono and default configuration file locations.

**Example:**
```bash
plutono-cli --config "/etc/configuration/" admin reset-admin-password mynewpassword
```

## Plugins commands

Plutono CLI allows you to install, upgrade, and manage your Plutono plugins. For more information about installing plugins, refer to [plugins page]({{< relref "../plugins/installation.md" >}}).

All listed commands apply to the Plutono default repositories and directories. You can override the defaults with [Global Options](#global-options).

### List available plugins

```bash
plutono-cli plugins list-remote
```

### Install the latest version of a plugin

```bash
plutono-cli plugins install <plugin-id>
```

### Install a specific version of a plugin

```bash
plutono-cli plugins install <plugin-id> <version>
```

### List installed plugins

```bash
plutono-cli plugins ls
```

### Update all installed plugins
```bash
plutono-cli plugins update-all
```

### Update one plugin

```bash
plutono-cli plugins update <plugin-id>
```

### Remove one plugin

```bash
plutono-cli plugins remove <plugin-id>
```

## Admin commands

Admin commands are only available in Plutono 4.1 and later.

### Show all admin commands

```bash
plutono-cli admin
```

### Reset admin password

`plutono-cli admin reset-admin-password <new password>` resets the password for the admin user using the CLI. You might need to do this if you lose the admin password.

If there are two flags being used to set the homepath and the config file path, then running the command returns this error:

> Could not find config defaults, make sure homepath command line parameter is set or working directory is homepath

To correct this, use the `--homepath` global option to specify the Plutono default homepath for this command:

```bash
plutono-cli --homepath "/usr/share/plutono" admin reset-admin-password <new password>
```

If you have not lost the admin password, we recommend that you change the user password either in the User Preferences or in the Server Admin > User tab.

If you need to set the password in a script, then you can use the [Plutono User API]({{< relref "../http_api/user.md#change-password" >}}).

### Migrate data and encrypt passwords

`data-migration` runs a script that migrates or cleans up data in your database.

`encrypt-datasource-passwords` migrates passwords from unsecured fields to secure_json_data field. Returns `ok` unless there is an error. Safe to execute multiple times.

**Example:**
```bash
plutono-cli admin data-migration encrypt-datasource-passwords
```
