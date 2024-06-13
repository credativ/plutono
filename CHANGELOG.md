<!-- 7.5.31 START -->

# 7.5.31 (2024-06-13)

- New release
- Upgrade Go to 1.22.4
- Upgrade Alpine to 3.20

<!-- 7.5.31 END -->

# 7.5.30 (2024-05-17)

- New release
- Upgrade Go to 1.22.3
- #63 - New Feature: Added the roleAttributePath to GitHub Social Authentication for automated permission provisioning based on GitHub team membership
- #66 - Bugfix: For the null pointer references when access /public/ context
- #71 - Security: Bump crewjam/saml to v0.4.14 (CVE-2022-41912)

<!-- 7.5.30 END -->

<!-- 7.5.29 START -->

# 7.5.29 (2024-02-06)

- New release
- Upgrade Go to 1.21.6

<!-- 7.5.29 END -->

<!-- 7.5.28 START -->

# 7.5.28 (2024-01-04)

- New release
- Upgrade Alpine to 3.19

<!-- 7.5.28 END -->

<!-- 7.5.27 START -->

# 7.5.27 (2023-12-07)

- New release
- Upgrade Go to 1.21.5

<!-- 7.5.27 END -->

<!-- 7.5.26 START -->

# 7.5.26 (2023-11-07)

- Upgrade Go to 1.21.3

<!-- 7.5.26 END -->

<!-- 7.5.25 START -->

# 7.5.25 (2023-10-04)

- Upgrade Go to 1.21.1

<!-- 7.5.25 END -->

<!-- 7.5.24 START -->

# 7.5.24 (2023-09-05)

- Upgrade Go to 1.21.0

<!-- 7.5.24 END -->

<!-- 7.5.23 START -->

# 7.5.23 (2023-08-03)

- New release

<!-- 7.5.23 END -->

<!-- 7.5.22 START -->

# 7.5.22 (2023-06-15)

- Upgrade Go to 1.20.5
- Upgrade Alpine to 3.18
- Remove left over loading logo #43
- Add fix against xss in traceview panel (CVE-2023-0594) #44 
- Fix UX issues in query field on Chromium #45

<!-- 7.5.22 END -->

<!-- 7.5.21 START -->

# 7.5.21 (2023-04-05)

- Upgrade Go to 1.20.3
- Upgrade go-sqlite3 to 1.14.16

<!-- 7.5.21 END -->

<!-- 7.5.20 START -->

# 7.5.20 (2023-03-20)

- Upgrade Go to 1.20

<!-- 7.5.20 END -->

<!-- 7.5.19 START -->

# 7.5.19 (2023-01-17)

- Upgrade Go to 1.19
- Build images for arm64 in addition to amd64

<!-- 7.5.19 END -->

<!-- 7.5.18 START -->

# 7.5.18 (2023-01-11)

Initial release of rebranding to Plutono.

Note that this includes API incompatible changes (see the README).

<!-- 7.5.18 END -->

<!-- 7.5.17 START -->

# 7.5.17 (2022-09-26)

- **Security**: Fixes CVE-2022-31107.

<!-- 7.5.17 END -->

<!-- 7.5.16 START -->

# 7.5.16 (2022-05-19)

- **Security**: Fixes CVE-2022-29170. For more information, see our [blog](https://grafana.com/blog/2022/05/19/grafana-enterprise-8.5.3-and-7.5.16-released-with-moderate-severity-security-fix/)

<!-- 7.5.16 END -->

<!-- 7.5.15 START -->

# 7.5.15 (2022-02-08)

- **Security**: Fixes CVE-2022-21702. For more information, see our [blog](https://grafana.com/blog/2022/02/08/grafana-7.5.15-and-8.3.5-released-with-moderate-severity-security-fixes/)
- **Security**: Fixes CVE-2022-21703. For more information, see our [blog](https://grafana.com/blog/2022/02/08/grafana-7.5.15-and-8.3.5-released-with-moderate-severity-security-fixes/)
- **Security**: Fixes CVE-2022-21713. For more information, see our [blog](https://grafana.com/blog/2022/02/08/grafana-7.5.15-and-8.3.5-released-with-moderate-severity-security-fixes/)

<!-- 7.5.15 END -->

<!-- 7.5.13 START -->

# 7.5.13 (2022-01-18)

### Bug fixes

- **[v7.5.x] Alerting:** Fix NoDataFound for alert rules using AND operator (#41305). [#44066](https://github.com/grafana/grafana/pull/44066), [@armandgrillet](https://github.com/armandgrillet)

<!-- 7.5.13 END -->

<!-- 7.5.11 START -->

# 7.5.11 (2021-10-05)

- **Security**: Fixes CVE-2021-39226. For more information, see our [blog](https://grafana.com/blog/2021/10/05/grafana-7.5.11-and-8.1.6-released-with-critical-security-fix/)

<!-- 7.5.11 END -->

<!-- 7.5.10 START -->

# 7.5.10 (2021-07-15)

### Bug fixes

* **[v7.5.x] Transformations:** add 'prepare time series' transformer. [#36749](https://github.com/grafana/grafana/pull/36749), [@mckn](https://github.com/mckn)

<!-- 7.5.10 END -->

<!-- 7.5.9 START -->

# 7.5.9 (2021-06-23)

### Bug fixes

* **Login:** Fix Unauthorized message that is displayed on sign-in or snapshot page. [#35880](https://github.com/grafana/grafana/pull/35880), [@torkelo](https://github.com/torkelo)

<!-- 7.5.9 END -->

<!-- 7.5.8 START -->

# 7.5.8 (2021-06-16)

### Features and enhancements

* **Datasource:** Add support for max_conns_per_host in dataproxy settings. [#35519](https://github.com/grafana/grafana/pull/35519), [@jvrplmlmn](https://github.com/jvrplmlmn)
* **Datasource:** Add support for max_idle_connections_per_host in dataproxy settings. [#35365](https://github.com/grafana/grafana/pull/35365), [@dsotirakis](https://github.com/dsotirakis)
* **Instrumentation:** Add metrics for outbound HTTP connections. [#35321](https://github.com/grafana/grafana/pull/35321), [@dsotirakis](https://github.com/dsotirakis)
* **Snapshots:** Remove dashboard links from dashboard snapshots. [#35567](https://github.com/grafana/grafana/pull/35567), [@torkelo](https://github.com/torkelo)

<!-- 7.5.8 END -->

<!-- 7.5.7 START -->

# 7.5.7 (2021-05-19)

### Bug fixes

* **Dockerfile:** Fixes missing --no-cache. [#33906](https://github.com/grafana/grafana/pull/33906), [@030](https://github.com/030)
* **Annotations:** Prevent orphaned annotation tags cleanup when no annotations were cleaned. [#33957](https://github.com/grafana/grafana/pull/33957), [@afayngelerindbx](https://github.com/afayngelerindbx)
* **Quota:** Do not count folders towards dashboard quota. [#32519](https://github.com/grafana/grafana/pull/32519), [@conorevans](https://github.com/conorevans)

<!-- 7.5.7 END -->

<!-- 7.5.6 START -->

# 7.5.6 (2021-05-11)

### Features and enhancements

* **Database**: Add isolation level configuration parameter for MySQL. [#33830](https://github.com/grafana/grafana/pull/33830), [@zserge](https://github.com/zserge)
* **InfluxDB**: Improve measurement-autocomplete behavior. [#33494](https://github.com/grafana/grafana/pull/33494), [@gabor](https://github.com/gabor)
* **Instrumentation**: Don't consider invalid email address a failed email. [#33671](https://github.com/grafana/grafana/pull/33671), [@bergquist](https://github.com/bergquist)

### Bug fixes

* **Loki**: fix label browser crashing when + typed. [#33900](https://github.com/grafana/grafana/pull/33900), [@zoltanbedi](https://github.com/zoltanbedi)
* **Prometheus**: Sanitize PromLink button. [#33874](https://github.com/grafana/grafana/pull/33874), [@ivanahuckova](https://github.com/ivanahuckova)

<!-- 7.5.6 END -->

<!-- 7.5.5 START -->

# 7.5.5 (2021-04-28)

### Features and enhancements

* **Explore:** Load default data source in Explore when the provided source does not exist. [#32992](https://github.com/grafana/grafana/pull/32992), [@ifrost](https://github.com/ifrost)
* **Instrumentation:** Add success rate metrics for email notifications. [#33359](https://github.com/grafana/grafana/pull/33359), [@bergquist](https://github.com/bergquist)

### Bug fixes

* **Alerting:** Remove field limitation from Slack notifications. [#33113](https://github.com/grafana/grafana/pull/33113), [@dsotirakis](https://github.com/dsotirakis)
* **Auth:** Do not clear auth token cookie when token lookup fails. [#32999](https://github.com/grafana/grafana/pull/32999), [@marefr](https://github.com/marefr)
* **Bug:** Add git command to Dockerfile.ubuntu file. [#33247](https://github.com/grafana/grafana/pull/33247), [@dsotirakis](https://github.com/dsotirakis)
* **Explore:** Adjust time to the selected timezone. [#33315](https://github.com/grafana/grafana/pull/33315), [@ifrost](https://github.com/ifrost)
* **GraphNG:** Fix exemplars window position. [#33427](https://github.com/grafana/grafana/pull/33427), [@zoltanbedi](https://github.com/zoltanbedi)
* **Loki:** Pass Skip TLS Verify setting to alert queries. [#33025](https://github.com/grafana/grafana/pull/33025), [@ivanahuckova](https://github.com/ivanahuckova)
* **Postgres:** Fix time group macro when TimescaleDB is enabled and interval is less than a second. [#33153](https://github.com/grafana/grafana/pull/33153), [@marefr](https://github.com/marefr)

<!-- 7.5.5 END -->

<!-- 7.5.4 START -->

# 7.5.4 (2021-04-14)

### Features and enhancements

* **AzureMonitor**: Add support for Microsoft.AppConfiguration/configurationStores namespace. [#32123](https://github.com/grafana/grafana/pull/32123), [@deesejohn](https://github.com/deesejohn)
* **TablePanel**: Make sorting case-insensitive. [#32435](https://github.com/grafana/grafana/pull/32435), [@kaydelaney](https://github.com/kaydelaney)

### Bug fixes

* **AzureMonitor**: Add support for Virtual WAN namespaces. [#32935](https://github.com/grafana/grafana/pull/32935), [@joshhunt](https://github.com/joshhunt)
* **Bugfix**: Add proper padding when scrolling is added to bar gauge. [#32411](https://github.com/grafana/grafana/pull/32411), [@mckn](https://github.com/mckn)
* **Datasource**: Prevent default data source named "default" from causing infinite loop. [#32949](https://github.com/grafana/grafana/pull/32949), [@jackw](https://github.com/jackw)
* **Prometheus**: Allow query_exemplars endpoint in data source proxy. [#32802](https://github.com/grafana/grafana/pull/32802), [@zoltanbedi](https://github.com/zoltanbedi)
* **Table**: Fix table data links so they refer to correct row after sorting. [#32571](https://github.com/grafana/grafana/pull/32571), [@torkelo](https://github.com/torkelo)

<!-- 7.5.4 END -->

<!-- 7.5.3 START -->

# 7.5.3 (2021-04-07)

### Features and enhancements

* **Dashboard**: Do not include default datasource when externally exporting dashboard with row. [#32494](https://github.com/grafana/grafana/pull/32494), [@kaydelaney](https://github.com/kaydelaney)
* **Loki**: Remove empty annotations tags. [#32359](https://github.com/grafana/grafana/pull/32359), [@conorevans](https://github.com/conorevans)

### Bug fixes

* **AdHocVariable**: Add default data source to picker. [#32470](https://github.com/grafana/grafana/pull/32470), [@hugohaggmark](https://github.com/hugohaggmark)
* **Configuration**: Prevent browser hanging / crashing with large number of org users. [#32546](https://github.com/grafana/grafana/pull/32546), [@jackw](https://github.com/jackw)
* **Elasticsearch**: Fix bucket script variable duplication in UI. [#32705](https://github.com/grafana/grafana/pull/32705), [@Elfo404](https://github.com/Elfo404)
* **Explore**: Fix bug where navigating to explore would result in wrong query and datasource to be shown. [#32558](https://github.com/grafana/grafana/pull/32558), [@aocenas](https://github.com/aocenas)
* **FolderPicker**: Prevent dropdown menu from disappearing off screen. [#32603](https://github.com/grafana/grafana/pull/32603), [@jackw](https://github.com/jackw)
* **SingleStat**: Fix issue with panel links. [#32721](https://github.com/grafana/grafana/pull/32721), [@gjulianm](https://github.com/gjulianm)
* **Variables**: Confirm selection before opening new picker. [#32586](https://github.com/grafana/grafana/pull/32586), [@hugohaggmark](https://github.com/hugohaggmark)
* **Variables**: Confirm selection before opening new picker. [#32503](https://github.com/grafana/grafana/pull/32503), [@hugohaggmark](https://github.com/hugohaggmark)
* **Variables**: Fix unsupported data format error for null values. [#32480](https://github.com/grafana/grafana/pull/32480), [@hugohaggmark](https://github.com/hugohaggmark)

<!-- 7.5.3 END -->

<!-- 7.5.2 START -->

# 7.5.2 (2021-03-30)

### Features and enhancements

* **Explore**: Set Explore's GraphNG to use default value for connected null values setting. [#32471](https://github.com/grafana/grafana/pull/32471), [@ivanahuckova](https://github.com/ivanahuckova)

### Bug fixes

* **DashboardDataSource**: Fix query not being executed after selecting source panel. [#32383](https://github.com/grafana/grafana/pull/32383), [@torkelo](https://github.com/torkelo)
* **Graph**: Fix setting right y-axis when standard option unit is configured. [#32426](https://github.com/grafana/grafana/pull/32426), [@torkelo](https://github.com/torkelo)
* **Table**: Fix links for image cells. [#32370](https://github.com/grafana/grafana/pull/32370), [@kaydelaney](https://github.com/kaydelaney)
* **Variables**: Fix data source variable when default data source is selected. [#32384](https://github.com/grafana/grafana/pull/32384), [@torkelo](https://github.com/torkelo)
* **Variables**: Fix manually entering non-matching custom value in variable input/picker error. [#32390](https://github.com/grafana/grafana/pull/32390), [@torkelo](https://github.com/torkelo)

<!-- 7.5.2 END -->

<!-- 7.5.1 START -->

# 7.5.1 (2021-03-26)

### Bug fixes

* **MSSQL**: Fix panic not implemented by upgrading go-mssqldb dependency. [#32347](https://github.com/grafana/grafana/pull/32347), [@aknuds1](https://github.com/aknuds1)

<!-- 7.5.1 END -->

<!-- 7.5.0 START -->

# 7.5.0 (2021-03-25)

### Features and enhancements

* **Alerting**: Add ability to include aliases with hyphen in InfluxDB. [#32262](https://github.com/grafana/grafana/pull/32262), [@grafanabot](https://github.com/grafanabot)
* **CloudWatch**: Use latest version of aws sdk. [#32217](https://github.com/grafana/grafana/pull/32217), [@sunker](https://github.com/sunker)

### Bug fixes

* **Alerting**: Add ability to include aliases with hyphen in InfluxDB. [#32224](https://github.com/grafana/grafana/pull/32224), [@dsotirakis](https://github.com/dsotirakis)
* **DashboardSettings**: Fixes issue with tags list not updating changes are made. [#32241](https://github.com/grafana/grafana/pull/32241), [@huynhsamha](https://github.com/huynhsamha)
* **DashboardSettings**: Fixes issue with tags list not updating changes are made. [#32189](https://github.com/grafana/grafana/pull/32189), [@huynhsamha](https://github.com/huynhsamha)
* **Loki**: Fix text search in Label browser. [#32293](https://github.com/grafana/grafana/pull/32293), [@ivanahuckova](https://github.com/ivanahuckova)

<!-- 7.5.0 END -->

<!-- 7.5.0-beta2 START -->

# 7.5.0-beta2 (2021-03-19)

### Features and enhancements

* **CloudWatch**: Add support for EC2 IAM role. [#31804](https://github.com/grafana/grafana/pull/31804), [@sunker](https://github.com/sunker)
* **CloudWatch**: Consume the grafana/aws-sdk. [#31807](https://github.com/grafana/grafana/pull/31807), [@sunker](https://github.com/sunker)
* **CloudWatch**: Restrict auth provider and assume role usage according to Grafana configuration. [#31805](https://github.com/grafana/grafana/pull/31805), [@sunker](https://github.com/sunker)
* **Cloudwatch**: ListMetrics API page limit. [#31788](https://github.com/grafana/grafana/pull/31788), [@sunker](https://github.com/sunker)
* **Cloudwatch**: Use shared library for aws auth. [#29550](https://github.com/grafana/grafana/pull/29550), [@ryantxu](https://github.com/ryantxu)
* **DataLinks**: Bring back single click links for Stat, Gauge and BarGauge panel. [#31692](https://github.com/grafana/grafana/pull/31692), [@dprokop](https://github.com/dprokop)
* **Docker**: Support pre-installed plugins from other sources in custom Dockerfiles. [#31234](https://github.com/grafana/grafana/pull/31234), [@sgnsys3](https://github.com/sgnsys3)
* **Elasticseach**: Add support for histogram fields. [#29079](https://github.com/grafana/grafana/pull/29079), [@simianhacker](https://github.com/simianhacker)
* **Exemplars**: Always query exemplars. [#31673](https://github.com/grafana/grafana/pull/31673), [@zoltanbedi](https://github.com/zoltanbedi)
* **Explore**: Support full inspect drawer. [#32005](https://github.com/grafana/grafana/pull/32005), [@ivanahuckova](https://github.com/ivanahuckova)
* **HttpServer**: Make read timeout configurable but disabled by default. [#31575](https://github.com/grafana/grafana/pull/31575), [@bergquist](https://github.com/bergquist)
* **SQLStore**: Close session in withDbSession. [#31775](https://github.com/grafana/grafana/pull/31775), [@aknuds1](https://github.com/aknuds1)
* **Templating**: Use dashboard timerange when variables are set to refresh 'On Dashboard Load'. [#31721](https://github.com/grafana/grafana/pull/31721), [@Elfo404](https://github.com/Elfo404)
* **Tempo**: Convert to backend data source. [#31618](https://github.com/grafana/grafana/pull/31618), [@zoltanbedi](https://github.com/zoltanbedi)

### Bug fixes

* **Admin**: Keeps expired api keys visible in table after delete. [#31636](https://github.com/grafana/grafana/pull/31636), [@hugohaggmark](https://github.com/hugohaggmark)
* **Data proxy**: Fix encoded characters in URL path should be proxied as encoded. [#30597](https://github.com/grafana/grafana/pull/30597), [@marefr](https://github.com/marefr)
* **Explore/Logs**: Fix escaping in ANSI logs. [#31731](https://github.com/grafana/grafana/pull/31731), [@ivanahuckova](https://github.com/ivanahuckova)
* **GraphNG**: Fix tooltip series color for multi data frame scenario. [#32098](https://github.com/grafana/grafana/pull/32098), [@dprokop](https://github.com/dprokop)
* **GraphNG**: Make sure data set and config are in sync when initializing and re-initializing uPlot. [#32106](https://github.com/grafana/grafana/pull/32106), [@dprokop](https://github.com/dprokop)
* **Loki**: Fix autocomplete when re-editing Loki label values. [#31828](https://github.com/grafana/grafana/pull/31828), [@ivanahuckova](https://github.com/ivanahuckova)
* **MixedDataSource**: Name is updated when data source variables change. [#32090](https://github.com/grafana/grafana/pull/32090), [@hugohaggmark](https://github.com/hugohaggmark)
* **PanelInspect**: Interpolates variables in CSV file name. [#31936](https://github.com/grafana/grafana/pull/31936), [@hugohaggmark](https://github.com/hugohaggmark)
* **ReduceTransform**: Include series with numeric string  names. [#31763](https://github.com/grafana/grafana/pull/31763), [@hugohaggmark](https://github.com/hugohaggmark)
* **Snapshots**: Fix usage of sign in link from the snapshot page. [#31986](https://github.com/grafana/grafana/pull/31986), [@marefr](https://github.com/marefr)
* **TimePicker**: Fixes hidden time picker shown in kiosk TV mode. [#32062](https://github.com/grafana/grafana/pull/32062), [@torkelo](https://github.com/torkelo)
* **ValueMappings**: Fixes value 0 not being mapped. [#31924](https://github.com/grafana/grafana/pull/31924), [@Willena](https://github.com/Willena)
* **Variables**: Fixes filtering in picker with null items. [#31979](https://github.com/grafana/grafana/pull/31979), [@hugohaggmark](https://github.com/hugohaggmark)
* **Variables**: Improves inspection performance and unknown filtering. [#31811](https://github.com/grafana/grafana/pull/31811), [@hugohaggmark](https://github.com/hugohaggmark)

### Plugin development fixes & changes

* **Auth**: Allow soft token revocation. [#31601](https://github.com/grafana/grafana/pull/31601), [@joanlopez](https://github.com/joanlopez)

<!-- 7.5.0-beta2 END -->

<!-- 7.5.0-beta1 START -->

# 7.5.0-beta1 (2021-03-04)

### Features and enhancements

* **Alerting**: Customise OK notification priorities for Pushover notifier. [#30169](https://github.com/grafana/grafana/pull/30169), [@acaire](https://github.com/acaire)
* **Alerting**: Improve default message for SensuGo notifier. [#31428](https://github.com/grafana/grafana/pull/31428), [@M4teo](https://github.com/M4teo)
* **Alerting**: PagerDuty: adding current state to the payload. [#29270](https://github.com/grafana/grafana/pull/29270), [@Eraac](https://github.com/Eraac)
* **AzureMonitor**: Add deprecation message for App Insights/Insights Analytics. [#30633](https://github.com/grafana/grafana/pull/30633), [@joshhunt](https://github.com/joshhunt)
* **CloudMonitoring**: Allow free text input for GCP project on dashboard variable query. [#28048](https://github.com/grafana/grafana/issues/28048)
* **CloudMonitoring**: Increase service api page size. [#30892](https://github.com/grafana/grafana/pull/30892), [@sunker](https://github.com/sunker)
* **CloudMonitoring**: Show service and SLO display name in SLO Query editor. [#30900](https://github.com/grafana/grafana/pull/30900), [@sunker](https://github.com/sunker)
* **CloudWatch**: Add AWS Ground Station metrics and dimensions. [#31362](https://github.com/grafana/grafana/pull/31362), [@ilyastoli](https://github.com/ilyastoli)
* **CloudWatch**: Add AWS Network Firewall metrics and dimensions. [#31498](https://github.com/grafana/grafana/pull/31498), [@ilyastoli](https://github.com/ilyastoli)
* **CloudWatch**: Add AWS Timestream Metrics and Dimensions. [#31624](https://github.com/grafana/grafana/pull/31624), [@ilyastoli](https://github.com/ilyastoli)
* **CloudWatch**: Add RDS Proxy metrics. [#31595](https://github.com/grafana/grafana/pull/31595), [@sunker](https://github.com/sunker)
* **CloudWatch**: Add eu-south-1 Cloudwatch region. [#31198](https://github.com/grafana/grafana/pull/31198), [@rubycut](https://github.com/rubycut)
* **CloudWatch**: Make it possible to specify custom api endpoint. [#31402](https://github.com/grafana/grafana/pull/31402), [@sunker](https://github.com/sunker)
* **Cloudwatch**: Add AWS/DDoSProtection metrics and dimensions. [#31297](https://github.com/grafana/grafana/pull/31297), [@relvira](https://github.com/relvira)
* **Dashboard**: Remove template variables option from ShareModal. [#30395](https://github.com/grafana/grafana/pull/30395), [@oscarkilhed](https://github.com/oscarkilhed)
* **Docs**: Define TLS/SSL terminology. [#30533](https://github.com/grafana/grafana/pull/30533), [@aknuds1](https://github.com/aknuds1)
* **Elasticsearch**: Add word highlighting to search results. [#30293](https://github.com/grafana/grafana/pull/30293), [@simianhacker](https://github.com/simianhacker)
* **Folders**: Editors should be able to edit name and delete folders. [#31242](https://github.com/grafana/grafana/pull/31242), [@torkelo](https://github.com/torkelo)
* **Graphite/SSE**: update graphite to work with server side expressions. [#31455](https://github.com/grafana/grafana/pull/31455), [@kylebrandt](https://github.com/kylebrandt)
* **InfluxDB**: Improve maxDataPoints error-message in Flux-mode, raise limits. [#31259](https://github.com/grafana/grafana/pull/31259), [@gabor](https://github.com/gabor)
* **InfluxDB**: In flux query editor, do not run query when disabled. [#31324](https://github.com/grafana/grafana/pull/31324), [@gabor](https://github.com/gabor)
* **LogsPanel**: Add deduplication option for logs. [#31019](https://github.com/grafana/grafana/pull/31019), [@ivanahuckova](https://github.com/ivanahuckova)
* **Loki**: Add line limit for annotations. [#31183](https://github.com/grafana/grafana/pull/31183), [@ivanahuckova](https://github.com/ivanahuckova)
* **Loki**: Add support for alerting. [#31424](https://github.com/grafana/grafana/pull/31424), [@ivanahuckova](https://github.com/ivanahuckova)
* **Loki**: Label browser. [#30351](https://github.com/grafana/grafana/pull/30351), [@davkal](https://github.com/davkal)
* **PieChart**: Add color changing options to pie chart. [#31588](https://github.com/grafana/grafana/pull/31588), [@oscarkilhed](https://github.com/oscarkilhed)
* **PostgreSQL**: Allow providing TLS/SSL certificates as text in addition to file paths. [#30353](https://github.com/grafana/grafana/pull/30353), [@ying-jeanne](https://github.com/ying-jeanne)
* **Postgres**: SSL certification. [#30352](https://github.com/grafana/grafana/pull/30352), [@ying-jeanne](https://github.com/ying-jeanne)
* **Profile**: Prevent OAuth users from changing user details or password. [#27886](https://github.com/grafana/grafana/pull/27886), [@dupondje](https://github.com/dupondje)
* **Prometheus**: Change default httpMethod for new instances to POST. [#31292](https://github.com/grafana/grafana/pull/31292), [@ivanahuckova](https://github.com/ivanahuckova)
* **Prometheus**: Min step defaults to seconds when no unit is set. [#30966](https://github.com/grafana/grafana/pull/30966), [@nutmos](https://github.com/nutmos)
* **Stats**: Exclude folders from total dashboard count. [#31320](https://github.com/grafana/grafana/pull/31320), [@bergquist](https://github.com/bergquist)
* **Tracing**: Specify type of data frame that is expected for TraceView. [#31465](https://github.com/grafana/grafana/pull/31465), [@aocenas](https://github.com/aocenas)
* **Transformers**: Add search to transform selection. [#30854](https://github.com/grafana/grafana/pull/30854), [@ryantxu](https://github.com/ryantxu)

### Bug fixes

* **Alerting**: Ensure Discord notification is sent when metric name is absent. [#31257](https://github.com/grafana/grafana/pull/31257), [@LeviHarrison](https://github.com/LeviHarrison)
* **Alerting**: Fix case when Alertmanager notifier fails if a URL is not working. [#31079](https://github.com/grafana/grafana/pull/31079), [@kurokochin](https://github.com/kurokochin)
* **CloudMonitoring**: Prevent resource type variable function from crashing. [#30901](https://github.com/grafana/grafana/pull/30901), [@sunker](https://github.com/sunker)
* **Color**: Fix issue where colors are reset to gray when switching panels. [#31611](https://github.com/grafana/grafana/pull/31611), [@torkelo](https://github.com/torkelo)
* **Explore**: Show ANSI colored logs in logs context. [#31510](https://github.com/grafana/grafana/pull/31510), [@ivanahuckova](https://github.com/ivanahuckova)
* **Explore**: keep enabled/disabled state in angular based QueryEditors correctly. [#31558](https://github.com/grafana/grafana/pull/31558), [@gabor](https://github.com/gabor)
* **Graph**: Fix tooltip not being displayed when close to edge of viewport. [#31493](https://github.com/grafana/grafana/pull/31493), [@msober](https://github.com/msober)
* **Heatmap**: Fix missing value in legend. [#31430](https://github.com/grafana/grafana/pull/31430), [@kurokochin](https://github.com/kurokochin)
* **InfluxDB**: Handle columns named "table". [#30985](https://github.com/grafana/grafana/pull/30985), [@gabor](https://github.com/gabor)
* **Prometheus**: Use configured HTTP method for /series and /labels endpoints. [#31401](https://github.com/grafana/grafana/pull/31401), [@ivanahuckova](https://github.com/ivanahuckova)
* **RefreshPicker**: Make valid intervals in url visible in RefreshPicker. [#30474](https://github.com/grafana/grafana/pull/30474), [@hugohaggmark](https://github.com/hugohaggmark)
* **TimeSeriesPanel**: Fix overlapping time axis ticks. [#31332](https://github.com/grafana/grafana/pull/31332), [@torkelo](https://github.com/torkelo)
* **TraceViewer**: Fix show log marker in spanbar. [#30742](https://github.com/grafana/grafana/pull/30742), [@zoltanbedi](https://github.com/zoltanbedi)

### Plugin development fixes & changes

* **Plugins**: Add autoEnabled plugin JSON field to auto enable App plugins and add configuration link to menu by default. [#31354](https://github.com/grafana/grafana/pull/31354), [@torkelo](https://github.com/torkelo)
* **Pagination**: Improve pagination for large number of pages. [#30151](https://github.com/grafana/grafana/pull/30151), [@nathanrodman](https://github.com/nathanrodman)

<!-- 7.5.0-beta1 END -->

<!-- 7.4.5 START -->

# 7.4.5 (2021-03-18)

### Bug fixes

* **Security**: Fix API permissions issues related to team-sync CVE-2021-28146, CVE-2021-28147. (Enterprise)
* **Security**: Usage insights requires signed in users CVE-2021-28148. (Enterprise)
* **Security**: Do not allow editors to incorrectly bypass permissions on the default data source. CVE-2021-27962. (Enterprise)

<!-- 7.4.5 END -->

<!-- 7.4.3 START -->

# 7.4.3 (2021-02-24)

### Bug fixes

* **AdHocVariables**: Fixes crash when values are stored as numbers. [#31382](https://github.com/grafana/grafana/pull/31382), [@hugohaggmark](https://github.com/hugohaggmark)
* **DashboardLinks**: Fix an issue where the dashboard links were causing a full page reload. [#31334](https://github.com/grafana/grafana/pull/31334), [@torkelo](https://github.com/torkelo)
* **Elasticsearch**: Fix query initialization logic & query transformation from Prometheus/Loki. [#31322](https://github.com/grafana/grafana/pull/31322), [@Elfo404](https://github.com/Elfo404)
* **QueryEditor**: Fix disabling queries in dashboards. [#31336](https://github.com/grafana/grafana/pull/31336), [@gabor](https://github.com/gabor)
* **Streaming**: Fix an issue with the time series panel and streaming data source when scrolling back from being out of view. [#31431](https://github.com/grafana/grafana/pull/31431), [@torkelo](https://github.com/torkelo)
* **Table**: Fix an issue regarding the fixed min and auto max values in bar gauge cell. [#31316](https://github.com/grafana/grafana/pull/31316), [@torkelo](https://github.com/torkelo)

<!-- 7.4.3 END -->

<!-- 7.4.2 START -->

# 7.4.2 (2021-02-17)

### Features and enhancements

* **Explore**: Do not show non queryable data sources in data source picker. [#31144](https://github.com/grafana/grafana/pull/31144), [@torkelo](https://github.com/torkelo)
* **Snapshots**: Do not allow an anonymous user to create snapshots. [#31263](https://github.com/grafana/grafana/pull/31263), [@marefr](https://github.com/marefr)

### Bug fixes

* **CloudWatch**: Ensure empty query row errors are not passed to the panel. [#31172](https://github.com/grafana/grafana/pull/31172), [@sunker](https://github.com/sunker)
* **DashboardLinks**: Fix the links that always cause a full page to reload. [#31178](https://github.com/grafana/grafana/pull/31178), [@torkelo](https://github.com/torkelo)
* **DashboardListPanel**: Fix issue with folder picker always showing All and using old form styles. [#31160](https://github.com/grafana/grafana/pull/31160), [@torkelo](https://github.com/torkelo)
* **IPv6**: Support host address configured with enclosing square brackets. [#31226](https://github.com/grafana/grafana/pull/31226), [@aknuds1](https://github.com/aknuds1)
* **Permissions**: Fix team and role permissions on folders/dashboards not displayed for non Grafana Admin users. [#31132](https://github.com/grafana/grafana/pull/31132), [@AgnesToulet](https://github.com/AgnesToulet)
* **Postgres**: Fix timeGroup macro converts long intervals to invalid numbers when TimescaleDB is enabled. [#31179](https://github.com/grafana/grafana/pull/31179), [@kurokochin](https://github.com/kurokochin)
* **Prometheus**: Fix enabling of disabled queries when editing in dashboard. [#31055](https://github.com/grafana/grafana/pull/31055), [@ivanahuckova](https://github.com/ivanahuckova)
* **QueryEditors**: Fix an issue that happens after moving queries then editing would update other queries. [#31193](https://github.com/grafana/grafana/pull/31193), [@torkelo](https://github.com/torkelo)
* **SqlDataSources**: Fix the Show Generated SQL button in query editors. [#31236](https://github.com/grafana/grafana/pull/31236), [@torkelo](https://github.com/torkelo)
* **StatPanels**: Fix an issue where the palette color scheme is not cleared when loading panel. [#31126](https://github.com/grafana/grafana/pull/31126), [@torkelo](https://github.com/torkelo)
* **Variables**: Add the default option back for the data source variable. [#31208](https://github.com/grafana/grafana/pull/31208), [@hugohaggmark](https://github.com/hugohaggmark)
* **Variables**: Fix missing empty elements from regex filters. [#31156](https://github.com/grafana/grafana/pull/31156), [@hugohaggmark](https://github.com/hugohaggmark)

<!-- 7.4.2 END -->

<!-- v7.4.2 START -->


<!-- v7.4.2 END -->

<!-- 7.4.1 START -->

# 7.4.1 (2021-02-11)

### Features and enhancements

* **Influx**: Make max series limit configurable and show the limiting message if applied. [#31025](https://github.com/grafana/grafana/pull/31025), [@aocenas](https://github.com/aocenas)
* **Make value mappings correctly interpret numeric-like strings**. [#30893](https://github.com/grafana/grafana/pull/30893), [@dprokop](https://github.com/dprokop)
* **Variables**: Adds queryparam formatting option. [#30858](https://github.com/grafana/grafana/pull/30858), [@hugohaggmark](https://github.com/hugohaggmark)

### Bug fixes

* **Alerting**: Fixes so notification channels are properly deleted. [#31040](https://github.com/grafana/grafana/pull/31040), [@hugohaggmark](https://github.com/hugohaggmark)
* **BarGauge**: Improvements to value sizing and table inner width calculations. [#30990](https://github.com/grafana/grafana/pull/30990), [@torkelo](https://github.com/torkelo)
* **DashboardLinks**: Fixes crash when link has no title. [#31008](https://github.com/grafana/grafana/pull/31008), [@hugohaggmark](https://github.com/hugohaggmark)
* **Elasticsearch**: Fix alias field value not being shown in query editor. [#30992](https://github.com/grafana/grafana/pull/30992), [@Elfo404](https://github.com/Elfo404)
* **Elasticsearch**: Fix log row context errors. [#31088](https://github.com/grafana/grafana/pull/31088), [@Elfo404](https://github.com/Elfo404)
* **Elasticsearch**: Show Size setting for raw_data metric. [#30980](https://github.com/grafana/grafana/pull/30980), [@Elfo404](https://github.com/Elfo404)
* **Graph**: Fixes so graph is shown for non numeric time values. [#30972](https://github.com/grafana/grafana/pull/30972), [@hugohaggmark](https://github.com/hugohaggmark)
* **Logging**: Ignore 'file already closed' error when closing file. [#31119](https://github.com/grafana/grafana/pull/31119), [@aknuds1](https://github.com/aknuds1)
* **Plugins**: Fix plugin signature validation for manifest v2 on Windows. [#31045](https://github.com/grafana/grafana/pull/31045), [@wbrowne](https://github.com/wbrowne)
* **TextPanel**: Fixes so panel title is updated when variables change. [#30884](https://github.com/grafana/grafana/pull/30884), [@hugohaggmark](https://github.com/hugohaggmark)
* **Transforms**: Fixes Outer join issue with duplicate field names not getting the same unique field names as before. [#31121](https://github.com/grafana/grafana/pull/31121), [@torkelo](https://github.com/torkelo)

<!-- 7.4.1 END -->

<!-- 7.4.0 START -->

# 7.4.0 (2021-02-04)

### Features and enhancements

* **CDN**: Adds support for serving assets over a CDN. [#30691](https://github.com/grafana/grafana/pull/30691), [@torkelo](https://github.com/torkelo)
* **DashboardLinks**: Support variable expression in to tooltip - Issue #30409. [#30569](https://github.com/grafana/grafana/pull/30569), [@huynhsamha](https://github.com/huynhsamha)
* **Explore**: Set Explore's GraphNG to be connected. [#30707](https://github.com/grafana/grafana/pull/30707), [@ivanahuckova](https://github.com/ivanahuckova)
* **InfluxDB**: Add http configuration when selecting InfluxDB v2 flavor. [#30827](https://github.com/grafana/grafana/pull/30827), [@aocenas](https://github.com/aocenas)
* **InfluxDB**: Show all datapoints for dynamically windowed flux query. [#30688](https://github.com/grafana/grafana/pull/30688), [@davkal](https://github.com/davkal)
* **Loki**: Improve live tailing errors. [#30517](https://github.com/grafana/grafana/pull/30517), [@ivanahuckova](https://github.com/ivanahuckova)

### Bug fixes

* **Admin**: Fixes so form values are filled in from backend. [#30544](https://github.com/grafana/grafana/pull/30544), [@hugohaggmark](https://github.com/hugohaggmark)
* **Admin**: Fixes so whole org drop down is visible when adding users to org. [#30481](https://github.com/grafana/grafana/pull/30481), [@hugohaggmark](https://github.com/hugohaggmark)
* **Alerting**: Hides threshold handle for percentual thresholds. [#30431](https://github.com/grafana/grafana/pull/30431), [@hugohaggmark](https://github.com/hugohaggmark)
* **CloudWatch**: Prevent field config from being overwritten. [#30437](https://github.com/grafana/grafana/pull/30437), [@sunker](https://github.com/sunker)
* **Decimals**: Big Improvements to auto decimals and fixes to auto decimals bug found in 7.4-beta1. [#30519](https://github.com/grafana/grafana/pull/30519), [@torkelo](https://github.com/torkelo)
* **Explore**: Fix jumpy live tailing. [#30650](https://github.com/grafana/grafana/pull/30650), [@ivanahuckova](https://github.com/ivanahuckova)
* **Explore**: Fix loading visualisation on the top of the new time series panel. [#30553](https://github.com/grafana/grafana/pull/30553), [@ivanahuckova](https://github.com/ivanahuckova)
* **Footer**: Fixes layout issue in footer. [#30443](https://github.com/grafana/grafana/pull/30443), [@torkelo](https://github.com/torkelo)
* **Graph**: Fixes so only users with correct permissions can add annotations. [#30419](https://github.com/grafana/grafana/pull/30419), [@hugohaggmark](https://github.com/hugohaggmark)
* **Mobile**: Fixes issue scrolling on mobile in chrome. [#30746](https://github.com/grafana/grafana/pull/30746), [@torkelo](https://github.com/torkelo)
* **PanelEdit**: Trigger refresh when changing data source. [#30744](https://github.com/grafana/grafana/pull/30744), [@torkelo](https://github.com/torkelo)
* **Panels**: Fixes so panels are refreshed when scrolling past them fast. [#30784](https://github.com/grafana/grafana/pull/30784), [@hugohaggmark](https://github.com/hugohaggmark)
* **Prometheus**: Fix show query instead of Value if no __name__ and metric. [#30511](https://github.com/grafana/grafana/pull/30511), [@zoltanbedi](https://github.com/zoltanbedi)
* **TimeSeriesPanel**: Fixes default value for Gradient mode. [#30484](https://github.com/grafana/grafana/pull/30484), [@torkelo](https://github.com/torkelo)
* **Variables**: Clears drop down state when leaving dashboard. [#30810](https://github.com/grafana/grafana/pull/30810), [@hugohaggmark](https://github.com/hugohaggmark)
* **Variables**: Fixes display value when using capture groups in regex. [#30636](https://github.com/grafana/grafana/pull/30636), [@hugohaggmark](https://github.com/hugohaggmark)
* **Variables**: Fixes so queries work for numbers values too. [#30602](https://github.com/grafana/grafana/pull/30602), [@hugohaggmark](https://github.com/hugohaggmark)
* **Variables**: Fixes so text format will show All instead of custom all value. [#30730](https://github.com/grafana/grafana/pull/30730), [@hugohaggmark](https://github.com/hugohaggmark)

### Plugin development fixes & changes

* **Plugins**: Fix failing plugin builds because of wrong internal import. [#30439](https://github.com/grafana/grafana/pull/30439), [@aocenas](https://github.com/aocenas)

<!-- 7.4.0 END -->

<!-- 7.4.0-beta1 START -->

# 7.4.0-beta1 (2021-01-20)

### Features and enhancements

* **API**: Add ID to snapshot API responses. [#29600](https://github.com/grafana/grafana/pull/29600), [@AgnesToulet](https://github.com/AgnesToulet)
* **AlertListPanel**: Add options to sort by Time(asc) and Time(desc). [#29764](https://github.com/grafana/grafana/pull/29764), [@dboslee](https://github.com/dboslee)
* **AlertListPanel**: Changed alert url to to go the panel view instead of panel edit. [#29060](https://github.com/grafana/grafana/pull/29060), [@zakiharis](https://github.com/zakiharis)
* **Alerting**: Add support for Sensu Go notification channel. [#28012](https://github.com/grafana/grafana/pull/28012), [@nixwiz](https://github.com/nixwiz)
* **Alerting**: Add support for alert notification query label interpolation. [#29908](https://github.com/grafana/grafana/pull/29908), [@wbrowne](https://github.com/wbrowne)
* **Annotations**: Remove annotation_tag entries as part of annotations cleanup. [#29534](https://github.com/grafana/grafana/pull/29534), [@dafydd-t](https://github.com/dafydd-t)
* **Azure Monitor**: Add Microsoft.Network/natGateways. [#29479](https://github.com/grafana/grafana/pull/29479), [@JoeyLemur](https://github.com/JoeyLemur)
* **Backend plugins**: Support Forward OAuth Identity for backend data source plugins. [#27055](https://github.com/grafana/grafana/pull/27055), [@billoley](https://github.com/billoley)
* **Cloud Monitoring**: MQL support. [#26551](https://github.com/grafana/grafana/pull/26551), [@mtanda](https://github.com/mtanda)
* **CloudWatch**: Add 'EventBusName' dimension to CloudWatch 'AWS/Events' namespace. [#28402](https://github.com/grafana/grafana/pull/28402), [@tomdaly](https://github.com/tomdaly)
* **CloudWatch**: Add support for AWS DirectConnect ConnectionErrorCount metric. [#29583](https://github.com/grafana/grafana/pull/29583), [@haeringer](https://github.com/haeringer)
* **CloudWatch**: Add support for AWS/ClientVPN metrics and dimensions. [#29055](https://github.com/grafana/grafana/pull/29055), [@marefr](https://github.com/marefr)
* **CloudWatch**: Added HTTP API Gateway specific metrics and dimensions. [#28780](https://github.com/grafana/grafana/pull/28780), [@karlatkinson](https://github.com/karlatkinson)
* **Configuration**: Add an option to hide certain users in the UI. [#28942](https://github.com/grafana/grafana/pull/28942), [@AgnesToulet](https://github.com/AgnesToulet)
* **Currency**: Adds Indonesian IDR currency. [#28363](https://github.com/grafana/grafana/pull/28363), [@hiddenrebel](https://github.com/hiddenrebel)
* **Dashboards**: Delete related data (permissions, stars, tags, versions, annotations) when deleting a dashboard or a folder. [#28826](https://github.com/grafana/grafana/pull/28826), [@AgnesToulet](https://github.com/AgnesToulet)
* **Dependencies**: Update angularjs to 1.8.2. [#28736](https://github.com/grafana/grafana/pull/28736), [@torkelo](https://github.com/torkelo)
* **Docker**:  Use root group in the custom Dockerfile. [#28639](https://github.com/grafana/grafana/pull/28639), [@chugunov](https://github.com/chugunov)
* **Elasticsearch**: Add Moving Function Pipeline Aggregation. [#28131](https://github.com/grafana/grafana/pull/28131), [@simianhacker](https://github.com/simianhacker)
* **Elasticsearch**: Add Support for Serial Differencing Pipeline Aggregation. [#28618](https://github.com/grafana/grafana/pull/28618), [@simianhacker](https://github.com/simianhacker)
* **Elasticsearch**: Deprecate browser access mode. [#29649](https://github.com/grafana/grafana/pull/29649), [@Elfo404](https://github.com/Elfo404)
* **Elasticsearch**: Interpolate variables in Filters Bucket Aggregation. [#28969](https://github.com/grafana/grafana/pull/28969), [@Elfo404](https://github.com/Elfo404)
* **Elasticsearch**: Support extended stats and percentiles in terms order by. [#28910](https://github.com/grafana/grafana/pull/28910), [@simianhacker](https://github.com/simianhacker)
* **Elasticsearch**: View in context feature for logs. [#28764](https://github.com/grafana/grafana/pull/28764), [@simianhacker](https://github.com/simianhacker)
* **Explore/Logs**: Alphabetically sort unique labels, labels and parsed fields. [#29030](https://github.com/grafana/grafana/pull/29030), [@ivanahuckova](https://github.com/ivanahuckova)
* **Explore/Logs**: Update Parsed fields to Detected fields. [#28881](https://github.com/grafana/grafana/pull/28881), [@ivanahuckova](https://github.com/ivanahuckova)
* **Field overrides**: Added matcher to match all fields returned by a specific query. [#28872](https://github.com/grafana/grafana/pull/28872), [@mckn](https://github.com/mckn)
* **Graph**: Add support for spline interpolation (smoothing) added in new time series panel. [#4303](https://github.com/grafana/grafana/issues/4303)
* **Instrumentation**: Add histograms for database queries. [#29662](https://github.com/grafana/grafana/pull/29662), [@dafydd-t](https://github.com/dafydd-t)
* **Jaeger**: Remove browser access mode. [#30349](https://github.com/grafana/grafana/pull/30349), [@zoltanbedi](https://github.com/zoltanbedi)
* **LogsPanel**: Don't show scroll bars when not needed. [#28972](https://github.com/grafana/grafana/pull/28972), [@aocenas](https://github.com/aocenas)
* **Loki**: Add query type and line limit to query editor in dashboard. [#29356](https://github.com/grafana/grafana/pull/29356), [@ivanahuckova](https://github.com/ivanahuckova)
* **Loki**: Add query type selector to query editor in Explore. [#28817](https://github.com/grafana/grafana/pull/28817), [@ivanahuckova](https://github.com/ivanahuckova)
* **Loki**: Retry web socket connection when connection is closed abnormally. [#29438](https://github.com/grafana/grafana/pull/29438), [@ivanahuckova](https://github.com/ivanahuckova)
* **MS SQL**: Integrated security. [#30369](https://github.com/grafana/grafana/pull/30369), [@daniellee](https://github.com/daniellee)
* **Middleware**: Add CSP support. [#29740](https://github.com/grafana/grafana/pull/29740), [@aknuds1](https://github.com/aknuds1)
* **OAuth**: Configurable user name attribute. [#28286](https://github.com/grafana/grafana/pull/28286), [@alexanderzobnin](https://github.com/alexanderzobnin)
* **PanelEditor**: Render panel field config categories as separate option group sections. [#30301](https://github.com/grafana/grafana/pull/30301), [@dprokop](https://github.com/dprokop)
* **Postgres**: SSL certification. [#30352](https://github.com/grafana/grafana/pull/30352), [@ying-jeanne](https://github.com/ying-jeanne)
* **Prometheus**: Add support for Exemplars. [#28057](https://github.com/grafana/grafana/pull/28057), [@zoltanbedi](https://github.com/zoltanbedi)
* **Prometheus**: Improve autocomplete performance and remove disabling of dynamic label lookup. [#30199](https://github.com/grafana/grafana/pull/30199), [@ivanahuckova](https://github.com/ivanahuckova)
* **Prometheus**: Update default query type option to "Both" in Explore query editor. [#28935](https://github.com/grafana/grafana/pull/28935), [@ivanahuckova](https://github.com/ivanahuckova)
* **Prometheus**: Use customQueryParameters for all queries. [#28949](https://github.com/grafana/grafana/pull/28949), [@alexbumbacea](https://github.com/alexbumbacea)
* **Security**: Prefer server cipher suites for http2. [#29379](https://github.com/grafana/grafana/pull/29379), [@bergquist](https://github.com/bergquist)
* **Security**: Remove insecure cipher suit as default option. [#29378](https://github.com/grafana/grafana/pull/29378), [@bergquist](https://github.com/bergquist)
* **StatPanels**: Add new calculation option for percentage difference. [#26369](https://github.com/grafana/grafana/pull/26369), [@jedstar](https://github.com/jedstar)
* **StatPanels**: Change default stats option to "Last (not null)". [#28617](https://github.com/grafana/grafana/pull/28617), [@ryantxu](https://github.com/ryantxu)
* **Table**: migrate old-table config to new table config. [#30142](https://github.com/grafana/grafana/pull/30142), [@jackw](https://github.com/jackw)
* **Templating**: Custom variable edit UI, change options input into textarea. [#28322](https://github.com/grafana/grafana/pull/28322), [@darrylsepeda](https://github.com/darrylsepeda)
* **TimeSeriesPanel**: The new graph panel now supports y-axis value mapping. [#30272](https://github.com/grafana/grafana/pull/30272), [@torkelo](https://github.com/torkelo)
* **Tracing**: Tag spans with user login and datasource name instead of id. [#29183](https://github.com/grafana/grafana/pull/29183), [@bergquist](https://github.com/bergquist)
* **Transformations**: Add "Rename By Regex" transformer. [#29281](https://github.com/grafana/grafana/pull/29281), [@simianhacker](https://github.com/simianhacker)
* **Transformations**: Added new transform for excluding and including rows based on their values. [#26884](https://github.com/grafana/grafana/pull/26884), [@Totalus](https://github.com/Totalus)
* **Transforms**: Add sort by transformer. [#30370](https://github.com/grafana/grafana/pull/30370), [@ryantxu](https://github.com/ryantxu)
* **Variables**: Add deprecation warning for value group tags. [#30160](https://github.com/grafana/grafana/pull/30160), [@torkelo](https://github.com/torkelo)
* **Variables**: Added __user.email to global variable. [#28853](https://github.com/grafana/grafana/pull/28853), [@mckn](https://github.com/mckn)
* **Variables**: Adds description field. [#29332](https://github.com/grafana/grafana/pull/29332), [@hugohaggmark](https://github.com/hugohaggmark)
* **Variables**: Adds variables inspection. [#25214](https://github.com/grafana/grafana/pull/25214), [@hugohaggmark](https://github.com/hugohaggmark)
* **Variables**: New Variables are stored immediately. [#29178](https://github.com/grafana/grafana/pull/29178), [@hugohaggmark](https://github.com/hugohaggmark)
* **Zipkin**: Remove browser access mode. [#30360](https://github.com/grafana/grafana/pull/30360), [@zoltanbedi](https://github.com/zoltanbedi)

### Bug fixes

* **API**: Query database from /api/health endpoint. [#28349](https://github.com/grafana/grafana/pull/28349), [@ceh](https://github.com/ceh)
* **Alerting**: Return proper status code when trying to create alert notification channel with duplicate name or uid. [#28043](https://github.com/grafana/grafana/pull/28043), [@jgulick48](https://github.com/jgulick48)
* **Auth**: Fix default maximum lifetime  an authenticated user can be logged in. [#30030](https://github.com/grafana/grafana/pull/30030), [@papagian](https://github.com/papagian)
* **Backend**: Fix IPv6 address parsing erroneous. [#28585](https://github.com/grafana/grafana/pull/28585), [@taciomcosta](https://github.com/taciomcosta)
* **CloudWatch**:  Make sure stats grow horizontally and not vertically in the Query Editor. [#30106](https://github.com/grafana/grafana/pull/30106), [@sunker](https://github.com/sunker)
* **Cloudwatch**: Fix issue with field calculation transform not working properly with Cloudwatch data. [#28761](https://github.com/grafana/grafana/pull/28761), [@torkelo](https://github.com/torkelo)
* **Dashboards**: Hide playlist edit functionality from viewers and snapshots link from unauthenticated users. [#28992](https://github.com/grafana/grafana/pull/28992), [@jackw](https://github.com/jackw)
* **Data source proxy**: Convert 401 HTTP status code from data source to 400. [#28962](https://github.com/grafana/grafana/pull/28962), [@aknuds1](https://github.com/aknuds1)
* **Decimals**: Improving auto decimals logic for high numbers and scaled units. [#30262](https://github.com/grafana/grafana/pull/30262), [@torkelo](https://github.com/torkelo)
* **Elasticsearch**: Fix date histogram auto interval handling for alert queries. [#30049](https://github.com/grafana/grafana/pull/30049), [@simianhacker](https://github.com/simianhacker)
* **Elasticsearch**: Fix index pattern not working with multiple base sections. [#28348](https://github.com/grafana/grafana/pull/28348), [@tomdaly](https://github.com/tomdaly)
* **Explore**: Clear errors after running a new query. [#30367](https://github.com/grafana/grafana/pull/30367), [@ivanahuckova](https://github.com/ivanahuckova)
* **Graph**: Fixes stacking issues like floating bars when data is not aligned. [#29051](https://github.com/grafana/grafana/pull/29051), [@torkelo](https://github.com/torkelo)
* **Graph**: Staircase and null value=null calculates auto Y-Min incorrectly (fixed in new Time series panel). [#12995](https://github.com/grafana/grafana/issues/12995)
* **Graph**: Staircase mode, do now draw line segment from zero when drawing null values as null (Fixed in new Time series panel). [#17838](https://github.com/grafana/grafana/issues/17838)
* **Image uploader**: Fix uploading of images to GCS. [#26493](https://github.com/grafana/grafana/pull/26493), [@gastonqiu](https://github.com/gastonqiu)
* **Influx**: Fixes issue with many queries being issued as you type in the variable query field. [#29968](https://github.com/grafana/grafana/pull/29968), [@dprokop](https://github.com/dprokop)
* **Logs Panel**: Fix inconsistent highlighting. [#28971](https://github.com/grafana/grafana/pull/28971), [@ivanahuckova](https://github.com/ivanahuckova)
* **Logs Panel**: Fixes problem dragging scrollbar inside logs panel. [#28974](https://github.com/grafana/grafana/pull/28974), [@aocenas](https://github.com/aocenas)
* **Loki**: Fix hiding of series in table if labels have number values. [#30185](https://github.com/grafana/grafana/pull/30185), [@ivanahuckova](https://github.com/ivanahuckova)
* **Loki**: Lower min step to 1ms. [#30135](https://github.com/grafana/grafana/pull/30135), [@ivanahuckova](https://github.com/ivanahuckova)
* **Loki**: Remove showing of unique labels with the empty string value. [#30363](https://github.com/grafana/grafana/pull/30363), [@ivanahuckova](https://github.com/ivanahuckova)
* **Loki**: Timeseries should not produce 0-values for missing data. [#30116](https://github.com/grafana/grafana/pull/30116), [@davkal](https://github.com/davkal)
* **Plugins**: Fix panic when using complex dynamic URLs in app plugin routes. [#27977](https://github.com/grafana/grafana/pull/27977), [@cinaglia](https://github.com/cinaglia)
* **Prometheus**: Fix link to Prometheus graph in dashboard. [#29543](https://github.com/grafana/grafana/pull/29543), [@ivanahuckova](https://github.com/ivanahuckova)
* **Provisioning**: Build paths in an os independent way. [#29143](https://github.com/grafana/grafana/pull/29143), [@amattheisen](https://github.com/amattheisen)
* **Provisioning**: Fixed problem with getting started panel being added to custom home dashboard. [#28750](https://github.com/grafana/grafana/pull/28750), [@torkelo](https://github.com/torkelo)
* **SAML**: Fixes bug in processing SAML response with empty <Issuer> element by updating saml library (Enterprise). [#29991](https://github.com/grafana/grafana/pull/29991), [@alexanderzobnin](https://github.com/alexanderzobnin)
* **SQL**: Define primary key for tables without it. [#22255](https://github.com/grafana/grafana/pull/22255), [@azhiltsov](https://github.com/azhiltsov)
* **Tracing**: Fix issue showing more than 300 spans. [#29377](https://github.com/grafana/grafana/pull/29377), [@zoltanbedi](https://github.com/zoltanbedi)
* **Units**: Changes FLOP/s to FLOPS and some other rates per second units get /s suffix. [#28825](https://github.com/grafana/grafana/pull/28825), [@Berbe](https://github.com/Berbe)
* **Variables**: Fixes Constant variable persistence confusion. [#29407](https://github.com/grafana/grafana/pull/29407), [@hugohaggmark](https://github.com/hugohaggmark)
* **Variables**: Fixes Textbox current value persistence. [#29481](https://github.com/grafana/grafana/pull/29481), [@hugohaggmark](https://github.com/hugohaggmark)
* **Variables**: Fixes loading with a custom all value in url. [#28958](https://github.com/grafana/grafana/pull/28958), [@hugohaggmark](https://github.com/hugohaggmark)
* **Variables**: Fixes so clicking on Selected in drop down will exclude All value from selection. [#29844](https://github.com/grafana/grafana/pull/29844), [@hugohaggmark](https://github.com/hugohaggmark)

### Breaking changes


#### Constant variables
 
In order to minimize the confusion with Constant variable usage, we've removed the ability to make Constant variables visible. This change will also migrate __`all`__ existing __`visible`__ Constant variables to Textbox variables because which we think this is a more appropriate type of variable for this use case.
 Issue [#29407](https://github.com/grafana/grafana/issues/29407)


#### Plugin compatibility 

We have upgraded AngularJS from version 1.6.6 to 1.8.2. Due to this upgrade some old angular plugins might stop working and will require a small update. This is due to the deprecation and removal of pre-assigned bindings. So if your custom angular controllers expect component bindings in the controller constructor you need to move this code to an `$onInit` function. For more details on how to migrate AngularJS code open the [migration guide](https://docs.angularjs.org/guide/migration) and search for **pre-assigning bindings**.

In order not to break all angular panel plugins and data sources we have some custom [angular inject behavior](https://github.com/grafana/grafana/blob/master/public/app/core/injectorMonkeyPatch.ts) that makes sure that bindings for these controllers are still set before constructor is called so many old angular panels and data source plugins will still work.  Issue [#28736](https://github.com/grafana/grafana/issues/28736)

### Deprecations


#### Query variable value group tags

This option to group query variable values into groups by tags has been an experimental feature since it was introduced. It was introduced to work around the lack of tags support in time series databases at the time. Now that tags (ie. labels) are the norm there is no longer any great need for this feature. This feature will be removed in Grafana v8 later this year.  Issue [#30160](https://github.com/grafana/grafana/issues/30160)

### Plugin development fixes & changes

* **AngularPlugins**: Angular controller events emitter is now a separate emitter and not the same as PanelModel events emitter. [#30379](https://github.com/grafana/grafana/pull/30379), [@torkelo](https://github.com/torkelo)
* **FieldConfig API**: Add ability to hide field option or disable it from the overrides. [#29879](https://github.com/grafana/grafana/pull/29879), [@dprokop](https://github.com/dprokop)
* **Select**: Changes default menu placement for Select from auto to bottom. [#29837](https://github.com/grafana/grafana/pull/29837), [@hugohaggmark](https://github.com/hugohaggmark)
* **Collapse**: Allow component children to use height: 100% styling. [#29776](https://github.com/grafana/grafana/pull/29776), [@aocenas](https://github.com/aocenas)
* **DataSourceWithBackend**: Throw error if health check fails in DataSourceWithBackend. [#29743](https://github.com/grafana/grafana/pull/29743), [@aocenas](https://github.com/aocenas)
* **NodeGraph**: Add node graph visualization. [#29706](https://github.com/grafana/grafana/pull/29706), [@aocenas](https://github.com/aocenas)
* **FieldColor**: Handling color changes when switching panel types. [#28875](https://github.com/grafana/grafana/pull/28875), [@dprokop](https://github.com/dprokop)
* **CodeEditor**: Added support for javascript language. [#28818](https://github.com/grafana/grafana/pull/28818), [@ae3e](https://github.com/ae3e)
* **grafana/toolkit**: Allow builds with lint warnings. [#28810](https://github.com/grafana/grafana/pull/28810), [@dprokop](https://github.com/dprokop)
* **grafana/toolkit**: Drop console and debugger statements by default when building plugin. [#28776](https://github.com/grafana/grafana/pull/28776), [@dprokop](https://github.com/dprokop)
* **Card**: Add new Card component. [#28216](https://github.com/grafana/grafana/pull/28216), [@Clarity-89](https://github.com/Clarity-89)
* **FieldConfig**: Implementation slider editor (#27592). [#28007](https://github.com/grafana/grafana/pull/28007), [@isaozlerfm](https://github.com/isaozlerfm)
* **MutableDataFrame**: Remove unique field name constraint and values field index and unused/seldom used stuff. [#27573](https://github.com/grafana/grafana/pull/27573), [@torkelo](https://github.com/torkelo)

<!-- 7.4.0-beta1 END -->

<!-- 7.3.7 START -->

# 7.3.7 (2021-01-14)

### Bug fixes

* **Auth**: Add missing request headers to SigV4 middleware allowlist. [#30115](https://github.com/grafana/grafana/pull/30115), [@wbrowne](https://github.com/wbrowne)
* **Elasticsearch**: Sort results by index order as well as @timestamp. [#29761](https://github.com/grafana/grafana/pull/29761), [@STEELBADGE](https://github.com/STEELBADGE)
* **SAML**: Fixes bug in processing SAML response with empty <Issuer> element by updating saml library (Enterprise). [#30179](https://github.com/grafana/grafana/pull/30179), [@alexanderzobnin](https://github.com/alexanderzobnin)
* **SeriesToRows**: Fixes issue in transform so that value field is always named Value. [#30054](https://github.com/grafana/grafana/pull/30054), [@torkelo](https://github.com/torkelo)

<!-- 7.3.7 END -->

<!-- 7.3.6 START -->

# 7.3.6 (2020-12-17)

### Security

* **SAML**: Fixes encoding/xml SAML vulnerability in Grafana Enterprise. [#29875](https://github.com/grafana/grafana/issues/29875)

<!-- 7.3.6 END -->

<!-- 7.3.5 START -->

# 7.3.5 (2020-12-10)

### Features and enhancements

* **Alerting**: Improve Prometheus Alert Rule error message. [#29390](https://github.com/grafana/grafana/pull/29390), [@wbrowne](https://github.com/wbrowne)

### Bug fixes

* **Alerting**: Fix alarm message formatting in Dingding. [#29482](https://github.com/grafana/grafana/pull/29482), [@tomowang](https://github.com/tomowang)
* **AzureMonitor**: Fix unit translation for MilliSeconds. [#29399](https://github.com/grafana/grafana/pull/29399), [@secustor](https://github.com/secustor)
* **Instrumentation**: Fix bug with invalid handler label value for HTTP request metrics. [#29529](https://github.com/grafana/grafana/pull/29529), [@bergquist](https://github.com/bergquist)
* **Prometheus**: Fixes problem where changing display name in Field tab had no effect. [#29441](https://github.com/grafana/grafana/pull/29441), [@zoltanbedi](https://github.com/zoltanbedi)
* **Tracing**: Fixed issue showing more than 300 spans. [#29377](https://github.com/grafana/grafana/pull/29377), [@zoltanbedi](https://github.com/zoltanbedi)

<!-- 7.3.5 END -->

<!-- 7.3.4 START -->

# 7.3.4 (2020-11-24)

### Bug fixes

* **Dashboard**: Fixes kiosk state after being redirected to login page and back. [#29273](https://github.com/grafana/grafana/pull/29273), [@torkelo](https://github.com/torkelo)
* **InfluxDB**: Update flux library to fix support for boolean label values. [#29310](https://github.com/grafana/grafana/pull/29310), [@ryantxu](https://github.com/ryantxu)
* **Security**: Fixes minor security issue with alert notification webhooks that allowed GET & DELETE requests. [#29330](https://github.com/grafana/grafana/pull/29330), [@wbrowne](https://github.com/wbrowne)
* **Table**: Fixes issues with phantom extra 0 for zero values. [#29165](https://github.com/grafana/grafana/pull/29165), [@dprokop](https://github.com/dprokop)

<!-- 7.3.4 END -->

<!-- 7.3.3 START -->

# 7.3.3 (2020-11-17)

### Bug fixes

* **Cloud monitoring**: Fix for multi-value template variable for project selector. [#29042](https://github.com/grafana/grafana/pull/29042), [@papagian](https://github.com/papagian)
* **LogsPanel**: Fixes problem dragging scrollbar inside logs panel. [#28974](https://github.com/grafana/grafana/pull/28974), [@aocenas](https://github.com/aocenas)
* **Provisioning**: Fixes application not pinned to the sidebar when it's enabled. [#29084](https://github.com/grafana/grafana/pull/29084), [@alexanderzobnin](https://github.com/alexanderzobnin)
* **StatPanel**: Fixes hanging issue when all values are zero. [#29077](https://github.com/grafana/grafana/pull/29077), [@torkelo](https://github.com/torkelo)
* **Thresholds**: Fixes color assigned to null values. [#29010](https://github.com/grafana/grafana/pull/29010), [@torkelo](https://github.com/torkelo)

<!-- 7.3.3 END -->

<!-- 7.3.2 START -->

# 7.3.2 (2020-11-11)

  ### Features / Enhancements
  * **CloudWatch Logs**: Change how we measure query progress. [#28912](https://github.com/grafana/grafana/pull/28912), [@aocenas](https://github.com/aocenas)
  * **Dashboards / Folders**: delete related data (permissions, stars, tags, versions, annotations) when deleting a dashboard or a folder. [#28826](https://github.com/grafana/grafana/pull/28826), [@AgnesToulet](https://github.com/AgnesToulet)
  * **Gauge**: Improve font size auto sizing. [#28797](https://github.com/grafana/grafana/pull/28797), [@torkelo](https://github.com/torkelo)
  * **Short URL**: Cleanup unvisited/stale short URLs. [#28867](https://github.com/grafana/grafana/pull/28867), [@wbrowne](https://github.com/wbrowne)
  * **Templating**: Custom variable edit UI, change options input into textarea. [#28322](https://github.com/grafana/grafana/pull/28322), [@darrylsepeda](https://github.com/darrylsepeda)

  ### Bug Fixes
  * **Cloudwatch**: Fix issue with field calculation transform not working properly with Cloudwatch data. [#28761](https://github.com/grafana/grafana/pull/28761), [@torkelo](https://github.com/torkelo)
  * **Dashboard**: fix view panel mode for Safari / iOS. [#28702](https://github.com/grafana/grafana/pull/28702), [@jackw](https://github.com/jackw)
  * **Elasticsearch**: Exclude pipeline aggregations from order by options. [#28620](https://github.com/grafana/grafana/pull/28620), [@simianhacker](https://github.com/simianhacker)
  * **Panel inspect**: Interpolate variables in panel inspect title. [#28779](https://github.com/grafana/grafana/pull/28779), [@dprokop](https://github.com/dprokop)
  * **Prometheus**: Fix copy paste behaving as cut and paste. [#28622](https://github.com/grafana/grafana/pull/28622), [@aocenas](https://github.com/aocenas)
  * **StatPanels**: Fixes auto min max when latest value is zero. [#28982](https://github.com/grafana/grafana/pull/28982), [@torkelo](https://github.com/torkelo)
  * **TableFilters**: Fixes filtering with field overrides. [#28690](https://github.com/grafana/grafana/pull/28690), [@hugohaggmark](https://github.com/hugohaggmark)
  * **Templating**: Speeds up certain variable queries for Postgres MySql MSSql. [#28686](https://github.com/grafana/grafana/pull/28686), [@hugohaggmark](https://github.com/hugohaggmark)
  * **Units**: added support to handle negative fractional numbers. [#28849](https://github.com/grafana/grafana/pull/28849), [@mckn](https://github.com/mckn)
  * **Variables**: Fix backward compatibility in custom variable options that contain colon. [#28896](https://github.com/grafana/grafana/pull/28896), [@mckn](https://github.com/mckn)

<!-- 7.3.2 END -->

# 7.3.1 (2020-10-30)

  ### Bug Fixes
  * **Cloudwatch**: Fix duplicate metric data. [#28642](https://github.com/grafana/grafana/pull/28642), [@zoltanbedi](https://github.com/zoltanbedi)
  * **Loki**: Fix error when some queries return zero results. [#28645](https://github.com/grafana/grafana/pull/28645), [@ivanahuckova](https://github.com/ivanahuckova)
  * **PanelMenu**: Fix panel submenu not being accessible for panels close to the right edge of the screen. [#28666](https://github.com/grafana/grafana/pull/28666), [@torkelo](https://github.com/torkelo)
  * **Plugins**: Fix descendent frontend plugin signature validation. [#28638](https://github.com/grafana/grafana/pull/28638), [@wbrowne](https://github.com/wbrowne)
  * **StatPanel**: Fix value being under graph and reduced likelihood for white and dark value text mixing. [#28641](https://github.com/grafana/grafana/pull/28641), [@torkelo](https://github.com/torkelo)
  * **TextPanel**: Fix problems where text panel would show old content. [#28643](https://github.com/grafana/grafana/pull/28643), [@torkelo](https://github.com/torkelo)

# 7.3.0 (2020-10-28)

  ### Features / Enhancements
  * **AzureMonitor**: Support decimal (as float64) type in analytics/logs. [#28480](https://github.com/grafana/grafana/pull/28480), [@kylebrandt](https://github.com/kylebrandt)
  * **Plugins signing**: UI information. [#28469](https://github.com/grafana/grafana/pull/28469), [@dprokop](https://github.com/dprokop)
  * **Short URL**: Update last seen at when visiting a short URL. [#28565](https://github.com/grafana/grafana/pull/28565), [@marefr](https://github.com/marefr)

  ### Bug Fixes
  * **Alerting**: Log warnings for obsolete notifiers when extracting alerts and remove frequent error log messages. [#28162](https://github.com/grafana/grafana/pull/28162), [@papagian](https://github.com/papagian)
  * **Auth**: Fix SigV4 request verification step for Amazon Elasticsearch Service. [#28481](https://github.com/grafana/grafana/pull/28481), [@wbrowne](https://github.com/wbrowne)
  * **Auth**: Should redirect to login when anonymous enabled and URL with different org than anonymous specified. [#28158](https://github.com/grafana/grafana/pull/28158), [@marefr](https://github.com/marefr)
  * **Elasticsearch**: Fix handling of errors when testing data source. [#28498](https://github.com/grafana/grafana/pull/28498), [@marefr](https://github.com/marefr)
  * **Graphite**: Fix default version to be 1.1. [#28471](https://github.com/grafana/grafana/pull/28471), [@ivanahuckova](https://github.com/ivanahuckova)
  * **StatPanel**: Fixes BizChart error max: yyy should not be less than min zzz. [#28587](https://github.com/grafana/grafana/pull/28587), [@hugohaggmark](https://github.com/hugohaggmark)


# 7.3.0-beta2 (2020-10-22)

### Features / Enhancements
* **Add monitoring mixing for Grafana**. [#28285](https://github.com/grafana/grafana/pull/28285), [@bergquist](https://github.com/bergquist)
* **CloudWatch**: Missing Namespace AWS/EC2CapacityReservations. [#28309](https://github.com/grafana/grafana/pull/28309), [@nonamef](https://github.com/nonamef)
* **Explore**: Support wide data frames. [#28393](https://github.com/grafana/grafana/pull/28393), [@aocenas](https://github.com/aocenas)
* **Instrumentation**: Add counters and histograms for database queries. [#28236](https://github.com/grafana/grafana/pull/28236), [@bergquist](https://github.com/bergquist)
* **Loki**: Visually distinguish error logs for LogQL2. [#28359](https://github.com/grafana/grafana/pull/28359), [@ivanahuckova](https://github.com/ivanahuckova)

### Bug Fixes
* **API**: Fix short URLs. [#28300](https://github.com/grafana/grafana/pull/28300), [@aknuds1](https://github.com/aknuds1)
* **BackendSrv**: Fixes queue countdown when unsubscribe is before response. [#28323](https://github.com/grafana/grafana/pull/28323), [@hugohaggmark](https://github.com/hugohaggmark)
* **CloudWatch/Athena - valid metrics and dimensions.**. [#28436](https://github.com/grafana/grafana/pull/28436), [@kwarunek](https://github.com/kwarunek)
* **Dashboard links**: Places drop down list so it's always visible. [#28330](https://github.com/grafana/grafana/pull/28330), [@maknik](https://github.com/maknik)
* **Graph**: Fix for graph size not taking up full height or width. [#28314](https://github.com/grafana/grafana/pull/28314), [@jackw](https://github.com/jackw)
* **Loki**: Base maxDataPoints limits on query type. [#28298](https://github.com/grafana/grafana/pull/28298), [@aocenas](https://github.com/aocenas)
* **Loki**: Run instant query only when doing metric query. [#28325](https://github.com/grafana/grafana/pull/28325), [@aocenas](https://github.com/aocenas)
* **Plugins**: Don't exit on duplicate plugin. [#28390](https://github.com/grafana/grafana/pull/28390), [@aknuds1](https://github.com/aknuds1)

# 7.3.0-beta1 (2020-10-15)

### Breaking changes

- **CloudWatch**: The AWS CloudWatch data source's authentication scheme has changed. See the [upgrade notes](https://grafana.com/docs/grafana/latest/installation/upgrading/#upgrading-to-v73) for details and how this may affect you.
- **Docker**: The Grafana docker image will run with the root group instead of the Grafana group. This may break builds for users who extend the official Docker images. Refer to the [upgrade notes](https://grafana.com/docs/grafana/latest/installation/upgrading/#upgrading-to-v73) for details.

### Features / Enhancements
* **Alerting**: Add labels to name when converting data frame to series. [#28085](https://github.com/grafana/grafana/pull/28085), [@kylebrandt](https://github.com/kylebrandt)
* **Alerting**: Ensuring LINE Notify notifications are sent for all alert states. [#27639](https://github.com/grafana/grafana/pull/27639), [@haraldkubota](https://github.com/haraldkubota)
* **Auth**: Add SigV4 auth option to datasources. [#27552](https://github.com/grafana/grafana/pull/27552), [@wbrowne](https://github.com/wbrowne)
* **AzureMonitor**: Pass through null values instead of setting 0. [#28126](https://github.com/grafana/grafana/pull/28126), [@kylebrandt](https://github.com/kylebrandt)
* **Cloud Monitoring**: Out-of-the-box dashboards. [#27864](https://github.com/grafana/grafana/pull/27864), [@papagian](https://github.com/papagian)
* **CloudWatch**: Add support for AWS DirectConnect virtual interface metrics and add missing dimensions. [#28008](https://github.com/grafana/grafana/pull/28008), [@jgulick48](https://github.com/jgulick48)
* **CloudWatch**: Adding support for Amazon ElastiCache Redis metrics. [#28040](https://github.com/grafana/grafana/pull/28040), [@jgulick48](https://github.com/jgulick48)
* **CloudWatch**: Adding support for additional Amazon CloudFront metrics. [#28069](https://github.com/grafana/grafana/pull/28069), [@darrylsepeda](https://github.com/darrylsepeda)
* **CloudWatch**: Re-implement authentication. [#25548](https://github.com/grafana/grafana/pull/25548), [@aknuds1](https://github.com/aknuds1),[@patstrom](https://github.com/patstrom)
* **Dashboard**: Allow shortlink generation. [#27409](https://github.com/grafana/grafana/pull/27409), [@MisterSquishy](https://github.com/MisterSquishy)
* **Docker**: OpenShift compatability. [#27813](https://github.com/grafana/grafana/pull/27813), [@xlson](https://github.com/xlson)
* **Elasticsearch**: Support multiple pipeline aggregations for a query. [#27945](https://github.com/grafana/grafana/pull/27945), [@simianhacker](https://github.com/simianhacker)
* **Explore**:  Allow shortlink generation. [#28222](https://github.com/grafana/grafana/pull/28222), [@ivanahuckova](https://github.com/ivanahuckova)
* **Explore**: Remove collapsing of visualisations. [#27026](https://github.com/grafana/grafana/pull/27026), [@ivanahuckova](https://github.com/ivanahuckova)
* **FieldColor**: Adds new standard color option for color. [#28039](https://github.com/grafana/grafana/pull/28039), [@torkelo](https://github.com/torkelo)
* **Gauge**: Improve text sizing and support non threshold color modes. [#28256](https://github.com/grafana/grafana/pull/28256), [@torkelo](https://github.com/torkelo)
* **NamedColors**: Named colors refactors. [#28235](https://github.com/grafana/grafana/pull/28235), [@torkelo](https://github.com/torkelo)
* **Panel Inspect**: Allow CSV download for Excel. [#27284](https://github.com/grafana/grafana/pull/27284), [@tomdaly](https://github.com/tomdaly)
* **Prometheus**: Add time range parameters to labels API. [#27548](https://github.com/grafana/grafana/pull/27548), [@kakkoyun](https://github.com/kakkoyun)
* **Snapshots**: Store dashboard data encrypted in the database. [#28129](https://github.com/grafana/grafana/pull/28129), [@wbrowne](https://github.com/wbrowne)
* **Table**: New cell hover behavior and image cell display mode. [#27669](https://github.com/grafana/grafana/pull/27669), [@torkelo](https://github.com/torkelo)
* **Timezones**: Include IANA timezone canonical name in TimeZoneInfo. [#27591](https://github.com/grafana/grafana/pull/27591), [@dprokop](https://github.com/dprokop)
* **Tracing**: Add Tempo data source. [#28204](https://github.com/grafana/grafana/pull/28204), [@aocenas](https://github.com/aocenas)
* **Transformations**: Add Concatenate fields transformer. [#28237](https://github.com/grafana/grafana/pull/28237), [@ryantxu](https://github.com/ryantxu)
* **Transformations**: improve the reduce transformer. [#27875](https://github.com/grafana/grafana/pull/27875), [@ryantxu](https://github.com/ryantxu)
* **Users**: Expire old user invites. [#27361](https://github.com/grafana/grafana/pull/27361), [@wbrowne](https://github.com/wbrowne)
* **Variables**: Adds loading state and indicators. [#27917](https://github.com/grafana/grafana/pull/27917), [@hugohaggmark](https://github.com/hugohaggmark)
* **Variables**: Adds support for key/value mapping in Custom variable. [#27829](https://github.com/grafana/grafana/pull/27829), [@sartaj10](https://github.com/sartaj10)
* **grafana/toolkit**: expose Jest maxWorkers arg for plugin test & build tasks. [#27724](https://github.com/grafana/grafana/pull/27724), [@domasx2](https://github.com/domasx2)

### Bug Fixes
* **Azure Analytics**: FormatAs Time series groups bool columns wrong. [#27713](https://github.com/grafana/grafana/issues/27713)
* **Azure**: Fixes cancellation of requests with different Azure sources. [#28180](https://github.com/grafana/grafana/pull/28180), [@hugohaggmark](https://github.com/hugohaggmark)
* **BackendSrv**: Reloads page instead of redirect on Unauthorized Error. [#28276](https://github.com/grafana/grafana/pull/28276), [@hugohaggmark](https://github.com/hugohaggmark)
* **Dashboard**: Do not allow users without edit permission to a folder to see new dashboard page. [#28249](https://github.com/grafana/grafana/pull/28249), [@torkelo](https://github.com/torkelo)
* **Dashboard**: Fixed issue accessing horizontal table scrollbar when placed at bottom of dashboard. [#28250](https://github.com/grafana/grafana/pull/28250), [@torkelo](https://github.com/torkelo)
* **DataProxy**: Add additional settings for dataproxy to help with network proxy timeouts. [#27841](https://github.com/grafana/grafana/pull/27841), [@kahinton](https://github.com/kahinton)
* **Database**: Adds new indices to alert_notification_state and alert_rule_tag tables. [#28166](https://github.com/grafana/grafana/pull/28166), [@KarineValenca](https://github.com/KarineValenca)
* **Explore**: Fix showing of Prometheus data in Query inspector. [#28128](https://github.com/grafana/grafana/pull/28128), [@ivanahuckova](https://github.com/ivanahuckova)
* **Explore**: Show results of Prometheus instant queries in formatted table. [#27767](https://github.com/grafana/grafana/pull/27767), [@ivanahuckova](https://github.com/ivanahuckova)
* **Graph**: Prevent legend from overflowing container. [#28254](https://github.com/grafana/grafana/pull/28254), [@jackw](https://github.com/jackw)
* **OAuth**: Fix token refresh failure when custom SSL settings are configured for OAuth provider. [#27523](https://github.com/grafana/grafana/pull/27523), [@billoley](https://github.com/billoley)
* **Plugins**: Let descendant plugins inherit their root's signature. [#27970](https://github.com/grafana/grafana/pull/27970), [@aknuds1](https://github.com/aknuds1)
* **Runtime**: Fix handling of short-lived background services. [#28025](https://github.com/grafana/grafana/pull/28025), [@ahlaw](https://github.com/ahlaw)
* **TemplateSrv**: Fix interpolating strings with object variables. [#28171](https://github.com/grafana/grafana/pull/28171), [@torkelo](https://github.com/torkelo)
* **Variables**: Fixes so constants set from url get completed state. [#28257](https://github.com/grafana/grafana/pull/28257), [@hugohaggmark](https://github.com/hugohaggmark)
* **Variables**: Prevent adhoc filters from crashing when they are not loaded properly. [#28226](https://github.com/grafana/grafana/pull/28226), [@mckn](https://github.com/mckn)

<!-- 7.2.3 START -->

# 7.2.3 (2020-12-17)

### Security

- **SAML**: Fixes encoding/xml SAML vulnerability in Grafana Enterprise [#29875](https://github.com/grafana/grafana/issues/29875), [@bergquist](https://github.com/bergquist)

<!-- 7.2.3 END -->

# 7.2.2 (2020-10-21)

### Features / Enhancements

**Caution:** Please do not use/enable the `database_metrics` feature flag. It will corrupt MySQL database tables. See [#28440](https://github.com/grafana/grafana/issues/28440) for more information.

~~**Instrumentation**: Add counters and histograms for database queries. [#28236](https://github.com/grafana/grafana/pull/28236), [@bergquist](https://github.com/bergquist)~~

- **Instrumentation**: Add histogram for request duration. [#28364](https://github.com/grafana/grafana/pull/28364), [@bergquist](https://github.com/bergquist)
- **Instrumentation**: Adds environment_info metric. [#28355](https://github.com/grafana/grafana/pull/28355), [@bergquist](https://github.com/bergquist)

### Bug Fixes

- **CloudWatch**: Fix custom metrics. [#28391](https://github.com/grafana/grafana/pull/28391), [@aknuds1](https://github.com/aknuds1)

# 7.2.1 (2020-10-08)

### Features / Enhancements
* **Api**: Add /healthz endpoint for health checks. [#27536](https://github.com/grafana/grafana/pull/27536), [@bergquist](https://github.com/bergquist)
* **Api**: Healthchecks should not be rejected due to domain enforcement checks. [#27981](https://github.com/grafana/grafana/pull/27981), [@bergquist](https://github.com/bergquist)
* **Instrumentation**: Removes invalid chars from label names. [#27921](https://github.com/grafana/grafana/pull/27921), [@bergquist](https://github.com/bergquist)
* **Orgs**: Remove organisations deprecation notice from backend. [#27788](https://github.com/grafana/grafana/pull/27788), [@wbrowne](https://github.com/wbrowne)
* **grafana/toolkit**: Add --coverage flag to plugin build command. [#27743](https://github.com/grafana/grafana/pull/27743), [@gassiss](https://github.com/gassiss)

### Bug Fixes
* **BarGauge**: Fixed scrollbar showing for bar gauge in Firefox. [#27784](https://github.com/grafana/grafana/pull/27784), [@torkelo](https://github.com/torkelo)
* **Dashboard**: Honour root_url for Explore link. [#27654](https://github.com/grafana/grafana/pull/27654), [@tiagomotasantos](https://github.com/tiagomotasantos)
* **DashboardLinks**: values in links are updated when variables change. [#27926](https://github.com/grafana/grafana/pull/27926), [@hugohaggmark](https://github.com/hugohaggmark)
* **Elasticsearch**: Add query's refId to each series returned by a query. [#27614](https://github.com/grafana/grafana/pull/27614), [@Elfo404](https://github.com/Elfo404)
* **Elasticsearch**: Fix ad-hoc filter support for Raw Data query and new table panel. [#28064](https://github.com/grafana/grafana/pull/28064), [@Elfo404](https://github.com/Elfo404)
* **Graph**: Fixed histogram bucket calculations to avoid missing buckets. [#27883](https://github.com/grafana/grafana/pull/27883), [@torkelo](https://github.com/torkelo)
* **Loki**: Run instant query only in Explore. [#27974](https://github.com/grafana/grafana/pull/27974), [@ivanahuckova](https://github.com/ivanahuckova)
* **Units**: bps & Bps default scale remains decimal (backwards-compatibility). [#27838](https://github.com/grafana/grafana/pull/27838), [@Berbe](https://github.com/Berbe)
* **ValueMappings**: Fix issue with value mappings in override applying to all columns. [#27718](https://github.com/grafana/grafana/pull/27718), [@torkelo](https://github.com/torkelo)

# 7.2.0 (2020-09-23)

### Features / Enhancements
- **Alerting**: Ensuring notifications displayed correctly in mobile device with Google Chat. [#27578](https://github.com/grafana/grafana/pull/27578), [@alvarolmedo](https://github.com/alvarolmedo)
- **TraceView**: Show full traceID and better discern multiple stackTraces in span details. [#27710](https://github.com/grafana/grafana/pull/27710), [@aocenas](https://github.com/aocenas)

### Bug Fixes
- **DataLinks**: Fixes issue with data links not interpolating values with correct field config. [#27622](https://github.com/grafana/grafana/pull/27622), [@torkelo](https://github.com/torkelo)
- **DataProxy**: Ignore empty URL's in plugin routes. [#27653](https://github.com/grafana/grafana/pull/27653), [@domasx2](https://github.com/domasx2)
- **Field config**: Respect config paths when rendering default value of field config property. [#27652](https://github.com/grafana/grafana/pull/27652), [@dprokop](https://github.com/dprokop)
- **Field config**: Fix mismatch in field config editor types. [#27657](https://github.com/grafana/grafana/pull/27657), [@dprokop](https://github.com/dprokop)
- **Panel editor**: Prevents adding transformations in panels with alerts. [#27706](https://github.com/grafana/grafana/pull/27706), [@hugohaggmark](https://github.com/hugohaggmark)
- **Stat panel**: Fix problem where string values where always green. [#27656](https://github.com/grafana/grafana/pull/27656), [@peterholmberg](https://github.com/peterholmberg)

# 7.2.0-beta2 (2020-09-17)

### Features / Enhancements
  - **API**: Enrich add user to org endpoints with user ID in the response. [#27551](https://github.com/grafana/grafana/pull/27551), [@AgnesToulet](https://github.com/AgnesToulet)
  - **API**: Enrich responses and improve error handling for alerting API endpoints. [#27550](https://github.com/grafana/grafana/pull/27550), [@AgnesToulet](https://github.com/AgnesToulet)
  - **Auth**: Replace maximum inactive/lifetime settings of days to duration. [#27150](https://github.com/grafana/grafana/pull/27150), [@Hansuuuuuuuuuu](https://github.com/Hansuuuuuuuuuu)
  - **Dashboard**: Support configuring default timezone via config file. [#27404](https://github.com/grafana/grafana/pull/27404), [@woutersmeenk](https://github.com/woutersmeenk)
  - **Elasticsearch**: Add support for date_nanos type. [#27538](https://github.com/grafana/grafana/pull/27538), [@Elfo404](https://github.com/Elfo404)
  - **Elasticsearch**: Allow fields starting with underscore. [#27520](https://github.com/grafana/grafana/pull/27520), [@Elfo404](https://github.com/Elfo404)
  - **Elasticsearch**: Increase maximum geohash aggregation precision to 12. [#27539](https://github.com/grafana/grafana/pull/27539), [@Elfo404](https://github.com/Elfo404)
  - **Field config**: Add support for paths in default field config setup. [#27570](https://github.com/grafana/grafana/pull/27570), [@dprokop](https://github.com/dprokop)
  - **Postgres**: Support request cancellation properly (Uses new backendSrv.fetch Observable request API). [#27478](https://github.com/grafana/grafana/pull/27478), [@hugohaggmark](https://github.com/hugohaggmark)
  - **Provisioning**: Remove provisioned dashboards without parental reader. [#26143](https://github.com/grafana/grafana/pull/26143), [@nabokihms](https://github.com/nabokihms)
  - **Variables**: Limit rendering of options in dropdown to improve search performance. [#27525](https://github.com/grafana/grafana/pull/27525), [@guoqn](https://github.com/guoqn)
  - **Units**: Binary-prefixed data rates. [#27022](https://github.com/grafana/grafana/pull/27022), [@Berbe](https://github.com/Berbe)

  ### Bug Fixes
  - **Admin**: Fixes close('X') button layout issue in API keys page. [#27625](https://github.com/grafana/grafana/pull/27625), [@nikasvan](https://github.com/nikasvan)
  - **Alerting**: Fix integration key so it's stored encrypted for Pagerduty notifier. [#27484](https://github.com/grafana/grafana/pull/27484), [@marefr](https://github.com/marefr)
  - **Annotations**: Fixes issue with showing error notice for cancelled annotation queries. [#27557](https://github.com/grafana/grafana/pull/27557), [@torkelo](https://github.com/torkelo)
  - **Azure/Insights**: Fix handling of legacy dimension values. [#27513](https://github.com/grafana/grafana/pull/27513), [@marefr](https://github.com/marefr)
  - **DataLinks**: Respects display name and adds field quoting. [#27616](https://github.com/grafana/grafana/pull/27616), [@hugohaggmark](https://github.com/hugohaggmark)
  - **ImageRendering**: Fix rendering panel using shared query in png, PDF reports and embedded scenarios. [#27628](https://github.com/grafana/grafana/pull/27628), [@torkelo](https://github.com/torkelo)
  - **InputControl**: Fixed using InputControl in unit tests from plugins. [#27615](https://github.com/grafana/grafana/pull/27615), [@torkelo](https://github.com/torkelo)
  - **NewsPanel**: Fixed XSS issue when rendering rss links. [#27612](https://github.com/grafana/grafana/pull/27612), [@torkelo](https://github.com/torkelo)
  - **Transforms**: Fix for issue in labels to fields transform where the new option value field name did not work properly. [#27501](https://github.com/grafana/grafana/pull/27501), [@torkelo](https://github.com/torkelo)

# 7.2.0-beta1 (2020-09-09)

### Breaking changes

- **Units**: The date time units `YYYY-MM-DD HH:mm:ss` and `MM/DD/YYYY h:mm:ss a` have been renamed to `Datetime ISO`
 and `Datetime US` respectively. This is no breaking change just a visual name change (the unit id is unchanged). The
 unit behavior is different however, it no longer hides the date part if the date is today. If you want this old
 behavior you need to change unit to `Datetime ISO (No date if today)` or  `Datetime US (No date if today)`.

### Features / Enhancements

- **API**: Return ID of the deleted resource for dashboard, datasource and folder DELETE endpoints. [#26691](https://github.com/grafana/grafana/pull/26691), [@AgnesToulet](https://github.com/AgnesToulet)
- **API**: Support paging in the admin orgs list API. [#26932](https://github.com/grafana/grafana/pull/26932), [@benjaminjb](https://github.com/benjaminjb)
- **API**: return resource ID for auth key creation, folder permissions update and user invite complete endpoints. [#27419](https://github.com/grafana/grafana/pull/27419), [@AgnesToulet](https://github.com/AgnesToulet)
- **Alerting**: Add toggle to disable alert threshold visibility in graph panel. [#25705](https://github.com/grafana/grafana/pull/25705), [@jpalpant](https://github.com/jpalpant)
- **Alerting**: Adds support for overriding 'dedup_key' via alert tags when using the Pagerduty notifier. [#27356](https://github.com/grafana/grafana/pull/27356), [@alavrovinfb](https://github.com/alavrovinfb)
- **Alerting**: Change alert rule link in alert notifications to open panel in view mode. [#27378](https://github.com/grafana/grafana/pull/27378), [@robertlestak](https://github.com/robertlestak)
- **Alerting**: Support storing sensitive notifier settings securely/encrypted. [#25114](https://github.com/grafana/grafana/pull/25114), [@mtanda](https://github.com/mtanda)
- **Annotation**: Add clean up job for old annotations. [#26156](https://github.com/grafana/grafana/pull/26156), [@bergquist](https://github.com/bergquist)
- **AzureMonitor**: select plugin route from cloudname. [#27273](https://github.com/grafana/grafana/pull/27273), [@kylebrandt](https://github.com/kylebrandt)
- **BackendSrv**: Uses credentials, deprecates withCredentials & defaults to same-origin. [#27385](https://github.com/grafana/grafana/pull/27385), [@hugohaggmark](https://github.com/hugohaggmark)
- **Chore**: Upgrade to Go 1.15.1. [#27326](https://github.com/grafana/grafana/pull/27326), [@aknuds1](https://github.com/aknuds1)
- **CloudWatch**: Update list of AmazonMQ metrics and dimensions. [#27332](https://github.com/grafana/grafana/pull/27332), [@szymonpk](https://github.com/szymonpk)
- **Cloudwatch**: Add Support for external ID in assume role. [#23685](https://github.com/grafana/grafana/pull/23685), [@gdhananjay](https://github.com/gdhananjay)
- **Cloudwatch**: Add af-south-1 region. [#26513](https://github.com/grafana/grafana/pull/26513), [@ruanbekker](https://github.com/ruanbekker)
- **Dashboard**: Add Duplicate dashboard links button to links list. [#26600](https://github.com/grafana/grafana/pull/26600), [@Hmerac](https://github.com/Hmerac)
- **Dashboard**: Adds folder name and link to the dashboard overview on the homepage. [#27214](https://github.com/grafana/grafana/pull/27214), [@michelengelen](https://github.com/michelengelen)
- **Database**: Set 0640 permissions on SQLite database file. [#26339](https://github.com/grafana/grafana/pull/26339), [@aknuds1](https://github.com/aknuds1)
- **DateFormats**: Default ISO & US formats never omit date part even if date is today (breaking change). [#27300](https://github.com/grafana/grafana/pull/27300), [@torkelo](https://github.com/torkelo)
- **Explore/Loki**: POC for toggling parsed fields in the list view. [#26178](https://github.com/grafana/grafana/pull/26178), [@fredr](https://github.com/fredr)
- **Explore**: Sort order of log results. [#26669](https://github.com/grafana/grafana/pull/26669), [@ivanahuckova](https://github.com/ivanahuckova)
- **Explore**: Transform prometheus query to elasticsearch query. [#23670](https://github.com/grafana/grafana/pull/23670), [@melchiormoulin](https://github.com/melchiormoulin)
- **Field overrides**: Overrides UI improvements. [#27073](https://github.com/grafana/grafana/pull/27073), [@dprokop](https://github.com/dprokop)
- **Heatmap**: Reduce the aggressiveness of hiding ticks/labels when panel is small. [#27016](https://github.com/grafana/grafana/pull/27016), [@lrstanley](https://github.com/lrstanley)
- **Image Store**: Add support for using signed URLs when uploading images to GCS. [#26840](https://github.com/grafana/grafana/pull/26840), [@marcosrmendezthd](https://github.com/marcosrmendezthd)
- **Image Store**: Fallback to application default credentials when no key file is specified for GCS. [#25948](https://github.com/grafana/grafana/pull/25948), [@Eraac](https://github.com/Eraac)
- **InfluxDB/Flux**: Increase series limit for Flux datasource. [#26746](https://github.com/grafana/grafana/pull/26746), [@sneddrs](https://github.com/sneddrs)
- **InfluxDB**: exclude result and table column from Flux table results. [#27081](https://github.com/grafana/grafana/pull/27081), [@ryantxu](https://github.com/ryantxu)
- **InfluxDB**: return a table rather than an error when timeseries is missing time. [#27320](https://github.com/grafana/grafana/pull/27320), [@ryantxu](https://github.com/ryantxu)
- **Instrumentation**: Adds instrumentation for outgoing datasource requests. [#27427](https://github.com/grafana/grafana/pull/27427), [@bergquist](https://github.com/bergquist)
- **Loki**: Add scopedVars support in legend formatting for repeated variables. [#27046](https://github.com/grafana/grafana/pull/27046), [@ivanahuckova](https://github.com/ivanahuckova)
- **Loki**: Re-introduce running of instant queries. [#27213](https://github.com/grafana/grafana/pull/27213), [@ivanahuckova](https://github.com/ivanahuckova)
- **Loki**: Support request cancellation properly (Uses new backendSrv.fetch Observable request API). [#27265](https://github.com/grafana/grafana/pull/27265), [@hugohaggmark](https://github.com/hugohaggmark)
- **MixedDatasource**: Shows retrieved data even if a data source fails. [#27024](https://github.com/grafana/grafana/pull/27024), [@hugohaggmark](https://github.com/hugohaggmark)
- **OAuth**: Handle DEFLATE compressed payloads in JWT for Generic OAuth. [#26969](https://github.com/grafana/grafana/pull/26969), [@billoley](https://github.com/billoley)
- **OAuth**: Increase state cookie max age. [#27258](https://github.com/grafana/grafana/pull/27258), [@bergquist](https://github.com/bergquist)
- **Orgs**: Remove org deprecation notice as we have decided to preserve multi-org support. [#26853](https://github.com/grafana/grafana/pull/26853), [@torkelo](https://github.com/torkelo)
- **PanelInspector**: Adds a Raw display mode but defaults to Formatted display mode. [#27306](https://github.com/grafana/grafana/pull/27306), [@hugohaggmark](https://github.com/hugohaggmark)
- **Postgres**: Support Unix socket for host. [#25778](https://github.com/grafana/grafana/pull/25778), [@aknuds1](https://github.com/aknuds1)
- **Prometheus**: Add scopedVars support in legend formatting for repeated variables. [#27047](https://github.com/grafana/grafana/pull/27047), [@ivanahuckova](https://github.com/ivanahuckova)
- **Prometheus**: Support request cancellation properly (Uses new backendSrv.fetch Observable request API). [#27090](https://github.com/grafana/grafana/pull/27090), [@hugohaggmark](https://github.com/hugohaggmark)
- **Prometheus**: add $__rate_interval variable. [#26937](https://github.com/grafana/grafana/pull/26937), [@zoltanbedi](https://github.com/zoltanbedi)
- **Provisioning**: Validate that datasource access field equals to direct or proxy. [#26440](https://github.com/grafana/grafana/pull/26440), [@nabokihms](https://github.com/nabokihms)
- **RangeUtils**: migrate logic from kbn to grafana/data. [#27347](https://github.com/grafana/grafana/pull/27347), [@ryantxu](https://github.com/ryantxu)
- **Table**: Adds column filtering. [#27225](https://github.com/grafana/grafana/pull/27225), [@hugohaggmark](https://github.com/hugohaggmark)
- **Table**: Support showing numbers in strings with full original value. [#27097](https://github.com/grafana/grafana/pull/27097), [@torkelo](https://github.com/torkelo)
- **TablePanel**: Add support for basic gauge as a cell display mode. [#26595](https://github.com/grafana/grafana/pull/26595), [@jutley](https://github.com/jutley)
- **Transformations**: Group by and aggregate on multiple fields. [#25498](https://github.com/grafana/grafana/pull/25498), [@Totalus](https://github.com/Totalus)
- **Transformations**: enable transformations reorder. [#27197](https://github.com/grafana/grafana/pull/27197), [@dprokop](https://github.com/dprokop)
- **Units**: Allow re-scaling nanoseconds up to days. [#26458](https://github.com/grafana/grafana/pull/26458), [@kaydelaney](https://github.com/kaydelaney)
- **grafana-cli**: Add ability to read password from stdin to reset admin password. [#26016](https://github.com/grafana/grafana/pull/26016), [@nabokihms](https://github.com/nabokihms)
- **Reporting**: add branding options. (Enterprise)
- **Reporting**: allow setting custom timerange. (Enterprise)

### Bug Fixes
- **Auth**: Fix signup workflow and UI when verify email is enabled. [#26263](https://github.com/grafana/grafana/pull/26263), [@KamalGalrani](https://github.com/KamalGalrani)
- **AzureMonitor**: Change filterDimensions property to match what is stored. [#27459](https://github.com/grafana/grafana/pull/27459), [@kylebrandt](https://github.com/kylebrandt)
- **Cloud Monitoring**: Fix missing title and text from cloud monitoring annotations. [#27187](https://github.com/grafana/grafana/pull/27187), [@atotto](https://github.com/atotto)
- **CloudWatch**: Fix error message returned from tag:GetResources. [#27205](https://github.com/grafana/grafana/pull/27205), [@kichik](https://github.com/kichik)
- **Cloudwatch**: Update AWS/MediaConnect metrics and dimensions. [#26093](https://github.com/grafana/grafana/pull/26093), [@papagian](https://github.com/papagian)
- **DashboardSettings**: Fixes auto refresh crash with space in interval. [#27438](https://github.com/grafana/grafana/pull/27438), [@hugohaggmark](https://github.com/hugohaggmark)
- **Elasticsearch**: Fix localized dates in index pattern. [#27351](https://github.com/grafana/grafana/pull/27351), [@domasx2](https://github.com/domasx2)
- **Elasticsearch**: Fix using multiple bucket script aggregations when only grouping by terms. [#24064](https://github.com/grafana/grafana/pull/24064), [@MarceloNunesAlves](https://github.com/MarceloNunesAlves)
- **Explore**: Expand template variables when redirecting from dashboard panel. [#27354](https://github.com/grafana/grafana/pull/27354), [@Elfo404](https://github.com/Elfo404)
- **FolderPicker**: Fixes not being able to create new folder. [#27092](https://github.com/grafana/grafana/pull/27092), [@hugohaggmark](https://github.com/hugohaggmark)
- **Graphite**: Show and hide query editor function popup on click. [#26923](https://github.com/grafana/grafana/pull/26923), [@ivanahuckova](https://github.com/ivanahuckova)
- **InfluxDB/Flux**: Fix for Alerts on InfluxDB Flux datasources only use the first series. [#27463](https://github.com/grafana/grafana/pull/27463), [@ryantxu](https://github.com/ryantxu)
- **Loki**: Send current time range when fetching labels and values. [#26622](https://github.com/grafana/grafana/pull/26622), [@ivanahuckova](https://github.com/ivanahuckova)
- **Prometheus**: Add backslash escaping for template variables. [#26205](https://github.com/grafana/grafana/pull/26205), [@ivanahuckova](https://github.com/ivanahuckova)
- **Prometheus**: Correctly format multi values variables in queries. [#26896](https://github.com/grafana/grafana/pull/26896), [@ivanahuckova](https://github.com/ivanahuckova)
- **Provisioning**: Add validation for missing organisations in datasource, dashboard, and notifier configurations. [#26601](https://github.com/grafana/grafana/pull/26601), [@nabokihms](https://github.com/nabokihms)
- **Rendering**: Fixed issue rendering text panel to image via image renderer plugin. [#27083](https://github.com/grafana/grafana/pull/27083), [@torkelo](https://github.com/torkelo)
- **Stats**: Use more efficient SQL and add timeouts. [#27390](https://github.com/grafana/grafana/pull/27390), [@sakjur](https://github.com/sakjur)
- **Table**: Support date unit formats on string values. [#26879](https://github.com/grafana/grafana/pull/26879), [@torkelo](https://github.com/torkelo)
- **Thresholds**: Fixed issue with thresholds in overrides not working after save and reload. [#27297](https://github.com/grafana/grafana/pull/27297), [@torkelo](https://github.com/torkelo)
- **Transformations**: Fixes outer join transformation when frames are missing field to join by. [#27453](https://github.com/grafana/grafana/pull/27453), [@hugohaggmark](https://github.com/hugohaggmark)
- **Transformations**: merge will properly handle empty frames and frames with multiple rows where values are overlapping. [#27362](https://github.com/grafana/grafana/pull/27362), [@mckn](https://github.com/mckn)
- **grafana-cli**: Fix installing of plugins missing directory entries in zip. [#26945](https://github.com/grafana/grafana/pull/26945), [@adrianlzt](https://github.com/adrianlzt)

# 7.1.5 (2020-08-25)

### Features / Enhancements
- **Stats**: Stop counting the same user multiple times. [#26777](https://github.com/grafana/grafana/pull/26777), [@sakjur](https://github.com/sakjur)

### Bug Fixes
- **Alerting**: remove LongToWide call in alerting. [#27140](https://github.com/grafana/grafana/pull/27140), [@kylebrandt](https://github.com/kylebrandt)
- **AzureMonitor**: fix panic introduced in 7.1.4 when unit was unspecified and alias was used. [#27113](https://github.com/grafana/grafana/pull/27113), [@kylebrandt](https://github.com/kylebrandt)
- **Variables**: Fixes issue with All variable not being resolved. [#27151](https://github.com/grafana/grafana/pull/27151), [@hugohaggmark](https://github.com/hugohaggmark)

# 7.1.4 (2020-08-20)

### Features / Enhancements
- **Azure App Insights Alert error - tsdb.HandleRequest() failed to convert dataframe "" to tsdb.TimeSeriesSlice**. [#26897](https://github.com/grafana/grafana/issues/26897)
- **AzureMonitor**: map more units. [#26990](https://github.com/grafana/grafana/pull/26990), [@kylebrandt](https://github.com/kylebrandt)
- **Azuremonitor**: do not set unit if literal "Unspecified". [#26839](https://github.com/grafana/grafana/pull/26839), [@kylebrandt](https://github.com/kylebrandt)
- **Dataframe/Alerting**: to tsdb.TimeSeriesSlice - accept "empty" time series. [#26903](https://github.com/grafana/grafana/pull/26903), [@kylebrandt](https://github.com/kylebrandt)
- **Field overrides**: Filter by field name using regex. [#27070](https://github.com/grafana/grafana/pull/27070), [@dprokop](https://github.com/dprokop)
- **Overrides**: expose byType matcher UI. [#27056](https://github.com/grafana/grafana/pull/27056), [@ryantxu](https://github.com/ryantxu)

### Bug Fixes
- **CloudWatch**: Add FreeStorageCapacity metric. [#26503](https://github.com/grafana/grafana/pull/26503), [@waqark3389](https://github.com/waqark3389)
- **CloudWatch**: Fix sorting of metrics results. [#26835](https://github.com/grafana/grafana/pull/26835), [@aknuds1](https://github.com/aknuds1)
- **Cloudwatch**: Add FileSystemId as a dimension key for the AWS/FSx namespace. [#26662](https://github.com/grafana/grafana/pull/26662), [@waqark3389](https://github.com/waqark3389)
- **InfluxDB**: Update Flux placeholder URL with respect to latest Go client. [#27086](https://github.com/grafana/grafana/pull/27086), [@aknuds1](https://github.com/aknuds1)
- **InfluxDB**: Upgrade Go client, use data source HTTP client. [#27012](https://github.com/grafana/grafana/pull/27012), [@aknuds1](https://github.com/aknuds1)
- **Proxy**: Fix updating refresh token in OAuth pass-thru. [#26885](https://github.com/grafana/grafana/pull/26885), [@seanlaff](https://github.com/seanlaff)
- **Templating**: Fixes so texts show in picker not the values. [#27002](https://github.com/grafana/grafana/pull/27002), [@hugohaggmark](https://github.com/hugohaggmark)

# 7.1.3 (2020-08-06)

### Bug Fixes
  - **Templating**: Templating: Fix undefined result when using raw interpolation format [#26818](https://github.com/grafana/grafana/pull/26818)

# 7.1.2 (2020-08-05)

### Features / Enhancements
  - **Explore**: Don't run queries on datasource change. [#26033](https://github.com/grafana/grafana/pull/26033), [@davkal](https://github.com/davkal)
  - **TemplateSrv**: Formatting options for ${__from} and ${__to}, unix seconds epoch, ISO 8601/RFC 3339. [#26466](https://github.com/grafana/grafana/pull/26466), [@torkelo](https://github.com/torkelo)
  - **Toolkit/Plugin**: throw an Error instead of a string. [#26618](https://github.com/grafana/grafana/pull/26618), [@leventebalogh](https://github.com/leventebalogh)

  ### Bug Fixes
  - **Dashbard**: Fix refresh interval settings to allow setting it to equal min_refresh_interval. [#26615](https://github.com/grafana/grafana/pull/26615), [@torkelo](https://github.com/torkelo)
  - **Flux**: Ensure connections to InfluxDB are closed. [#26735](https://github.com/grafana/grafana/pull/26735), [@sneddrs](https://github.com/sneddrs)
  - **Query history**: Fix search filtering if null value. [#26768](https://github.com/grafana/grafana/pull/26768), [@ivanahuckova](https://github.com/ivanahuckova)
  - **QueryOptions**: Fix not being able to change cache timeout setting. [#26614](https://github.com/grafana/grafana/pull/26614), [@torkelo](https://github.com/torkelo)
  - **StatPanel**: Fix stat panel display name not showing when explicitly set. [#26616](https://github.com/grafana/grafana/pull/26616), [@torkelo](https://github.com/torkelo)
  - **Templating**: Fixed access to system variables like __dashboard, __user & __org during dashboard load & variable queries. [#26637](https://github.com/grafana/grafana/pull/26637), [@torkelo](https://github.com/torkelo)
  - **TextPanel**: Fix content overflowing panel boundaries. [#26612](https://github.com/grafana/grafana/pull/26612), [@torkelo](https://github.com/torkelo)
  - **TimePicker**: Fix position and responsive behavior. [#26570](https://github.com/grafana/grafana/pull/26570), [@torkelo](https://github.com/torkelo)
  - **TimePicker**: Fixes app crash when changing custom range to nothing. [#26775](https://github.com/grafana/grafana/pull/26775), [@hugohaggmark](https://github.com/hugohaggmark)
  - **Units**: Remove duplicate SI prefix from mSv and µSv. [#26598](https://github.com/grafana/grafana/pull/26598), [@tofurky](https://github.com/tofurky)

# 7.1.1 (2020-07-24)

### Features / Enhancements
- **Graph**: Support setting field unit & override data source (automatic) unit. [#26529](https://github.com/grafana/grafana/pull/26529), [@ryantxu](https://github.com/ryantxu)
- **Tracing**: Add errorIconColor prop to TraceSpanData. [#26509](https://github.com/grafana/grafana/pull/26509), [@zoltanbedi](https://github.com/zoltanbedi)

### Bug Fixes
- **Branding**: Fix login app title. [#26425](https://github.com/grafana/grafana/pull/26425), [@benrubson](https://github.com/benrubson)
- **Bring back scripts evaluation in TextPanel**. [#26413](https://github.com/grafana/grafana/pull/26413), [@dprokop](https://github.com/dprokop)
- **Dashboard**: Fix empty panels after scrolling on Safari/iOS. [#26495](https://github.com/grafana/grafana/pull/26495), [@torkelo](https://github.com/torkelo)
- **Dashboard**: Fix for viewer can enter panel edit mode by modifying url (but cannot not save anything). [#26556](https://github.com/grafana/grafana/pull/26556), [@torkelo](https://github.com/torkelo)
- **Elasticsearch**: Fix displaying of bucket script input. [#26552](https://github.com/grafana/grafana/pull/26552), [@ivanahuckova](https://github.com/ivanahuckova)
- **Explore**: parse queryType from explore url. [#26349](https://github.com/grafana/grafana/pull/26349), [@zoltanbedi](https://github.com/zoltanbedi)
- **Tracing**: upstream fix for hovering on log lines. [#26426](https://github.com/grafana/grafana/pull/26426), [@zoltanbedi](https://github.com/zoltanbedi)

# 7.1.0 (2020-07-16)

### Features / Enhancements
- **Backend**: Use latest go plugin sdk (0.74.0) to sort wide frames. [#26207](https://github.com/grafana/grafana/pull/26207), [@kylebrandt](https://github.com/kylebrandt)
- **Elasticsearch**: Create Raw Doc metric to render raw JSON docs in columns in the new table panel. [#26233](https://github.com/grafana/grafana/pull/26233), [@ivanahuckova](https://github.com/ivanahuckova)
- **PluginsListPage**: More plugins button should open in new window. [#26305](https://github.com/grafana/grafana/pull/26305), [@zoltanbedi](https://github.com/zoltanbedi)

### Bug Fixes
- **AdminUsers**: Reset page to zero on query change. [#26293](https://github.com/grafana/grafana/pull/26293), [@hshoff](https://github.com/hshoff)
- **CloudWatch Logs**: Fixes grouping of results by numeric field. [#26298](https://github.com/grafana/grafana/pull/26298), [@kaydelaney](https://github.com/kaydelaney)
- **DashboardLinks**: Do not over-query search endpoint. [#26311](https://github.com/grafana/grafana/pull/26311), [@torkelo](https://github.com/torkelo)
- **Docker**: Make sure to create default plugin provisioning directory. [#26017](https://github.com/grafana/grafana/pull/26017), [@marefr](https://github.com/marefr)
- **Elastic**: Fix error "e.buckets[Symbol.iterator] is not a function" when using filter. [#26217](https://github.com/grafana/grafana/pull/26217), [@ivanahuckova](https://github.com/ivanahuckova)
- **Explore/Loki**: Escape \ in labels for show context queries. [#26116](https://github.com/grafana/grafana/pull/26116), [@ivanahuckova](https://github.com/ivanahuckova)
- **Jaeger/Zipkin**: URL-encode service names and trace ids for API calls. [#26115](https://github.com/grafana/grafana/pull/26115), [@ivanahuckova](https://github.com/ivanahuckova)
- **Prometheus**: Fix prom links in mixed mode. [#26244](https://github.com/grafana/grafana/pull/26244), [@zoltanbedi](https://github.com/zoltanbedi)
- **Provisioning**: Fix bug when provision app plugins using Enterprise edition. [#26340](https://github.com/grafana/grafana/pull/26340), [@marefr](https://github.com/marefr)
- **Sign In** Use correct url for the Sign In button. [#26239](https://github.com/grafana/grafana/pull/26239), [@dprokop](https://github.com/dprokop)

# 7.1.0-beta3 (2020-07-13)

### Features / Enhancements
- **Explore**: Unification of logs/metrics/traces user interface. [#25890](https://github.com/grafana/grafana/pull/25890), [@aocenas](https://github.com/aocenas)
- **Graph panel**: Move Stacking and null values before Hover tooltip options (#26035). [#26037](https://github.com/grafana/grafana/pull/26037), [@jsoref](https://github.com/jsoref)
- **LDAP**: Get all groups for all group base search DNs. [#25825](https://github.com/grafana/grafana/pull/25825), [@Annegies](https://github.com/Annegies)
- **Table**: JSON Cell should try to convert strings to JSON. [#26024](https://github.com/grafana/grafana/pull/26024), [@ryantxu](https://github.com/ryantxu)
- **Transform**: adding missing "table"-transform and "series to rows"-transform to Grafana v7-transforms. [#26042](https://github.com/grafana/grafana/pull/26042), [@mckn](https://github.com/mckn)

### Bug Fixes
- **AdminUsersTable**: Fix width issues. [#26019](https://github.com/grafana/grafana/pull/26019), [@tskarhed](https://github.com/tskarhed)
- **BarGauge**: Fix space bug in single series mode. [#26176](https://github.com/grafana/grafana/pull/26176), [@torkelo](https://github.com/torkelo)
- **Dashboard**: Allow removing min refresh interval from refresh options (5s or other). [#26150](https://github.com/grafana/grafana/pull/26150), [@torkelo](https://github.com/torkelo)
- **DataLinks**: Fixed interpolation of repeated variables used in Graph data links. [#26147](https://github.com/grafana/grafana/pull/26147), [@torkelo](https://github.com/torkelo)
- **Do not break dashboard settings UI when time intervals end with trailing comma**. [#26126](https://github.com/grafana/grafana/pull/26126), [@dprokop](https://github.com/dprokop)
- **Elastic**: Display correct log message based on selected log field. [#26020](https://github.com/grafana/grafana/pull/26020), [@ivanahuckova](https://github.com/ivanahuckova)
- **InfluxDB**: Fixed new group by dropdown now showing after first use. [#26031](https://github.com/grafana/grafana/pull/26031), [@torkelo](https://github.com/torkelo)
- **StatPanel**: Fixes issue with name showing for single series / field results. [#26070](https://github.com/grafana/grafana/pull/26070), [@torkelo](https://github.com/torkelo)
- **Templating**: Fix recursive loop of template variable queries when changing ad-hoc-variable. [#26191](https://github.com/grafana/grafana/pull/26191), [@torkelo](https://github.com/torkelo)

# 7.0.6 (2020-07-09)

### Bug fixes

- **Templating**: Fixed recursive queries triggered when switching dashboard settings view [#26137](https://github.com/grafana/grafana/pull/26137)
- **Templating**: Fix recursive loop of template variable queries when changing ad-hoc-variable [#26191](https://github.com/grafana/grafana/pull/26191)
- **Auth**: Add support for forcing authentication in anonymous mode and modify SignIn to use it instead of redirect [#25567](https://github.com/grafana/grafana/pull/25567)
- **Auth**: Fix POST request failures with anonymous access [#26049](https://github.com/grafana/grafana/pull/26049)

# 7.1.0-beta 2 (2020-07-02)

### Features / Enhancements
- **Loki**: Allow aliasing Loki queries in dashboard. [#25706](https://github.com/grafana/grafana/pull/25706), [@bastjan](https://github.com/bastjan)

### Bug Fixes
- **Explore**: Fix href when jumping from Explore to Add data source. [#25991](https://github.com/grafana/grafana/pull/25991), [@ivanahuckova](https://github.com/ivanahuckova)
- **Fix**: Build-in plugins failed to load in windows. [#25982](https://github.com/grafana/grafana/pull/25982), [@papagian](https://github.com/papagian)

# 7.1.0-beta 1 (2020-07-01)

### Features / Enhancements
- **Alerting**: Adds support for multiple URLs in Alertmanager notifier. [#24196](https://github.com/grafana/grafana/pull/24196), [@alistarle](https://github.com/alistarle)
- **Alerting**: updating the victorops alerter to handle the no_data alert type. [#23761](https://github.com/grafana/grafana/pull/23761), [@rrusso1982](https://github.com/rrusso1982)
- **Azure**: Application Insights metrics to Frame and support multiple query dimensions. [#25849](https://github.com/grafana/grafana/pull/25849), [@kylebrandt](https://github.com/kylebrandt)
- **Azure**: Multiple dimension support for Azure Monitor Service. [#25947](https://github.com/grafana/grafana/pull/25947), [@kylebrandt](https://github.com/kylebrandt)
- **Azure**: Split Insights into two services. [#25410](https://github.com/grafana/grafana/pull/25410), [@kylebrandt](https://github.com/kylebrandt)
- **Backend plugins**: Refactor to allow shared contract between core and external backend plugins. [#25472](https://github.com/grafana/grafana/pull/25472), [@marefr](https://github.com/marefr)
- **Branding**: Use AppTitle as document title. [#25271](https://github.com/grafana/grafana/pull/25271), [@benrubson](https://github.com/benrubson)
- **Chore**: upgrade to typescript 3.9.3. [#25154](https://github.com/grafana/grafana/pull/25154), [@ryantxu](https://github.com/ryantxu)
- **CloudWatch**: Add Route53 DNSQueries metric and dimension. [#25125](https://github.com/grafana/grafana/pull/25125), [@erkolson](https://github.com/erkolson)
- **CloudWatch**: Added AWS DataSync metrics and dimensions. [#25054](https://github.com/grafana/grafana/pull/25054), [@ilyastoli](https://github.com/ilyastoli)
- **CloudWatch**: Added AWS MediaStore metrics and dimensions. [#25492](https://github.com/grafana/grafana/pull/25492), [@ilyastoli](https://github.com/ilyastoli)
- **CloudWatch**: Added AWS RoboMaker metrics and dimensions. [#25090](https://github.com/grafana/grafana/pull/25090), [@ilyastoli](https://github.com/ilyastoli)
- **CloudWatch**: Added AWS SDKMetrics metrics and dimensions. [#25150](https://github.com/grafana/grafana/pull/25150), [@ilyastoli](https://github.com/ilyastoli)
- **CloudWatch**: Added AWS ServiceCatalog metrics and dimensions. [#25812](https://github.com/grafana/grafana/pull/25812), [@ilyastoli](https://github.com/ilyastoli)
- **CloudWatch**: Added AWS WAFV2 metrics. [#24048](https://github.com/grafana/grafana/pull/24048), [@mikkokupsu](https://github.com/mikkokupsu)
- **Dashboards**: Make path to default dashboard configurable. [#25595](https://github.com/grafana/grafana/pull/25595), [@bergquist](https://github.com/bergquist)
- **Elastic**: Internal data links. [#25942](https://github.com/grafana/grafana/pull/25942), [@aocenas](https://github.com/aocenas)
- **Elasticsearch**: Add support for template variable in date histogram min_doc_count. [#21064](https://github.com/grafana/grafana/pull/21064), [@faxm0dem](https://github.com/faxm0dem)
- **Elasticsearch**: Adds cumulative sum aggregation support. [#24820](https://github.com/grafana/grafana/pull/24820), [@retzkek](https://github.com/retzkek)
- **Elasticsearch**: Support using a variable for histogram and terms min doc count. [#25392](https://github.com/grafana/grafana/pull/25392), [@marefr](https://github.com/marefr)
- **Explore/Loki**: Show results of instant queries only in table and time series only in graph. [#25845](https://github.com/grafana/grafana/pull/25845), [@ivanahuckova](https://github.com/ivanahuckova)
- **Explore**: Remove legend formatting when switching from panel to Explore. [#25848](https://github.com/grafana/grafana/pull/25848), [@ivanahuckova](https://github.com/ivanahuckova)
- **Footer**: Add back footer to login page. [#25656](https://github.com/grafana/grafana/pull/25656), [@torkelo](https://github.com/torkelo)
- **ForgottenPassword**: Move view to login screen. [#25366](https://github.com/grafana/grafana/pull/25366), [@tskarhed](https://github.com/tskarhed)
- **Gauge**: Hide orientation option in panel options. [#25511](https://github.com/grafana/grafana/pull/25511), [@torkelo](https://github.com/torkelo)
- **Grafana-UI**: Add FileUpload. [#25835](https://github.com/grafana/grafana/pull/25835), [@Clarity-89](https://github.com/Clarity-89)
- **GraphPanel**: Make legend values clickable series toggles. [#25581](https://github.com/grafana/grafana/pull/25581), [@hshoff](https://github.com/hshoff)
- **Influx**: Support flux in the influx datasource. [#25308](https://github.com/grafana/grafana/pull/25308), [@ryantxu](https://github.com/ryantxu)
- **Migration**: Select org. [#24739](https://github.com/grafana/grafana/pull/24739), [@tskarhed](https://github.com/tskarhed)
- **Migration**: Settings forms. [#24741](https://github.com/grafana/grafana/pull/24741), [@tskarhed](https://github.com/tskarhed)
- **Panel Inspect**: use Monaco editor for json display. [#25251](https://github.com/grafana/grafana/pull/25251), [@ryantxu](https://github.com/ryantxu)
- **Panel edit**: Clicking twice on a visualization closes the VizPicker. [#25739](https://github.com/grafana/grafana/pull/25739), [@peterholmberg](https://github.com/peterholmberg)
- **PanelInspect**: Update UI for Data display options. [#25478](https://github.com/grafana/grafana/pull/25478), [@tskarhed](https://github.com/tskarhed)
- **Plugins**: move jaeger trace type to grafana data. [#25403](https://github.com/grafana/grafana/pull/25403), [@zoltanbedi](https://github.com/zoltanbedi)
- **Provisioning**: Adds support for enabling app plugins. [#25649](https://github.com/grafana/grafana/pull/25649), [@marefr](https://github.com/marefr)
- **Provisioning**: Use folders structure from the file system to create desired folders in dashboard provisioning. [#23117](https://github.com/grafana/grafana/pull/23117), [@nabokihms](https://github.com/nabokihms)
- **Query history**: Add keyboard shortcut support for commenting. [#24736](https://github.com/grafana/grafana/pull/24736), [@ivanahuckova](https://github.com/ivanahuckova)
- **Query history**: Add search for query history and starred queries. [#25747](https://github.com/grafana/grafana/pull/25747), [@ivanahuckova](https://github.com/ivanahuckova)
- **Rich history**: Updates for default settings and starred queries deletion. [#25732](https://github.com/grafana/grafana/pull/25732), [@ivanahuckova](https://github.com/ivanahuckova)
- **Search**: support URL query params. [#25541](https://github.com/grafana/grafana/pull/25541), [@Clarity-89](https://github.com/Clarity-89)
- **Stackdriver**: Deep linking from Grafana panels to the Metrics Explorer. [#25858](https://github.com/grafana/grafana/pull/25858), [@papagian](https://github.com/papagian)
- **Stackdriver**: Rename Stackdriver to Google Cloud Monitoring. [#25807](https://github.com/grafana/grafana/pull/25807), [@papagian](https://github.com/papagian)
- **StatPanel**: Option showing name instead of value and more. [#25676](https://github.com/grafana/grafana/pull/25676), [@torkelo](https://github.com/torkelo)
- **Switch**: Deprecate checked prop in favor of value. [#25862](https://github.com/grafana/grafana/pull/25862), [@tskarhed](https://github.com/tskarhed)
- **Tab**: Make active tab clickable and add hyperlink functionality. [#25546](https://github.com/grafana/grafana/pull/25546), [@tskarhed](https://github.com/tskarhed)
- **Table**: Adds adhoc filtering. [#25467](https://github.com/grafana/grafana/pull/25467), [@hugohaggmark](https://github.com/hugohaggmark)
- **Teams**: Add index for permission check. [#25736](https://github.com/grafana/grafana/pull/25736), [@sakjur](https://github.com/sakjur)
- **Template variable filters**: Hide overflowing text. [#25801](https://github.com/grafana/grafana/pull/25801), [@tskarhed](https://github.com/tskarhed)
- **Templating**: Add bult in __user {name, id, login, email} variable to templating system. [#23378](https://github.com/grafana/grafana/pull/23378), [@aidanmountford](https://github.com/aidanmountford)
- **Templating**: removes old Angular variable system and featureToggle. [#24779](https://github.com/grafana/grafana/pull/24779), [@hugohaggmark](https://github.com/hugohaggmark)
- **TextPanel**: Adds proper editor for markdown and html. [#25618](https://github.com/grafana/grafana/pull/25618), [@hugohaggmark](https://github.com/hugohaggmark)
- **TextPanel**: Removes Angular Text Panel. [#25504](https://github.com/grafana/grafana/pull/25504), [@hugohaggmark](https://github.com/hugohaggmark)
- **TextPanel**: Removes text mode. [#25589](https://github.com/grafana/grafana/pull/25589), [@hugohaggmark](https://github.com/hugohaggmark)
- **TimeZone**: unify the time zone pickers to one that can rule them all. [#24803](https://github.com/grafana/grafana/pull/24803), [@mckn](https://github.com/mckn)
- **Transform**: added merge transform that will merge multiple series/tables into one table. [#25490](https://github.com/grafana/grafana/pull/25490), [@mckn](https://github.com/mckn)
- **Units**: add base-pascals and rotational speed units. [#22879](https://github.com/grafana/grafana/pull/22879), [@sakjur](https://github.com/sakjur)
- **Units**: add new unit for duration, it is optimized for displaying days, hours, minutes and seconds. [#24175](https://github.com/grafana/grafana/pull/24175), [@pabigot](https://github.com/pabigot)
- **Variables**: enables cancel for slow query variables queries. [#24430](https://github.com/grafana/grafana/pull/24430), [@hugohaggmark](https://github.com/hugohaggmark)
- **switches default value for security settings**. [#25175](https://github.com/grafana/grafana/pull/25175), [@bergquist](https://github.com/bergquist)
- **Reporting:** add monthly schedule option. (Enterprise)

### Bug Fixes
- **DatatLinks**: Fix open in new tab state mismatch. [#25826](https://github.com/grafana/grafana/pull/25826), [@tskarhed](https://github.com/tskarhed)
- **Explore/Loki**: Fix field type in table for instant queries. [#25907](https://github.com/grafana/grafana/pull/25907), [@ivanahuckova](https://github.com/ivanahuckova)
- **Explore/Loki**: Fix scrolling of context when leaving context window. [#25838](https://github.com/grafana/grafana/pull/25838), [@ivanahuckova](https://github.com/ivanahuckova)
- **Explore/SQL data sources**: Show correctly interpolated queries. [#25110](https://github.com/grafana/grafana/pull/25110), [@ivanahuckova](https://github.com/ivanahuckova)
- **Explore/Tooltip**: Fix label value in tooltip. [#25940](https://github.com/grafana/grafana/pull/25940), [@ivanahuckova](https://github.com/ivanahuckova)
- **Explore**: Fix query editors on mobile. [#25148](https://github.com/grafana/grafana/pull/25148), [@ivanahuckova](https://github.com/ivanahuckova)
- **Explore**: adds an ability to exit log row context with ESC key. [#24205](https://github.com/grafana/grafana/pull/24205), [@Estrax](https://github.com/Estrax)
- **Fix**: Value mappings match against string values. [#25929](https://github.com/grafana/grafana/pull/25929), [@peterholmberg](https://github.com/peterholmberg)
- **GraphPanel**: Fix annotations overflowing panels. [#25606](https://github.com/grafana/grafana/pull/25606), [@hshoff](https://github.com/hshoff)
- **Instrumentation**: Fix setting Jaeger tracing address through Grafana config. [#25768](https://github.com/grafana/grafana/pull/25768), [@marefr](https://github.com/marefr)
- **Prometheus**: Fix performance issue in processing of histogram labels. [#25813](https://github.com/grafana/grafana/pull/25813), [@bsherrod](https://github.com/bsherrod)
- **Provisioning**: Makes file the default dashboard provisioner type. [#24856](https://github.com/grafana/grafana/pull/24856), [@bergquist](https://github.com/bergquist)
- **Templating**: fixes variables not being interpolated after dashboard refresh. [#25698](https://github.com/grafana/grafana/pull/25698), [@hugohaggmark](https://github.com/hugohaggmark)
- **Units**: Custom unit suffix and docs for custom units. [#25710](https://github.com/grafana/grafana/pull/25710), [@torkelo](https://github.com/torkelo)
- **ValueFormats**: Fix byte-format data rates. [#25424](https://github.com/grafana/grafana/pull/25424), [@mueslo](https://github.com/mueslo)
- **Variables**: Fixes maximum call stack bug for empty value. [#25503](https://github.com/grafana/grafana/pull/25503), [@hugohaggmark](https://github.com/hugohaggmark)

### Security fixes
- **Graph**: Fix XSS vulnerability with series overrides [#25401](https://github.com/grafana/grafana/pull/25401). Thanks to Rotem Reiss for reporting this.

# 7.0.5 (2020-06-30)

### Bug Fixes

- **Datasource**: Make sure data proxy timeout applies to HTTP client. [#25865](https://github.com/grafana/grafana/pull/25865), [@marefr](https://github.com/marefr)
- **Graphite**: Fix tag value dropdowns not showing in query editor. [#25889](https://github.com/grafana/grafana/pull/25889), [@torkelo](https://github.com/torkelo)

# 7.0.4 (2020-06-25)

### Features / Enhancements

- **Dashboard**: Redirects for old (pre 7.0) edit & view panel urls. [#25653](https://github.com/grafana/grafana/pull/25653), [@torkelo](https://github.com/torkelo)
- **Stackdriver**: Use default project name if project name isn't set on the query. [#25413](https://github.com/grafana/grafana/pull/25413), [@alexashley](https://github.com/alexashley)
- **TablePanel**: Sort numbers correctly. [#25421](https://github.com/grafana/grafana/pull/25421), [@speakyourcode](https://github.com/speakyourcode)
- **Update Bitcoin currency to use proper symbol, add mBTC and μBTC**. [#24182](https://github.com/grafana/grafana/pull/24182), [@overcookedpanda](https://github.com/overcookedpanda)
- **Variables**: Links that update variables on current dashboard does not trigger refresh / update. [#25192](https://github.com/grafana/grafana/pull/25192), [@torkelo](https://github.com/torkelo)

### Bug Fixes

- **Azure Monitor**: fixes undefined is not iterable. [#25586](https://github.com/grafana/grafana/pull/25586), [@hugohaggmark](https://github.com/hugohaggmark)
- **Datasources**: Handle URL parsing error. [#25742](https://github.com/grafana/grafana/pull/25742), [@marefr](https://github.com/marefr)
- **InfluxDB**: Fix invalid memory address or nil pointer dereference when schema is missing in URL. [#25565](https://github.com/grafana/grafana/pull/25565), [@marefr](https://github.com/marefr)
- **Security**: Use Header.Set and Header.Del for X-Grafana-User header. [#25495](https://github.com/grafana/grafana/pull/25495), [@beardhatcode](https://github.com/beardhatcode)
- **Stackdriver**: Fix creating Label Values datasource query variable. [#25633](https://github.com/grafana/grafana/pull/25633), [@papagian](https://github.com/papagian)
- **Table**: Support custom date formats via custom unit. [#25195](https://github.com/grafana/grafana/pull/25195), [@torkelo](https://github.com/torkelo)
- **Templating**: Fixes query variable with \${\_\_searchFilter} value selection not causing refresh & url update. [#25770](https://github.com/grafana/grafana/pull/25770), [@torkelo](https://github.com/torkelo)

# 7.0.3 (2020-06-03)

### Features / Enhancements

- **Stats**: include all fields. [#24829](https://github.com/grafana/grafana/pull/24829), [@ryantxu](https://github.com/ryantxu)
- **Variables**: change VariableEditorList row action Icon to IconButton. [#25217](https://github.com/grafana/grafana/pull/25217), [@hshoff](https://github.com/hshoff)

### Bug Fixes

- **Cloudwatch**: Fix dimensions of DDoSProtection. [#25317](https://github.com/grafana/grafana/pull/25317), [@papagian](https://github.com/papagian)
- **Configuration**: Fix env var override of sections containing hyphen. [#25178](https://github.com/grafana/grafana/pull/25178), [@marefr](https://github.com/marefr)
- **Dashboard**: Get panels in collapsed rows. [#25079](https://github.com/grafana/grafana/pull/25079), [@peterholmberg](https://github.com/peterholmberg)
- **Do not show alerts tab when alerting is disabled**. [#25285](https://github.com/grafana/grafana/pull/25285), [@dprokop](https://github.com/dprokop)
- **Jaeger**: fixes cascader option label duration value. [#25129](https://github.com/grafana/grafana/pull/25129), [@Estrax](https://github.com/Estrax)
- **Transformations**: Fixed Transform tab crash & no update after adding first transform. [#25152](https://github.com/grafana/grafana/pull/25152), [@torkelo](https://github.com/torkelo)

# 7.0.2 (2020-06-03)

- **Security**: Urgent security patch release. Please read more in our [blog](https://grafana.com/blog/2020/06/03/grafana-6.7.4-and-7.0.2-released-with-important-security-fix/)

# 7.0.1 (2020-05-26)

### Features / Enhancements

- **Datasource/CloudWatch**: Makes CloudWatch Logs query history more readable. [#24795](https://github.com/grafana/grafana/pull/24795), [@kaydelaney](https://github.com/kaydelaney)
- **Download CSV**: Add date and time formatting. [#24992](https://github.com/grafana/grafana/pull/24992), [@ryantxu](https://github.com/ryantxu)
- **Table**: Make last cell value visible when right aligned. [#24921](https://github.com/grafana/grafana/pull/24921), [@peterholmberg](https://github.com/peterholmberg)
- **TablePanel**: Adding sort order persistance. [#24705](https://github.com/grafana/grafana/pull/24705), [@torkelo](https://github.com/torkelo)
- **Transformations**: Display correct field name when using reduce transformation. [#25068](https://github.com/grafana/grafana/pull/25068), [@peterholmberg](https://github.com/peterholmberg)
- **Transformations**: Allow custom number input for binary operations. [#24752](https://github.com/grafana/grafana/pull/24752), [@ryantxu](https://github.com/ryantxu)

### Bug Fixes

- **Cloudwatch**: Fix AWS WAF and AWS DDoSProtection metrics. [#25071](https://github.com/grafana/grafana/pull/25071), [@papagian](https://github.com/papagian)
- **Dashboard/Links**: Fixes dashboard links by tags not working. [#24773](https://github.com/grafana/grafana/pull/24773), [@KamalGalrani](https://github.com/KamalGalrani)
- **Dashboard/Links**: Fixes open in new window for dashboard link. [#24772](https://github.com/grafana/grafana/pull/24772), [@KamalGalrani](https://github.com/KamalGalrani)
- **Dashboard/Links**: Variables are resolved and limits to 100. [#25076](https://github.com/grafana/grafana/pull/25076), [@hugohaggmark](https://github.com/hugohaggmark)
- **DataLinks**: Bring back variables interpolation in title. [#24970](https://github.com/grafana/grafana/pull/24970), [@dprokop](https://github.com/dprokop)
- **Datasource/CloudWatch**: Field suggestions no longer limited to prefix-only. [#24855](https://github.com/grafana/grafana/pull/24855), [@kaydelaney](https://github.com/kaydelaney)
- **Explore/Table**: Keep existing field types if possible. [#24944](https://github.com/grafana/grafana/pull/24944), [@kaydelaney](https://github.com/kaydelaney)
- **Explore**: Fix wrap lines toggle for results of queries with filter expression. [#24915](https://github.com/grafana/grafana/pull/24915), [@ivanahuckova](https://github.com/ivanahuckova)
- **Explore**: fix undo in query editor. [#24797](https://github.com/grafana/grafana/pull/24797), [@zoltanbedi](https://github.com/zoltanbedi)
- **Explore**: fix word break in type head info. [#25014](https://github.com/grafana/grafana/pull/25014), [@zoltanbedi](https://github.com/zoltanbedi)
- **Graph**: Legend decimals now work as expected. [#24931](https://github.com/grafana/grafana/pull/24931), [@torkelo](https://github.com/torkelo)
- **LoginPage**: Fix hover color for service buttons. [#25009](https://github.com/grafana/grafana/pull/25009), [@tskarhed](https://github.com/tskarhed)
- **LogsPanel**: Fix scrollbar. [#24850](https://github.com/grafana/grafana/pull/24850), [@ivanahuckova](https://github.com/ivanahuckova)
- **MoveDashboard**: Fix for moving dashboard caused all variables to be lost. [#25005](https://github.com/grafana/grafana/pull/25005), [@torkelo](https://github.com/torkelo)
- **Organize transformer**: Use display name in field order comparer. [#24984](https://github.com/grafana/grafana/pull/24984), [@dprokop](https://github.com/dprokop)
- **Panel**: shows correct panel menu items in view mode. [#24912](https://github.com/grafana/grafana/pull/24912), [@hugohaggmark](https://github.com/hugohaggmark)
- **PanelEditor Fix missing labels and description if there is only single option in category**. [#24905](https://github.com/grafana/grafana/pull/24905), [@dprokop](https://github.com/dprokop)
- **PanelEditor**: Overrides name matcher still show all original field names even after Field default display name is specified. [#24933](https://github.com/grafana/grafana/pull/24933), [@torkelo](https://github.com/torkelo)
- **PanelInspector**: Makes sure Data display options are visible. [#24902](https://github.com/grafana/grafana/pull/24902), [@hugohaggmark](https://github.com/hugohaggmark)
- **PanelInspector**: Hides unsupported data display options for Panel type. [#24918](https://github.com/grafana/grafana/pull/24918), [@hugohaggmark](https://github.com/hugohaggmark)
- **PanelMenu**: Make menu disappear on button press. [#25015](https://github.com/grafana/grafana/pull/25015), [@tskarhed](https://github.com/tskarhed)
- **Postgres**: Fix add button. [#25087](https://github.com/grafana/grafana/pull/25087), [@phemmer](https://github.com/phemmer)
- **Prometheus**: Fix recording rules expansion. [#24977](https://github.com/grafana/grafana/pull/24977), [@ivanahuckova](https://github.com/ivanahuckova)
- **Stackdriver**: Fix creating Service Level Objectives (SLO) datasource query variable. [#25023](https://github.com/grafana/grafana/pull/25023), [@papagian](https://github.com/papagian)

# 7.0.0 (2020-05-18)

## Breaking changes

- **Removed PhantomJS**: PhantomJS was deprecated in [Grafana v6.4](https://grafana.com/docs/grafana/latest/guides/whats-new-in-v6-4/#phantomjs-deprecation) and starting from Grafana v7.0.0, all PhantomJS support has been removed. This means that Grafana no longer ships with a built-in image renderer, and we advise you to install the [Grafana Image Renderer plugin](https://grafana.com/grafana/plugins/grafana-image-renderer).
- **Dashboard**: A global minimum dashboard refresh interval is now enforced and defaults to 5 seconds.
- **Interval calculation**: There is now a new option `Max data points` that controls the auto interval `$__interval` calculation. Interval was previously calculated by dividing the panel width by the time range. With the new max data points option it is now easy to set `$__interval` to a dynamic value that is time range agnostic. For example if you set `Max data points` to 10 Grafana will dynamically set `$__interval` by dividing the current time range by 10.
- **Datasource/Loki**: Support for [deprecated Loki endpoints](https://github.com/grafana/loki/blob/master/docs/api.md#lokis-http-api) has been removed.
- **Backend plugins**: Grafana now requires backend plugins to be signed, otherwise Grafana will not load/start them. This is an additional security measure to make sure backend plugin binaries and files haven't been tampered with. Refer to [Upgrade Grafana](https://grafana.com/docs/grafana/latest/installation/upgrading/#upgrading-to-v7-0) for more information.
- **Docker**: Our Ubuntu based images have been upgraded to Ubuntu [20.04 LTS](https://releases.ubuntu.com/20.04/).
- **@grafana/ui**: Forms migration notice, see [@grafana/ui changelog](https://github.com/grafana/grafana/blob/master/packages/grafana-ui/CHANGELOG.md)
- **@grafana/ui**: Select API change for creating custom values, see [@grafana/ui changelog](https://github.com/grafana/grafana/blob/master/packages/grafana-ui/CHANGELOG.md)

**Deprecation warnings**

- Scripted dashboards is now deprecated. The feature is not removed but will be in a future release. We hope to address the underlying requirement of dynamic dashboards in a different way. [#24059](https://github.com/grafana/grafana/issues/24059)
- The unofficial first version of backend plugins together with usage of [grafana/grafana-plugin-model](https://github.com/grafana/grafana-plugin-model) is now deprecated and support for that will be removed in a future release. Please refer to [backend plugins documentation](https://grafana.com/docs/grafana/latest/developers/plugins/backend/) for information about the new officially supported backend plugins.

## 7.0 Feature highlights

### Data transformations

Not just visualizing data from anywhere, in Grafana 7 you can transform it too. By chaining a simple set of point and click transformations users will be able join, pivot, filter, re-name and calculate to get the results they need. Perfect for operations across queries or data sources missing essential data transformations.

Data transformations will provide a common set of data operations that were previously duplicated as custom features in many panels or data sources but are now an integral part of the Grafana data processing pipeline and something all data sources and panels can take advantage of.

In Grafana 7.0 we have a shared data model for both time series and table data that we call [DataFrame](https://github.com/grafana/grafana/blob/master/docs/sources/plugins/developing/dataframe.md). A DataFrame is like a table with columns but we refer to columns as fields. A time series is simply a DataFrame with two fields (time & value).

**Transformations shipping in 7.0**

- **Reduce**: Reduce many rows / data points to a single value
- **Filter by name**: Filter fields by name or regex
- **Filter by refId**: Filter by query letter
- **Organize fields**: Reorder, rename and hide fields.
- **Labels to fields**: Transform time series with labels into a table where labels get's converted to fields and the result is joined by time
- **Join by field**: Join many result sets (series) together using for example the time field. Useful for transforming time series into a table with a shared time column and where each series get it's own column.
- **Add field from calculation**: This is a powerful transformation that allows you perform many different types of math operations and add the result as a new field. Examples:
  - Calculate the difference between two series or fields and add the result to a new field
  - Multiply one field with another another and add the result to a new field

### New panel edit experience

In Grafana 7 we have redesigned the UI for editing panels. The first visible change is that we have separated panel display settings to a right hand side pane that you can collapse or expand depending on what your focus is on. With this change we are also introducing our new unified option model & UI for defining data configuration and display options. This unified data configuration system powers a consistent UI for setting data options across visualizations as well as making all data display settings data driven and overridable.

This new option architecture and UI will make all panels have a consistent set of options and behaviors for attributes like `unit`, `min`, `max`, `thresholds`, `links`, `decimals`. Not only that but all these options will share a consistent UI for specifying override rules and is extensible for custom panel specific options.

We have yet to migrate all core panels to this new architecture so in 7.0 there will sadly be some big inconsistencies in the UI between panels. Hopefully this will be fixed soon in future releases as we update all the core panels and help the community update the community panel plugins.

### New table panel

Grafana 7.0 comes with a new table panel (and deprecates the old one). This new table panel supports horizontal scrolling and column resize. Paired with the new `Organize fields` transformation detailed above you can reorder, hide & rename columns. This new panel also supports new cell display modes, like showing a bar gauge inside a cell.

### Panel inspector

The panel inspector is a feature that every panel will support, including internal as well as external community plugins. In this new panel inspector, you can view the raw data in a table format, apply some pre-defined transformations, and download as CSV. You can find the **Inspect** setting in the panel menu. Use the keyboard shortcut `i` when hovering over a panel to get the panel inspector to appear.

### Improved time zone support

Starting in version 7.0, you can override the time zone used to display date and time values in a dashboard.

With this feature, you can specify the local time zone of the service or system that you are monitoring. This can be helpful when monitoring a system or service that operates across several time zones.

We have also extended the time zone options so you can select any of the standard [ISO 8601 time zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).

### Features / Enhancements

- **Azure Monitor**: Deep linking from Log Analytic queries to the Azure Portal. [#24417](https://github.com/grafana/grafana/pull/24417), [@daniellee](https://github.com/daniellee)
- **Backend plugins**: Log deprecation warning when using the unofficial first version of backend plugins. [#24675](https://github.com/grafana/grafana/pull/24675), [@marefr](https://github.com/marefr)
- **CloudWatch/Logs**: Add data links to CloudWatch logs for deep linking to AWS. [#24334](https://github.com/grafana/grafana/pull/24334), [@kaydelaney](https://github.com/kaydelaney)
- **CloudWatch**: Unify look of query mode select between dashboard and explore. [#24648](https://github.com/grafana/grafana/pull/24648), [@aocenas](https://github.com/aocenas)
- **Docker**: Adds tzdata package to Ubuntu image. [#24422](https://github.com/grafana/grafana/pull/24422), [@xlson](https://github.com/xlson)
- **Editor**: New line on Enter, run query on Shift+Enter. [#24654](https://github.com/grafana/grafana/pull/24654), [@davkal](https://github.com/davkal)
- **Loki**: Allow multiple derived fields with the same name. [#24437](https://github.com/grafana/grafana/pull/24437), [@aocenas](https://github.com/aocenas)
- **Orgs**: Add future deprecation notice. [#24502](https://github.com/grafana/grafana/pull/24502), [@torkelo](https://github.com/torkelo)

### Bug Fixes

- **@grafana/toolkit**: Use process.cwd() instead of PWD to get directory. [#24677](https://github.com/grafana/grafana/pull/24677), [@zoltanbedi](https://github.com/zoltanbedi)
- **Admin**: Makes long settings values line break in settings page. [#24559](https://github.com/grafana/grafana/pull/24559), [@hugohaggmark](https://github.com/hugohaggmark)
- **Azure Monitor**: Fix failure when using table join in Log Analytics queries. [#24528](https://github.com/grafana/grafana/pull/24528), [@daniellee](https://github.com/daniellee)
- **CloudWatch/Logs**: Add error message when log groups are not selected. [#24361](https://github.com/grafana/grafana/pull/24361), [@aocenas](https://github.com/aocenas)
- **CloudWatch/Logs**: Allows a user to search for log groups that aren't there initially. [#24695](https://github.com/grafana/grafana/pull/24695), [@kaydelaney](https://github.com/kaydelaney)
- **CloudWatch/Logs**: Correctly interpolate variables in logs queries. [#24619](https://github.com/grafana/grafana/pull/24619), [@kaydelaney](https://github.com/kaydelaney)
- **CloudWatch/Logs**: Fix autocomplete after by keyword. [#24644](https://github.com/grafana/grafana/pull/24644), [@aocenas](https://github.com/aocenas)
- **CloudWatch/Logs**: Fix field autocomplete suggestions inside function. [#24406](https://github.com/grafana/grafana/pull/24406), [@aocenas](https://github.com/aocenas)
- **CloudWatch/Logs**: Fix fields not being refetched when log group changed. [#24529](https://github.com/grafana/grafana/pull/24529), [@aocenas](https://github.com/aocenas)
- **CloudWatch/Logs**: Fix panic on multiple aggregations queries. [#24683](https://github.com/grafana/grafana/pull/24683), [@aocenas](https://github.com/aocenas)
- **CloudWatch/Logs**: Fix query error when results were sparse. [#24702](https://github.com/grafana/grafana/pull/24702), [@aocenas](https://github.com/aocenas)
- **CloudWatch/Logs**: Fix suggestion for already inserted field. [#24581](https://github.com/grafana/grafana/pull/24581), [@aocenas](https://github.com/aocenas)
- **CloudWatch/Logs**: Fix suggestions of fields after comma. [#24520](https://github.com/grafana/grafana/pull/24520), [@aocenas](https://github.com/aocenas)
- **CloudWatch/Logs**: Fixes various autocomplete issues. [#24583](https://github.com/grafana/grafana/pull/24583), [@kaydelaney](https://github.com/kaydelaney)
- **CloudWatch/Logs**: Handle errors that are not awserr.Error instances. [#24641](https://github.com/grafana/grafana/pull/24641), [@aknuds1](https://github.com/aknuds1)
- **CloudWatch/Logs**: Handle invalidation of log groups when switching data source. [#24703](https://github.com/grafana/grafana/pull/24703), [@kaydelaney](https://github.com/kaydelaney)
- **CloudWatch/Logs**: Make stats hint show consistently. [#24392](https://github.com/grafana/grafana/pull/24392), [@aocenas](https://github.com/aocenas)
- **CloudWatch/Logs**: Prevents hidden data frame fields from displaying in tables. [#24580](https://github.com/grafana/grafana/pull/24580), [@kaydelaney](https://github.com/kaydelaney)
- **CloudWatch/Logs**: Results of stats queries are now grouped. [#24396](https://github.com/grafana/grafana/pull/24396), [@kaydelaney](https://github.com/kaydelaney)
- **CloudWatch/Logs**: Usability improvements. [#24447](https://github.com/grafana/grafana/pull/24447), [@kaydelaney](https://github.com/kaydelaney)
- **Dashboard**: Allow editing provisioned dashboard JSON and add confirmation when JSON is copied to dashboard. [#24680](https://github.com/grafana/grafana/pull/24680), [@dprokop](https://github.com/dprokop)
- **Dashboard**: Fix for strange "dashboard not found" errors when opening links in dashboard settings. [#24416](https://github.com/grafana/grafana/pull/24416), [@torkelo](https://github.com/torkelo)
- **Dashboard**: Fix so default data source is selected when data source can't be found in panel editor. [#24526](https://github.com/grafana/grafana/pull/24526), [@mckn](https://github.com/mckn)
- **Dashboard**: Fixed issue changing a panel from transparent back to normal in panel editor. [#24483](https://github.com/grafana/grafana/pull/24483), [@torkelo](https://github.com/torkelo)
- **Dashboard**: Make header names reflect the field name when exporting to CSV file from the panel inspector. [#24624](https://github.com/grafana/grafana/pull/24624), [@peterholmberg](https://github.com/peterholmberg)
- **Dashboard**: Make sure side pane is displayed with tabs by default in panel editor. [#24636](https://github.com/grafana/grafana/pull/24636), [@dprokop](https://github.com/dprokop)
- **Data source**: Fix query/annotation help content formatting. [#24687](https://github.com/grafana/grafana/pull/24687), [@AgnesToulet](https://github.com/AgnesToulet)
- **Data source**: Fixes async mount errors. [#24579](https://github.com/grafana/grafana/pull/24579), [@Estrax](https://github.com/Estrax)
- **Data source**: Fixes saving a data source without failure when URL doesn't specify a protocol. [#24497](https://github.com/grafana/grafana/pull/24497), [@aknuds1](https://github.com/aknuds1)
- **Explore/Prometheus**: Show results of instant queries only in table. [#24508](https://github.com/grafana/grafana/pull/24508), [@ivanahuckova](https://github.com/ivanahuckova)
- **Explore**: Fix rendering of react query editors. [#24593](https://github.com/grafana/grafana/pull/24593), [@ivanahuckova](https://github.com/ivanahuckova)
- **Explore**: Fixes loading more logs in logs context view. [#24135](https://github.com/grafana/grafana/pull/24135), [@Estrax](https://github.com/Estrax)
- **Graphite**: Fix schema and dedupe strategy in rollup indicators for Metrictank queries. [#24685](https://github.com/grafana/grafana/pull/24685), [@torkelo](https://github.com/torkelo)
- **Graphite**: Makes query annotations work again. [#24556](https://github.com/grafana/grafana/pull/24556), [@hugohaggmark](https://github.com/hugohaggmark)
- **Logs**: Clicking "Load more" from context overlay doesn't expand log row. [#24299](https://github.com/grafana/grafana/pull/24299), [@kaydelaney](https://github.com/kaydelaney)
- **Logs**: Fix total bytes process calculation. [#24691](https://github.com/grafana/grafana/pull/24691), [@davkal](https://github.com/davkal)
- **Org/user/team preferences**: Fixes so UI Theme can be set back to Default. [#24628](https://github.com/grafana/grafana/pull/24628), [@AgnesToulet](https://github.com/AgnesToulet)
- **Plugins**: Fix manifest validation. [#24573](https://github.com/grafana/grafana/pull/24573), [@aknuds1](https://github.com/aknuds1)
- **Provisioning**: Use proxy as default access mode in provisioning. [#24669](https://github.com/grafana/grafana/pull/24669), [@bergquist](https://github.com/bergquist)
- **Search**: Fix select item when pressing enter and Grafana is served using a sub path. [#24634](https://github.com/grafana/grafana/pull/24634), [@tskarhed](https://github.com/tskarhed)
- **Search**: Save folder expanded state. [#24496](https://github.com/grafana/grafana/pull/24496), [@Clarity-89](https://github.com/Clarity-89)
- **Security**: Tag value sanitization fix in OpenTSDB data source. [#24539](https://github.com/grafana/grafana/pull/24539), [@rotemreiss](https://github.com/rotemreiss)
- **Table**: Do not include angular options in options when switching from angular panel. [#24684](https://github.com/grafana/grafana/pull/24684), [@torkelo](https://github.com/torkelo)
- **Table**: Fixed persisting column resize for time series fields. [#24505](https://github.com/grafana/grafana/pull/24505), [@torkelo](https://github.com/torkelo)
- **Table**: Fixes Cannot read property subRows of null. [#24578](https://github.com/grafana/grafana/pull/24578), [@hugohaggmark](https://github.com/hugohaggmark)
- **Time picker**: Fixed so you can enter a relative range in the time picker without being converted to absolute range. [#24534](https://github.com/grafana/grafana/pull/24534), [@mckn](https://github.com/mckn)
- **Transformations**: Make transform dropdowns not cropped. [#24615](https://github.com/grafana/grafana/pull/24615), [@dprokop](https://github.com/dprokop)
- **Transformations**: Sort order should be preserved as entered by user when using the reduce transformation. [#24494](https://github.com/grafana/grafana/pull/24494), [@hugohaggmark](https://github.com/hugohaggmark)
- **Units**: Adds scale symbol for currencies with suffixed symbol. [#24678](https://github.com/grafana/grafana/pull/24678), [@hugohaggmark](https://github.com/hugohaggmark)
- **Variables**: Fixes filtering options with more than 1000 entries. [#24614](https://github.com/grafana/grafana/pull/24614), [@hugohaggmark](https://github.com/hugohaggmark)
- **Variables**: Fixes so Textbox variables read value from url. [#24623](https://github.com/grafana/grafana/pull/24623), [@hugohaggmark](https://github.com/hugohaggmark)
- **Zipkin**: Fix error when span contains remoteEndpoint. [#24524](https://github.com/grafana/grafana/pull/24524), [@aocenas](https://github.com/aocenas)
- **SAML**: Switch from email to login for user login attribute mapping (Enterprise)

# 7.0.0-beta3 (2020-05-08)

### Features / Enhancements

- **Docker**: Upgrade to Alpine 3.11. [#24056](https://github.com/grafana/grafana/pull/24056), [@aknuds1](https://github.com/aknuds1)
- **Forms**: Remove Forms namespace [BREAKING]. Will cause all `Forms` imports to stop working. See migration guide in [@grafana/ui changelog](https://github.com/grafana/grafana/blob/master/packages/grafana-ui/CHANGELOG.md)[#24378](https://github.com/grafana/grafana/pull/24378), [@tskarhed](https://github.com/tskarhed)

### Bug Fixes

- **CloudWatch**: Fix error with expression only query. [#24362](https://github.com/grafana/grafana/pull/24362), [@aocenas](https://github.com/aocenas)
- **Elasticsearch**: Fix building of raw document queries resulting in error Unknown BaseAggregationBuilder error. [#24403](https://github.com/grafana/grafana/pull/24403), [@ivanahuckova](https://github.com/ivanahuckova)
- **Prometheus**: Fix for prometheus legend formats for instant time series queries. [#24407](https://github.com/grafana/grafana/pull/24407), [@torkelo](https://github.com/torkelo)

# 7.0.0-beta2 (2020-05-07)

## Breaking changes

- **Removed PhantomJS**: PhantomJS was deprecated in [Grafana v6.4](https://grafana.com/docs/grafana/latest/guides/whats-new-in-v6-4/#phantomjs-deprecation) and starting from Grafana v7.0.0, all PhantomJS support has been removed. This means that Grafana no longer ships with a built-in image renderer, and we advise you to install the [Grafana Image Renderer plugin](https://grafana.com/grafana/plugins/grafana-image-renderer).
- **Docker**: Our Ubuntu based images have been upgraded to Ubuntu [20.04 LTS](https://releases.ubuntu.com/20.04/).
- **Dashboard**: A global minimum dashboard refresh interval is now enforced and defaults to 5 seconds.
- **@grafana/ui**: Forms migration notice, see [@grafana/ui changelog](https://github.com/grafana/grafana/blob/master/packages/grafana-ui/CHANGELOG.md)
- **Interval calculation**: There is now a new option `Max data points` that controls the auto interval `$__interval` calculation. Interval was previously calculated by dividing the panel width by the time range. With the new max data points option it is now easy to set `$__interval` to a dynamic value that is time range agnostic. For example if you set `Max data points` to 10 Grafana will dynamically set `$__interval` by dividing the current time range by 10.
- **Datasource/Loki**: Support for [deprecated Loki endpoints](https://github.com/grafana/loki/blob/master/docs/api.md#lokis-http-api) has been removed.

**Deprecation warnings**

- Scripted dashboards are now deprecated. The feature is not removed but will be in a future release. We hope to address the underlying requirement of dynamic dashboards in a different way. [#24059](https://github.com/grafana/grafana/issues/24059)

### Features / Enhancements

- **CloudWatch**: Adds more examples to CloudWatch Logs cheatsheet. [#24288](https://github.com/grafana/grafana/pull/24288), [@kaydelaney](https://github.com/kaydelaney)
- **Elasticsearch**: Changes terms min_doc_count default from 1 to 0. [#24204](https://github.com/grafana/grafana/pull/24204), [@Estrax](https://github.com/Estrax)
- **Login Page**: New design. [#23892](https://github.com/grafana/grafana/pull/23892), [@torkelo](https://github.com/torkelo)
- **Logs**: Add log level Fatal. [#24185](https://github.com/grafana/grafana/pull/24185), [@davkal](https://github.com/davkal)
- **Loki**: Show loki datasource stats in panel inspector. [#24190](https://github.com/grafana/grafana/pull/24190), [@davkal](https://github.com/davkal)
- **Migration**: Dashboard links. [#23553](https://github.com/grafana/grafana/pull/23553), [@peterholmberg](https://github.com/peterholmberg)
- **Plugins**: Require signing of external back-end plugins. [#24075](https://github.com/grafana/grafana/pull/24075), [@aknuds1](https://github.com/aknuds1)
- **Prometheus**: Add off switch for metric/label name lookup. [#24034](https://github.com/grafana/grafana/pull/24034), [@s-h-a-d-o-w](https://github.com/s-h-a-d-o-w)
- **Search**: Bring back open search by clicking dashboard name. [#24151](https://github.com/grafana/grafana/pull/24151), [@torkelo](https://github.com/torkelo)
- **Tracing**: Header updates. [#24153](https://github.com/grafana/grafana/pull/24153), [@aocenas](https://github.com/aocenas)
- **Transformations**: Improve time series support. [#23978](https://github.com/grafana/grafana/pull/23978), [@ryantxu](https://github.com/ryantxu)

### Bug Fixes

- **CloudWatch logs**: Fix default region interpolation and reset log groups on region change. [#24346](https://github.com/grafana/grafana/pull/24346), [@aocenas](https://github.com/aocenas)
- **Dashboard**: Fix for folder picker menu not being visible outside modal when saving dashboard. [#24296](https://github.com/grafana/grafana/pull/24296), [@tskarhed](https://github.com/tskarhed)
- **Dashboard**: Go to explore now works even after discarding dashboard changes. [#24149](https://github.com/grafana/grafana/pull/24149), [@torkelo](https://github.com/torkelo)
- **Dashboard**: Only show cache timeout option in panel edit if enabled in data source plugin json. [#24095](https://github.com/grafana/grafana/pull/24095), [@peterholmberg](https://github.com/peterholmberg)
- **Dashboard**: Propagate unhandled errors when saving dashboard. [#24081](https://github.com/grafana/grafana/pull/24081), [@peterholmberg](https://github.com/peterholmberg)
- **Dashboard**: Variable without a current value in json model causes crash on load. [#24261](https://github.com/grafana/grafana/pull/24261), [@torkelo](https://github.com/torkelo)
- **DashboardManager**: Disable editing if there are no folder permissions. [#24237](https://github.com/grafana/grafana/pull/24237), [@tskarhed](https://github.com/tskarhed)
- **DataLinks**: Do not add empty links. [#24088](https://github.com/grafana/grafana/pull/24088), [@dprokop](https://github.com/dprokop)
- **Explore/Loki**: Removes old query syntax support for regex filter. [#24281](https://github.com/grafana/grafana/pull/24281), [@Estrax](https://github.com/Estrax)
- **Explore**: Fix showing of results of queries in table. [#24018](https://github.com/grafana/grafana/pull/24018), [@ivanahuckova](https://github.com/ivanahuckova)
- **Field options**: show field name when title option config is empty. [#24335](https://github.com/grafana/grafana/pull/24335), [@dprokop](https://github.com/dprokop)
- **Graph**: Fixed graph tooltip getting stuck / not being cleared when leaving dashboard. [#24162](https://github.com/grafana/grafana/pull/24162), [@torkelo](https://github.com/torkelo)
- **Graph**: Fixed issue with x-axis labels showing "MM/DD" after viewing dashboard with pie chart. [#24341](https://github.com/grafana/grafana/pull/24341), [@mckn](https://github.com/mckn)
- **Jaeger**: Fix how label is created in cascader. [#24164](https://github.com/grafana/grafana/pull/24164), [@aocenas](https://github.com/aocenas)
- **Loki**: Fix label matcher for log metrics queries. [#24238](https://github.com/grafana/grafana/pull/24238), [@ivanahuckova](https://github.com/ivanahuckova)
- **Panel inspect**: hides Query tab for plugins without Query ability. [#24216](https://github.com/grafana/grafana/pull/24216), [@hugohaggmark](https://github.com/hugohaggmark)
- **Prometheus**: Refresh query field metrics on data source change. [#24116](https://github.com/grafana/grafana/pull/24116), [@s-h-a-d-o-w](https://github.com/s-h-a-d-o-w)
- **Select**: Fixes so component loses focus on selecting value or pressing outside of input. [#24008](https://github.com/grafana/grafana/pull/24008), [@mckn](https://github.com/mckn)
- **Stat/Gauge/BarGauge**: Shows default cursor when missing links. [#24284](https://github.com/grafana/grafana/pull/24284), [@hugohaggmark](https://github.com/hugohaggmark)
- **Tracing**: Fix view bounds after trace change. [#23994](https://github.com/grafana/grafana/pull/23994), [@aocenas](https://github.com/aocenas)
- **Variables**: Migrates old tags format for consistency. [#24276](https://github.com/grafana/grafana/pull/24276), [@hugohaggmark](https://github.com/hugohaggmark)
- **Reporting**: Update report schedule as soon as a report is updated (Enterprise)
- **White-labeling**: Makes login title and subtitle configurable (Enterprise)

# 7.0.0-beta1 (2020-04-28)

## Breaking changes

- **Removed PhantomJS**: PhantomJS was deprecated in [Grafana v6.4](https://grafana.com/docs/grafana/latest/guides/whats-new-in-v6-4/#phantomjs-deprecation) and starting from Grafana v7.0.0, all PhantomJS support has been removed. This means that Grafana no longer ships with a built-in image renderer, and we advise you to install the [Grafana Image Renderer plugin](https://grafana.com/grafana/plugins/grafana-image-renderer).
- **Docker**: Our Ubuntu based images have been upgraded to Ubuntu [20.04 LTS](https://releases.ubuntu.com/20.04/).
- **Dashboard**: A global minimum dashboard refresh interval is now enforced and defaults to 5 seconds.
- **@grafana/ui**: Forms migration notice, see [@grafana/ui changelog](https://github.com/grafana/grafana/blob/master/packages/grafana-ui/CHANGELOG.md)
- **@grafana/ui**: Select API change for creating custom values, see [@grafana/ui changelog](https://github.com/grafana/grafana/blob/master/packages/grafana-ui/CHANGELOG.md)
- **Interval calculation**: There is now a new option `Max data points` that controls the auto interval `$__interval` calculation. Interval was previously calculated by dividing the panel width by the time range. With the new max data points option it is now easy to set `$__interval` to a dynamic value that is time range agnostic. For example if you set `Max data points` to 10 Grafana will dynamically set `$__interval` by dividing the current time range by 10.
- **Datasource/Loki**: Support for [deprecated Loki endpoints](https://github.com/grafana/loki/blob/master/docs/api.md#lokis-http-api) has been removed.

### Features / Enhancements

- **@grafana/ui**: Create Icon component and replace icons. [#23402](https://github.com/grafana/grafana/pull/23402), [@ivanahuckova](https://github.com/ivanahuckova)
- **@grafana/ui**: Create slider component. [#22275](https://github.com/grafana/grafana/pull/22275), [@ivanahuckova](https://github.com/ivanahuckova)
- **@grafana/ui**: Remove ColorPalette component. [#23592](https://github.com/grafana/grafana/pull/23592), [@ivanahuckova](https://github.com/ivanahuckova)
- **AWS IAM**: Support for AWS EKS ServiceAccount roles for CloudWatch and S3 image upload. [#21594](https://github.com/grafana/grafana/pull/21594), [@patstrom](https://github.com/patstrom)
- **Alerting**: Adds support for basic auth in Alertmanager notifier. [#23231](https://github.com/grafana/grafana/pull/23231), [@melchiormoulin](https://github.com/melchiormoulin)
- **Alerting**: Enable Alert rule tags to override PagerDuty Severity setting. [#22736](https://github.com/grafana/grafana/pull/22736), [@AndrewBurian](https://github.com/AndrewBurian)
- **Alerting**: Handle image renderer unavailable when edit notifiers. [#23711](https://github.com/grafana/grafana/pull/23711), [@marefr](https://github.com/marefr)
- **Alerting**: Upload error image when image renderer unavailable. [#23713](https://github.com/grafana/grafana/pull/23713), [@marefr](https://github.com/marefr)
- **Alerting**: support alerting on data.Frame (that can be time series). [#22812](https://github.com/grafana/grafana/pull/22812), [@kylebrandt](https://github.com/kylebrandt)
- **Azure Monitor**: Add alerting support - Port Azure log analytics to the backend. [#23839](https://github.com/grafana/grafana/pull/23839), [@daniellee](https://github.com/daniellee)
- **Backend plugins**: Support alerting in external data source plugins. [#6841](https://github.com/grafana/grafana/issues/6841)
- **Build**: Bundle plugins. [#23787](https://github.com/grafana/grafana/pull/23787), [@aknuds1](https://github.com/aknuds1)
- **Build**: Remove usage of Go vendoring. [#23796](https://github.com/grafana/grafana/pull/23796), [@kylebrandt](https://github.com/kylebrandt)
- **Build**: Upgrade to Go 1.14. [#23371](https://github.com/grafana/grafana/pull/23371), [@aknuds1](https://github.com/aknuds1)
- **CloudWatch**: Added AWS Chatbot metrics and dimensions. [#23516](https://github.com/grafana/grafana/pull/23516), [@ilyastoli](https://github.com/ilyastoli)
- **CloudWatch**: Added Cassandra namespace. [#23299](https://github.com/grafana/grafana/pull/23299), [@vikkyomkar](https://github.com/vikkyomkar)
- **CloudWatch**: Added missing Cassandra metrics. [#23467](https://github.com/grafana/grafana/pull/23467), [@ilyastoli](https://github.com/ilyastoli)
- **CloudWatch**: Adds support for Cloudwatch Logs. [#23566](https://github.com/grafana/grafana/pull/23566), [@kaydelaney](https://github.com/kaydelaney)
- **CloudWatch**: Prefer webIdentity over EC2 role. [#23452](https://github.com/grafana/grafana/pull/23452), [@dnascimento](https://github.com/dnascimento)
- **CloudWatch**: Prefer webIdentity over EC2 role also when assuming a role. [#23807](https://github.com/grafana/grafana/pull/23807), [@bruecktech](https://github.com/bruecktech)
- **Components**: IconButton. [#23510](https://github.com/grafana/grafana/pull/23510), [@torkelo](https://github.com/torkelo)
- **Dashboard**: Add failsafe for slug generation. [#23709](https://github.com/grafana/grafana/pull/23709), [@sakjur](https://github.com/sakjur)
- **Dashboard**: Enforce minimum dashboard refresh interval to 5 seconds per default. [#23929](https://github.com/grafana/grafana/pull/23929), [@marefr](https://github.com/marefr)
- **Dashboard**: Handle no renderer available in panel share dialog. [#23856](https://github.com/grafana/grafana/pull/23856), [@marefr](https://github.com/marefr)
- **Dashboard**: Support additional variable format options (singlequote, doublequote, sqlstring). [#21622](https://github.com/grafana/grafana/pull/21622), [@xiaobeiyang](https://github.com/xiaobeiyang)
- **Dashboard**: Support data links via field overrides. [#23590](https://github.com/grafana/grafana/pull/23590), [@dprokop](https://github.com/dprokop)
- **Data source**: Max data points now used in interval calculation for all data sources. [#23915](https://github.com/grafana/grafana/pull/23915), [@torkelo](https://github.com/torkelo)
- **Database**: Order results in UserSearch by username/email. [#23328](https://github.com/grafana/grafana/pull/23328), [@aknuds1](https://github.com/aknuds1)
- **Database**: Update the xorm dependency to v0.8.1. [#22376](https://github.com/grafana/grafana/pull/22376), [@novalagung](https://github.com/novalagung)
- **Docker**: Upgrade to Ubuntu 20.04 in Dockerfiles. [#23852](https://github.com/grafana/grafana/pull/23852), [@aknuds1](https://github.com/aknuds1)
- **Docs**: Adding API reference documentation support for the packages libraries. [#21931](https://github.com/grafana/grafana/pull/21931), [@mckn](https://github.com/mckn)
- **Tracing**: Add trace UI to show traces from tracing datasources and Jaeger datasource. [#23047](https://github.com/grafana/grafana/pull/23047), [@aocenas](https://github.com/aocenas)
- **Frontend**: Adding support to select preferred timezone for presentation of date and time values. [#23586](https://github.com/grafana/grafana/pull/23586), [@mckn](https://github.com/mckn)
- **Grafana Toolkit**: Adds template for backend data source. [#23864](https://github.com/grafana/grafana/pull/23864), [@bergquist](https://github.com/bergquist)
- **Graphite**: Rollup indicator and custom meta data inspector. [#22738](https://github.com/grafana/grafana/pull/22738), [@torkelo](https://github.com/torkelo)
- **HTTP API**: Allow assigning a specific organization when creating a new user. [#21775](https://github.com/grafana/grafana/pull/21775), [@Sytten](https://github.com/Sytten)
- **Image Rendering**: New setting to control render request concurrency. [#23950](https://github.com/grafana/grafana/pull/23950), [@marefr](https://github.com/marefr)
- **Image Rendering**: Remove PhantomJS support. [#23460](https://github.com/grafana/grafana/pull/23460), [@marefr](https://github.com/marefr)
- **Logs**: Derived fields link design. [#23695](https://github.com/grafana/grafana/pull/23695), [@aocenas](https://github.com/aocenas)
- **Metrics**: Add image rendering metrics. [#23827](https://github.com/grafana/grafana/pull/23827), [@alexanderzobnin](https://github.com/alexanderzobnin)
- **Metrics**: Instrument backend plugin requests. [#23346](https://github.com/grafana/grafana/pull/23346), [@bergquist](https://github.com/bergquist)
- **Migration**: Add old Input to legacy namespace. [#23286](https://github.com/grafana/grafana/pull/23286), [@tskarhed](https://github.com/tskarhed)
- **Migration**: Admin org edit page. [#23866](https://github.com/grafana/grafana/pull/23866), [@tskarhed](https://github.com/tskarhed)
- **Migration**: Alerting - notifications list. [#22548](https://github.com/grafana/grafana/pull/22548), [@tskarhed](https://github.com/tskarhed)
- **Migration**: Change password. [#23623](https://github.com/grafana/grafana/pull/23623), [@tskarhed](https://github.com/tskarhed)
- **Migration**: Create org. [#22542](https://github.com/grafana/grafana/pull/22542), [@tskarhed](https://github.com/tskarhed)
- **Migration**: Data/Panel link editor. [#23778](https://github.com/grafana/grafana/pull/23778), [@tskarhed](https://github.com/tskarhed)
- **Migration**: Final components to LegacyForms. [#23707](https://github.com/grafana/grafana/pull/23707), [@tskarhed](https://github.com/tskarhed)
- **Migration**: Layout Selector. [#23790](https://github.com/grafana/grafana/pull/23790), [@tskarhed](https://github.com/tskarhed)
- **Migration**: Migrate admin/users. [#22759](https://github.com/grafana/grafana/pull/22759), [@mckn](https://github.com/mckn)
- **Migration**: Migrates ad hoc variable type to react/redux. [#22784](https://github.com/grafana/grafana/pull/22784), [@mckn](https://github.com/mckn)
- **Migration**: Move Switch from Forms namespace. [#23386](https://github.com/grafana/grafana/pull/23386), [@tskarhed](https://github.com/tskarhed)
- **Migration**: Move last components from Forms namespace. [#23556](https://github.com/grafana/grafana/pull/23556), [@tskarhed](https://github.com/tskarhed)
- **Migration**: Move old Switch to legacy namespace. [#23357](https://github.com/grafana/grafana/pull/23357), [@tskarhed](https://github.com/tskarhed)
- **Migration**: New datasource. [#23221](https://github.com/grafana/grafana/pull/23221), [@tskarhed](https://github.com/tskarhed)
- **Migration**: Org users page. [#23372](https://github.com/grafana/grafana/pull/23372), [@tskarhed](https://github.com/tskarhed)
- **Migration**: Orgs list. [#23821](https://github.com/grafana/grafana/pull/23821), [@tskarhed](https://github.com/tskarhed)
- **Migration**: Remove Button from Forms namespace. [#23105](https://github.com/grafana/grafana/pull/23105), [@tskarhed](https://github.com/tskarhed)
- **Migration**: Teams and alert list. [#23810](https://github.com/grafana/grafana/pull/23810), [@tskarhed](https://github.com/tskarhed)
- **Migration**: TextArea from Forms namespace. [#23436](https://github.com/grafana/grafana/pull/23436), [@tskarhed](https://github.com/tskarhed)
- **Migration**: User edit. [#23110](https://github.com/grafana/grafana/pull/23110), [@tskarhed](https://github.com/tskarhed)
- **OAuth**: Adds Okta provider. [#22972](https://github.com/grafana/grafana/pull/22972), [@alexanderzobnin](https://github.com/alexanderzobnin)
- **OAuth**: Introduce new setting for configuring max age of OAuth state cookie. [#23195](https://github.com/grafana/grafana/pull/23195), [@rtrompier](https://github.com/rtrompier)
- **Plugins**: Add deprecation notice to setEditor method in PanelPlugin. [#23895](https://github.com/grafana/grafana/pull/23895), [@dprokop](https://github.com/dprokop)
- **Plugins**: Adds support for URL params in plugin routes. [#23503](https://github.com/grafana/grafana/pull/23503), [@daniellee](https://github.com/daniellee)
- **Plugins**: Fluent API for custom field config and panel options creation for PanelPlugin. [#23070](https://github.com/grafana/grafana/pull/23070), [@dprokop](https://github.com/dprokop)
- **Plugins**: Hide plugins page from viewers, and limit /api/plugins to only core plugins when called by viewer role. [#21901](https://github.com/grafana/grafana/pull/21901), [@dprokop](https://github.com/dprokop)
- **Postgres**: Add SSL support for datasource. [#21341](https://github.com/grafana/grafana/pull/21341), [@ryankurte](https://github.com/ryankurte)
- **Prometheus**: Render missing labels in legend formats as an empty string. [#22355](https://github.com/grafana/grafana/pull/22355), [@Hixon10](https://github.com/Hixon10)
- **Provisioning**: Allows specifying uid for datasource and use that in derived fields. [#23585](https://github.com/grafana/grafana/pull/23585), [@aocenas](https://github.com/aocenas)
- **Provisioning**: Validate that dashboard providers have unique names. [#22898](https://github.com/grafana/grafana/pull/22898), [@youshy](https://github.com/youshy)
- **Search**: Replace search implementation. [#23855](https://github.com/grafana/grafana/pull/23855), [@sakjur](https://github.com/sakjur)
- **Search**: migrate dashboard search to react. [#23274](https://github.com/grafana/grafana/pull/23274), [@Clarity-89](https://github.com/Clarity-89)
- **Server**: Don't include trailing slash in cookie path when hosting Grafana in a sub path. [#22265](https://github.com/grafana/grafana/pull/22265), [@consideRatio](https://github.com/consideRatio)
- **Stackdriver**: Support for SLO queries. [#22917](https://github.com/grafana/grafana/pull/22917), [@sunker](https://github.com/sunker)
- **Table**: Add support for organizing fields/columns. [#23135](https://github.com/grafana/grafana/pull/23135), [@mckn](https://github.com/mckn)
- **Table**: Improvements to column resizing, style and alignment. [#23663](https://github.com/grafana/grafana/pull/23663), [@torkelo](https://github.com/torkelo)
- **Table**: upgrades react-table to 7.0.0 and typings. [#23247](https://github.com/grafana/grafana/pull/23247), [@hugohaggmark](https://github.com/hugohaggmark)
- **Table**: Handle column overflow and horizontal scrolling in table panel. [#4157](https://github.com/grafana/grafana/issues/4157)
- **Tracing**: Dark theme styling for TraceView. [#23406](https://github.com/grafana/grafana/pull/23406), [@aocenas](https://github.com/aocenas)
- **Tracing**: Zipkin datasource. [#23829](https://github.com/grafana/grafana/pull/23829), [@aocenas](https://github.com/aocenas)
- **Transformations**: Adds labels as fields transformer. [#23703](https://github.com/grafana/grafana/pull/23703), [@hugohaggmark](https://github.com/hugohaggmark)
- **Transformations**: Improve UI and add some love to filter by name. [#23751](https://github.com/grafana/grafana/pull/23751), [@dprokop](https://github.com/dprokop)
- **Transformations**: calculate a new field based on the row values. [#23675](https://github.com/grafana/grafana/pull/23675), [@ryantxu](https://github.com/ryantxu)
- **Units**: add (IEC) and (Metric) to bits and bytes. [#23175](https://github.com/grafana/grafana/pull/23175), [@flopp999](https://github.com/flopp999)
- **Usagestats**: Add usage stats about what type of data source is used in alerting. [#23125](https://github.com/grafana/grafana/pull/23125), [@bergquist](https://github.com/bergquist)
- **delete old dashboard versions in multiple batches**. [#23348](https://github.com/grafana/grafana/pull/23348), [@DanCech](https://github.com/DanCech)
- **grafana/data**: PanelTypeChangedHandler API update to use PanelModel instead of panel options object [BREAKING]. [#22754](https://github.com/grafana/grafana/pull/22754), [@dprokop](https://github.com/dprokop)
- **grafana/ui**: Add basic horizontal and vertical layout components. [#22303](https://github.com/grafana/grafana/pull/22303), [@dprokop](https://github.com/dprokop)
- **Auth** SAML Role and Team Sync (Enterprise)
- **Presence Indicators**: Display the avatars of active users on dashboards (Enterprise)
- **Reporting**: Makes it possible to disable the scheduler (Enterprise)
- **Dashboard**: Dashboard usage view (Enterprise)
- **Reporting** Makes it possible to trigger report emails without scheduler (Enterprise)
- **Search**: Sorting based on dashboard views and errors (Enterprise)
- **Reporting**: Improved landscape mode and panel image quality (Enterprise)
- **Reporting**: Adds config setting for image_scale_factor of panel images (Enterprise)

### Bug Fixes

- **@grafana/ui**: Fix time range when only partial datetime is provided. [#23122](https://github.com/grafana/grafana/pull/23122), [@ivanahuckova](https://github.com/ivanahuckova)
- **Alerting**: Only include image in notifier when enabled. [#23194](https://github.com/grafana/grafana/pull/23194), [@marefr](https://github.com/marefr)
- **Alerting**: Basic auth should not be required in the Alertmanager notifier. [#23691](https://github.com/grafana/grafana/pull/23691), [@bergquist](https://github.com/bergquist)
- **Alerting**: Translate notification IDs to UIDs when extracting alert rules. [#19882](https://github.com/grafana/grafana/pull/19882), [@aSapien](https://github.com/aSapien)
- **Azure Monitor**: Fix for application insights Azure China plugin route. [#23877](https://github.com/grafana/grafana/pull/23877), [@daniellee](https://github.com/daniellee)
- **CloudWatch**: Add ServerlessDatabaseCapacity to AWS/RDS metrics. [#23635](https://github.com/grafana/grafana/pull/23635), [@jackstevenson](https://github.com/jackstevenson)
- **Dashboard**: Fix global variable "\_\_org.id". [#23362](https://github.com/grafana/grafana/pull/23362), [@vikkyomkar](https://github.com/vikkyomkar)
- **Dashboard**: Handle min refresh interval when importing dashboard. [#23959](https://github.com/grafana/grafana/pull/23959), [@marefr](https://github.com/marefr)
- **DataSourceProxy**: Handle URL parsing error. [#23731](https://github.com/grafana/grafana/pull/23731), [@aknuds1](https://github.com/aknuds1)
- **Frontend**: Fix sorting of organization popup in alphabetical order. [#22259](https://github.com/grafana/grafana/pull/22259), [@vikkyomkar](https://github.com/vikkyomkar)
- **Image Rendering**: Make it work using serve_from_sub_path configured. [#23706](https://github.com/grafana/grafana/pull/23706), [@marefr](https://github.com/marefr)
- **Image rendering**: Fix missing icon on plugins list. [#23958](https://github.com/grafana/grafana/pull/23958), [@marefr](https://github.com/marefr)
- **Logs**: Fix error when non-string log level supplied. [#23654](https://github.com/grafana/grafana/pull/23654), [@ivanahuckova](https://github.com/ivanahuckova)
- **Rich history**: Fix create url and run query for various datasources. [#23627](https://github.com/grafana/grafana/pull/23627), [@ivanahuckova](https://github.com/ivanahuckova)
- **Security**: Fix XSS vulnerability in table panel. [#23816](https://github.com/grafana/grafana/pull/23816), [@torkelo](https://github.com/torkelo)

<!-- 6.7.5 START -->

# 6.7.5 (2020-12-17)

### Security

- **SAML**: Fixes encoding/xml SAML vulnerability in Grafana Enterprise [#29875](https://github.com/grafana/grafana/issues/29875), [@bergquist](https://github.com/bergquist)

<!-- 6.7.5 END -->



# 6.7.4 (2020-06-03)

- **Security**: Urgent security patch release. Please read more in our [blog](https://grafana.com/blog/2020/06/03/grafana-6.7.4-and-7.0.2-released-with-important-security-fix/)

# 6.7.3 (2020-04-23)

### Bug Fixes

- **Admin**: Fix Synced via LDAP message for non-LDAP external users. [#23477](https://github.com/grafana/grafana/pull/23477), [@alexanderzobnin](https://github.com/alexanderzobnin)
- **Alerting**: Fixes notifications for alerts with empty message in Google Hangouts notifier. [#23559](https://github.com/grafana/grafana/pull/23559), [@hugohaggmark](https://github.com/hugohaggmark)
- **AuthProxy**: Fixes bug where long username could not be cached.. [#22926](https://github.com/grafana/grafana/pull/22926), [@jcmcken](https://github.com/jcmcken)
- **Dashboard**: Fix saving dashboard when editing raw dashboard JSON model. [#23314](https://github.com/grafana/grafana/pull/23314), [@peterholmberg](https://github.com/peterholmberg)
- **Dashboard**: Try to parse 8 and 15 digit numbers as timestamps if parsing of time range as date fails. [#21694](https://github.com/grafana/grafana/pull/21694), [@jessetan](https://github.com/jessetan)
- **DashboardListPanel**: Fixed problem with empty panel after going into edit mode (General folder filter being automatically added) . [#23426](https://github.com/grafana/grafana/pull/23426), [@torkelo](https://github.com/torkelo)
- **Data source**: Handle datasource withCredentials option properly. [#23380](https://github.com/grafana/grafana/pull/23380), [@hvtuananh](https://github.com/hvtuananh)
- **Security**: Fix annotation popup XSS vulnerability [#23813](https://github.com/grafana/grafana/pull/23813), [@torkelo](https://github.com/torkelo). Big thanks to Juha Laaksonen for reporting this issue.
- **Security**: Fix XSS vulnerability in table panel [#23816](https://github.com/grafana/grafana/pull/23816), [@torkelo](https://github.com/torkelo). Big thanks to Rotem Reiss for reporting this issue.
- **Server**: Exit Grafana with status code 0 if no error. [#23312](https://github.com/grafana/grafana/pull/23312), [@aknuds1](https://github.com/aknuds1)
- **TablePanel**: Fix XSS issue in header column rename (backport). [#23814](https://github.com/grafana/grafana/pull/23814), [@torkelo](https://github.com/torkelo)
- **Variables**: Fixes error when setting adhoc variable values. [#23580](https://github.com/grafana/grafana/pull/23580), [@hugohaggmark](https://github.com/hugohaggmark)

# 6.7.2 (2020-04-02)

### Bug Fixes

- **BackendSrv**: Adds config to response to fix issue for external plugins that used this property . [#23032](https://github.com/grafana/grafana/pull/23032), [@torkelo](https://github.com/torkelo)
- **Dashboard**: Fixed issue with saving new dashboard after changing title . [#23104](https://github.com/grafana/grafana/pull/23104), [@dprokop](https://github.com/dprokop)
- **DataLinks**: make sure we use the correct datapoint when dataset contains null value.. [#22981](https://github.com/grafana/grafana/pull/22981), [@mckn](https://github.com/mckn)
- **Plugins**: Fixed issue for plugins that imported dateMath util . [#23069](https://github.com/grafana/grafana/pull/23069), [@mckn](https://github.com/mckn)
- **Security**: Fix for dashboard snapshot original dashboard link could contain XSS vulnerability in url. [#23254](https://github.com/grafana/grafana/pull/23254), [@torkelo](https://github.com/torkelo). Big thanks to Ahmed A. Sherif for reporting this issue.
- **Variables**: Fixes issue with too many queries being issued for nested template variables after value change. [#23220](https://github.com/grafana/grafana/pull/23220), [@torkelo](https://github.com/torkelo)
- **Plugins**: Expose promiseToDigest. [#23249](https://github.com/grafana/grafana/pull/23249), [@torkelo](https://github.com/torkelo)
- **Reporting**: Fixes issue updating a report created by someone else (Enterprise)

# 6.7.1 (2020-03-20)

### Bug Fixes

- **Azure**: Fixed dropdowns not showing current value. [#22914](https://github.com/grafana/grafana/pull/22914), [@torkelo](https://github.com/torkelo)
- **BackendSrv**: only add content-type on POST, PUT requests. [#22910](https://github.com/grafana/grafana/pull/22910), [@hugohaggmark](https://github.com/hugohaggmark)
- **Panels**: Fixed size issue with panel internal size when exiting panel edit mode. [#22912](https://github.com/grafana/grafana/pull/22912), [@torkelo](https://github.com/torkelo)
- **Reporting**: fixes migrations compatibility with mysql (Enterprise)
- **Reporting**: Reduce default concurrency limit to 4 (Enterprise)

# 6.7.0 (2020-03-19)

### Features / Enhancements

- **AzureMonitor**: support workspaces function for template variables. [#22882](https://github.com/grafana/grafana/pull/22882), [@daniellee](https://github.com/daniellee)
- **SQLStore**: Add migration for adding index on annotation.alert_id. [#22876](https://github.com/grafana/grafana/pull/22876), [@aknuds1](https://github.com/aknuds1)
- **TablePanel**: Enable new units picker . [#22833](https://github.com/grafana/grafana/pull/22833), [@dprokop](https://github.com/dprokop)

### Bug Fixes

- **AngularPanels**: Fixed inner height calculation for angular panels . [#22796](https://github.com/grafana/grafana/pull/22796), [@torkelo](https://github.com/torkelo)
- **BackendSrv**: makes sure provided headers are correctly recognized and set. [#22778](https://github.com/grafana/grafana/pull/22778), [@hugohaggmark](https://github.com/hugohaggmark)
- **Forms**: Fix input suffix position (caret-down in Select) . [#22780](https://github.com/grafana/grafana/pull/22780), [@torkelo](https://github.com/torkelo)
- **Graphite**: Fixed issue with query editor and next select metric now showing after selecting metric node . [#22856](https://github.com/grafana/grafana/pull/22856), [@torkelo](https://github.com/torkelo)
- **Rich History**: UX adjustments and fixes. [#22729](https://github.com/grafana/grafana/pull/22729), [@ivanahuckova](https://github.com/ivanahuckova)

# 6.7.0-beta1 (2020-03-12)

## Breaking changes

- **Slack**: Removed _Mention_ setting and instead introduce _Mention Users_, _Mention Groups_, and _Mention Channel_. The first two settings require user and group IDs, respectively. This change was necessary because the way of mentioning via the Slack API [changed](https://api.slack.com/changelog/2017-09-the-one-about-usernames) and mentions in Slack notifications no longer worked.
- **Alerting**: Reverts the behavior of `diff` and `percent_diff` to not always be absolute. Something we introduced by mistake in [6.1.0](https://github.com/grafana/grafana/commit/28eaac3a9c7082e8c496005c1cb66b4b70a4f82f). Alerting now support `diff()`, `diff_abs()`, `percent_diff()` and `percent_diff_abs()`. [#21338](https://github.com/grafana/grafana/pull/21338)

### Notice about changes in backendSrv for plugin authors

In our mission to migrate away from AngularJS to React we have removed all AngularJS dependencies in the core data retrieval service `backendSrv`.

Removing the AngularJS dependencies in `backendSrv` has the unfortunate side effect of AngularJS digest no longer being triggered for any request made with `backendSrv`. Because of this, external plugins using `backendSrv` directly may suffer from strange behaviour in the UI.

To remedy this issue, as a plugin author you need to trigger the digest after a direct call to `backendSrv`.

Example:

```js
backendSrv.get(‘http://your.url/api’).then(result => {
    this.result = result;
    this.$scope.$digest();
});
```

Another unfortunate outcome from this work in `backendSrv` is that the response format for `.headers()` changed from a function to an object.

To make your plugin work on 6.7.x as well as on previous versions you should add something like the following:

```typescript
let responseHeaders = response.headers;
if (!responseHeaders) {
  return null;
}

// Support pre 6.7 angular HTTP rather than fetch
if (typeof responseHeaders === 'function') {
  responseHeaders = responseHeaders();
}
```

You can test your plugin with the `master` branch version of Grafana.

### Features / Enhancements

- **API**: Include IP address when logging request error. [#21596](https://github.com/grafana/grafana/pull/21596), [@thedeveloperr](https://github.com/thedeveloperr)
- **Alerting**: Support passing tags to Pagerduty and allow notification on specific event categories . [#21335](https://github.com/grafana/grafana/pull/21335), [@johntdyer](https://github.com/johntdyer)
- **Chore**: Remove angular dependency from backendSrv. [#20999](https://github.com/grafana/grafana/pull/20999), [@kaydelaney](https://github.com/kaydelaney)
- **CloudWatch**: Surround dimension names with double quotes. [#22222](https://github.com/grafana/grafana/pull/22222), [@jeet-parekh](https://github.com/jeet-parekh)
- **CloudWatch**: updated metrics and dimensions for Athena, DocDB, and Route53Resolver. [#22604](https://github.com/grafana/grafana/pull/22604), [@jeet-parekh](https://github.com/jeet-parekh)
- **Cloudwatch**: add Usage Metrics. [#22179](https://github.com/grafana/grafana/pull/22179), [@passing](https://github.com/passing)
- **Dashboard**: Adds support for a global minimum dashboard refresh interval. [#19416](https://github.com/grafana/grafana/pull/19416), [@lfroment0](https://github.com/lfroment0)
- **DatasourceEditor**: Add UI to edit custom HTTP headers. [#17846](https://github.com/grafana/grafana/pull/17846), [@adrien-f](https://github.com/adrien-f)
- **Elastic**: To get fields, start with today's index and go backwards. [#22318](https://github.com/grafana/grafana/pull/22318), [@ChadiEM](https://github.com/ChadiEM)
- **Explore**: Rich history. [#22570](https://github.com/grafana/grafana/pull/22570), [@ivanahuckova](https://github.com/ivanahuckova)
- **Graph**: canvas's Stroke is executed after loop. [#22610](https://github.com/grafana/grafana/pull/22610), [@merturl](https://github.com/merturl)
- **Graphite**: Don't issue empty "select metric" queries. [#22699](https://github.com/grafana/grafana/pull/22699), [@papagian](https://github.com/papagian)
- **Image Rendering**: Store render key in remote cache to enable renderer to callback to public/load balancer URL when running in HA mode. [#22031](https://github.com/grafana/grafana/pull/22031), [@marefr](https://github.com/marefr)
- **LDAP**: Add fallback to search_base_dns if group_search_base_dns is undefined.. [#21263](https://github.com/grafana/grafana/pull/21263), [@bb-Ricardo](https://github.com/bb-Ricardo)
- **OAuth**: Implement Azure AD provide. [#20030](https://github.com/grafana/grafana/pull/20030), [@twendt](https://github.com/twendt)
- **Prometheus**: Implement region annotation. [#22225](https://github.com/grafana/grafana/pull/22225), [@secustor](https://github.com/secustor)
- **Prometheus**: make \$\_\_range more precise. [#21722](https://github.com/grafana/grafana/pull/21722), [@bmerry](https://github.com/bmerry)
- **Prometheus**: Do not show rate hint when increase function is used in query. [#21955](https://github.com/grafana/grafana/pull/21955), [@fredwangwang](https://github.com/fredwangwang)
- **Stackdriver**: Project selector. [#22447](https://github.com/grafana/grafana/pull/22447), [@sunker](https://github.com/sunker)
- **TablePanel**: display multi-line text. [#20210](https://github.com/grafana/grafana/pull/20210), [@michael-az](https://github.com/michael-az)
- **Templating**: Add new global built-in variables. [#21790](https://github.com/grafana/grafana/pull/21790), [@dcastanier](https://github.com/dcastanier)
- **Reporting**: add concurrent render limit to settings (Enterprise)
- **Reporting**: Add rendering timeout in settings (Enterprise)

### Bug Fixes

- **API**: Fix redirect issues. [#22285](https://github.com/grafana/grafana/pull/22285), [@papagian](https://github.com/papagian)
- **Alerting**: Don't include image_url field with Slack message if empty. [#22372](https://github.com/grafana/grafana/pull/22372), [@aknuds1](https://github.com/aknuds1)
- **Alerting**: Fixed bad background color for default notifications in alert tab . [#22660](https://github.com/grafana/grafana/pull/22660), [@krvajal](https://github.com/krvajal)
- **Annotations**: In table panel when setting transform to annotation, they will now show up right away without a manual refresh. [#22323](https://github.com/grafana/grafana/pull/22323), [@krvajal](https://github.com/krvajal)
- **Azure Monitor**: Fix app insights source to allow for new **timeFrom and **timeTo. [#21879](https://github.com/grafana/grafana/pull/21879), [@ChadNedzlek](https://github.com/ChadNedzlek)
- **BackendSrv**: Fixes POST body for form data. [#21714](https://github.com/grafana/grafana/pull/21714), [@hugohaggmark](https://github.com/hugohaggmark)
- **CloudWatch**: Credentials cache invalidation fix. [#22473](https://github.com/grafana/grafana/pull/22473), [@sunker](https://github.com/sunker)
- **CloudWatch**: Expand alias variables when query yields no result. [#22695](https://github.com/grafana/grafana/pull/22695), [@sunker](https://github.com/sunker)
- **Dashboard**: Fix bug with NaN in alerting. [#22053](https://github.com/grafana/grafana/pull/22053), [@a-melnyk](https://github.com/a-melnyk)
- **Explore**: Fix display of multiline logs in log panel and explore. [#22057](https://github.com/grafana/grafana/pull/22057), [@thomasdraebing](https://github.com/thomasdraebing)
- **Heatmap**: Legend color range is incorrect when using custom min/max. [#21748](https://github.com/grafana/grafana/pull/21748), [@sv5d](https://github.com/sv5d)
- **Security**: Fixed XSS issue in dashboard history diff . [#22680](https://github.com/grafana/grafana/pull/22680), [@torkelo](https://github.com/torkelo)
- **StatPanel**: Fixes base color is being used for null values .
  [#22646](https://github.com/grafana/grafana/pull/22646), [@torkelo](https://github.com/torkelo)

# 6.6.2 (2020-02-20)

### Features / Enhancements

- **Data proxy**: Log proxy errors using Grafana logger. [#22174](https://github.com/grafana/grafana/pull/22174), [@bergquist](https://github.com/bergquist)
- **Metrics**: Add gauge for requests currently in flight. [#22168](https://github.com/grafana/grafana/pull/22168), [@bergquist](https://github.com/bergquist)

### Bug Fixes

- **@grafana/ui**: Fix displaying of bars in React Graph. [#21968](https://github.com/grafana/grafana/pull/21968), [@ivanahuckova](https://github.com/ivanahuckova)
- **API**: Fix redirect issue when configured to use a subpath. [#21652](https://github.com/grafana/grafana/pull/21652), [@briangann](https://github.com/briangann)
- **API**: Improve recovery middleware when response already been written. [#22256](https://github.com/grafana/grafana/pull/22256), [@marefr](https://github.com/marefr)
- **Auth**: Don't rotate auth token when requests are cancelled by client. [#22106](https://github.com/grafana/grafana/pull/22106), [@bergquist](https://github.com/bergquist)
- **Docker**: Downgrade to 18.04 LTS base image. [#22313](https://github.com/grafana/grafana/pull/22313), [@aknuds1](https://github.com/aknuds1)
- **Elasticsearch**: Fix auto interval for date histogram in explore logs mode. [#21937](https://github.com/grafana/grafana/pull/21937), [@ivanahuckova](https://github.com/ivanahuckova)
- **Image Rendering**: Fix PhantomJS compatibility with es2016 node dependencies. [#21677](https://github.com/grafana/grafana/pull/21677), [@dprokop](https://github.com/dprokop)
- **Links**: Assure base url when single stat, panel and data links are built. [#21956](https://github.com/grafana/grafana/pull/21956), [@dprokop](https://github.com/dprokop)
- **Loki, Prometheus**: Fix PromQL and LogQL syntax highlighting. [#21944](https://github.com/grafana/grafana/pull/21944), [@ivanahuckova](https://github.com/ivanahuckova)
- **OAuth**: Enforce auto_assign_org_id setting when role mapping enabled using Generic OAuth. [#22268](https://github.com/grafana/grafana/pull/22268), [@aknuds1](https://github.com/aknuds1)
- **Prometheus**: Updates explore query editor to prevent it from throwing error on edit. [#21605](https://github.com/grafana/grafana/pull/21605), [@Estrax](https://github.com/Estrax)
- **Server**: Reorder cipher suites for better security. [#22101](https://github.com/grafana/grafana/pull/22101), [@tofu-rocketry](https://github.com/tofu-rocketry)
- **TimePicker**: fixing weird behavior with calendar when switching between months/years . [#22253](https://github.com/grafana/grafana/pull/22253), [@mckn](https://github.com/mckn)

# 6.6.1 (2020-02-06)

### Bug Fixes

- **Annotations**: Change indices and rewrites annotation find query to improve database query performance. [#21915](https://github.com/grafana/grafana/pull/21915), [@papagian](https://github.com/papagian), [@marefr](https://github.com/marefr), [@kylebrandt](https://github.com/kylebrandt)
- **Azure Monitor**: Fix Application Insights API key field to allow input. [#21738](https://github.com/grafana/grafana/pull/21738), [@shavonn](https://github.com/shavonn)
- **BarGauge**: Fix so we properly display the "no result" value when query returns empty result. [#21791](https://github.com/grafana/grafana/pull/21791), [@mckn](https://github.com/mckn)
- **Datasource**: Show access (Browser/Server) select on the Prometheus datasource. [#21833](https://github.com/grafana/grafana/pull/21833), [@jorgelbg](https://github.com/jorgelbg)
- **DatasourceSettings**: Fixed issue navigating away from data source settings page. [#21841](https://github.com/grafana/grafana/pull/21841), [@torkelo](https://github.com/torkelo)
- **Graph Panel**: Fix typo in thresholds form. [#21903](https://github.com/grafana/grafana/pull/21903), [@orendain](https://github.com/orendain)
- **Graphite**: Fixed issue with functions with multiple required params and no defaults caused params that could not be edited (groupByNodes groupByTags). [#21814](https://github.com/grafana/grafana/pull/21814), [@torkelo](https://github.com/torkelo)
- **Image Rendering**: Fix render of graph panel legend aligned to the right using Grafana image renderer plugin/service. [#21854](https://github.com/grafana/grafana/pull/21854), [@marefr](https://github.com/marefr)
- **Metrics**: Adds back missing summary quantiles. [#21858](https://github.com/grafana/grafana/pull/21858), [@kogent](https://github.com/kogent)
- **OpenTSDB**: Adds back missing ngInject to make it work again. [#21796](https://github.com/grafana/grafana/pull/21796), [@marefr](https://github.com/marefr)
- **Plugins**: Fix routing in app plugin pages. [#21847](https://github.com/grafana/grafana/pull/21847), [@dprokop](https://github.com/dprokop)
- **Prometheus**: Fixes default step value for annotation query. [#21934](https://github.com/grafana/grafana/pull/21934), [@hugohaggmark](https://github.com/hugohaggmark)
- **Quota**: Makes LDAP + Quota work for the first login of a new user. [#21949](https://github.com/grafana/grafana/pull/21949), [@xlson](https://github.com/xlson)
- **StatPanels**: Fixed change from singlestat to Gauge / BarGauge / Stat where default min & max (0, 100) was copied . [#21820](https://github.com/grafana/grafana/pull/21820), [@torkelo](https://github.com/torkelo)
- **TimePicker**: Should display in kiosk mode. [#21816](https://github.com/grafana/grafana/pull/21816), [@evgbibko](https://github.com/evgbibko)
- **grafana/toolkit**: Fix failing linter when there were lint issues. [#21849](https://github.com/grafana/grafana/pull/21849), [@dprokop](https://github.com/dprokop)

# 6.6.0 (2020-01-27)

### Features / Enhancements

- **CloudWatch**: Add DynamoDB Accelerator (DAX) metrics & dimensions. [#21644](https://github.com/grafana/grafana/pull/21644), [@kenju](https://github.com/kenju)
- **CloudWatch**: Auto period snap to next higher period. [#21659](https://github.com/grafana/grafana/pull/21659), [@sunker](https://github.com/sunker)
- **Template variables**: Add error for failed query variable on time range update. [#21731](https://github.com/grafana/grafana/pull/21731), [@tskarhed](https://github.com/tskarhed)
- **XSS**: Sanitize column link. [#21735](https://github.com/grafana/grafana/pull/21735), [@tskarhed](https://github.com/tskarhed)

### Bug Fixes

- **Elasticsearch**: Fix adhoc variable filtering for logs query. [#21346](https://github.com/grafana/grafana/pull/21346), [@ceh](https://github.com/ceh)
- **Explore**: Fix colors for log level when level value is capitalised. [#21646](https://github.com/grafana/grafana/pull/21646), [@ivanahuckova](https://github.com/ivanahuckova)
- **Explore**: Fix context view in logs, where some rows may have been filtered out.. [#21729](https://github.com/grafana/grafana/pull/21729), [@aocenas](https://github.com/aocenas)
- **Loki**: Fix Loki with repeated panels and interpolation for Explore. [#21685](https://github.com/grafana/grafana/pull/21685), [@ivanahuckova](https://github.com/ivanahuckova)
- **SQLStore**: Fix PostgreSQL failure to create organisation for first time. [#21648](https://github.com/grafana/grafana/pull/21648), [@papagian](https://github.com/papagian)

# 6.6.0-beta1 (2020-01-20)

## Breaking changes

- **PagerDuty**: Change `payload.custom_details` field in PagerDuty notification to be a JSON object instead of a string.
- **Security**: The `[security]` setting `cookie_samesite` configured to `none` now renders cookies with `SameSite=None` attribute compared to before where no `SameSite` attribute was added to cookies. To get the old behavior, use value `disabled` instead of `none`. Refer to [Upgrade Grafana](https://grafana.com/docs/grafana/latest/installation/upgrading/#upgrading-to-v6-6) for more information.

### Features / Enhancements

- **Graphite**: Add Metrictank dashboard to Graphite datasource
- **Admin**: Show name of user in users table view. [#18108](https://github.com/grafana/grafana/pull/18108), [@eleijonmarck](https://github.com/eleijonmarck)
- **Alerting**: Add configurable severity support for PagerDuty notifier. [#19425](https://github.com/grafana/grafana/pull/19425), [@yemble](https://github.com/yemble)
- **Alerting**: Add more information to webhook notifications. [#20420](https://github.com/grafana/grafana/pull/20420), [@michael-az](https://github.com/michael-az)
- **Alerting**: Add support for sending tags in OpsGenie notifier. [#20810](https://github.com/grafana/grafana/pull/20810), [@aSapien](https://github.com/aSapien)
- **Alerting**: Added fallbackText to Google Chat notifier. [#21464](https://github.com/grafana/grafana/pull/21464), [@alvarolmedo](https://github.com/alvarolmedo)
- **Alerting**: Adds support for sending a single email to all recipients in email notifier. [#21091](https://github.com/grafana/grafana/pull/21091), [@marefr](https://github.com/marefr)
- **Alerting**: Enable setting of OpsGenie priority via a tag. [#21298](https://github.com/grafana/grafana/pull/21298), [@zabullet](https://github.com/zabullet)
- **Alerting**: Use fully qualified status emoji in Threema notifier. [#21305](https://github.com/grafana/grafana/pull/21305), [@dbrgn](https://github.com/dbrgn)
- **Alerting**: new min_interval_seconds option to enforce a minimum evaluation frequency . [#21188](https://github.com/grafana/grafana/pull/21188), [@papagian](https://github.com/papagian)
- **CloudWatch**: Calculate period based on time range. [#21471](https://github.com/grafana/grafana/pull/21471), [@sunker](https://github.com/sunker)
- **CloudWatch**: Display partial result in graph when max DP/call limit is reached . [#21533](https://github.com/grafana/grafana/pull/21533), [@sunker](https://github.com/sunker)
- **CloudWatch**: ECS/ContainerInsights metrics support. [#21125](https://github.com/grafana/grafana/pull/21125), [@briancurt](https://github.com/briancurt)
- **CloudWatch**: Upgrade aws-sdk-go. [#20510](https://github.com/grafana/grafana/pull/20510), [@mtanda](https://github.com/mtanda)
- **DataLinks**: allow using values from other fields in the same row (cells). [#21478](https://github.com/grafana/grafana/pull/21478), [@ryantxu](https://github.com/ryantxu)
- **Editor**: Ignore closing brace when it was added by editor. [#21172](https://github.com/grafana/grafana/pull/21172), [@davkal](https://github.com/davkal)
- **Explore**: Context tooltip to copy labels and values from graph. [#21405](https://github.com/grafana/grafana/pull/21405), [@ivanahuckova](https://github.com/ivanahuckova)
- **Explore**: Log message line wrapping options for logs. [#20360](https://github.com/grafana/grafana/pull/20360), [@ivanahuckova](https://github.com/ivanahuckova)
- **Forms**: introduce RadioButtonGroup. [#20828](https://github.com/grafana/grafana/pull/20828), [@dprokop](https://github.com/dprokop)
- **Frontend**: Changes in Redux location should not strip subpath from location url. [#20161](https://github.com/grafana/grafana/pull/20161), [@wybczu](https://github.com/wybczu)
- **Graph**: Add fill gradient option to series override line fill. [#20941](https://github.com/grafana/grafana/pull/20941), [@hendrikvh](https://github.com/hendrikvh)
- **Graphite**: Add metrictank dashboard to Graphite datasource. [#20776](https://github.com/grafana/grafana/pull/20776), [@Dieterbe](https://github.com/Dieterbe)
- **Graphite**: Do not change query when opening the query editor and there is no data. [#21588](https://github.com/grafana/grafana/pull/21588), [@daniellee](https://github.com/daniellee)
- **Gravatar**: Use HTTPS by default. [#20964](https://github.com/grafana/grafana/pull/20964), [@jiajunhuang](https://github.com/jiajunhuang)
- **Loki**: Support for template variable queries. [#20697](https://github.com/grafana/grafana/pull/20697), [@ivanahuckova](https://github.com/ivanahuckova)
- **NewsPanel**: Add news as a builtin panel. [#21128](https://github.com/grafana/grafana/pull/21128), [@ryantxu](https://github.com/ryantxu)
- **OAuth**: Removes send_client_credentials_via_post setting . [#20044](https://github.com/grafana/grafana/pull/20044), [@LK4D4](https://github.com/LK4D4)
- **OpenTSDB**: Adding lookup limit to OpenTSDB datasource settings. [#20647](https://github.com/grafana/grafana/pull/20647), [@itamarst](https://github.com/itamarst)
- **Postgres/MySQL/MSSQL**: Adds support for region annotations. [#20752](https://github.com/grafana/grafana/pull/20752), [@Bercon](https://github.com/Bercon)
- **Prometheus**: Field to specify step in Explore. [#20195](https://github.com/grafana/grafana/pull/20195), [@Estrax](https://github.com/Estrax)
- **Prometheus**: User metrics metadata to inform query hints. [#21304](https://github.com/grafana/grafana/pull/21304), [@davkal](https://github.com/davkal)
- **Renderer**: Add user-agent to remote rendering service requests. [#20956](https://github.com/grafana/grafana/pull/20956), [@kfdm](https://github.com/kfdm)
- **Security**: Add disabled option for cookie samesite attribute. [#21472](https://github.com/grafana/grafana/pull/21472), [@marefr](https://github.com/marefr)
- **Stackdriver**: Support meta labels. [#21373](https://github.com/grafana/grafana/pull/21373), [@sunker](https://github.com/sunker)
- **TablePanel, GraphPanel**: Exclude hidden columns from CSV. [#19925](https://github.com/grafana/grafana/pull/19925), [@literalplus](https://github.com/literalplus)
- **Templating**: Update variables on location changed. [#21480](https://github.com/grafana/grafana/pull/21480), [@ryantxu](https://github.com/ryantxu)
- **Tracing**: Support configuring Jaeger client from environment. [#21103](https://github.com/grafana/grafana/pull/21103), [@hairyhenderson](https://github.com/hairyhenderson)
- **Units**: Add currency and energy units. [#20428](https://github.com/grafana/grafana/pull/20428), [@anirudh-ramesh](https://github.com/anirudh-ramesh)
- **Units**: Support dynamic count and currency units. [#21279](https://github.com/grafana/grafana/pull/21279), [@ryantxu](https://github.com/ryantxu)
- **grafana/toolkit**: Add option to override webpack config. [#20872](https://github.com/grafana/grafana/pull/20872), [@sebimarkgraf](https://github.com/sebimarkgraf)
- **grafana/ui**: ConfirmModal component. [#20965](https://github.com/grafana/grafana/pull/20965), [@alexanderzobnin](https://github.com/alexanderzobnin)
- **grafana/ui**: Create Tabs component. [#21328](https://github.com/grafana/grafana/pull/21328), [@peterholmberg](https://github.com/peterholmberg)
- **grafana/ui**: New table component. [#20991](https://github.com/grafana/grafana/pull/20991), [@peterholmberg](https://github.com/peterholmberg)
- **grafana/ui**: New updated time picker. [#20931](https://github.com/grafana/grafana/pull/20931), [@mckn](https://github.com/mckn)
- **White-labeling**: Makes it possible to customize the footer and login background (Enterprise)

### Bug Fixes

- **API**: Optionally list expired API keys. [#20468](https://github.com/grafana/grafana/pull/20468), [@papagian](https://github.com/papagian)
- **Alerting**: Fix custom_details to be a JSON object instead of a string in PagerDuty notifier. [#21150](https://github.com/grafana/grafana/pull/21150), [@tehGoti](https://github.com/tehGoti)
- **Alerting**: Fix image rendering and uploading timeout preventing to send alert notifications. [#21536](https://github.com/grafana/grafana/pull/21536), [@marefr](https://github.com/marefr)
- **Alerting**: Fix panic in dingding notifier . [#20378](https://github.com/grafana/grafana/pull/20378), [@csyangchen](https://github.com/csyangchen)
- **Alerting**: Fix template query validation logic. [#20721](https://github.com/grafana/grafana/pull/20721), [@okhowang](https://github.com/okhowang)
- **Alerting**: If no permission to clear history, keep the historical data. [#19007](https://github.com/grafana/grafana/pull/19007), [@lzdw](https://github.com/lzdw)
- **Alerting**: Unpausing a non-paused alert rule should not change status to Unknown. [#21375](https://github.com/grafana/grafana/pull/21375), [@vikkyomkar](https://github.com/vikkyomkar)
- **Api**: Fix returned message when enabling, disabling and deleting a non-existing user. [#21391](https://github.com/grafana/grafana/pull/21391), [@dpavlos](https://github.com/dpavlos)
- **Auth**: Rotate auth tokens at the end of requests. [#21347](https://github.com/grafana/grafana/pull/21347), [@woodsaj](https://github.com/woodsaj)
- **Azure Monitor**: Fixes error using azure monitor credentials with log analytics and non-default cloud. [#21032](https://github.com/grafana/grafana/pull/21032), [@shavonn](https://github.com/shavonn)
- **CLI**: Return error and aborts when plugin file extraction fails. [#20849](https://github.com/grafana/grafana/pull/20849), [@marefr](https://github.com/marefr)
- **CloudWatch**: Multi-valued template variable dimension alias fix. [#21541](https://github.com/grafana/grafana/pull/21541), [@sunker](https://github.com/sunker)
- **Dashboard**: Disable draggable panels on small devices. [#20629](https://github.com/grafana/grafana/pull/20629), [@peterholmberg](https://github.com/peterholmberg)
- **DataLinks**: Links with \${\_\_value.time} do not work when clicking on first result . [#20019](https://github.com/grafana/grafana/pull/20019), [@dweineha](https://github.com/dweineha)
- **Explore**: Fix showing of results in selected timezone (UTC/local). [#20812](https://github.com/grafana/grafana/pull/20812), [@ivanahuckova](https://github.com/ivanahuckova)
- **Explore**: Fix timepicker when browsing back after switching datasource. [#21454](https://github.com/grafana/grafana/pull/21454), [@ivanahuckova](https://github.com/ivanahuckova)
- **Explore**: Sync timepicker and logs after live-tailing stops. [#20979](https://github.com/grafana/grafana/pull/20979), [@ivanahuckova](https://github.com/ivanahuckova)
- **Graph**: Fix when clicking a plot on a touch device we won't display the annotation menu. [#21479](https://github.com/grafana/grafana/pull/21479), [@mckn](https://github.com/mckn)
- **OAuth**: Fix role mapping from id token. [#20300](https://github.com/grafana/grafana/pull/20300), [@seanson](https://github.com/seanson)
- **Plugins**: Add appSubUrl string to config pages. [#21414](https://github.com/grafana/grafana/pull/21414), [@Maddin-619](https://github.com/Maddin-619)
- **Provisioning**: Start provision dashboards after Grafana server have started. [#21564](https://github.com/grafana/grafana/pull/21564), [@marefr](https://github.com/marefr)
- **Render**: Use https as protocol when rendering if HTTP2 enabled. [#21600](https://github.com/grafana/grafana/pull/21600), [@marefr](https://github.com/marefr)
- **Security**: Use same cookie settings for all cookies. [#19787](https://github.com/grafana/grafana/pull/19787), [@jeffdesc](https://github.com/jeffdesc)
- **Singlestat**: Support empty value map texts. [#20952](https://github.com/grafana/grafana/pull/20952), [@hendrikvh](https://github.com/hendrikvh)
- **Units**: Custom suffix and prefix units can now be specified, for example custom currency & SI & time formats. [#20763](https://github.com/grafana/grafana/pull/20763), [@ryantxu](https://github.com/ryantxu)
- **grafana/ui**: Do not build grafana/ui in strict mode as it depends on non-strict libs. [#21319](https://github.com/grafana/grafana/pull/21319), [@dprokop](https://github.com/dprokop)

# 6.5.3 (2020-01-15)

### Features / Enhancements

- **API**: Validate redirect_to cookie has valid (Grafana) url . [#21057](https://github.com/grafana/grafana/pull/21057), [@papagian](https://github.com/papagian), Thanks Habi S Ravi for reporting this issue.

### Bug Fixes

- **AdHocFilter**: Shows SubMenu when filtering directly from table. [#21017](https://github.com/grafana/grafana/pull/21017), [@hugohaggmark](https://github.com/hugohaggmark)
- **Cloudwatch**: Fixed crash when switching from cloudwatch data source. [#21376](https://github.com/grafana/grafana/pull/21376), [@torkelo](https://github.com/torkelo)
- **DataLinks**: Sanitize data/panel link URLs. [#21140](https://github.com/grafana/grafana/pull/21140), [@dprokop](https://github.com/dprokop)
- **Elastic**: Fix multiselect variable interpolation for logs. [#20894](https://github.com/grafana/grafana/pull/20894), [@ivanahuckova](https://github.com/ivanahuckova)
- **Prometheus**: Fixes so user can change HTTP Method in config settings. [#21055](https://github.com/grafana/grafana/pull/21055), [@hugohaggmark](https://github.com/hugohaggmark)
- **Prometheus**: Prevents validation of inputs when clicking in them without changing the value. [#21059](https://github.com/grafana/grafana/pull/21059), [@hugohaggmark](https://github.com/hugohaggmark)
- **Rendering**: Fix panel PNG rendering when using sub url & serve_from_sub_path = true. [#21306](https://github.com/grafana/grafana/pull/21306), [@bgranvea](https://github.com/bgranvea)
- **Table**: Matches column names with unescaped regex characters. [#21164](https://github.com/grafana/grafana/pull/21164), [@hugohaggmark](https://github.com/hugohaggmark)

# 6.5.2 (2019-12-11)

### Bug Fixes

- **Alerting**: Improve alert threshold handle dragging behavior. [#20922](https://github.com/grafana/grafana/pull/20922), [@torkelo](https://github.com/torkelo)
- **AngularPanels**: Fixed loading spinner being stuck in some rare cases. [#20878](https://github.com/grafana/grafana/pull/20878), [@torkelo](https://github.com/torkelo)
- **CloudWatch**: Fix query editor does not render in Explore. [#20909](https://github.com/grafana/grafana/pull/20909), [@davkal](https://github.com/davkal)
- **CloudWatch**: Remove illegal character escaping in inferred expressions. [#20915](https://github.com/grafana/grafana/pull/20915), [@sunker](https://github.com/sunker)
- **CloudWatch**: Remove template variable error message. [#20864](https://github.com/grafana/grafana/pull/20864), [@sunker](https://github.com/sunker)
- **CloudWatch**: Use datasource template variable in curated dashboards. [#20917](https://github.com/grafana/grafana/pull/20917), [@sunker](https://github.com/sunker)
- **Elasticsearch**: Set default port to 9200 in ConfigEditor. [#20948](https://github.com/grafana/grafana/pull/20948), [@papagian](https://github.com/papagian)
- **Gauge/BarGauge**: Added support for value mapping of "no data"-state to text/value. [#20842](https://github.com/grafana/grafana/pull/20842), [@mckn](https://github.com/mckn)
- **Graph**: Prevent tooltip from being displayed outside of window. [#20874](https://github.com/grafana/grafana/pull/20874), [@mckn](https://github.com/mckn)
- **Graphite**: Fixes error with annotation metric queries . [#20857](https://github.com/grafana/grafana/pull/20857), [@dprokop](https://github.com/dprokop)
- **Login**: Fix fatal error when navigating from reset password page. [#20747](https://github.com/grafana/grafana/pull/20747), [@peterholmberg](https://github.com/peterholmberg)
- **MixedDatasources**: Do not filter out all mixed data sources in add mixed query dropdown. [#20990](https://github.com/grafana/grafana/pull/20990), [@torkelo](https://github.com/torkelo)
- **Prometheus**: Fix caching for default labels request. [#20718](https://github.com/grafana/grafana/pull/20718), [@aocenas](https://github.com/aocenas)
- **Prometheus**: Run default labels query only once. [#20898](https://github.com/grafana/grafana/pull/20898), [@aocenas](https://github.com/aocenas)
- **Security**: Fix invite link still accessible after completion or revocation. [#20863](https://github.com/grafana/grafana/pull/20863), [@aknuds1](https://github.com/aknuds1)
- **Server**: Fail when unable to create log directory. [#20804](https://github.com/grafana/grafana/pull/20804), [@aknuds1](https://github.com/aknuds1)
- **TeamPicker**: Increase size limit from 10 to 100. [#20882](https://github.com/grafana/grafana/pull/20882), [@hendrikvh](https://github.com/hendrikvh)
- **Units**: Remove SI prefix symbol from new milli/microSievert(/h) units. [#20650](https://github.com/grafana/grafana/pull/20650), [@zegelin](https://github.com/zegelin)

# 6.5.1 (2019-11-28)

### Bug Fixes

- **CloudWatch**: Region template query fix. [#20661](https://github.com/grafana/grafana/pull/20661), [@sunker](https://github.com/sunker)
- **CloudWatch**: Fix annotations query editor loading. [#20687](https://github.com/grafana/grafana/pull/20687), [@sunker](https://github.com/sunker)
- **Panel**: Fixes undefined services/dependencies in plugins without `/**@ngInject*/`. [#20696](https://github.com/grafana/grafana/pull/20696), [@hugohaggmark](https://github.com/hugohaggmark)
- **Server**: Fix failure to start with "bind: address already in use" when using socket as protocol. [#20679](https://github.com/grafana/grafana/pull/20679), [@aknuds1](https://github.com/aknuds1)
- **Stats**: Fix active admins/editors/viewers stats are counted more than once if the user is part of more than one org. [#20711](https://github.com/grafana/grafana/pull/20711), [@papagian](https://github.com/papagian)

# 6.5.0 (2019-11-25)

### Features / Enhancements

- **CloudWatch**: Add curated dashboards for most popular amazon services. [#20486](https://github.com/grafana/grafana/pull/20486), [@sunker](https://github.com/sunker)
- **CloudWatch**: Enable Min time interval. [#20260](https://github.com/grafana/grafana/pull/20260), [@mtanda](https://github.com/mtanda)
- **Explore**: UI improvements for log details. [#20485](https://github.com/grafana/grafana/pull/20485), [@ivanahuckova](https://github.com/ivanahuckova)
- **Server**: Improve grafana-server diagnostics configuration for profiling and tracing. [#20593](https://github.com/grafana/grafana/pull/20593), [@papagian](https://github.com/papagian)

### Bug Fixes

- **BarGauge/Gauge**: Add back missing title option field display options. [#20616](https://github.com/grafana/grafana/pull/20616), [@torkelo](https://github.com/torkelo)
- **CloudWatch**: Fix high CPU load. [#20579](https://github.com/grafana/grafana/pull/20579), [@marefr](https://github.com/marefr)
- **CloudWatch**: Fix high resolution mode without expression. [#20459](https://github.com/grafana/grafana/pull/20459), [@mtanda](https://github.com/mtanda)
- **CloudWatch**: Make sure period variable is being interpreted correctly. [#20447](https://github.com/grafana/grafana/pull/20447), [@sunker](https://github.com/sunker)
- **CloudWatch**: Remove HighResolution toggle since it's not being used. [#20440](https://github.com/grafana/grafana/pull/20440), [@sunker](https://github.com/sunker)
- **Cloudwatch**: Fix LaunchTime attribute tag bug. [#20237](https://github.com/grafana/grafana/pull/20237), [@sunker](https://github.com/sunker)
- **Data links**: Fix URL field turns read-only for graph panels. [#20381](https://github.com/grafana/grafana/pull/20381), [@dprokop](https://github.com/dprokop)
- **Explore**: Keep logQL filters when selecting labels in log row details. [#20570](https://github.com/grafana/grafana/pull/20570), [@ivanahuckova](https://github.com/ivanahuckova)
- **MySQL**: Fix TLS auth settings in config page. [#20501](https://github.com/grafana/grafana/pull/20501), [@peterholmberg](https://github.com/peterholmberg)
- **Provisioning**: Fix unmarshaling nested jsonData values. [#20399](https://github.com/grafana/grafana/pull/20399), [@aocenas](https://github.com/aocenas)
- **Server**: Should fail when server is unable to bind port. [#20409](https://github.com/grafana/grafana/pull/20409), [@aknuds1](https://github.com/aknuds1)
- **Templating**: Prevents crash when \$\_\_searchFilter is not a string. [#20526](https://github.com/grafana/grafana/pull/20526), [@hugohaggmark](https://github.com/hugohaggmark)
- **TextPanel**: Fixes issue with template variable value not properly html escaped [#20588](https://github.com/grafana/grafana/pull/20588), [@torkelo](https://github.com/torkelo)
- **TimePicker**: Should update after location change. [#20466](https://github.com/grafana/grafana/pull/20466), [@torkelo](https://github.com/torkelo)

## Breaking changes

- **CloudWatch**: Pre Grafana 6.5.0, the CloudWatch datasource used the GetMetricStatistics API for all queries that did not have an ´id´ and did not have an ´expression´ defined in the query editor. The GetMetricStatistics API has a limit of 400 transactions per second. In this release, all queries use the GetMetricData API. The GetMetricData API has a limit of 50 transactions per second and 100 metrics per transaction. For API pricing information, please refer to the CloudWatch pricing page (https://aws.amazon.com/cloudwatch/pricing/).

- **CloudWatch**: The GetMetricData API does not return metric unit, so unit auto detection in panels is no longer supported.

- **CloudWatch**: The `HighRes` switch has been removed from the query editor. Read more about this in [upgrading to 6.5](https://grafana.com/docs/installation/upgrading/#upgrading-to-v6-5).

- **CloudWatch**: In previous versions of Grafana, there was partial support for using multi-valued template variables as dimension values. When a multi-valued template variable is being used for dimension values in Grafana 6.5, a [search expression](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-search-expressions.html) will be generated. In the GetMetricData API, expressions are limited to 1024 characters, so you might reach this limit if you are using a large number of values. Read our [upgrading to 6.5](https://grafana.com/docs/installation/upgrading/#upgrading-to-v6-5) guide to see how you can use the `*` wildcard for this use case.

# 6.5.0-beta1 (2019-11-14)

### Features / Enhancements

- **API**: Add `createdAt` and `updatedAt` to api/users/lookup. [#19496](https://github.com/grafana/grafana/pull/19496), [@gotjosh](https://github.com/gotjosh)
- **API**: Add createdAt field to /api/users/:id. [#19475](https://github.com/grafana/grafana/pull/19475), [@cored](https://github.com/cored)
- **Admin**: Adds setting to disable creating initial admin user. [#19505](https://github.com/grafana/grafana/pull/19505), [@shavonn](https://github.com/shavonn)
- **Alerting**: Include alert_state in Kafka notifier payload. [#20099](https://github.com/grafana/grafana/pull/20099), [@arnaudlemaignen](https://github.com/arnaudlemaignen)
- **AuthProxy**: Can now login with auth proxy and get a login token. [#20175](https://github.com/grafana/grafana/pull/20175), [@torkelo](https://github.com/torkelo)
- **AuthProxy**: replaces setting ldap_sync_ttl with sync_ttl. [#20191](https://github.com/grafana/grafana/pull/20191), [@jongyllen](https://github.com/jongyllen)
- **AzureMonitor**: Alerting for Azure Application Insights. [#19381](https://github.com/grafana/grafana/pull/19381), [@ChadNedzlek](https://github.com/ChadNedzlek)
- **Build**: Upgrade to Go 1.13. [#19502](https://github.com/grafana/grafana/pull/19502), [@aknuds1](https://github.com/aknuds1)
- **CLI**: Reduce memory usage for plugin installation. [#19639](https://github.com/grafana/grafana/pull/19639), [@olivierlemasle](https://github.com/olivierlemasle)
- **CloudWatch**: Add ap-east-1 to hard-coded region lists. [#19523](https://github.com/grafana/grafana/pull/19523), [@Nessworthy](https://github.com/Nessworthy)
- **CloudWatch**: ContainerInsights metrics support. [#18971](https://github.com/grafana/grafana/pull/18971), [@francopeapea](https://github.com/francopeapea)
- **CloudWatch**: Support dynamic queries using dimension wildcards [#20058](https://github.com/grafana/grafana/issues/20058), [@sunker](https://github.com/sunker)
- **CloudWatch**: Stop using GetMetricStatistics and use GetMetricData for all time series requests [#20057](https://github.com/grafana/grafana/issues/20057), [@sunker](https://github.com/sunker)
- **CloudWatch**: Convert query editor from Angular to React [#19880](https://github.com/grafana/grafana/issues/19880), [@sunker](https://github.com/sunker)
- **CloudWatch**: Convert config editor from Angular to React [#19881](https://github.com/grafana/grafana/issues/19881), [@shavonn](https://github.com/shavonn)
- **CloudWatch**: Improved error handling when throttling occurs [#20348](https://github.com/grafana/grafana/issues/20348), [@sunker](https://github.com/sunker)
- **CloudWatch**: Deep linking from Grafana panel to CloudWatch console [#20279](https://github.com/grafana/grafana/issues/20279), [@sunker](https://github.com/sunker)
- **CloudWatch**: Add Grafana user agent to GMD calls [#20277](https://github.com/grafana/grafana/issues/20277), [@sunker](https://github.com/sunker)
- **Dashboard**: Allows the d-solo route to be used without slug. [#19640](https://github.com/grafana/grafana/pull/19640), [@97amarnathk](https://github.com/97amarnathk)
- **Docker**: Build and publish an additional Ubuntu based docker image. [#20196](https://github.com/grafana/grafana/pull/20196), [@aknuds1](https://github.com/aknuds1)
- **Elasticsearch**: Adds support for region annotations. [#17602](https://github.com/grafana/grafana/pull/17602), [@fangel](https://github.com/fangel)
- **Explore**: Add custom DataLinks on datasource level (like tracing links). [#20060](https://github.com/grafana/grafana/pull/20060), [@aocenas](https://github.com/aocenas)
- **Explore**: Add functionality to show/hide query row results. [#19794](https://github.com/grafana/grafana/pull/19794), [@ivanahuckova](https://github.com/ivanahuckova)
- **Explore**: Synchronise time ranges in split mode. [#19274](https://github.com/grafana/grafana/pull/19274), [@ivanahuckova](https://github.com/ivanahuckova)
- **Explore**: UI change for log row details . [#20034](https://github.com/grafana/grafana/pull/20034), [@ivanahuckova](https://github.com/ivanahuckova)
- **Frontend**: Migrate DataSource HTTP Settings to React. [#19452](https://github.com/grafana/grafana/pull/19452), [@dprokop](https://github.com/dprokop)
- **Frontend**: Show browser not supported notification. [#19904](https://github.com/grafana/grafana/pull/19904), [@peterholmberg](https://github.com/peterholmberg)
- **Graph**: Added series override option to have hidden series be persisted on save. [#20124](https://github.com/grafana/grafana/pull/20124), [@Gauravshah](https://github.com/Gauravshah)
- **Graphite**: Add Metrictank option to settings to view Metrictank request processing info in new inspect feature. [#20138](https://github.com/grafana/grafana/pull/20138), [@ryantxu](https://github.com/ryantxu)
- **LDAP**: Enable single user sync. [#19446](https://github.com/grafana/grafana/pull/19446), [@gotjosh](https://github.com/gotjosh)
- **LDAP**: Last org admin can login but wont be removed. [#20326](https://github.com/grafana/grafana/pull/20326), [@xlson](https://github.com/xlson)
- **LDAP**: Support env variable expressions in ldap.toml file. [#20173](https://github.com/grafana/grafana/pull/20173), [@torkelo](https://github.com/torkelo)
- **OAuth**: Generic OAuth role mapping support. [#17149](https://github.com/grafana/grafana/pull/17149), [@hypery2k](https://github.com/hypery2k)
- **Prometheus**: Custom query parameters string for Thanos downsampling. [#19121](https://github.com/grafana/grafana/pull/19121), [@seuf](https://github.com/seuf)
- **Provisioning**: Allow saving of provisioned dashboards. [#19820](https://github.com/grafana/grafana/pull/19820), [@jongyllen](https://github.com/jongyllen)
- **Security**: Minor XSS issue resolved by angularjs upgrade from 1.6.6 -> 1.6.9. [#19849](https://github.com/grafana/grafana/pull/19849), [@peterholmberg](https://github.com/peterholmberg)
- **TablePanel**: Prevents crash when data contains mixed data formats. [#20202](https://github.com/grafana/grafana/pull/20202), [@hugohaggmark](https://github.com/hugohaggmark)
- **Templating**: Introduces \$\_\_searchFilter to Query Variables. [#19858](https://github.com/grafana/grafana/pull/19858), [@hugohaggmark](https://github.com/hugohaggmark)
- **Templating**: Made default template variable query editor field a textarea with automatic height. [#20288](https://github.com/grafana/grafana/pull/20288), [@torkelo](https://github.com/torkelo)
- **Units**: Add milli/microSievert, milli/microSievert/h and pixels. [#20144](https://github.com/grafana/grafana/pull/20144), [@ryantxu](https://github.com/ryantxu)
- **Units**: Added mega ampere and watt-hour per kg. [#19922](https://github.com/grafana/grafana/pull/19922), [@Karan96Kaushik](https://github.com/Karan96Kaushik)
- **Enterprise**: Enterprise without a license behaves like OSS (Enterprise)

### Bug Fixes

- **API**: Added dashboardId and slug in response to dashboard import api. [#19692](https://github.com/grafana/grafana/pull/19692), [@jongyllen](https://github.com/jongyllen)
- **API**: Fix logging of dynamic listening port. [#19644](https://github.com/grafana/grafana/pull/19644), [@oleggator](https://github.com/oleggator)
- **BarGauge**: Fix so that default thresholds not keeps resetting. [#20190](https://github.com/grafana/grafana/pull/20190), [@lzdw](https://github.com/lzdw)
- **CloudWatch**: Fix incorrect casing of Redshift dimension entry for service class and stage. [#19897](https://github.com/grafana/grafana/pull/19897), [@nlsdfnbch](https://github.com/nlsdfnbch)
- **CloudWatch**: Fixing AWS Kafka dimension names. [#19986](https://github.com/grafana/grafana/pull/19986), [@skuxy](https://github.com/skuxy)
- **CloudWatch**: Metric math broken when using multi template variables [#18337](https://github.com/grafana/grafana/issues/18337), [@sunker](https://github.com/sunker)
- **CloudWatch**: Graphs with multiple multi-value dimension variables don't work [#17949](https://github.com/grafana/grafana/issues/17949), [@sunker](https://github.com/sunker)
- **CloudWatch**: Variables' values surrounded with braces in request sent to AWS [#14451](https://github.com/grafana/grafana/issues/14451), [@sunker](https://github.com/sunker)
- **CloudWatch**: Cloudwatch Query for a list of instances for which data is available in the selected time interval [#12784](https://github.com/grafana/grafana/issues/12784), [@sunker](https://github.com/sunker)
- **CloudWatch**: Dimension's positioning/order should be stored in the json dashboard [#11062](https://github.com/grafana/grafana/issues/11062), [@sunker](https://github.com/sunker)
- **CloudWatch**: Batch CloudWatch API call support in backend [#7991](https://github.com/grafana/grafana/issues/7991), [@sunker](https://github.com/sunker)
- **ColorPicker**: Fixes issue with ColorPicker disappearing too quickly . [#20289](https://github.com/grafana/grafana/pull/20289), [@dprokop](https://github.com/dprokop)
- **Datasource**: Add custom headers on alerting queries. [#19508](https://github.com/grafana/grafana/pull/19508), [@weeco](https://github.com/weeco)
- **Docker**: Add additional glibc dependencies to support certain backend plugins in alpine. [#20214](https://github.com/grafana/grafana/pull/20214), [@briangann](https://github.com/briangann)
- **Docker**: Build and use musl-based binaries in alpine images to resolve glibc incompatibility issues. [#19798](https://github.com/grafana/grafana/pull/19798), [@aknuds1](https://github.com/aknuds1)
- **Elasticsearch**: Fix template variables interpolation when redirecting to Explore. [#20314](https://github.com/grafana/grafana/pull/20314), [@ivanahuckova](https://github.com/ivanahuckova)
- **Elasticsearch**: Support rendering in logs panel. [#20229](https://github.com/grafana/grafana/pull/20229), [@davkal](https://github.com/davkal)
- **Explore**: Expand template variables when redirecting from dashboard panel. [#19582](https://github.com/grafana/grafana/pull/19582), [@ivanahuckova](https://github.com/ivanahuckova)
- **OAuth**: Make the login button display name of custom OAuth provider. [#20209](https://github.com/grafana/grafana/pull/20209), [@dprokop](https://github.com/dprokop)
- **ReactPanels**: Adds Explore menu item. [#20236](https://github.com/grafana/grafana/pull/20236), [@hugohaggmark](https://github.com/hugohaggmark)
- **Team Sync**: Fix URL encode Group IDs for external team sync. [#20280](https://github.com/grafana/grafana/pull/20280), [@gotjosh](https://github.com/gotjosh)

## Breaking changes

- **CloudWatch**: Pre Grafana 6.5.0, the CloudWatch datasource used the GetMetricStatistics API for all queries that did not have an ´id´ and did not have an ´expression´ defined in the query editor. The GetMetricStatistics API has a limit of 400 transactions per second. In this release, all queries use the GetMetricData API. The GetMetricData API has a limit of 50 transactions per second and 100 metrics per transaction. For API pricing information, please refer to the CloudWatch pricing page (https://aws.amazon.com/cloudwatch/pricing/).

- **CloudWatch**: The GetMetricData API does not return metric unit, so unit auto detection in panels is no longer supported.

- **CloudWatch**: The `HighRes` switch has been removed from the query editor. Read more about this in [upgrading to 6.5](https://grafana.com/docs/installation/upgrading/#upgrading-to-v6-5).

- **CloudWatch**: In previous versions of Grafana, there was partial support for using multi-valued template variables as dimension values. When a multi-valued template variable is being used for dimension values in Grafana 6.5, a [search expression](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-search-expressions.html) will be generated. In the GetMetricData API, expressions are limited to 1024 characters, so you might reach this limit if you are using a large number of values. Read our [upgrading to 6.5](https://grafana.com/docs/installation/upgrading/#upgrading-to-v6-5) guide to see how you can use the `*` wildcard for this use case.

# 6.4.5 (2019-11-25)

### Bug Fixes

- **CloudWatch**: Fix high CPU load [#20579](https://github.com/grafana/grafana/pull/20579)

# 6.4.4 (2019-11-06)

### Bug Fixes

- **MySQL**: Fix encoding in connection string [#20192](https://github.com/grafana/grafana/pull/20192)
- **DataLinks**: Fix blur issues. [#19883](https://github.com/grafana/grafana/pull/19883), [@aocenas](https://github.com/aocenas)
- **Docker**: Makes it possible to parse timezones in the docker image. [#20081](https://github.com/grafana/grafana/pull/20081), [@xlson](https://github.com/xlson)
- **LDAP**: All LDAP servers should be tried even if one of them returns a connection error. [#20077](https://github.com/grafana/grafana/pull/20077), [@jongyllen](https://github.com/jongyllen)
- **LDAP**: No longer shows incorrectly matching groups based on role in debug page. [#20018](https://github.com/grafana/grafana/pull/20018), [@xlson](https://github.com/xlson)
- **Singlestat**: Fix no data / null value mapping . [#19951](https://github.com/grafana/grafana/pull/19951), [@ryantxu](https://github.com/ryantxu)

#### Security vulnerability

The MySQL data source connection string fix, [#20192](https://github.com/grafana/grafana/pull/20192), that was part of this release
also fixed a security vulnerability. Thanks Yuriy Dyachenko for discovering and notifying us about this.

# 6.4.3 (2019-10-16)

### Bug Fixes

- **Alerting**: All notification channels should send even if one fails to send. [#19807](https://github.com/grafana/grafana/pull/19807), [@jan25](https://github.com/jan25)
- **AzureMonitor**: Fix slate interference with dropdowns. [#19799](https://github.com/grafana/grafana/pull/19799), [@aocenas](https://github.com/aocenas)
- **ContextMenu**: make ContextMenu positioning aware of the viewport width. [#19699](https://github.com/grafana/grafana/pull/19699), [@krvajal](https://github.com/krvajal)
- **DataLinks**: Fix context menu not showing in singlestat-ish visualisations. [#19809](https://github.com/grafana/grafana/pull/19809), [@dprokop](https://github.com/dprokop)
- **DataLinks**: Fix url field not releasing focus. [#19804](https://github.com/grafana/grafana/pull/19804), [@aocenas](https://github.com/aocenas)
- **Datasource**: Fixes clicking outside of some query editors required 2 clicks. [#19822](https://github.com/grafana/grafana/pull/19822), [@aocenas](https://github.com/aocenas)
- **Panels**: Fixes default tab for visualizations without Queries Tab. [#19803](https://github.com/grafana/grafana/pull/19803), [@hugohaggmark](https://github.com/hugohaggmark)
- **Singlestat**: Fixed issue with mapping null to text. [#19689](https://github.com/grafana/grafana/pull/19689), [@torkelo](https://github.com/torkelo)
- **@grafana/toolkit**: Don't fail plugin creation when git user.name config is not set. [#19821](https://github.com/grafana/grafana/pull/19821), [@dprokop](https://github.com/dprokop)
- **@grafana/toolkit**: TSLint line number off by 1. [#19782](https://github.com/grafana/grafana/pull/19782), [@fredwangwang](https://github.com/fredwangwang)

# 6.4.2 (2019-10-08)

### Bug Fixes

- **CloudWatch**: Changes incorrect dimension wmlid to wlmid . [#19679](https://github.com/grafana/grafana/pull/19679), [@ATTron](https://github.com/ATTron)
- **Grafana Image Renderer**: Fixes plugin page. [#19664](https://github.com/grafana/grafana/pull/19664), [@hugohaggmark](https://github.com/hugohaggmark)
- **Graph**: Fixes auto decimals logic for y axis ticks that results in too many decimals for high values. [#19618](https://github.com/grafana/grafana/pull/19618), [@torkelo](https://github.com/torkelo)
- **Graph**: Switching to series mode should re-render graph. [#19623](https://github.com/grafana/grafana/pull/19623), [@torkelo](https://github.com/torkelo)
- **Loki**: Fix autocomplete on label values. [#19579](https://github.com/grafana/grafana/pull/19579), [@aocenas](https://github.com/aocenas)
- **Loki**: Removes live option for logs panel. [#19533](https://github.com/grafana/grafana/pull/19533), [@davkal](https://github.com/davkal)
- **Profile**: Fix issue with user profile not showing more than sessions sessions in some cases. [#19578](https://github.com/grafana/grafana/pull/19578), [@huynhsamha](https://github.com/huynhsamha)
- **Prometheus**: Fixes so results in Panel always are sorted by query order. [#19597](https://github.com/grafana/grafana/pull/19597), [@hugohaggmark](https://github.com/hugohaggmark)
- **ShareQuery**: Fixed issue when using -- Dashboard -- datasource (to share query result) when dashboard had rows. [#19610](https://github.com/grafana/grafana/pull/19610), [@torkelo](https://github.com/torkelo)
- **Show SAML login button if SAML is enabled**. [#19591](https://github.com/grafana/grafana/pull/19591), [@papagian](https://github.com/papagian)
- **SingleStat**: Fixes \$\_\_name postfix/prefix usage. [#19687](https://github.com/grafana/grafana/pull/19687), [@hugohaggmark](https://github.com/hugohaggmark)
- **Table**: Proper handling of json data with dataframes. [#19596](https://github.com/grafana/grafana/pull/19596), [@marefr](https://github.com/marefr)
- **Units**: Fixed wrong id for Terabits/sec. [#19611](https://github.com/grafana/grafana/pull/19611), [@andreaslangnevyjel](https://github.com/andreaslangnevyjel)

# 6.4.1 (2019-10-02)

### Bug Fixes

- **Provisioning**: Fixed issue where empty nested keys in YAML provisioning caused server crash, [#19547](https://github.com/grafana/grafana/pull/19547)
- **ImageRendering**: Fixed issue with image rendering in enterprise build (Enterprise)
- **Reporting**: Fixed issue with reporting service when STMP disabled (Enterprise).

# 6.4.0 (2019-10-01)

### Features / Enhancements

- **Build**: Upgrade go to 1.12.10. [#19499](https://github.com/grafana/grafana/pull/19499), [@marefr](https://github.com/marefr)
- **DataLinks**: Suggestions menu improvements. [#19396](https://github.com/grafana/grafana/pull/19396), [@dprokop](https://github.com/dprokop)
- **Explore**: Take root_url setting into account when redirecting from dashboard to explore. [#19447](https://github.com/grafana/grafana/pull/19447), [@ivanahuckova](https://github.com/ivanahuckova)
- **Explore**: Update broken link to logql docs. [#19510](https://github.com/grafana/grafana/pull/19510), [@ivanahuckova](https://github.com/ivanahuckova)
- **Logs**: Adds Logs Panel as a visualization. [#19504](https://github.com/grafana/grafana/pull/19504), [@davkal](https://github.com/davkal)
- **Reporting**: Generate and email PDF reports based on Dashboards (Enterprise)

### Bug Fixes

- **CLI**: Fix version selection for plugin install. [#19498](https://github.com/grafana/grafana/pull/19498), [@aocenas](https://github.com/aocenas)
- **Graph**: Fixes minor issue with series override color picker and custom color . [#19516](https://github.com/grafana/grafana/pull/19516), [@torkelo](https://github.com/torkelo)

## Plugins that need updating when upgrading from 6.3 to 6.4

- [Splunk](https://grafana.com/grafana/plugins/grafana-splunk-datasource)

# 6.4.0-beta2 (2019-09-25)

### Features / Enhancements

- **Azure Monitor**: Remove support for cross resource queries (#19115)". [#19346](https://github.com/grafana/grafana/pull/19346), [@sunker](https://github.com/sunker)
- **Docker**: Upgrade packages to resolve reported vulnerabilities. [#19188](https://github.com/grafana/grafana/pull/19188), [@marefr](https://github.com/marefr)
- **Graphite**: Time range expansion reduced from 1 minute to 1 second. [#19246](https://github.com/grafana/grafana/pull/19246), [@torkelo](https://github.com/torkelo)
- **grafana/toolkit**: Add plugin creation task. [#19207](https://github.com/grafana/grafana/pull/19207), [@dprokop](https://github.com/dprokop)

### Bug Fixes

- **Alerting**: Prevents creating alerts from unsupported queries. [#19250](https://github.com/grafana/grafana/pull/19250), [@hugohaggmark](https://github.com/hugohaggmark)
- **Alerting**: Truncate PagerDuty summary when greater than 1024 characters. [#18730](https://github.com/grafana/grafana/pull/18730), [@nvllsvm](https://github.com/nvllsvm)
- **Cloudwatch**: Fix autocomplete for Gamelift dimensions. [#19146](https://github.com/grafana/grafana/pull/19146), [@kevinpz](https://github.com/kevinpz)
- **Dashboard**: Fix export for sharing when panels use default data source. [#19315](https://github.com/grafana/grafana/pull/19315), [@torkelo](https://github.com/torkelo)
- **Database**: Rewrite system statistics query to perform better. [#19178](https://github.com/grafana/grafana/pull/19178), [@papagian](https://github.com/papagian)
- **Gauge/BarGauge**: Fix issue with [object Object] in titles . [#19217](https://github.com/grafana/grafana/pull/19217), [@ryantxu](https://github.com/ryantxu)
- **MSSQL**: Revert usage of new connectionstring format introduced by #18384. [#19203](https://github.com/grafana/grafana/pull/19203), [@marefr](https://github.com/marefr)
- **Multi-LDAP**: Do not fail-fast on invalid credentials. [#19261](https://github.com/grafana/grafana/pull/19261), [@gotjosh](https://github.com/gotjosh)
- **MySQL, Postgres, MSSQL**: Fix validating query with template variables in alert . [#19237](https://github.com/grafana/grafana/pull/19237), [@marefr](https://github.com/marefr)
- **MySQL, Postgres**: Update raw sql when query builder updates. [#19209](https://github.com/grafana/grafana/pull/19209), [@marefr](https://github.com/marefr)
- **MySQL**: Limit datasource error details returned from the backend. [#19373](https://github.com/grafana/grafana/pull/19373), [@marefr](https://github.com/marefr)

# 6.4.0-beta1 (2019-09-17)

### Features / Enhancements

- **Reporting**: Created scheduled PDF reports for any dashboard (Enterprise).
- **API**: Readonly datasources should not be created via the API. [#19006](https://github.com/grafana/grafana/pull/19006), [@papagian](https://github.com/papagian)
- **Alerting**: Include configured AlertRuleTags in Webhooks notifier. [#18233](https://github.com/grafana/grafana/pull/18233), [@dominic-miglar](https://github.com/dominic-miglar)
- **Annotations**: Add annotations support to Loki. [#18949](https://github.com/grafana/grafana/pull/18949), [@aocenas](https://github.com/aocenas)
- **Annotations**: Use a single row to represent a region. [#17673](https://github.com/grafana/grafana/pull/17673), [@ryantxu](https://github.com/ryantxu)
- **Auth**: Allow inviting existing users when login form is disabled. [#19048](https://github.com/grafana/grafana/pull/19048), [@548017](https://github.com/548017)
- **Azure Monitor**: Add support for cross resource queries. [#19115](https://github.com/grafana/grafana/pull/19115), [@sunker](https://github.com/sunker)
- **CLI**: Allow installing custom binary plugins. [#17551](https://github.com/grafana/grafana/pull/17551), [@aocenas](https://github.com/aocenas)
- **Dashboard**: Adds Logs Panel (alpha) as visualization option for Dashboards. [#18641](https://github.com/grafana/grafana/pull/18641), [@hugohaggmark](https://github.com/hugohaggmark)
- **Dashboard**: Reuse query results between panels . [#16660](https://github.com/grafana/grafana/pull/16660), [@ryantxu](https://github.com/ryantxu)
- **Dashboard**: Set time to to 23:59:59 when setting To time using calendar. [#18595](https://github.com/grafana/grafana/pull/18595), [@simPod](https://github.com/simPod)
- **DataLinks**: Add DataLinks support to Gauge, BarGauge and stat panel. [#18605](https://github.com/grafana/grafana/pull/18605), [@ryantxu](https://github.com/ryantxu)
- **DataLinks**: Enable access to labels & field names. [#18918](https://github.com/grafana/grafana/pull/18918), [@torkelo](https://github.com/torkelo)
- **DataLinks**: Enable multiple data links per panel. [#18434](https://github.com/grafana/grafana/pull/18434), [@dprokop](https://github.com/dprokop)
- **Docker**: switch docker image to alpine base with phantomjs support. [#18468](https://github.com/grafana/grafana/pull/18468), [@DanCech](https://github.com/DanCech)
- **Elasticsearch**: allow templating queries to order by doc_count. [#18870](https://github.com/grafana/grafana/pull/18870), [@hackery](https://github.com/hackery)
- **Explore**: Add throttling when doing live queries. [#19085](https://github.com/grafana/grafana/pull/19085), [@aocenas](https://github.com/aocenas)
- **Explore**: Adds ability to go back to dashboard, optionally with query changes. [#17982](https://github.com/grafana/grafana/pull/17982), [@kaydelaney](https://github.com/kaydelaney)
- **Explore**: Reduce default time range to last hour. [#18212](https://github.com/grafana/grafana/pull/18212), [@davkal](https://github.com/davkal)
- **Gauge/BarGauge**: Support decimals for min/max. [#18368](https://github.com/grafana/grafana/pull/18368), [@ryantxu](https://github.com/ryantxu)
- **Graph**: New series override transform constant that renders a single point as a line across the whole graph. [#19102](https://github.com/grafana/grafana/pull/19102), [@davkal](https://github.com/davkal)
- **Image rendering**: Add deprecation warning when PhantomJS is used for rendering images. [#18933](https://github.com/grafana/grafana/pull/18933), [@papagian](https://github.com/papagian)
- **InfluxDB**: Enable interpolation within ad-hoc filter values. [#18077](https://github.com/grafana/grafana/pull/18077), [@kvc-code](https://github.com/kvc-code)
- **LDAP**: Allow an user to be synchronized against LDAP. [#18976](https://github.com/grafana/grafana/pull/18976), [@gotjosh](https://github.com/gotjosh)
- **Ldap**: Add ldap debug page. [#18759](https://github.com/grafana/grafana/pull/18759), [@peterholmberg](https://github.com/peterholmberg)
- **Loki**: Remove prefetching of default label values. [#18213](https://github.com/grafana/grafana/pull/18213), [@davkal](https://github.com/davkal)
- **Metrics**: Add failed alert notifications metric. [#18089](https://github.com/grafana/grafana/pull/18089), [@koorgoo](https://github.com/koorgoo)
- **OAuth**: Support JMES path lookup when retrieving user email. [#14683](https://github.com/grafana/grafana/pull/14683), [@bobmshannon](https://github.com/bobmshannon)
- **OAuth**: return GitLab groups as a part of user info (enable team sync). [#18388](https://github.com/grafana/grafana/pull/18388), [@alexanderzobnin](https://github.com/alexanderzobnin)
- **Panels**: Add unit for electrical charge - ampere-hour. [#18950](https://github.com/grafana/grafana/pull/18950), [@anirudh-ramesh](https://github.com/anirudh-ramesh)
- **Plugin**: AzureMonitor - Reapply MetricNamespace support. [#17282](https://github.com/grafana/grafana/pull/17282), [@raphaelquati](https://github.com/raphaelquati)
- **Plugins**: better warning when plugins fail to load. [#18671](https://github.com/grafana/grafana/pull/18671), [@ryantxu](https://github.com/ryantxu)
- **Postgres**: Add support for scram sha 256 authentication. [#18397](https://github.com/grafana/grafana/pull/18397), [@nonamef](https://github.com/nonamef)
- **RemoteCache**: Support SSL with Redis. [#18511](https://github.com/grafana/grafana/pull/18511), [@kylebrandt](https://github.com/kylebrandt)
- **SingleStat**: The gauge option in now disabled/hidden (unless it's an old panel with it already enabled) . [#18610](https://github.com/grafana/grafana/pull/18610), [@ryantxu](https://github.com/ryantxu)
- **Stackdriver**: Add extra alignment period options. [#18909](https://github.com/grafana/grafana/pull/18909), [@sunker](https://github.com/sunker)
- **Units**: Add South African Rand (ZAR) to currencies. [#18893](https://github.com/grafana/grafana/pull/18893), [@jeteon](https://github.com/jeteon)
- **Units**: Adding T,P,E,Z,and Y bytes. [#18706](https://github.com/grafana/grafana/pull/18706), [@chiqomar](https://github.com/chiqomar)

### Bug Fixes

- **Alerting**: Notification is sent when state changes from no_data to ok. [#18920](https://github.com/grafana/grafana/pull/18920), [@papagian](https://github.com/papagian)
- **Alerting**: fix duplicate alert states when the alert fails to save to the database. [#18216](https://github.com/grafana/grafana/pull/18216), [@kylebrandt](https://github.com/kylebrandt)
- **Alerting**: fix response popover prompt when add notification channels. [#18967](https://github.com/grafana/grafana/pull/18967), [@lzdw](https://github.com/lzdw)
- **CloudWatch**: Fix alerting for queries with Id (using GetMetricData). [#17899](https://github.com/grafana/grafana/pull/17899), [@alex-berger](https://github.com/alex-berger)
- **Explore**: Fix auto completion on label values for Loki. [#18988](https://github.com/grafana/grafana/pull/18988), [@aocenas](https://github.com/aocenas)
- **Explore**: Fixes crash using back button with a zoomed in graph. [#19122](https://github.com/grafana/grafana/pull/19122), [@hugohaggmark](https://github.com/hugohaggmark)
- **Explore**: Fixes so queries in Explore are only run if Graph/Table is shown. [#19000](https://github.com/grafana/grafana/pull/19000), [@hugohaggmark](https://github.com/hugohaggmark)
- **MSSQL**: Change connectionstring to URL format to fix using passwords with semicolon. [#18384](https://github.com/grafana/grafana/pull/18384), [@Russiancold](https://github.com/Russiancold)
- **MSSQL**: Fix memory leak when debug enabled. [#19049](https://github.com/grafana/grafana/pull/19049), [@briangann](https://github.com/briangann)
- **Provisioning**: Allow escaping literal '$' with '$\$' in configs to avoid interpolation. [#18045](https://github.com/grafana/grafana/pull/18045), [@kylebrandt](https://github.com/kylebrandt)
- **TimePicker**: Fixes hiding time picker dropdown in FireFox. [#19154](https://github.com/grafana/grafana/pull/19154), [@hugohaggmark](https://github.com/hugohaggmark)

## Breaking changes

### Annotations

There are some breaking changes in the annotations HTTP API for region annotations. Region annotations are now represented
using a single event instead of two separate events. Check breaking changes in HTTP API [below](#http-api) and [HTTP API documentation](https://grafana.com/docs/http_api/annotations/) for more details.

### Docker

Grafana is now using Alpine 3.10 as docker base image.

### HTTP API

- `GET /api/alert-notifications` now requires at least editor access. New `/api/alert-notifications/lookup` returns less information than `/api/alert-notifications` and can be access by any authenticated user.
- `GET /api/alert-notifiers` now requires at least editor access
- `GET /api/org/users` now requires org admin role. New `/api/org/users/lookup` returns less information than `/api/org/users` and can be access by users that are org admins, admin in any folder or admin of any team.
- `GET /api/annotations` no longer returns `regionId` property.
- `POST /api/annotations` no longer supports `isRegion` property.
- `PUT /api/annotations/:id` no longer supports `isRegion` property.
- `PATCH /api/annotations/:id` no longer supports `isRegion` property.
- `DELETE /api/annotations/region/:id` has been removed.

## Deprecation notes

### PhantomJS

[PhantomJS](https://phantomjs.org/), which is used for rendering images of dashboards and panels, is deprecated and will be removed in a future Grafana release. A deprecation warning will from now on be logged when Grafana starts up if PhantomJS is in use.

Please consider migrating from PhantomJS to the [Grafana Image Renderer plugin](https://grafana.com/grafana/plugins/grafana-image-renderer).

# 6.3.7 (2019-11-22)

### Bug Fixes

- **CloudWatch**: Fix high CPU load [#20579](https://github.com/grafana/grafana/pull/20579)

# 6.3.6 (2019-09-23)

### Features / Enhancements

- **Metrics**: Adds setting for turning off total stats metrics. [#19142](https://github.com/grafana/grafana/pull/19142), [@marefr](https://github.com/marefr)

### Bug Fixes

- **Database**: Rewrite system statistics query to perform better. [#19178](https://github.com/grafana/grafana/pull/19178), [@papagian](https://github.com/papagian)
- **Explore**: Fixes error when switching from prometheus to loki data sources. [#18599](https://github.com/grafana/grafana/pull/18599), [@kaydelaney](https://github.com/kaydelaney)

# 6.3.5 (2019-09-02)

### Upgrades

- **Build**: Upgrade to go 1.12.9. [#18638](https://github.com/grafana/grafana/pull/18638), [@marcusolsson](https://github.com/marcusolsson)

### Bug Fixes

- **Dashboard**: Fixes dashboards init failed loading error for dashboards with panel links that had missing properties. [#18786](https://github.com/grafana/grafana/pull/18786), [@torkelo](https://github.com/torkelo)
- **Editor**: Fixes issue where only entire lines were being copied. [#18806](https://github.com/grafana/grafana/pull/18806), [@kaydelaney](https://github.com/kaydelaney)
- **Explore**: Fixes query field layout in splitted view for Safari browsers. [#18654](https://github.com/grafana/grafana/pull/18654), [@hugohaggmark](https://github.com/hugohaggmark)
- **LDAP**: multildap + ldap integration. [#18588](https://github.com/grafana/grafana/pull/18588), [@markelog](https://github.com/markelog)
- **Profile/UserAdmin**: Fix for user agent parser crashes grafana-server on 32-bit builds. [#18788](https://github.com/grafana/grafana/pull/18788), [@marcusolsson](https://github.com/marcusolsson)
- **Prometheus**: Prevents panel editor crash when switching to Prometheus data source. [#18616](https://github.com/grafana/grafana/pull/18616), [@hugohaggmark](https://github.com/hugohaggmark)
- **Prometheus**: Changes brace-insertion behavior to be less annoying. [#18698](https://github.com/grafana/grafana/pull/18698), [@kaydelaney](https://github.com/kaydelaney)

# 6.3.4 (2019-08-29)

- **Security**: Urgent security patch release. Please read more in our [blog](https://grafana.com/blog/2019/08/29/grafana-5.4.5-and-6.3.4-released-with-important-security-fix/)

# 6.3.3 (2019-08-15)

### Bug Fixes

- **Annotations**: Fix failing annotation query when time series query is cancelled. [#18532](https://github.com/grafana/grafana/pull/18532), [@dprokop](https://github.com/dprokop)
- **Auth**: Do not set SameSite cookie attribute if cookie_samesite is none. [#18462](https://github.com/grafana/grafana/pull/18462), [@papagian](https://github.com/papagian)
- **DataLinks**: Apply scoped variables to data links correctly. [#18454](https://github.com/grafana/grafana/pull/18454), [@dprokop](https://github.com/dprokop)
- **DataLinks**: Respect timezone when displaying datapoint's timestamp in graph context menu. [#18461](https://github.com/grafana/grafana/pull/18461), [@dprokop](https://github.com/dprokop)
- **DataLinks**: Use datapoint timestamp correctly when interpolating variables. [#18459](https://github.com/grafana/grafana/pull/18459), [@dprokop](https://github.com/dprokop)
- **Explore**: Fix loading error for empty queries. [#18488](https://github.com/grafana/grafana/pull/18488), [@davkal](https://github.com/davkal)
- **Graph**: Fixes legend issue clicking on series line icon and issue with horizontal scrollbar being visible on windows. [#18563](https://github.com/grafana/grafana/pull/18563), [@torkelo](https://github.com/torkelo)
- **Graphite**: Avoid glob of single-value array variables . [#18420](https://github.com/grafana/grafana/pull/18420), [@gotjosh](https://github.com/gotjosh)
- **Prometheus**: Fix queries with label_replace remove the \$1 match when loading query editor. [#18480](https://github.com/grafana/grafana/pull/18480), [@hugohaggmark](https://github.com/hugohaggmark)
- **Prometheus**: More consistently allows for multi-line queries in editor. [#18362](https://github.com/grafana/grafana/pull/18362), [@kaydelaney](https://github.com/kaydelaney)
- **TimeSeries**: Assume values are all numbers. [#18540](https://github.com/grafana/grafana/pull/18540), [@ryantxu](https://github.com/ryantxu)

# 6.3.2 (2019-08-07)

### Bug Fixes

- **Gauge/BarGauge**: Fixes issue with lost thresholds and an issue loading Gauge with avg stat. [#18375](https://github.com/grafana/grafana/pull/18375)

# 6.3.1 (2019-08-07)

### Bug Fixes

- **PanelLinks**: Fixes crash issue with Gauge & Bar Gauge panels with panel links (drill down links). [#18430](https://github.com/grafana/grafana/pull/18430)

# 6.3.0 (2019-08-06)

### Features / Enhancements

- **OAuth**: Do not set SameSite OAuth cookie if cookie_samesite is None. [#18392](https://github.com/grafana/grafana/pull/18392), [@papagian](https://github.com/papagian)

### Bug Fixes

- **PanelLinks**: Fix render issue when there is no panel description. [#18408](https://github.com/grafana/grafana/pull/18408), [@dehrax](https://github.com/dehrax)

# 6.3.0-beta4 (2019-08-02)

### Features / Enhancements

- **Auth Proxy**: Include additional headers as part of the cache key. [#18298](https://github.com/grafana/grafana/pull/18298), [@gotjosh](https://github.com/gotjosh)

# 6.3.0-beta3 (2019-08-02)

### Bug Fixes

- **OAuth**: Fix "missing saved state" OAuth login failure due to SameSite cookie policy. [#18332](https://github.com/grafana/grafana/pull/18332), [@papagian](https://github.com/papagian)
- **cli**: fix for recognizing when in dev mode.. [#18334](https://github.com/grafana/grafana/pull/18334), [@xlson](https://github.com/xlson)

# 6.3.0-beta2 (2019-07-26)

### Features / Enhancements

- **Build grafana images consistently**. [#18224](https://github.com/grafana/grafana/pull/18224), [@hassanfarid](https://github.com/hassanfarid)
- **Docs**: SAML. [#18069](https://github.com/grafana/grafana/pull/18069), [@gotjosh](https://github.com/gotjosh)
- **Permissions**: Show plugins in nav for non admin users but hide plugin configuration. [#18234](https://github.com/grafana/grafana/pull/18234), [@aocenas](https://github.com/aocenas)
- **TimePicker**: Increase max height of quick range dropdown. [#18247](https://github.com/grafana/grafana/pull/18247), [@torkelo](https://github.com/torkelo)

### Bug Fixes

- **DataLinks**: Fixes incorrect interpolation of \${\_\_series_name} . [#18251](https://github.com/grafana/grafana/pull/18251), [@torkelo](https://github.com/torkelo)
- **Loki**: Display live tailed logs in correct order in Explore. [#18031](https://github.com/grafana/grafana/pull/18031), [@kaydelaney](https://github.com/kaydelaney)
- **PhantomJS**: Fixes rendering on Debian Buster. [#18162](https://github.com/grafana/grafana/pull/18162), [@xlson](https://github.com/xlson)
- **TimePicker**: Fixed style issue for custom range popover. [#18244](https://github.com/grafana/grafana/pull/18244), [@torkelo](https://github.com/torkelo)
- **Timerange**: Fixes a bug where custom time ranges didn't respect UTC. [#18248](https://github.com/grafana/grafana/pull/18248), [@kaydelaney](https://github.com/kaydelaney)
- **remote_cache**: Fix redis connstr parsing. [#18204](https://github.com/grafana/grafana/pull/18204), [@mblaschke](https://github.com/mblaschke)

# 6.3.0-beta1 (2019-07-10)

### Features / Enhancements

- **Alerting**: Add tags to alert rules. [#10989](https://github.com/grafana/grafana/pull/10989), [@Thib17](https://github.com/Thib17)
- **Alerting**: Attempt to send email notifications to all given email addresses. [#16881](https://github.com/grafana/grafana/pull/16881), [@zhulongcheng](https://github.com/zhulongcheng)
- **Alerting**: Improve alert rule testing. [#16286](https://github.com/grafana/grafana/pull/16286), [@marefr](https://github.com/marefr)
- **Alerting**: Support for configuring content field for Discord alert notifier. [#17017](https://github.com/grafana/grafana/pull/17017), [@jan25](https://github.com/jan25)
- **Alertmanager**: Replace illegal chars with underscore in label names. [#17002](https://github.com/grafana/grafana/pull/17002), [@bergquist](https://github.com/bergquist)
- **Auth**: Allow expiration of API keys. [#17678](https://github.com/grafana/grafana/pull/17678), [@papagian](https://github.com/papagian)
- **Auth**: Return device, os and browser when listing user auth tokens in HTTP API. [#17504](https://github.com/grafana/grafana/pull/17504), [@shavonn](https://github.com/shavonn)
- **Auth**: Support list and revoke of user auth tokens in UI. [#17434](https://github.com/grafana/grafana/pull/17434), [@shavonn](https://github.com/shavonn)
- **AzureMonitor**: change clashing built-in Grafana variables/macro names for Azure Logs. [#17140](https://github.com/grafana/grafana/pull/17140), [@shavonn](https://github.com/shavonn)
- **CloudWatch**: Made region visible for AWS Cloudwatch Expressions. [#17243](https://github.com/grafana/grafana/pull/17243), [@utkarshcmu](https://github.com/utkarshcmu)
- **Cloudwatch**: Add AWS DocDB metrics. [#17241](https://github.com/grafana/grafana/pull/17241), [@utkarshcmu](https://github.com/utkarshcmu)
- **Dashboard**: Use timezone dashboard setting when exporting to CSV. [#18002](https://github.com/grafana/grafana/pull/18002), [@dehrax](https://github.com/dehrax)
- **Data links**. [#17267](https://github.com/grafana/grafana/pull/17267), [@torkelo](https://github.com/torkelo)
- **Docker**: Switch base image to ubuntu:latest from debian:stretch to avoid security issues.. [#17066](https://github.com/grafana/grafana/pull/17066), [@bergquist](https://github.com/bergquist)
- **Elasticsearch**: Support for visualizing logs in Explore . [#17605](https://github.com/grafana/grafana/pull/17605), [@marefr](https://github.com/marefr)
- **Explore**: Adds Live option for supported data sources. [#17062](https://github.com/grafana/grafana/pull/17062), [@hugohaggmark](https://github.com/hugohaggmark)
- **Explore**: Adds orgId to URL for sharing purposes. [#17895](https://github.com/grafana/grafana/pull/17895), [@kaydelaney](https://github.com/kaydelaney)
- **Explore**: Adds support for new loki 'start' and 'end' params for labels endpoint. [#17512](https://github.com/grafana/grafana/pull/17512), [@kaydelaney](https://github.com/kaydelaney)
- **Explore**: Adds support for toggling raw query mode in explore. [#17870](https://github.com/grafana/grafana/pull/17870), [@kaydelaney](https://github.com/kaydelaney)
- **Explore**: Allow switching between metrics and logs . [#16959](https://github.com/grafana/grafana/pull/16959), [@marefr](https://github.com/marefr)
- **Explore**: Combines the timestamp and local time columns into one. [#17775](https://github.com/grafana/grafana/pull/17775), [@hugohaggmark](https://github.com/hugohaggmark)
- **Explore**: Display log lines context . [#17097](https://github.com/grafana/grafana/pull/17097), [@dprokop](https://github.com/dprokop)
- **Explore**: Don't parse log levels if provided by field or label. [#17180](https://github.com/grafana/grafana/pull/17180), [@marefr](https://github.com/marefr)
- **Explore**: Improves performance of Logs element by limiting re-rendering. [#17685](https://github.com/grafana/grafana/pull/17685), [@kaydelaney](https://github.com/kaydelaney)
- **Explore**: Support for new LogQL filtering syntax. [#16674](https://github.com/grafana/grafana/pull/16674), [@davkal](https://github.com/davkal)
- **Explore**: Use new TimePicker from Grafana/UI. [#17793](https://github.com/grafana/grafana/pull/17793), [@hugohaggmark](https://github.com/hugohaggmark)
- **Explore**: handle newlines in LogRow Highlighter. [#17425](https://github.com/grafana/grafana/pull/17425), [@rrfeng](https://github.com/rrfeng)
- **Graph**: Added new fill gradient option. [#17528](https://github.com/grafana/grafana/pull/17528), [@torkelo](https://github.com/torkelo)
- **GraphPanel**: Don't sort series when legend table & sort column is not visible . [#17095](https://github.com/grafana/grafana/pull/17095), [@shavonn](https://github.com/shavonn)
- **InfluxDB**: Support for visualizing logs in Explore. [#17450](https://github.com/grafana/grafana/pull/17450), [@hugohaggmark](https://github.com/hugohaggmark)
- **Logging**: Login and Logout actions (#17760). [#17883](https://github.com/grafana/grafana/pull/17883), [@ATTron](https://github.com/ATTron)
- **Logging**: Move log package to pkg/infra. [#17023](https://github.com/grafana/grafana/pull/17023), [@zhulongcheng](https://github.com/zhulongcheng)
- **Metrics**: Expose stats about roles as metrics. [#17469](https://github.com/grafana/grafana/pull/17469), [@bergquist](https://github.com/bergquist)
- **MySQL/Postgres/MSSQL**: Add parsing for day, weeks and year intervals in macros. [#13086](https://github.com/grafana/grafana/pull/13086), [@bernardd](https://github.com/bernardd)
- **MySQL**: Add support for periodically reloading client certs. [#14892](https://github.com/grafana/grafana/pull/14892), [@tpetr](https://github.com/tpetr)
- **Plugins**: replace dataFormats list with skipDataQuery flag in plugin.json. [#16984](https://github.com/grafana/grafana/pull/16984), [@ryantxu](https://github.com/ryantxu)
- **Prometheus**: Take timezone into account for step alignment. [#17477](https://github.com/grafana/grafana/pull/17477), [@fxmiii](https://github.com/fxmiii)
- **Prometheus**: Use overridden panel range for \$\_\_range instead of dashboard range. [#17352](https://github.com/grafana/grafana/pull/17352), [@patrick246](https://github.com/patrick246)
- **Prometheus**: added time range filter to series labels query. [#16851](https://github.com/grafana/grafana/pull/16851), [@FUSAKLA](https://github.com/FUSAKLA)
- **Provisioning**: Support folder that doesn't exist yet in dashboard provisioning. [#17407](https://github.com/grafana/grafana/pull/17407), [@Nexucis](https://github.com/Nexucis)
- **Refresh picker**: Handle empty intervals. [#17585](https://github.com/grafana/grafana/pull/17585), [@dehrax](https://github.com/dehrax)
- **Singlestat**: Add y min/max config to singlestat sparklines. [#17527](https://github.com/grafana/grafana/pull/17527), [@pitr](https://github.com/pitr)
- **Snapshot**: use given key and deleteKey. [#16876](https://github.com/grafana/grafana/pull/16876), [@zhulongcheng](https://github.com/zhulongcheng)
- **Templating**: Correctly display \_\_text in multi-value variable after page reload. [#17840](https://github.com/grafana/grafana/pull/17840), [@EduardSergeev](https://github.com/EduardSergeev)
- **Templating**: Support selecting all filtered values of a multi-value variable. [#16873](https://github.com/grafana/grafana/pull/16873), [@r66ad](https://github.com/r66ad)
- **Tracing**: allow propagation with Zipkin headers. [#17009](https://github.com/grafana/grafana/pull/17009), [@jrockway](https://github.com/jrockway)
- **Users**: Disable users removed from LDAP. [#16820](https://github.com/grafana/grafana/pull/16820), [@alexanderzobnin](https://github.com/alexanderzobnin)
- **SAML**: Add SAML as an authentication option (Enterprise)

### Bug Fixes

- **AddPanel**: Fix issue when removing moved add panel widget . [#17659](https://github.com/grafana/grafana/pull/17659), [@dehrax](https://github.com/dehrax)
- **CLI**: Fix encrypt-datasource-passwords fails with sql error. [#18014](https://github.com/grafana/grafana/pull/18014), [@marefr](https://github.com/marefr)
- **Elasticsearch**: Fix default max concurrent shard requests. [#17770](https://github.com/grafana/grafana/pull/17770), [@marefr](https://github.com/marefr)
- **Explore**: Fix browsing back to dashboard panel. [#17061](https://github.com/grafana/grafana/pull/17061), [@jschill](https://github.com/jschill)
- **Explore**: Fix filter by series level in logs graph. [#17798](https://github.com/grafana/grafana/pull/17798), [@marefr](https://github.com/marefr)
- **Explore**: Fix issues when loading and both graph/table are collapsed. [#17113](https://github.com/grafana/grafana/pull/17113), [@marefr](https://github.com/marefr)
- **Explore**: Fix selection/copy of log lines. [#17121](https://github.com/grafana/grafana/pull/17121), [@marefr](https://github.com/marefr)
- **Fix**: Wrap value of multi variable in array when coming from URL. [#16992](https://github.com/grafana/grafana/pull/16992), [@aocenas](https://github.com/aocenas)
- **Frontend**: Fix for Json tree component not working. [#17608](https://github.com/grafana/grafana/pull/17608), [@srid12](https://github.com/srid12)
- **Graphite**: Fix for issue with alias function being moved last. [#17791](https://github.com/grafana/grafana/pull/17791), [@torkelo](https://github.com/torkelo)
- **Graphite**: Fixes issue with seriesByTag & function with variable param. [#17795](https://github.com/grafana/grafana/pull/17795), [@torkelo](https://github.com/torkelo)
- **Graphite**: use POST for /metrics/find requests. [#17814](https://github.com/grafana/grafana/pull/17814), [@papagian](https://github.com/papagian)
- **HTTP Server**: Serve Grafana with a custom URL path prefix. [#17048](https://github.com/grafana/grafana/pull/17048), [@jan25](https://github.com/jan25)
- **InfluxDB**: Fixes single quotes are not escaped in label value filters. [#17398](https://github.com/grafana/grafana/pull/17398), [@Panzki](https://github.com/Panzki)
- **Prometheus**: Correctly escape '|' literals in interpolated PromQL variables. [#16932](https://github.com/grafana/grafana/pull/16932), [@Limess](https://github.com/Limess)
- **Prometheus**: Fix when adding label for metrics which contains colons in Explore. [#16760](https://github.com/grafana/grafana/pull/16760), [@tolwi](https://github.com/tolwi)
- **SinglestatPanel**: Remove background color when value turns null. [#17552](https://github.com/grafana/grafana/pull/17552), [@druggieri](https://github.com/druggieri)

# 6.2.5 (2019-06-25)

### Features / Enhancements

- **Grafana-CLI**: Wrapper for `grafana-cli` within RPM/DEB packages and config/homepath are now global flags. [#17695](https://github.com/grafana/grafana/pull/17695), [@gotjosh](https://github.com/gotjosh)
- **Panel**: Fully escape html in drilldown links (was only sanitized before) . [#17731](https://github.com/grafana/grafana/pull/17731), [@dehrax](https://github.com/dehrax)

### Bug Fixes

- **Config**: Fix connectionstring for remote_cache in defaults.ini. [#17675](https://github.com/grafana/grafana/pull/17675), [@kylebrandt](https://github.com/kylebrandt)
- **Elasticsearch**: Fix empty query (via template variable) should be sent as wildcard. [#17488](https://github.com/grafana/grafana/pull/17488), [@davewat](https://github.com/davewat)
- **HTTP-Server**: Fix Strict-Transport-Security header. [#17644](https://github.com/grafana/grafana/pull/17644), [@kylebrandt](https://github.com/kylebrandt)
- **TablePanel**: fix annotations display. [#17646](https://github.com/grafana/grafana/pull/17646), [@ryantxu](https://github.com/ryantxu)

# 6.2.4 (2019-06-18)

### Bug Fixes

- **Grafana-CLI**: Fix receiving flags via command line . [#17617](https://github.com/grafana/grafana/pull/17617), [@gotjosh](https://github.com/gotjosh)
- **HTTPServer**: Fix X-XSS-Protection header formatting. [#17620](https://github.com/grafana/grafana/pull/17620), [@yverry](https://github.com/yverry)

# 6.2.3 (2019-06-17)

### Known issues

- **grafana-cli**: The argument `--pluginsDir` is not working.
- **docker**: Due to above problem with grafana-cli the docker run will fail to start the container if you're installing plugins using the `PL_INSTALL_PLUGINS` environment variable. We have removed 6.2.3 tag from docker hub and latest tag now points to 6.2.2.

More details in bug report: https://github.com/grafana/grafana/issues/17613

### Features / Enhancements

- **AuthProxy**: Optimistic lock pattern for remote cache Set. [#17485](https://github.com/grafana/grafana/pull/17485), [@papagian](https://github.com/papagian)
- **HTTPServer**: Options for returning new headers X-Content-Type-Options, X-XSS-Protection and Strict-Transport-Security. [#17522](https://github.com/grafana/grafana/pull/17522), [@kylebrandt](https://github.com/kylebrandt)

### Bug Fixes

- **Auth Proxy**: Fix non-negative cache TTL. [#17495](https://github.com/grafana/grafana/pull/17495), [@kylebrandt](https://github.com/kylebrandt)
- **Grafana-CLI**: Fix receiving configuration flags from the command line. [#17606](https://github.com/grafana/grafana/pull/17606), [@gotjosh](https://github.com/gotjosh)
- **OAuth**: Fix for wrong user token updated on OAuth refresh in DS proxy. [#17541](https://github.com/grafana/grafana/pull/17541), [@redbaron](https://github.com/redbaron)
- **remote_cache**: Fix redis. [#17483](https://github.com/grafana/grafana/pull/17483), [@kylebrandt](https://github.com/kylebrandt)

# 6.2.2 (2019-06-05)

### Features / Enhancements

- **Security**: Prevent CSV formula injection attack when exporting data. [#17363](https://github.com/grafana/grafana/pull/17363), [@DanCech](https://github.com/DanCech)

### Bug Fixes

- **CloudWatch**: Fixes error when hiding/disabling queries . [#17283](https://github.com/grafana/grafana/pull/17283), [@jpiccari](https://github.com/jpiccari)
- **Database**: Fixed slow permission query in folder/dashboard search. [#17427](https://github.com/grafana/grafana/pull/17427), [@aocenas](https://github.com/aocenas)
- **Explore**: Fixed updating time range before running queries. [#17349](https://github.com/grafana/grafana/pull/17349), [@marefr](https://github.com/marefr)
- **Plugins**: Fixed plugin config page navigation when using subpath. [#17364](https://github.com/grafana/grafana/pull/17364), [@torkelo](https://github.com/torkelo)

# 6.2.1 (2019-05-27)

### Features / Enhancements

- **CLI**: Add command to migrate all data sources to use encrypted password fields . [#17118](https://github.com/grafana/grafana/pull/17118), [@aocenas](https://github.com/aocenas)
- **Gauge/BarGauge**: Improvements to auto value font size . [#17292](https://github.com/grafana/grafana/pull/17292), [@torkelo](https://github.com/torkelo)

### Bug Fixes

- **Auth Proxy**: Resolve database is locked errors. [#17274](https://github.com/grafana/grafana/pull/17274), [@marefr](https://github.com/marefr)
- **Database**: Retry transaction if sqlite returns database is locked error. [#17276](https://github.com/grafana/grafana/pull/17276), [@marefr](https://github.com/marefr)
- **Explore**: Fixes so clicking in a Prometheus Table the query is filtered by clicked value. [#17083](https://github.com/grafana/grafana/pull/17083), [@hugohaggmark](https://github.com/hugohaggmark)
- **Singlestat**: Fixes issue with value placement and line wraps. [#17249](https://github.com/grafana/grafana/pull/17249), [@torkelo](https://github.com/torkelo)
- **Tech**: Update jQuery to 3.4.1 to fix issue on iOS 10 based browsers as well as Chrome 53.x . [#17290](https://github.com/grafana/grafana/pull/17290), [@timbutler](https://github.com/timbutler)

# 6.2.0 (2019-05-22)

### Bug Fixes

- **BarGauge**: Fix for negative min values. [#17192](https://github.com/grafana/grafana/pull/17192), [@torkelo](https://github.com/torkelo)
- **Gauge/BarGauge**: Fix for issues editing min & max options. [#17174](https://github.com/grafana/grafana/pull/17174)
- **Search**: Make only folder name only open search with current folder filter. [#17226](https://github.com/grafana/grafana/pull/17226)
- **AzureMonitor**: Revert to clearing chained dropdowns. [#17212](https://github.com/grafana/grafana/pull/17212)

### Breaking Changes

- **Plugins**: Data source plugins that process hidden queries need to add a "hiddenQueries: true" attribute in plugin.json. [#17124](https://github.com/grafana/grafana/pull/17124), [@ryantxu](https://github.com/ryantxu)

### Removal of old deprecated package repository

5 months ago we deprecated our old package cloud repository and [replaced it](https://grafana.com/blog/2019/01/05/moving-to-packages.grafana.com/) with our own. We will remove the old depreciated
repo on July 1st. Make sure you have switched to the new repo by then. The new repository has all our old releases so you are not required to upgrade just to switch package repository.

# 6.2.0-beta2 (2019-05-15)

### Features / Enhancements

- **Plugins**: Support templated urls in plugin routes. [#16599](https://github.com/grafana/grafana/pull/16599), [@briangann](https://github.com/briangann)
- **Packaging**: New MSI windows installer package\*\*. [#17073](https://github.com/grafana/grafana/pull/17073), [@briangann](https://github.com/briangann)

### Bug Fixes

- **Dashboard**: Fixes blank dashboard after window resize with panel without title. [#16942](https://github.com/grafana/grafana/pull/16942), [@torkelo](https://github.com/torkelo)
- **Dashboard**: Fixes lazy loading & expanding collapsed rows on mobile. [#17055](https://github.com/grafana/grafana/pull/17055), [@torkelo](https://github.com/torkelo)
- **Dashboard**: Fixes scrolling issues for Edge browser. [#17033](https://github.com/grafana/grafana/pull/17033), [@jschill](https://github.com/jschill)
- **Dashboard**: Show refresh button in first kiosk(tv) mode. [#17032](https://github.com/grafana/grafana/pull/17032), [@torkelo](https://github.com/torkelo)
- **Explore**: Fix empty result from data source should render logs container. [#16999](https://github.com/grafana/grafana/pull/16999), [@marefr](https://github.com/marefr)
- **Explore**: Fixes so clicking in a Prometheus Table the query is filtered by clicked value. [#17083](https://github.com/grafana/grafana/pull/17083), [@hugohaggmark](https://github.com/hugohaggmark)
- **Explore**: Makes it possible to zoom in Explore/Loki/Graph without exception. [#16991](https://github.com/grafana/grafana/pull/16991), [@hugohaggmark](https://github.com/hugohaggmark)
- **Gauge**: Fixes orientation issue after switching from BarGauge to Gauge. [#17064](https://github.com/grafana/grafana/pull/17064), [@torkelo](https://github.com/torkelo)
- **GettingStarted**: Fixes layout issues in getting started panel. [#16941](https://github.com/grafana/grafana/pull/16941), [@torkelo](https://github.com/torkelo)
- **InfluxDB**: Fix HTTP method should default to GET. [#16949](https://github.com/grafana/grafana/pull/16949), [@StephenSorriaux](https://github.com/StephenSorriaux)
- **Panels**: Fixed alert icon position in panel header. [#17070](https://github.com/grafana/grafana/pull/17070), [@torkelo](https://github.com/torkelo)
- **Panels**: Fixes panel error tooltip not showing. [#16993](https://github.com/grafana/grafana/pull/16993), [@torkelo](https://github.com/torkelo)
- **Plugins**: Fix how datemath utils are exposed to plugins. [#16976](https://github.com/grafana/grafana/pull/16976), [@marefr](https://github.com/marefr)
- **Singlestat**: fixed centering issue for very small panels. [#16944](https://github.com/grafana/grafana/pull/16944), [@torkelo](https://github.com/torkelo)
- **Search**: Scroll issue in dashboard search in latest Chrome. [#17054](https://github.com/grafana/grafana/pull/17054), [@jschill](https://github.com/jschill)
- **Docker**: Prevent a permission denied error when writing files to the default provisioning directory. [#16831](https://github.com/grafana/grafana/pull/16831), [@wmedlar](https://github.com/wmedlar)
- **Gauge**: Adds background shade to gauge track and improves height usage. [#17019](https://github.com/grafana/grafana/pull/17019), [@torkelo](https://github.com/torkelo)
- **RemoteCache**: Avoid race condition in Set causing error on insert. . [#17082](https://github.com/grafana/grafana/pull/17082), [@bergquist](https://github.com/bergquist)

# 6.2.0-beta1 (2019-05-07)

### Features / Enhancements

- **Admin**: Add more stats about roles. [#16667](https://github.com/grafana/grafana/pull/16667), [@bergquist](https://github.com/bergquist)
- **Alert list panel**: Support variables in filters. [#16892](https://github.com/grafana/grafana/pull/16892), [@psschand](https://github.com/psschand)
- **Alerting**: Adjust label for send on all alerts to default . [#16554](https://github.com/grafana/grafana/pull/16554), [@simPod](https://github.com/simPod)
- **Alerting**: Makes timeouts and retries configurable. [#16259](https://github.com/grafana/grafana/pull/16259), [@kobehaha](https://github.com/kobehaha)
- **Alerting**: No notification when going from no data to pending. [#16905](https://github.com/grafana/grafana/pull/16905), [@bergquist](https://github.com/bergquist)
- **Alerting**: Pushover alert, support for different sound for OK. [#16525](https://github.com/grafana/grafana/pull/16525), [@Hofls](https://github.com/Hofls)
- **Auth**: Enable retries and transaction for some db calls for auth tokens . [#16785](https://github.com/grafana/grafana/pull/16785), [@bergquist](https://github.com/bergquist)
- **AzureMonitor**: Adds support for multiple subscriptions per data source. [#16922](https://github.com/grafana/grafana/pull/16922), [@daniellee](https://github.com/daniellee)
- **Bar Gauge**: New multi series enabled gauge like panel with horizontal and vertical layouts and 3 display modes. [#16918](https://github.com/grafana/grafana/pull/16918), [@torkelo](https://github.com/torkelo)
- **Build**: Upgrades to golang 1.12.4. [#16545](https://github.com/grafana/grafana/pull/16545), [@bergquist](https://github.com/bergquist)
- **CloudWatch**: Update AWS/IoT metric and dimensions. [#16337](https://github.com/grafana/grafana/pull/16337), [@nonamef](https://github.com/nonamef)
- **Config**: Show user-friendly error message instead of stack trace. [#16564](https://github.com/grafana/grafana/pull/16564), [@Hofls](https://github.com/Hofls)
- **Dashboard**: Enable filtering dashboards in search by current folder. [#16790](https://github.com/grafana/grafana/pull/16790), [@dprokop](https://github.com/dprokop)
- **Dashboard**: Lazy load out of view panels . [#15554](https://github.com/grafana/grafana/pull/15554), [@ryantxu](https://github.com/ryantxu)
- **DataProxy**: Restore Set-Cookie header after proxy request. [#16838](https://github.com/grafana/grafana/pull/16838), [@marefr](https://github.com/marefr)
- **Data Sources**: Add pattern validation for time input on data source config pages. [#16837](https://github.com/grafana/grafana/pull/16837), [@aocenas](https://github.com/aocenas)
- **Elasticsearch**: Add 7.x version support. [#16646](https://github.com/grafana/grafana/pull/16646), [@alcidesv](https://github.com/alcidesv)
- **Explore**: Adds reconnect for failing data source. [#16226](https://github.com/grafana/grafana/pull/16226), [@hugohaggmark](https://github.com/hugohaggmark)
- **Explore**: Support user timezone. [#16469](https://github.com/grafana/grafana/pull/16469), [@marefr](https://github.com/marefr)
- **InfluxDB**: Add support for POST HTTP verb. [#16690](https://github.com/grafana/grafana/pull/16690), [@StephenSorriaux](https://github.com/StephenSorriaux)
- **Loki**: Search is now case insensitive. [#15948](https://github.com/grafana/grafana/pull/15948), [@steven-sheehy](https://github.com/steven-sheehy)
- **OAuth**: Update jwt regexp to include `=`. [#16521](https://github.com/grafana/grafana/pull/16521), [@DanCech](https://github.com/DanCech)
- **Panels**: No title will no longer make panel header take up space. [#16884](https://github.com/grafana/grafana/pull/16884), [@torkelo](https://github.com/torkelo)
- **Prometheus**: Adds tracing headers for Prometheus datasource. [#16724](https://github.com/grafana/grafana/pull/16724), [@svagner](https://github.com/svagner)
- **Provisioning**: Add API endpoint to reload provisioning configs. [#16579](https://github.com/grafana/grafana/pull/16579), [@aocenas](https://github.com/aocenas)
- **Provisioning**: Do not allow deletion of provisioned dashboards. [#16211](https://github.com/grafana/grafana/pull/16211), [@aocenas](https://github.com/aocenas)
- **Provisioning**: Interpolate env vars in provisioning files. [#16499](https://github.com/grafana/grafana/pull/16499), [@aocenas](https://github.com/aocenas)
- **Provisioning**: Support FolderUid in Dashboard Provisioning Config. [#16559](https://github.com/grafana/grafana/pull/16559), [@swtch1](https://github.com/swtch1)
- **Security**: Add new setting allow_embedding. [#16853](https://github.com/grafana/grafana/pull/16853), [@marefr](https://github.com/marefr)
- **Security**: Store data source passwords encrypted in secureJsonData. [#16175](https://github.com/grafana/grafana/pull/16175), [@aocenas](https://github.com/aocenas)
- **UX**: Improve Grafana usage for smaller screens. [#16783](https://github.com/grafana/grafana/pull/16783), [@torkelo](https://github.com/torkelo)
- **Units**: Add angle units, Arc Minutes and Seconds. [#16271](https://github.com/grafana/grafana/pull/16271), [@Dripoul](https://github.com/Dripoul)

### Bug Fixes

- **Build**: Fix bug where grafana didn't start after mysql on rpm packages. [#16917](https://github.com/grafana/grafana/pull/16917), [@bergquist](https://github.com/bergquist)
- **CloudWatch**: Fixes query order not affecting series ordering & color. [#16408](https://github.com/grafana/grafana/pull/16408), [@mtanda](https://github.com/mtanda)
- **CloudWatch**: Use default alias if there is no alias for metrics. [#16732](https://github.com/grafana/grafana/pull/16732), [@utkarshcmu](https://github.com/utkarshcmu)
- **Config**: Fixes bug where timeouts for alerting was not parsed correctly. [#16784](https://github.com/grafana/grafana/pull/16784), [@aocenas](https://github.com/aocenas)
- **Elasticsearch**: Fix view percentiles metric in table without date histogram. [#15686](https://github.com/grafana/grafana/pull/15686), [@Igor-Ratsuk](https://github.com/Igor-Ratsuk)
- **Explore**: Prevents histogram loading from killing Prometheus instance. [#16768](https://github.com/grafana/grafana/pull/16768), [@hugohaggmark](https://github.com/hugohaggmark)
- **Graph**: Allow override decimals to fully override. [#16414](https://github.com/grafana/grafana/pull/16414), [@torkelo](https://github.com/torkelo)
- **Mixed Data Source**: Fix error when one query is disabled. [#16409](https://github.com/grafana/grafana/pull/16409), [@marefr](https://github.com/marefr)
- **Search**: Fixes search limits and adds a page parameter. [#16458](https://github.com/grafana/grafana/pull/16458), [@torkelo](https://github.com/torkelo)
- **Security**: Responses from backend should not be cached. [#16848](https://github.com/grafana/grafana/pull/16848), [@marefr](https://github.com/marefr)

### Breaking changes

- **Gauge Panel**: The suffix / prefix options have been removed from the new Gauge Panel (introduced in v6.0). [#16870](https://github.com/grafana/grafana/issues/16870).

# 6.1.6 (2019-04-29)

### Features / Enhancements

- **Security**: Bump jQuery to 3.4.0 . [#16761](https://github.com/grafana/grafana/pull/16761), [@dprokop](https://github.com/dprokop)

### Bug Fixes

- **Playlist**: Fix loading dashboards by tag. [#16727](https://github.com/grafana/grafana/pull/16727), [@marefr](https://github.com/marefr)

# 6.1.5 (2019-04-29)

- **Security**: Urgent security patch release. Please read more in our [blog](https://grafana.com/blog/2019/04/29/grafana-5.4.4-and-6.1.6-released-with-important-security-fix/)

# 6.1.4 (2019-04-16)

### Bug Fixes

- **DataPanel**: Added missing built-in interval variables to scopedVars. [#16556](https://github.com/grafana/grafana/pull/16556), [@torkelo](https://github.com/torkelo)
- **Explore**: Adds maxDataPoints to data source query options . [#16513](https://github.com/grafana/grafana/pull/16513), [@hugohaggmark](https://github.com/hugohaggmark)
- **Explore**: Fixes so intervals are recalculated on run query. [#16510](https://github.com/grafana/grafana/pull/16510), [@hugohaggmark](https://github.com/hugohaggmark)
- **Heatmap**: Fix for empty graph when panel is too narrow (#16378). [#16460](https://github.com/grafana/grafana/pull/16460), [@alexanderzobnin](https://github.com/alexanderzobnin)
- **Heatmap**: Fixed auto decimals when bucket name is not number. [#16609](https://github.com/grafana/grafana/pull/16609), [@torkelo](https://github.com/torkelo)
- **QueryInspector**: Now shows error responses again. [#16514](https://github.com/grafana/grafana/pull/16514), [@torkelo](https://github.com/torkelo)

# 6.1.3 (2019-04-09)

### Bug Fixes

- **Graph**: Fixed auto decimals in legend values for some units like `ms` and `s`. [#16455](https://github.com/grafana/grafana/pull/16455), [@torkelo](https://github.com/torkelo)
- **Graph**: Fixed png rendering with legend to the right. [#16463](https://github.com/grafana/grafana/pull/16463), [@torkelo](https://github.com/torkelo)
- **Singlestat**: Use decimals when manually specified. [#16451](https://github.com/grafana/grafana/pull/16451), [@torkelo](https://github.com/torkelo)
- **UI Switch**: Fix broken UI switches. Fixes Default Data Source switch, Explore Logs switches, Gauge option switches. [#16303](https://github.com/grafana/grafana/pull/16303), [@dprokop](https://github.com/dprokop)

# 6.1.2 (2019-04-08)

### Bug Fixes

- **Graph**: Fixed series legend color for hidden series. [#16438](https://github.com/grafana/grafana/pull/16438), [@Ijin08](https://github.com/Ijin08)
- **Graph**: Fixed tooltip highlight on white theme. [#16429](https://github.com/grafana/grafana/pull/16429), [@torkelo](https://github.com/torkelo)
- **Styles**: Fixed menu hover highlight border. [#16431](https://github.com/grafana/grafana/pull/16431), [@torkelo](https://github.com/torkelo)
- **Singlestat Panel**: Correctly use the override decimals. [#16413](https://github.com/grafana/grafana/pull/16413), [@torkelo](https://github.com/torkelo)

# 6.1.1 (2019-04-05)

### Bug Fixes

- **Alerting**: Notification channel http api fixes. [#16379](https://github.com/grafana/grafana/pull/16379), [@marefr](https://github.com/marefr)
- **Graphite**: Editing graphite query function now works again. [#16390](https://github.com/grafana/grafana/pull/16390), [@torkelo](https://github.com/torkelo)
- **Playlist**: Kiosk & auto fit panels modes are working normally again . [#16403](https://github.com/grafana/grafana/pull/16403), [@torkelo](https://github.com/torkelo)
- **QueryEditors**: Toggle edit mode now always work on slower computers. [#16394](https://github.com/grafana/grafana/pull/16394), [@seanlaff](https://github.com/seanlaff)

# 6.1.0 (2019-04-03)

### Bug Fixes

- **CloudWatch**: Fix for dimension value list when changing dimension key. [#16356](https://github.com/grafana/grafana/pull/16356), [@mtanda](https://github.com/mtanda)
- **Graphite**: Editing function arguments now works again. [#16297](https://github.com/grafana/grafana/pull/16297), [@torkelo](https://github.com/torkelo)
- **InfluxDB**: Fix tag names with periods in alert evaluation. [#16255](https://github.com/grafana/grafana/pull/16255), [@floyd-may](https://github.com/floyd-may)
- **PngRendering**: Fix for panel height & title centering . [#16351](https://github.com/grafana/grafana/pull/16351), [@torkelo](https://github.com/torkelo)
- **Templating**: Fix for editing query variables. [#16299](https://github.com/grafana/grafana/pull/16299), [@torkelo](https://github.com/torkelo)

# 6.1.0-beta1 (2019-03-27)

### New Features

- **Prometheus**: adhoc filter support [#8253](https://github.com/grafana/grafana/issues/8253), thx [@mtanda](https://github.com/mtanda)
- **Permissions**: Editors can become admin for dashboards, folders and teams they create. [#15977](https://github.com/grafana/grafana/pull/15977), [@xlson](https://github.com/xlson)

### Minor

- **Auth**: Support listing and revoking auth tokens via API [#15836](https://github.com/grafana/grafana/issues/15836)
- **Alerting**: DingDing notification channel now includes alert values. [#13825](https://github.com/grafana/grafana/pull/13825), [@athurg](https://github.com/athurg)
- **Alerting**: Notification channel http api enhancements. [#16219](https://github.com/grafana/grafana/pull/16219), [@marefr](https://github.com/marefr)
- **CloudWatch**: Update metrics/dimensions list. [#16137](https://github.com/grafana/grafana/pull/16137), [@mtanda](https://github.com/mtanda)
- **Cloudwatch**: Add AWS RDS MaximumUsedTransactionIDs metric [#15077](https://github.com/grafana/grafana/pull/15077), thx [@activeshadow](https://github.com/activeshadow)
- **Cache**: Adds support for using out of proc caching in the backend [#10816](https://github.com/grafana/grafana/issues/10816)
- **Dashboard**: New keyboard shortcut `d l` toggles all Graph legends in a dashboard. [#15770](https://github.com/grafana/grafana/pull/15770), [@jsferrei](https://github.com/jsferrei)
- **Data Source**: Only log connection string in dev environment [#16001](https://github.com/grafana/grafana/issues/16001)
- **DataProxy**: Add custom header (X-Grafana-User) to data source requests with the current username. [#15998](https://github.com/grafana/grafana/pull/15998), [@aocenas](https://github.com/aocenas)
- **DataProxy**: Make it possible to add user details to requests sent to the dataproxy [#6359](https://github.com/grafana/grafana/issues/6359) and [#15931](https://github.com/grafana/grafana/issues/15931)
- **DataProxy**: Adds oauth pass-through option for data sources. [#15205](https://github.com/grafana/grafana/pull/15205), [@seanlaff](https://github.com/seanlaff)
- **Explore**: Hide empty duplicates column in logs viewer. [#15982](https://github.com/grafana/grafana/pull/15982), [@steven-sheehy](https://github.com/steven-sheehy)
- **Explore**: Make it possible to close left pane of split view. [#16155](https://github.com/grafana/grafana/pull/16155), [@dprokop](https://github.com/dprokop)
- **Explore**: Move back / forward with browser buttons now works. [#16150](https://github.com/grafana/grafana/pull/16150), [@hugohaggmark](https://github.com/hugohaggmark)
- **Explore**: Update Loki labels when label selector is opened. [#16131](https://github.com/grafana/grafana/pull/16131), [@dprokop](https://github.com/dprokop)
- **Graph Panel**: New options for X-axis Min & Max (for histograms). [#14877](https://github.com/grafana/grafana/pull/14877), [@papagian](https://github.com/papagian)
- **Heatmap**: You can now choose to hide buckets with zero value. [#15934](https://github.com/grafana/grafana/pull/15934), [@alexanderzobnin](https://github.com/alexanderzobnin)
- **Heatmap**: `Middle` bucket bound option [#15683](https://github.com/grafana/grafana/issues/15683)
- **Heatmap**: `Reverse order` option for changing order of buckets [#15683](https://github.com/grafana/grafana/issues/15683)
- **Prometheus**: Change alignment of range queries to end before now and not in future. [#16110](https://github.com/grafana/grafana/pull/16110), [@davkal](https://github.com/davkal)
- **Prometheus**: Dedup annotations events with same timestamp . [#16152](https://github.com/grafana/grafana/pull/16152), [@torkelo](https://github.com/torkelo)
- **SQL**: Use default min interval of 1m for all SQL data sources. [#15799](https://github.com/grafana/grafana/pull/15799), [@marefr](https://github.com/marefr)
- **TablePanel**: Column color style now works even after removing columns. [#16227](https://github.com/grafana/grafana/pull/16227), [@torkelo](https://github.com/torkelo)
- **Templating**: Custom variable value now escapes all backslashes properly. [#15980](https://github.com/grafana/grafana/pull/15980), [@srid12](https://github.com/srid12)
- **Templating**: Data source variable now supports multi-value for uses cases that involve repeating panels & rows. [#15914](https://github.com/grafana/grafana/pull/15914), [@torkelo](https://github.com/torkelo)
- **VictorOps**: Adds more information to the victor ops notifiers [#15744](https://github.com/grafana/grafana/issues/15744), thx [@zhulongcheng](https://github.com/zhulongcheng)

### Bug Fixes

- **Alerting**: Don't include non-existing image in MS Teams notifications. [#16116](https://github.com/grafana/grafana/pull/16116), [@SGI495](https://github.com/SGI495)
- **Api**: Invalid org invite code [#10506](https://github.com/grafana/grafana/issues/10506)
- **Annotations**: Fix for native annotations filtered by template variable with pipe. [#15515](https://github.com/grafana/grafana/pull/15515), [@marefr](https://github.com/marefr)
- **Dashboard**: Fix for time regions spanning across midnight. [#16201](https://github.com/grafana/grafana/pull/16201), [@marefr](https://github.com/marefr)
- **Data Source**: Handles nil jsondata field gracefully [#14239](https://github.com/grafana/grafana/issues/14239)
- **Data Source**: Empty user/password was not updated when updating data sources [#15608](https://github.com/grafana/grafana/pull/15608), thx [@Maddin-619](https://github.com/Maddin-619)
- **Elasticsearch**: Fixes using template variables in the alias field. [#16229](https://github.com/grafana/grafana/pull/16229), [@daniellee](https://github.com/daniellee)
- **Elasticsearch**: Fix incorrect index pattern padding in alerting queries. [#15892](https://github.com/grafana/grafana/pull/15892), [@sandlis](https://github.com/sandlis)
- **Explore**: Fix for Prometheus autocomplete not working in Firefox. [#16192](https://github.com/grafana/grafana/pull/16192), [@hugohaggmark](https://github.com/hugohaggmark)
- **Explore**: Fix for url does not keep query after browser refresh. [#16189](https://github.com/grafana/grafana/pull/16189), [@hugohaggmark](https://github.com/hugohaggmark)
- **Gauge**: Interpolate scoped variables in repeated gauges [#15739](https://github.com/grafana/grafana/issues/15739)
- **Graphite**: Fixed issue with using series ref and series by tag. [#16111](https://github.com/grafana/grafana/pull/16111), [@torkelo](https://github.com/torkelo)
- **Graphite**: Fixed variable quoting when variable value is numeric. [#16149](https://github.com/grafana/grafana/pull/16149), [@torkelo](https://github.com/torkelo)
- **Heatmap**: Fixes Y-axis tick labels being in wrong order for some Prometheus queries. [#15932](https://github.com/grafana/grafana/pull/15932), [@alexanderzobnin](https://github.com/alexanderzobnin)
- **Heatmap**: Negative values are now displayed correctly in graph & legend. [#15953](https://github.com/grafana/grafana/pull/15953), [@alexanderzobnin](https://github.com/alexanderzobnin)
- **Heatmap**: legend shows wrong colors for small values [#14019](https://github.com/grafana/grafana/issues/14019)
- **InfluxDB**: Always close request body even for error status codes. [#16207](https://github.com/grafana/grafana/pull/16207), [@ramongtx](https://github.com/ramongtx)
- **ManageDashboards**: Fix for checkboxes not appearing properly Firefox . [#15981](https://github.com/grafana/grafana/pull/15981), [@srid12](https://github.com/srid12)
- **Playlist**: Leaving playlist now always stops playlist . [#15791](https://github.com/grafana/grafana/pull/15791), [@peterholmberg](https://github.com/peterholmberg)
- **Prometheus**: fixes regex ad-hoc filters variables with wildcards. [#16234](https://github.com/grafana/grafana/pull/16234), [@daniellee](https://github.com/daniellee)
- **TablePanel**: Column color style now works even after removing columns. [#16227](https://github.com/grafana/grafana/pull/16227), [@torkelo](https://github.com/torkelo)
- **TablePanel**: Fix for white text on white background when value is null. [#16199](https://github.com/grafana/grafana/pull/16199), [@peterholmberg](https://github.com/peterholmberg)

# 6.0.2 (2019-03-19)

### Bug Fixes

- **Alerting**: Fixed issue with AlertList panel links resulting in panel not found errors. [#15975](https://github.com/grafana/grafana/pull/15975), [@torkelo](https://github.com/torkelo)
- **Dashboard**: Improved error handling when rendering dashboard panels. [#15970](https://github.com/grafana/grafana/pull/15970), [@torkelo](https://github.com/torkelo)
- **LDAP**: Fix allow anonymous server bind for ldap search. [#15872](https://github.com/grafana/grafana/pull/15872), [@marefr](https://github.com/marefr)
- **Discord**: Fix discord notifier so it doesn't crash when there are no image generated. [#15833](https://github.com/grafana/grafana/pull/15833), [@marefr](https://github.com/marefr)
- **Panel Edit**: Prevent search in VizPicker from stealing focus. [#15802](https://github.com/grafana/grafana/pull/15802), [@peterholmberg](https://github.com/peterholmberg)
- **Data Source admin**: Fixed url of back button in data source edit page, when root_url configured. [#15759](https://github.com/grafana/grafana/pull/15759), [@dprokop](https://github.com/dprokop)

# 6.0.1 (2019-03-06)

### Bug Fixes

- **Metrics**: Fixes broken usagestats metrics for /metrics [#15651](https://github.com/grafana/grafana/issues/15651)
- **Dashboard**: Fixes kiosk mode should have &kiosk appended to the url [#15765](https://github.com/grafana/grafana/issues/15765)
- **Dashboard**: Fixes kiosk=tv mode with autofitpanels should respect header [#15650](https://github.com/grafana/grafana/issues/15650)
- **Image rendering**: Fixed image rendering issue for dashboards with auto refresh, . [#15818](https://github.com/grafana/grafana/pull/15818), [@torkelo](https://github.com/torkelo)
- **Dashboard**: Fix only users that can edit a dashboard should be able to update panel json. [#15805](https://github.com/grafana/grafana/pull/15805), [@marefr](https://github.com/marefr)
- **LDAP**: fix allow anonymous initial bind for ldap search. [#15803](https://github.com/grafana/grafana/pull/15803), [@marefr](https://github.com/marefr)
- **UX**: Fixed scrollbar not visible initially (only after manual scroll). [#15798](https://github.com/grafana/grafana/pull/15798), [@torkelo](https://github.com/torkelo)
- **Data Source admin** TestData [#15793](https://github.com/grafana/grafana/pull/15793), [@hugohaggmark](https://github.com/hugohaggmark)
- **Dashboard**: Fixed scrolling issue that caused scroll to be locked to bottom. [#15792](https://github.com/grafana/grafana/pull/15792), [@torkelo](https://github.com/torkelo)
- **Explore**: Viewers with viewers_can_edit should be able to access /explore. [#15787](https://github.com/grafana/grafana/pull/15787), [@jschill](https://github.com/jschill)
- **Security** fix: limit access to org admin and alerting pages. [#15761](https://github.com/grafana/grafana/pull/15761), [@marefr](https://github.com/marefr)
- **Panel Edit** minInterval changes did not persist [#15757](https://github.com/grafana/grafana/pull/15757), [@hugohaggmark](https://github.com/hugohaggmark)
- **Teams**: Fixed bug when getting teams for user. [#15595](https://github.com/grafana/grafana/pull/15595), [@hugohaggmark](https://github.com/hugohaggmark)
- **Stackdriver**: fix for float64 bounds for distribution metrics [#14509](https://github.com/grafana/grafana/issues/14509)
- **Stackdriver**: no reducers available for distribution type [#15179](https://github.com/grafana/grafana/issues/15179)

# 6.0.0 stable (2019-02-25)

### Bug Fixes

- **Dashboard**: fixes click after scroll in series override menu [#15621](https://github.com/grafana/grafana/issues/15621)
- **MySQL**: fix mysql query using \_interval_ms variable throws error [#14507](https://github.com/grafana/grafana/issues/14507)

# 6.0.0-beta3 (2019-02-19)

### Minor

- **CLI**: Grafana CLI should preserve permissions for backend binaries for Linux and Darwin [#15500](https://github.com/grafana/grafana/issues/15500)
- **Alerting**: Allow image rendering 90 percent of alertTimeout [#15395](https://github.com/grafana/grafana/pull/15395)

### Bug fixes

- **Influxdb**: Add support for alerting on InfluxDB queries that use the non_negative_difference function [#15415](https://github.com/grafana/grafana/issues/15415), thx [@kiran3394](https://github.com/kiran3394)
- **Alerting**: Fix percent_diff calculation when points are nulls [#15443](https://github.com/grafana/grafana/issues/15443), thx [@max-neverov](https://github.com/max-neverov)
- **Alerting**: Fixed handling of alert urls with true flags [#15454](https://github.com/grafana/grafana/issues/15454)

# 6.0.0-beta2 (2019-02-11)

### New Features

- **AzureMonitor**: Enable alerting by converting Azure Monitor API to Go [#14623](https://github.com/grafana/grafana/issues/14623)

### Minor

- **Alerting**: Adds support for images in pushover notifier [#10780](https://github.com/grafana/grafana/issues/10780), thx [@jpenalbae](https://github.com/jpenalbae)
- **Graphite/InfluxDB/OpenTSDB**: Fix always take dashboard timezone into consideration when handle custom time ranges [#15284](https://github.com/grafana/grafana/issues/15284)
- **Stackdriver**: Template variables in filters using globbing format [#15182](https://github.com/grafana/grafana/issues/15182)
- **Cloudwatch**: Add `resource_arns` template variable query function [#8207](https://github.com/grafana/grafana/issues/8207), thx [@jeroenvollenbrock](https://github.com/jeroenvollenbrock)
- **Cloudwatch**: Add AWS/Neptune metrics [#14231](https://github.com/grafana/grafana/issues/14231), thx [@tcpatterson](https://github.com/tcpatterson)
- **Cloudwatch**: Add AWS/EC2/API metrics [#14233](https://github.com/grafana/grafana/issues/14233), thx [@tcpatterson](https://github.com/tcpatterson)
- **Cloudwatch**: Add AWS RDS ServerlessDatabaseCapacity metric [#15265](https://github.com/grafana/grafana/pull/15265), thx [@larsjoergensen](https://github.com/larsjoergensen)
- **MySQL**: Adds data source SSL CA/client certificates support [#8570](https://github.com/grafana/grafana/issues/8570), thx [@bugficks](https://github.com/bugficks)
- **MSSQL**: Timerange are now passed for template variable queries [#13324](https://github.com/grafana/grafana/issues/13324), thx [@thatsparesh](https://github.com/thatsparesh)
- **Annotations**: Support PATCH verb in annotations http api [#12546](https://github.com/grafana/grafana/issues/12546), thx [@SamuelToh](https://github.com/SamuelToh)
- **Templating**: Add json formatting to variable interpolation [#15291](https://github.com/grafana/grafana/issues/15291), thx [@mtanda](https://github.com/mtanda)
- **Login**: Anonymous usage stats for token auth [#15288](https://github.com/grafana/grafana/issues/15288)
- **AzureMonitor**: improve autocomplete for Log Analytics and App Insights editor [#15131](https://github.com/grafana/grafana/issues/15131)
- **LDAP**: Fix IPA/FreeIPA v4.6.4 does not allow LDAP searches with empty attributes [#14432](https://github.com/grafana/grafana/issues/14432)
- **Provisioning**: Allow testing data sources that were added by config [#12164](https://github.com/grafana/grafana/issues/12164)
- **Security**: Fix CSRF Token validation for POSTs [#1441](https://github.com/grafana/grafana/issues/1441)

### Breaking changes

- **Internal Metrics** Edition has been added to the build_info metric. This will break any Graphite queries using this metric. Edition will be a new label for the Prometheus metric. [#15363](https://github.com/grafana/grafana/pull/15363)

### Bug fixes

- **Gauge**: Fix issue with gauge requests being cancelled [#15366](https://github.com/grafana/grafana/issues/15366)
- **Gauge**: Accept decimal inputs for thresholds [#15372](https://github.com/grafana/grafana/issues/15372)
- **UI**: Fix error caused by named colors that are not part of named colors palette [#15373](https://github.com/grafana/grafana/issues/15373)
- **Search**: Bug pressing special regexp chars in input fields [#12972](https://github.com/grafana/grafana/issues/12972)
- **Permissions**: No need to have edit permissions to be able to "Save as" [#13066](https://github.com/grafana/grafana/issues/13066)

# 6.0.0-beta1 (2019-01-30)

### New Features

- **Alerting**: Adds support for Google Hangouts Chat notifications [#11221](https://github.com/grafana/grafana/issues/11221), thx [@PatrickSchuster](https://github.com/PatrickSchuster)
- **Elasticsearch**: Support bucket script pipeline aggregations [#5968](https://github.com/grafana/grafana/issues/5968)
- **Influxdb**: Add support for time zone (`tz`) clause [#10322](https://github.com/grafana/grafana/issues/10322), thx [@cykl](https://github.com/cykl)
- **Snapshots**: Enable deletion of public snapshot [#14109](https://github.com/grafana/grafana/issues/14109)
- **Provisioning**: Provisioning support for alert notifiers [#10487](https://github.com/grafana/grafana/issues/10487), thx [@pbakulev](https://github.com/pbakulev)
- **Explore**: A whole new way to do ad-hoc metric queries and exploration. Split view in half and compare metrics & logs and much much more. [Read more here](http://docs.grafana.org/features/explore/)
- **Auth**: Replace remember me cookie solution for Grafana's builtin, LDAP and OAuth authentication with a solution based on short-lived tokens [#15303](https://github.com/grafana/grafana/issues/15303)

### Minor

- **Templating**: Built in time range variables `$__from` and `$__to`, [#1909](https://github.com/grafana/grafana/issues/1909)
- **Alerting**: Use separate timeouts for alert evals and notifications [#14701](https://github.com/grafana/grafana/issues/14701), thx [@sharkpc0813](https://github.com/sharkpc0813)
- **Elasticsearch**: Add support for offset in date histogram aggregation [#12653](https://github.com/grafana/grafana/issues/12653), thx [@mattiarossi](https://github.com/mattiarossi)
- **Elasticsearch**: Add support for moving average and derivative using doc count (metric count) [#8843](https://github.com/grafana/grafana/issues/8843) [#11175](https://github.com/grafana/grafana/issues/11175)
- **Elasticsearch**: Add support for template variable interpolation in alias field [#4075](https://github.com/grafana/grafana/issues/4075), thx [@SamuelToh](https://github.com/SamuelToh)
- **Influxdb**: Fix autocomplete of measurements does not escape search string properly [#11503](https://github.com/grafana/grafana/issues/11503), thx [@SamuelToh](https://github.com/SamuelToh)
- **Stackdriver**: Aggregating series returns more than one series [#14581](https://github.com/grafana/grafana/issues/14581) and [#13914](https://github.com/grafana/grafana/issues/13914), thx [@kinok](https://github.com/kinok)
- **Cloudwatch**: Fix Assume Role Arn [#14722](https://github.com/grafana/grafana/issues/14722), thx [@jaken551](https://github.com/jaken551)
- **Postgres/MySQL/MSSQL**: Nanosecond timestamp support (`$__unixEpochNanoFilter`, `$__unixEpochNanoFrom`, `$__unixEpochNanoTo`) [#14711](https://github.com/grafana/grafana/pull/14711), thx [@ander26](https://github.com/ander26)
- **Provisioning**: Fixes bug causing infinite growth in dashboard_version table. [#12864](https://github.com/grafana/grafana/issues/12864)
- **Auth**: Prevent password reset when login form is disabled or either LDAP or Auth Proxy is enabled [#14246](https://github.com/grafana/grafana/issues/14246), thx [@SilverFire](https://github.com/SilverFire)
- **Admin**: Fix prevent removing last grafana admin permissions [#11067](https://github.com/grafana/grafana/issues/11067), thx [@danielbh](https://github.com/danielbh)
- **Admin**: When multiple user invitations, all links are the same as the first user who was invited [#14483](https://github.com/grafana/grafana/issues/14483)
- **LDAP**: Upgrade go-ldap to v3 [#14548](https://github.com/grafana/grafana/issues/14548)
- **OAuth**: Support OAuth providers that are not RFC6749 compliant [#14562](https://github.com/grafana/grafana/issues/14562), thx [@tdabasinskas](https://github.com/tdabasinskas)
- **Proxy whitelist**: Add CIDR capability to auth_proxy whitelist [#14546](https://github.com/grafana/grafana/issues/14546), thx [@jacobrichard](https://github.com/jacobrichard)
- **Dashboard**: `Min width` changed to `Max per row` for repeating panels. This lets you specify the maximum number of panels to show per row and by that repeated panels will always take up full width of row [#12991](https://github.com/grafana/grafana/pull/12991), thx [@pgiraud](https://github.com/pgiraud)
- **Dashboard**: Retain decimal precision when exporting CSV [#13929](https://github.com/grafana/grafana/issues/13929), thx [@cinaglia](https://github.com/cinaglia)
- **Templating**: Escaping "Custom" template variables [#13754](https://github.com/grafana/grafana/issues/13754), thx [@IntegersOfK](https://github.com/IntegersOfK)
- **Templating**: Add percentencode formatting to variable interpolation to be used mainly for url escaping [#12764](https://github.com/grafana/grafana/issues/12764), thx [@cxcv](https://github.com/cxcv)
- **Units**: Add blood glucose level units mg/dL and mmol/L [#14519](https://github.com/grafana/grafana/issues/14519), thx [@kjedamzik](https://github.com/kjedamzik)
- **Units**: Add Floating Point Operations per Second units [#14558](https://github.com/grafana/grafana/pull/14558), thx [@hahnjo](https://github.com/hahnjo)
- **Table**: Renders epoch string as date if date column style [#14484](https://github.com/grafana/grafana/issues/14484)
- **Dataproxy**: Override incoming Authorization header [#13815](https://github.com/grafana/grafana/issues/13815), thx [@kornholi](https://github.com/kornholi)
- **Dataproxy**: Add global data source proxy timeout setting [#5699](https://github.com/grafana/grafana/issues/5699), thx [@RangerRick](https://github.com/RangerRick)
- **Database**: Support specifying database host using IPV6 for backend database and sql data sources [#13711](https://github.com/grafana/grafana/issues/13711), thx [@ellisvlad](https://github.com/ellisvlad)
- **Database**: Support defining additional database connection string args when using `url` property in database settings [#14709](https://github.com/grafana/grafana/pull/14709), thx [@tpetr](https://github.com/tpetr)
- **Stackdriver**: crossSeriesAggregation not being sent with the query [#15129](https://github.com/grafana/grafana/issues/15129), thx [@Legogris](https://github.com/Legogris)

### Bug fixes

- **Search**: Fix for issue with scrolling the "tags filter" dropdown, fixes [#14486](https://github.com/grafana/grafana/issues/14486)
- **Prometheus**: Query for annotation always uses 60s step regardless of dashboard range, fixes [#14795](https://github.com/grafana/grafana/issues/14795)
- **Annotations**: Fix creating annotation when graph panel has no data points position the popup outside viewport [#13765](https://github.com/grafana/grafana/issues/13765), thx [@banjeremy](https://github.com/banjeremy)
- **Piechart/Flot**: Fixes multiple piechart instances with donut bug [#15062](https://github.com/grafana/grafana/pull/15062)
- **Postgres**: Fix default port not added when port not configured [#15189](https://github.com/grafana/grafana/issues/15189)
- **Alerting**: Fixes crash bug when alert notifier folders are missing [#15295](https://github.com/grafana/grafana/issues/15295)
- **Dashboard**: Fix save provisioned dashboard modal [#15219](https://github.com/grafana/grafana/pull/15219)
- **Dashboard**: Fix having a long query in prometheus dashboard query editor blocks 30% of the query field when on OSX and having native scrollbars [#15122](https://github.com/grafana/grafana/issues/15122)
- **Explore**: Fix issue with wrapping on long queries [#15222](https://github.com/grafana/grafana/issues/15222)
- **Explore**: Fix cut & paste adds newline before and after selection [#15223](https://github.com/grafana/grafana/issues/15223)
- **Dataproxy**: Fix global data source proxy timeout not added to correct http client [#15258](https://github.com/grafana/grafana/issues/15258) [#5699](https://github.com/grafana/grafana/issues/5699)

### Breaking changes

- **Text Panel**: The text panel does no longer by default allow unsanitized HTML. [#4117](https://github.com/grafana/grafana/issues/4117). This means that if you have text panels with scripts tags they will no longer work as before. To enable unsafe javascript execution in text panels enable the settings `disable_sanitize_html` under the section `[panels]` in your Grafana ini file, or set env variable `PL_PANELS_DISABLE_SANITIZE_HTML=true`.
- **Dashboard**: Panel property `minSpan` replaced by `maxPerRow`. Dashboard migration will automatically migrate all dashboard panels using the `minSpan` property to the new `maxPerRow` property [#12991](https://github.com/grafana/grafana/pull/12991)

For older release notes, refer to the [CHANGELOG_ARCHIVE.md](https://github.com/grafana/grafana/blob/master/CHANGELOG_ARCHIVE.md)
