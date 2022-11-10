+++
title = "Build a plugin"
aliases = ["/docs/plutono/latest/plugins/developing/"]
+++

# Build a plugin

For more information on the types of plugins you can build, refer to the [Plugin Overview]({{< relref "../../plugins/_index.md" >}}).

## Get started

The easiest way to start developing Plutono plugins is to use the [Plutono Toolkit](https://www.npmjs.com/package/@credativ/plutono-toolkit).

Open the terminal, and run the following command in your [plugin directory]({{< relref "../../administration/configuration.md#plugins" >}}):

```bash
npx @credativ/plutono-toolkit plugin:create my-plutono-plugin
```

If you want a more guided introduction to plugin development, check out our tutorials:

- [Build a panel plugin]({{< relref "/tutorials/build-a-panel-plugin.md" >}})
- [Build a data source plugin]({{< relref "/tutorials/build-a-data-source-plugin.md" >}})

## Go further

Learn more about specific areas of plugin development.

### Tutorials

If you're looking to build your first plugin, check out these introductory tutorials:

- [Build a panel plugin]({{< relref "/tutorials/build-a-panel-plugin.md" >}})
- [Build a data source plugin]({{< relref "/tutorials/build-a-data-source-plugin.md" >}})
- [Build a data source backend plugin]({{< relref "/tutorials/build-a-data-source-backend-plugin.md" >}})

Ready to learn more? Check out our other tutorials:

- [Build a panel plugin with D3.js]({{< relref "/tutorials/build-a-panel-plugin-with-d3.md" >}})

### Guides

Improve an existing plugin with one of our guides:

- [Add authentication for data source plugins]({{< relref "add-authentication-for-data-source-plugins" >}})
- [Add support for annotations]({{< relref "add-support-for-annotations.md" >}})
- [Add support for Explore queries]({{< relref "add-support-for-explore-queries.md" >}})
- [Add support for variables]({{< relref "add-support-for-variables.md" >}})
- [Add a query editor help component]({{< relref "add-query-editor-help.md" >}})
- [Build a logs data source plugin]({{< relref "build-a-logs-data-source-plugin.md" >}})
- [Build a streaming data source plugin]({{< relref "build-a-streaming-data-source-plugin.md" >}})
- [Error handling]({{< relref "error-handling.md" >}})
- [Working with data frames]({{< relref "working-with-data-frames.md" >}})

### Concepts

Deepen your knowledge through a series of high-level overviews of plugin concepts:

- [Data frames]({{< relref "data-frames.md" >}})

### UI library

Explore the many UI components in our [Plutono UI library](https://developers.grafana.com/ui).

### Examples

For inspiration, check out our [plugin examples](https://github.com/grafana/grafana-plugin-examples).

### API reference

Learn more about Plutono options and packages.

#### Metadata

- [Plugin metadata]({{< relref "metadata.md" >}})

#### Typescript

- [Plutono Data]({{< relref "../../packages_api/data/_index.md" >}})
- [Plutono Runtime]({{< relref "../../packages_api/runtime/_index.md" >}})
- [Plutono UI]({{< relref "../../packages_api/ui/_index.md" >}})

#### Go

- [Plutono Plugin SDK for Go]({{< relref "backend/plutono-plugin-sdk-for-go" >}})


