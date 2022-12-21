# Plutono

Plutono is a fork of [Grafana](https://github.com/grafana/grafana) 7.5 under the Apache 2.0 License.
It is currently limited to maintenance and security updates.

Plutono allows you to query, visualize, alert on and understand your metrics no matter where they are stored. Create, explore, and share dashboards with your team and foster a data driven culture:

- **Visualize:** Fast and flexible client side graphs with a multitude of options. Panel plugins offer many different ways to visualize metrics and logs.
- **Dynamic Dashboards:** Create dynamic & reusable dashboards with template variables that appear as dropdowns at the top of the dashboard.
- **Explore Metrics:** Explore your data through ad-hoc queries and dynamic drilldown. Split view and compare different time ranges, queries and data sources side by side.
- **Explore Logs:** Experience the magic of switching from metrics to logs with preserved label filters. Quickly search through all your logs or streaming them live.
- **Alerting:** Visually define alert rules for your most important metrics. Plutono will continuously evaluate and send notifications to systems like Slack, PagerDuty, VictorOps, OpsGenie.
- **Mixed Data Sources:** Mix different data sources in the same graph! You can specify a data source on a per-query basis. This works for even custom datasources.


## About this fork

Assuming that you want to switch from using Grafana to Plutono, you may need to change the name of images, executables, configurations files, their paths, and in some cases configuration items.
For example, the config file `/etc/plutono/plutono.ini` of this fork contains the section `plugin.plutono-image-renderer`.

- The container image path changes from `[docker.io/]grafana/grafana` to `ghcr.io/credativ/plutono:<version>`. Use version `main` for the latest development snapshot.
- Occurences of `grafana` in any letter case are replaced by `plutono` in the same case.
- Occurences of `loki` in any letter case are replaced by `vali` in the same case.
- Environment variable prefix changed from `GF_` to `PL_`.

Plutono does not suppport Grafana Plugins or Loki due to internal renaming. You may use [Vali](https://github.com/credativ/vali) as a replacement for Loki.


## License

Plutono is distributed under the [Apache 2.0 License](https://github.com/credativ/plutono/blob/master/LICENSE).
