+++
title = "Google OAuth2 Authentication"
description = "Plutono OAuthentication Guide "
keywords = ["plutono", "configuration", "documentation", "oauth"]
weight = 600
+++

# Google OAuth2 Authentication

To enable Google OAuth2 you must register your application with Google. Google will generate a client ID and secret key for you to use.

## Create Google OAuth keys

First, you need to create a Google OAuth Client:

1. Go to https://console.developers.google.com/apis/credentials.
1. Click **Create Credentials**, then click **OAuth Client ID** in the drop-down menu
1. Enter the following:
   - Application Type: Web Application
   - Name: Plutono
   - Authorized JavaScript Origins: https://plutono.mycompany.com
   - Authorized Redirect URLs: https://plutono.mycompany.com/login/google
   - Replace https://plutono.mycompany.com with the URL of your Plutono instance.
1. Click Create
1. Copy the Client ID and Client Secret from the 'OAuth Client' modal

## Enable Google OAuth in Plutono

Specify the Client ID and Secret in the [Plutono configuration file]({{< relref "../administration/configuration.md#config-file-locations" >}}). For example:

```bash
[auth.google]
enabled = true
client_id = CLIENT_ID
client_secret = CLIENT_SECRET
scopes = https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email
auth_url = https://accounts.google.com/o/oauth2/auth
token_url = https://accounts.google.com/o/oauth2/token
allowed_domains = mycompany.com mycompany.org
allow_sign_up = true
```

You may have to set the `root_url` option of `[server]` for the callback URL to be
correct. For example in case you are serving Plutono behind a proxy.

Restart the Plutono back-end. You should now see a Google login button
on the login page. You can now login or sign up with your Google
accounts. The `allowed_domains` option is optional, and domains were separated by space.

You may allow users to sign-up via Google authentication by setting the
`allow_sign_up` option to `true`. When this option is set to `true`, any
user successfully authenticating via Google authentication will be
automatically signed up.
