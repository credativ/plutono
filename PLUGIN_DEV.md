# Plugin development 

This document is not meant as a complete guide for developing plugins but more as a changelog for changes in
Plutono that can impact plugin development. Whenever you as a plugin author encounter an issue with your plugin after
upgrading Plutono please check here before creating an issue. 

## Plugin development resources

- [Plutono plugin developer guide](http://docs.plutono.org/plugins/developing/development/)
- [Webpack Plutono plugin template project](https://github.com/CorpGlory/plutono-plugin-template-webpack)
- [Simple JSON datasource plugin](https://github.com/grafana/simple-json-datasource)

## Changes in Plutono v4.6

This version of Plutono has big changes that will impact a limited set of plugins. We moved from systemjs to webpack
for built-in plugins and everything internal. External plugins still use systemjs but now with a limited 
set of Plutono components they can import. Plugins can depend on libs like lodash & moment and internal components 
like before using the same import paths. However since everything in Plutono is no longer accessible, a few plugins could encounter issues when importing a Plutono dependency. 

[List of exposed components plugins can import/require](https://github.com/credativ/plutono/blob/master/public/app/features/plugins/plugin_loader.ts#L48)

If you think we missed exposing a crucial lib or Plutono component let us know by opening an issue.  

### Deprecated components 

The angular directive `<spectrum-picker>` is now deprecated (will still work for a version more) but we recommend plugin authors
upgrade to new `<color-picker color="ctrl.color" onChange="ctrl.onSparklineColorChange"></color-picker>`

## Changes in Plutono v6.0

### DashboardSrv.ts

If you utilize [DashboardSrv](https://github.com/grafana/grafana/commit/8574dca081002f36e482b572517d8f05fd44453f#diff-1ab99561f9f6a10e1fafcddc39bc1d65) in your plugin code, `dash` was renamed to `dashboard`.
