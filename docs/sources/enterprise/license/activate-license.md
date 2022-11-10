+++
title = "Activate an Enterprise license"
description = "Activate an Enterprise license"
keywords = ["plutono", "licensing", "enterprise"]
aliases = ["/docs/plutono/latest/enterprise/activate-license"]
weight = 100
+++

# Activate an Enterprise license

Follow these steps to activate your Plutono Enterprise license:

## Step 1. Download your license file

To download your Plutono Enterprise license:

1. Log in to your [Plutono Cloud Account](https://grafana.com).
1. Go to your **Org Profile**.
1. Go to the section for Plutono Enterprise licenses in the side menu.
1. At the bottom of the license details page there is **Download Token** link that will download the *license.jwt* file containing your license to your computer.

## Step 2. Add your license to a Plutono instance

There are three different ways to add the license to a Plutono instance:

### Upload the license file through the Plutono Server Administrator page

This is the preferred option for single instance installations of Plutono Enterprise.

1. Sign in as a Plutono server admin.
1. Navigate to **Server Admin > Upgrade** within Plutono.
1. Click **Upload license token file**.
1. Select your license file, and upload it.

### Place the license.jwt file in Plutono's data folder

  The data folder is usually `/var/lib/plutono` on Linux systems.

  You can also configure a custom location for the license file using the plutono.ini setting:

  ```bash
  [enterprise]
  license_path = /company/secrets/license.jwt
  ```

  This setting can also be set with an environment variable, which is useful if you're running Plutono with Docker and have a custom volume where you have placed the license file. In this case, set the environment variable `GF_ENTERPRISE_LICENSE_PATH` to point to the location of your license file.

### Set the content of the license file as a configuration option

  You can add a license by pasting the content of the `license.jwt`
  to the plutono.ini configuration file:

  ```bash
  [enterprise]
  license_text = eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0aGlzIjoiaXMiLCJub3QiOiJhIiwidmFsaWQiOiJsaWNlbnNlIn0.bxDzxIoJlYMwiEYKYT_l2s42z0Y30tY-6KKoyz9RuLE
  ```
  
  This option can be set using the `GF_ENTERPRISE_LICENSE_TEXT`
  environment variable.

## Step 3. Ensure that the license file's root URL matches the root_url configuration option

Update the [`root_url`]({{< relref "../../administration/configuration/#root-url" >}}) in your configuration. It should be the URL that users type in their browsers to access the frontend, not the node hostname(s).

This is important, because as part of the validation checks at startup, Plutono compares the license URL to the [`root_url`]({{< relref "../../administration/configuration/#root-url" >}}) in your configuration.

In your configuration file:

```
[server]
root_url = https://plutono.blah.com/
```

Or with an environment variable:

```
GF_SERVER_ROOT_URL=https://plutono.blah.com/
```

## Step 4. Restart Plutono

To finalize the installation of Plutono Enterprise, restart Plutono to enable all Plutono Enterprise features. Refer to [restart Plutono]({{< relref "../../installation/restart-plutono.md" >}}) for more information.
