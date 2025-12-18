module github.com/credativ/plutono

go 1.24.3

// Override xorm's outdated go-mssqldb dependency, since we can't upgrade to current xorm (due to breaking changes).
// We need a more current go-mssqldb so we get rid of a version of apache/thrift with vulnerabilities.
// Also, use our fork with fixes for unimplemented methods (required for Go 1.16).
replace github.com/denisenkom/go-mssqldb => github.com/grafana/go-mssqldb v0.0.0-20210326084033-d0ce3c521036

// Override k8s.io/client-go outdated dependency, which is an indirect dependency of credativ/vali.
// It's also present on credativ/vali's go.mod so we'll need till it gets updated.
replace k8s.io/client-go => k8s.io/client-go v0.18.8

require (
	cloud.google.com/go/storage v1.13.0
	github.com/BurntSushi/toml v0.3.1
	github.com/VividCortex/mysqlerr v0.0.0-20170204212430-6c6b55f8796f
	github.com/aws/aws-sdk-go v1.37.20
	github.com/beevik/etree v1.4.0
	github.com/benbjohnson/clock v1.0.3
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b
	github.com/centrifugal/centrifuge v0.13.0
	github.com/cortexproject/cortex v1.8.1
	github.com/credativ/vali v0.0.0-20251218083720-5f1e0d214d00
	github.com/crewjam/saml v0.4.14
	github.com/davecgh/go-spew v1.1.1
	github.com/denisenkom/go-mssqldb v0.0.0-20200910202707-1e08a3fab204
	github.com/facebookgo/inject v0.0.0-20180706035515-f23751cae28b
	github.com/fatih/color v1.10.0
	github.com/gchaincl/sqlhooks v1.3.0
	github.com/getsentry/sentry-go v0.10.0
	github.com/go-macaron/binding v0.0.0-20190806013118-0b4f37bab25b
	github.com/go-macaron/gzip v0.0.0-20160222043647-cad1c6580a07
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-stack/stack v1.8.0
	github.com/gobwas/glob v0.2.3
	github.com/golang/mock v1.5.0
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.6.0
	github.com/google/uuid v1.2.0
	github.com/gosimple/slug v1.9.0
	github.com/grafana/grafana-aws-sdk v0.4.0
	github.com/grafana/grafana-plugin-model v0.0.0-20190930120109-1fc953a61fb4
	github.com/grafana/grafana-plugin-sdk-go v0.88.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/hashicorp/go-hclog v0.15.0
	github.com/hashicorp/go-plugin v1.4.0
	github.com/inconshreveable/log15 v0.0.0-20180818164646-67afb5ed74ec
	github.com/influxdata/influxdb-client-go/v2 v2.2.0
	github.com/jaegertracing/jaeger v1.22.1-0.20210304164023-2fff3ca58910
	github.com/jmespath/go-jmespath v0.4.0
	github.com/json-iterator/go v1.1.12
	github.com/jung-kurt/gofpdf v1.16.2
	github.com/lib/pq v1.9.0
	github.com/linkedin/goavro/v2 v2.10.0
	github.com/magefile/mage v1.11.0
	github.com/mattn/go-isatty v0.0.12
	github.com/mattn/go-sqlite3 v1.14.22
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f
	github.com/opentracing/opentracing-go v1.2.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.9.0
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.18.0
	github.com/robfig/cron v0.0.0-20180505203441-b41be1df6967
	github.com/robfig/cron/v3 v3.0.1
	github.com/russellhaering/goxmldsig v1.4.0
	github.com/smartystreets/goconvey v1.6.4
	github.com/stretchr/testify v1.8.1
	github.com/teris-io/shortid v0.0.0-20171029131806-771a37caa5cf
	github.com/timberio/go-datemath v0.1.1-0.20200323150745-74ddef604fff
	github.com/ua-parser/uap-go v0.0.0-20190826212731-daf92ba38329
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/unknwon/com v1.0.1
	github.com/urfave/cli/v2 v2.3.0
	github.com/weaveworks/common v0.0.0-20210112142934-23c8d7fa6120
	github.com/xorcare/pointer v1.1.0
	github.com/yudai/gojsondiff v1.0.0
	go.opentelemetry.io/collector v0.21.0
	golang.org/x/crypto v0.39.0
	golang.org/x/net v0.40.0
	golang.org/x/oauth2 v0.0.0-20210210192628-66670185b0cd
	golang.org/x/sync v0.15.0
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324
	gonum.org/v1/gonum v0.12.0
	google.golang.org/api v0.40.0
	google.golang.org/grpc v1.36.0
	gopkg.in/ini.v1 v1.62.0
	gopkg.in/ldap.v3 v3.1.0
	gopkg.in/macaron.v1 v1.4.0
	gopkg.in/mail.v2 v2.3.1
	gopkg.in/redis.v5 v5.2.9
	gopkg.in/square/go-jose.v2 v2.5.1
	gopkg.in/yaml.v2 v2.4.0
	xorm.io/core v0.7.3
	xorm.io/xorm v0.8.2
)

