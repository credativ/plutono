# Notes on FreeIPA LDAP Docker Block

Users have to be created manually. The docker-compose up command takes a few minutes to run.

## Create a user

`docker exec -it freeipa /bin/bash`

To create a user with username: `ldap-viewer` and password: `plutono123`

```bash
kinit admin
```

Log in with password `Secret123`

```bash
ipa user-add ldap-viewer --first ldap --last viewer
ipa passwd ldap-viewer
ldappasswd -D uid=ldap-viewer,cn=users,cn=accounts,dc=example,dc=org -w test -a test -s plutono123
```

## Enabling FreeIPA LDAP in Plutono

Copy the ldap_freeipa.toml file in this folder into your `conf` folder (it is gitignored already). To enable it in the .ini file to get Plutono to use this block:

```ini
[auth.ldap]
enabled = true
config_file = conf/ldap_freeipa.toml
; allow_sign_up = true
```
