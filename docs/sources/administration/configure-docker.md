+++
title = "Configure Plutono Docker image"
description = "Guide for configuring the Plutono Docker image"
keywords = ["plutono", "configuration", "documentation", "docker"]
aliases = ["/docs/plutono/latest/installation/configure-docker/"]
weight = 200
+++

# Configure a Plutono Docker image

If you are running Plutono in a Docker image, then you configure Plutono using [environment variables]({{< relref "../administration/configuration.md#configure-with-environment-variables" >}}) rather than directly editing the configuration file. If you want to save your data, then you also need to designate persistent storage or bind mounts for the Plutono container.

## Save your Plutono data

If you do not designate a location for information storage, then all your Plutono data disappears as soon as you stop your container. To save your data, you need to set up persistent storage or bind mounts for your container.

### Run Plutono container with persistent storage (recommended)

```bash
# create a persistent volume for your data in /var/lib/plutono (database and plugins)
docker volume create plutono-storage

# start plutono
docker run -d -p 3000:3000 --name=plutono -v plutono-storage:/var/lib/plutono plutono/plutono
```

### Run Plutono container using bind mounts

You may want to run Plutono in Docker but use folders on your host for the database or configuration. When doing so, it becomes important to start the container with a user that is able to access and write to the folder you map into the container.

```bash
mkdir data # creates a folder for your data
ID=$(id -u) # saves your user id in the ID variable

# starts plutono with your user id and using the data folder
docker run -d --user $ID --volume "$PWD/data:/var/lib/plutono" -p 3000:3000 plutono/plutono:7.2.1
```

## Default paths

The following settings are hard-coded when launching the Plutono Docker container and can only be overridden using environment variables, not in `conf/plutono.ini`.

Setting               | Default value
----------------------|---------------------------
GF_PATHS_CONFIG       | /etc/plutono/plutono.ini
GF_PATHS_DATA         | /var/lib/plutono
GF_PATHS_HOME         | /usr/share/plutono
GF_PATHS_LOGS         | /var/log/plutono
GF_PATHS_PLUGINS      | /var/lib/plutono/plugins
GF_PATHS_PROVISIONING | /etc/plutono/provisioning

## Logging

Logs in the Docker container go to standard out by default, as is common in the Docker world. Change this by setting a different [log mode]({{< relref "../administration/configuration.md#mode" >}}).

Example:

```bash
# Run Plutono while logging to both standard out and /var/log/plutono/plutono.log
docker run -p 3000:3000 -e "GF_LOG_MODE=console file" plutono/plutono
```

## Configure Plutono with Docker Secrets

> Only available in Plutono v5.2 and later.

It's possible to supply Plutono with configuration through files. This works well with [Docker Secrets](https://docs.docker.com/engine/swarm/secrets/) as the secrets by default gets mapped into `/run/secrets/<name of secret>` of the container.

You can do this with any of the configuration options in conf/plutono.ini by setting `GF_<SectionName>_<KeyName>__FILE` to the path of the file holding the secret.

For example, you could set the admin password this way:

- Admin password secret: `/run/secrets/admin_password`
- Environment variable: `GF_SECURITY_ADMIN_PASSWORD__FILE=/run/secrets/admin_password`

## Configure AWS credentials for CloudWatch Support

```bash
docker run -d \
-p 3000:3000 \
--name=plutono \
-e "GF_AWS_PROFILES=default" \
-e "GF_AWS_default_ACCESS_KEY_ID=YOUR_ACCESS_KEY" \
-e "GF_AWS_default_SECRET_ACCESS_KEY=YOUR_SECRET_KEY" \
-e "GF_AWS_default_REGION=us-east-1" \
plutono/plutono
```

You may also specify multiple profiles to `GF_AWS_PROFILES` (e.g.
`GF_AWS_PROFILES=default another`).

Supported variables:

- `GF_AWS_${profile}_ACCESS_KEY_ID`: AWS access key ID (required).
- `GF_AWS_${profile}_SECRET_ACCESS_KEY`: AWS secret access  key (required).
- `GF_AWS_${profile}_REGION`: AWS region (optional).