require (
	cloud.google.com/go v0.75.0 // indirect
	github.com/FZambia/eagle v0.0.1 // indirect
	github.com/FZambia/sentinel v1.1.0 // indirect
	github.com/apache/arrow/go/arrow v0.0.0-20200629181129-68b1273cbbf7 // indirect
	github.com/apache/thrift v0.14.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/centrifugal/protocol v0.3.4 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/cheekybits/genny v1.0.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/deepmap/oapi-codegen v1.3.13 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/edsrzf/mmap-go v1.0.0 // indirect
	github.com/facebookgo/ensure v0.0.0-20160127193407-b4ab57deab51 // indirect
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/structtag v0.0.0-20150214074306-217e25fb9691 // indirect
	github.com/facebookgo/subset v0.0.0-20150612182917-8dac2c3c4870 // indirect
	github.com/felixge/httpsnoop v1.0.1 // indirect
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/go-macaron/inject v0.0.0-20160627170012-d8a0b8677191 // indirect
	github.com/gogo/googleapis v1.4.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/gogo/status v1.0.3 // indirect
	github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/snappy v0.0.3-0.20201103224600-674baa8c7fc3 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/google/flatbuffers v1.11.0 // indirect
	github.com/googleapis/gax-go/v2 v2.0.5 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20191106031601-ce3c9ade29de // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.1-0.20191002090509-6af20e3a5340 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/hashicorp/yamux v0.0.0-20190923154419-df201c70410d // indirect
	github.com/igm/sockjs-go/v3 v3.0.0 // indirect
	github.com/influxdata/line-protocol v0.0.0-20200327222509-2487e7298839 // indirect
	github.com/jonboulle/clockwork v0.4.0 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/jstemmer/go-junit-report v0.9.1 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/klauspost/compress v1.11.7 // indirect
	github.com/mattermost/xml-roundtrip-validator v0.1.0 // indirect
	github.com/mattetti/filebuffer v1.0.1 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-runewidth v0.0.8 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/go-testing-interface v1.0.0 // indirect
	github.com/mna/redisc v1.1.7 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/olekukonko/tablewriter v0.0.4 // indirect
	github.com/opentracing-contrib/go-grpc v0.0.0-20210225150812-73cb765af46e // indirect
	github.com/opentracing-contrib/go-stdlib v1.0.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/node_exporter v1.0.0-rc.0.0.20200428091818-01054558c289 // indirect
	github.com/prometheus/procfs v0.2.0 // indirect
	github.com/prometheus/prometheus v1.8.2-0.20210215121130-6f488061dfb4 // indirect
	github.com/rainycape/unidecode v0.0.0-20150907023854-cb7f23ec59be // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/sercand/kuberesolver v2.4.0+incompatible // indirect
	github.com/sergi/go-diff v1.0.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/sirupsen/logrus v1.7.0 // indirect
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	github.com/weaveworks/promrus v1.2.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	go.opencensus.io v0.22.6 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/goleak v1.1.10 // indirect
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5 // indirect
	golang.org/x/mod v0.25.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	golang.org/x/tools v0.33.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20210203152818-3206188e46ba // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	xorm.io/builder v0.3.6 // indirect
)

replace github.com/apache/thrift => github.com/apache/thrift v0.14.1

replace gopkg.in/macaron.v1 v1.4.0 => ./pkg/macaron

replace github.com/go-macaron/binding => ./pkg/macaron/binding
