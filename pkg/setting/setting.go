// Copyright 2014 Unknwon
// Copyright 2014 Torkel Ödegaard

package setting

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/common/model"
	ini "gopkg.in/ini.v1"

	"github.com/credativ/plutono/pkg/components/gtime"
	"github.com/credativ/plutono/pkg/infra/log"
	"github.com/credativ/plutono/pkg/util"
	"github.com/grafana/grafana-aws-sdk/pkg/awsds"
)

type Scheme string

const (
	HTTPScheme   Scheme = "http"
	HTTPSScheme  Scheme = "https"
	HTTP2Scheme  Scheme = "h2"
	SocketScheme Scheme = "socket"
)

const (
	redactedPassword = "*********"
	DefaultHTTPAddr  = "0.0.0.0"
	Dev              = "development"
	Prod             = "production"
	Test             = "test"
)

// This constant corresponds to the default value for ldap_sync_ttl in .ini files
// it is used for comparison and has to be kept in sync
const (
	authProxySyncTTL = 60
)

// zoneInfo names environment variable for setting the path to look for the timezone database in go
const zoneInfo = "ZONEINFO"

var (
	// App settings.
	Env              = Dev
	AppUrl           string
	AppSubUrl        string
	ServeFromSubPath bool
	InstanceName     string

	// build
	BuildVersion    string
	BuildCommit     string
	BuildBranch     string
	BuildStamp      int64
	IsEnterprise    bool
	ApplicationName string

	// packaging
	Packaging = "unknown"

	// Paths
	HomePath       string
	CustomInitPath = "conf/custom.ini"

	// HTTP server options
	HttpAddr, HttpPort             string
	CertFile, KeyFile              string
	DataProxyLogging               bool
	DataProxyTimeout               int
	DataProxyTLSHandshakeTimeout   int
	DataProxyExpectContinueTimeout int
	DataProxyMaxConnsPerHost       int
	DataProxyMaxIdleConns          int
	DataProxyMaxIdleConnsPerHost   int
	DataProxyKeepAlive             int
	DataProxyIdleConnTimeout       int
	StaticRootPath                 string
	EnableGzip                     bool
	EnforceDomain                  bool

	// Security settings.
	SecretKey              string
	DisableGravatar        bool
	EmailCodeValidMinutes  int
	DataProxyWhiteList     map[string]bool
	CookieSecure           bool
	CookieSameSiteDisabled bool
	CookieSameSiteMode     http.SameSite

	// Snapshots
	ExternalSnapshotUrl   string
	ExternalSnapshotName  string
	ExternalEnabled       bool
	SnapShotRemoveExpired bool

	// Dashboard history
	DashboardVersionsToKeep int
	MinRefreshInterval      string

	// User settings
	AllowUserSignUp         bool
	AllowUserOrgCreate      bool
	AutoAssignOrg           bool
	AutoAssignOrgId         int
	AutoAssignOrgRole       string
	VerifyEmailEnabled      bool
	LoginHint               string
	PasswordHint            string
	DisableLoginForm        bool
	DisableSignoutMenu      bool
	SignoutRedirectUrl      string
	ExternalUserMngLinkUrl  string
	ExternalUserMngLinkName string
	ExternalUserMngInfo     string
	OAuthAutoLogin          bool
	ViewersCanEdit          bool

	// HTTP auth
	SigV4AuthEnabled bool

	AnonymousEnabled bool

	// Auth proxy settings
	AuthProxyEnabled        bool
	AuthProxyHeaderProperty string

	// Basic Auth
	BasicAuthEnabled bool

	// Global setting objects.
	Raw *ini.File

	// for logging purposes
	configFiles                  []string
	appliedCommandLineProperties []string
	appliedEnvOverrides          []string

	// analytics
	GoogleAnalyticsId  string
	GoogleTagManagerId string

	// LDAP
	LDAPEnabled           bool
	LDAPConfigFile        string
	LDAPSyncCron          string
	LDAPAllowSignup       bool
	LDAPActiveSyncEnabled bool

	// Quota
	Quota QuotaSettings

	// Alerting
	AlertingEnabled            bool
	ExecuteAlerts              bool
	AlertingRenderLimit        int
	AlertingErrorOrTimeout     string
	AlertingNoDataOrNullValues string

	AlertingEvaluationTimeout   time.Duration
	AlertingNotificationTimeout time.Duration
	AlertingMaxAttempts         int
	AlertingMinInterval         int64

	// Explore UI
	ExploreEnabled bool

	// Plutono.NET URL
	PlutonoComUrl string

	ImageUploadProvider string
)

// AddChangePasswordLink returns if login form is disabled or not since
// the same intention can be used to hide both features.
func AddChangePasswordLink() bool {
	return !DisableLoginForm
}

// TODO move all global vars to this struct
type Cfg struct {
	Raw    *ini.File
	Logger log.Logger

	// HTTP Server Settings
	AppURL           string
	AppSubURL        string
	ServeFromSubPath bool
	StaticRootPath   string
	Protocol         Scheme
	SocketPath       string
	RouterLogging    bool
	Domain           string
	CDNRootURL       *url.URL
	ReadTimeout      time.Duration
	EnableGzip       bool
	EnforceDomain    bool

	// build
	BuildVersion string
	BuildCommit  string
	BuildBranch  string
	BuildStamp   int64
	IsEnterprise bool

	// packaging
	Packaging string

	// Paths
	ProvisioningPath   string
	DataPath           string
	LogsPath           string
	PluginsPath        string
	BundledPluginsPath string

	// SMTP email settings
	Smtp SmtpSettings

	// Rendering
	ImagesDir                      string
	RendererUrl                    string
	RendererCallbackUrl            string
	RendererConcurrentRequestLimit int

	// Security
	DisableInitAdminCreation          bool
	DisableBruteForceLoginProtection  bool
	CookieSecure                      bool
	CookieSameSiteDisabled            bool
	CookieSameSiteMode                http.SameSite
	AllowEmbedding                    bool
	XSSProtectionHeader               bool
	ContentTypeProtectionHeader       bool
	StrictTransportSecurity           bool
	StrictTransportSecurityMaxAge     int
	StrictTransportSecurityPreload    bool
	StrictTransportSecuritySubDomains bool
	// CSPEnabled toggles Content Security Policy support.
	CSPEnabled bool
	// CSPTemplate contains the Content Security Policy template.
	CSPTemplate string

	TempDataLifetime         time.Duration
	PluginsEnableAlpha       bool
	PluginsAppsSkipVerifyTLS bool
	PluginSettings           PluginSettings
	PluginsAllowUnsigned     []string
	MarketplaceURL           string
	DisableSanitizeHtml      bool

	// Metrics
	MetricsEndpointEnabled           bool
	MetricsEndpointBasicAuthUsername string
	MetricsEndpointBasicAuthPassword string
	MetricsEndpointDisableTotalStats bool
	MetricsPlutonoEnvironmentInfo    map[string]string

	// Dashboards
	DefaultHomeDashboardPath string

	// Auth
	LoginCookieName              string
	LoginMaxInactiveLifetime     time.Duration
	LoginMaxLifetime             time.Duration
	TokenRotationIntervalMinutes int
	SigV4AuthEnabled             bool
	BasicAuthEnabled             bool
	AdminUser                    string
	AdminPassword                string

	// AWS Plugin Auth
	AWSAllowedAuthProviders []string
	AWSAssumeRoleEnabled    bool
	AWSListMetricsPageLimit int

	// Auth proxy settings
	AuthProxyEnabled          bool
	AuthProxyHeaderName       string
	AuthProxyHeaderProperty   string
	AuthProxyAutoSignUp       bool
	AuthProxyEnableLoginToken bool
	AuthProxyWhitelist        string
	AuthProxyHeaders          map[string]string
	AuthProxySyncTTL          int

	// OAuth
	OAuthCookieMaxAge int

	// SAML Auth
	SAMLEnabled             bool
	SAMLSingleLogoutEnabled bool

	// Dataproxy
	SendUserHeader bool

	// DistributedCache
	RemoteCacheOptions *RemoteCacheOptions

	EditorsCanAdmin bool

	ApiKeyMaxSecondsToLive int64

	// Use to enable new features which may still be in alpha/beta stage.
	FeatureToggles       map[string]bool
	AnonymousEnabled     bool
	AnonymousOrgName     string
	AnonymousOrgRole     string
	AnonymousHideVersion bool

	DateFormats DateFormats

	// User
	UserInviteMaxLifetime time.Duration
	HiddenUsers           map[string]struct{}

	// Annotations
	AnnotationCleanupJobBatchSize      int64
	AlertingAnnotationCleanupSetting   AnnotationCleanupSettings
	DashboardAnnotationCleanupSettings AnnotationCleanupSettings
	APIAnnotationCleanupSettings       AnnotationCleanupSettings

	// Sentry config
	Sentry Sentry

	// Data sources
	DataSourceLimit int
	OptionsLimit    int64
	MaxDataPoints   int64

	// Snapshots
	SnapshotPublicMode bool

	ErrTemplateName string

	Env string

	// LDAP
	LDAPEnabled     bool
	LDAPAllowSignup bool

	Quota QuotaSettings

	DefaultTheme string
	HomePage     string

	AutoAssignOrg     bool
	AutoAssignOrgId   int
	AutoAssignOrgRole string

	// ExpressionsEnabled specifies whether expressions are enabled.
	ExpressionsEnabled bool
}

// IsLiveEnabled returns if plutono live should be enabled
func (cfg Cfg) IsLiveEnabled() bool {
	return cfg.FeatureToggles["live"]
}

// IsNgAlertEnabled returns whether the standalone alerts feature is enabled.
func (cfg Cfg) IsNgAlertEnabled() bool {
	return cfg.FeatureToggles["ngalert"]
}

// IsDatabaseMetricsEnabled returns whether the database instrumentation feature is enabled.
func (cfg Cfg) IsDatabaseMetricsEnabled() bool {
	return cfg.FeatureToggles["database_metrics"]
}

// IsHTTPRequestHistogramEnabled returns whether the http_request_histogram feature is enabled.
func (cfg Cfg) IsHTTPRequestHistogramEnabled() bool {
	return cfg.FeatureToggles["http_request_histogram"]
}

// IsPanelLibraryEnabled returns whether the panel library feature is enabled.
func (cfg Cfg) IsPanelLibraryEnabled() bool {
	return cfg.FeatureToggles["panelLibrary"]
}

type CommandLineArgs struct {
	Config   string
	HomePath string
	Args     []string
}

func parseAppUrlAndSubUrl(section *ini.Section) (string, string, error) {
	appUrl := valueAsString(section, "root_url", "http://localhost:3000/")

	if appUrl[len(appUrl)-1] != '/' {
		appUrl += "/"
	}

	// Check if has app suburl.
	url, err := url.Parse(appUrl)
	if err != nil {
		log.Fatalf(4, "Invalid root_url(%s): %s", appUrl, err)
	}

	appSubUrl := strings.TrimSuffix(url.Path, "/")
	return appUrl, appSubUrl, nil
}

func ToAbsUrl(relativeUrl string) string {
	return AppUrl + relativeUrl
}

func RedactedValue(key, value string) string {
	uppercased := strings.ToUpper(key)
	// Sensitive information: password, secrets etc
	for _, pattern := range []string{
		"PASSWORD",
		"SECRET",
		"PROVIDER_CONFIG",
		"PRIVATE_KEY",
		"SECRET_KEY",
		"CERTIFICATE",
		"ACCOUNT_KEY",
		"ENCRYPTION_KEY",
		"VAULT_TOKEN",
	} {
		if match, err := regexp.MatchString(pattern, uppercased); match && err == nil {
			return redactedPassword
		}
	}

	for _, exception := range []string{
		"RUDDERSTACK",
		"APPLICATION_INSIGHTS",
		"SENTRY",
	} {
		if strings.Contains(uppercased, exception) {
			return value
		}
	}

	if u, err := RedactedURL(value); err == nil {
		return u
	}

	return value
}

func RedactedURL(value string) (string, error) {
	// Value could be a list of URLs
	chunks := util.SplitString(value)

	for i, chunk := range chunks {
		var hasTmpPrefix bool
		const tmpPrefix = "http://"

		if !strings.Contains(chunk, "://") {
			chunk = tmpPrefix + chunk
			hasTmpPrefix = true
		}

		u, err := url.Parse(chunk)
		if err != nil {
			return "", err
		}

		redacted := u.Redacted()
		if hasTmpPrefix {
			redacted = strings.Replace(redacted, tmpPrefix, "", 1)
		}

		chunks[i] = redacted
	}

	if strings.Contains(value, ",") {
		return strings.Join(chunks, ","), nil
	}

	return strings.Join(chunks, " "), nil
}

func applyEnvVariableOverrides(file *ini.File) error {
	appliedEnvOverrides = make([]string, 0)
	for _, section := range file.Sections() {
		for _, key := range section.Keys() {
			envKey := EnvKey(section.Name(), key.Name())
			envValue := os.Getenv(envKey)

			if len(envValue) > 0 {
				key.SetValue(envValue)
				appliedEnvOverrides = append(appliedEnvOverrides, fmt.Sprintf("%s=%s", envKey, RedactedValue(envKey, envValue)))
			}
		}
	}

	return nil
}

func (cfg *Cfg) readPlutonoEnvironmentMetrics() error {
	environmentMetricsSection := cfg.Raw.Section("metrics.environment_info")
	keys := environmentMetricsSection.Keys()
	cfg.MetricsPlutonoEnvironmentInfo = make(map[string]string, len(keys))

	for _, key := range keys {
		labelName := model.LabelName(key.Name())
		labelValue := model.LabelValue(key.Value())

		if !labelName.IsValid() {
			return fmt.Errorf("invalid label name in [metrics.environment_info] configuration. name %q", labelName)
		}

		if !labelValue.IsValid() {
			return fmt.Errorf("invalid label value in [metrics.environment_info] configuration. name %q value %q", labelName, labelValue)
		}

		cfg.MetricsPlutonoEnvironmentInfo[string(labelName)] = string(labelValue)
	}

	return nil
}

func (cfg *Cfg) readAnnotationSettings() {
	section := cfg.Raw.Section("annotations")
	cfg.AnnotationCleanupJobBatchSize = section.Key("cleanupjob_batchsize").MustInt64(100)

	dashboardAnnotation := cfg.Raw.Section("annotations.dashboard")
	apiIAnnotation := cfg.Raw.Section("annotations.api")
	alertingSection := cfg.Raw.Section("alerting")

	var newAnnotationCleanupSettings = func(section *ini.Section, maxAgeField string) AnnotationCleanupSettings {
		maxAge, err := gtime.ParseDuration(section.Key(maxAgeField).MustString(""))
		if err != nil {
			maxAge = 0
		}

		return AnnotationCleanupSettings{
			MaxAge:   maxAge,
			MaxCount: section.Key("max_annotations_to_keep").MustInt64(0),
		}
	}

	cfg.AlertingAnnotationCleanupSetting = newAnnotationCleanupSettings(alertingSection, "max_annotation_age")
	cfg.DashboardAnnotationCleanupSettings = newAnnotationCleanupSettings(dashboardAnnotation, "max_age")
	cfg.APIAnnotationCleanupSettings = newAnnotationCleanupSettings(apiIAnnotation, "max_age")
}

func (cfg *Cfg) readExpressionsSettings() {
	expressions := cfg.Raw.Section("expressions")
	cfg.ExpressionsEnabled = expressions.Key("enabled").MustBool(true)
}

type AnnotationCleanupSettings struct {
	MaxAge   time.Duration
	MaxCount int64
}

func EnvKey(sectionName string, keyName string) string {
	sN := strings.ToUpper(strings.ReplaceAll(sectionName, ".", "_"))
	sN = strings.ReplaceAll(sN, "-", "_")
	kN := strings.ToUpper(strings.ReplaceAll(keyName, ".", "_"))
	envKey := fmt.Sprintf("PL_%s_%s", sN, kN)
	return envKey
}

func applyCommandLineDefaultProperties(props map[string]string, file *ini.File) {
	appliedCommandLineProperties = make([]string, 0)
	for _, section := range file.Sections() {
		for _, key := range section.Keys() {
			keyString := fmt.Sprintf("default.%s.%s", section.Name(), key.Name())
			value, exists := props[keyString]
			if exists {
				key.SetValue(value)
				appliedCommandLineProperties = append(appliedCommandLineProperties,
					fmt.Sprintf("%s=%s", keyString, RedactedValue(keyString, value)))
			}
		}
	}
}

func applyCommandLineProperties(props map[string]string, file *ini.File) {
	for _, section := range file.Sections() {
		sectionName := section.Name() + "."
		if section.Name() == ini.DefaultSection {
			sectionName = ""
		}
		for _, key := range section.Keys() {
			keyString := sectionName + key.Name()
			value, exists := props[keyString]
			if exists {
				appliedCommandLineProperties = append(appliedCommandLineProperties, fmt.Sprintf("%s=%s", keyString, value))
				key.SetValue(value)
			}
		}
	}
}

func getCommandLineProperties(args []string) map[string]string {
	props := make(map[string]string)

	for _, arg := range args {
		if !strings.HasPrefix(arg, "cfg:") {
			continue
		}

		trimmed := strings.TrimPrefix(arg, "cfg:")
		parts := strings.Split(trimmed, "=")
		if len(parts) != 2 {
			log.Fatalf(3, "Invalid command line argument. argument: %v", arg)
			return nil
		}

		props[parts[0]] = parts[1]
	}
	return props
}

func makeAbsolute(path string, root string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(root, path)
}

func loadSpecifiedConfigFile(configFile string, masterFile *ini.File) error {
	if configFile == "" {
		configFile = filepath.Join(HomePath, CustomInitPath)
		// return without error if custom file does not exist
		if !pathExists(configFile) {
			return nil
		}
	}

	userConfig, err := ini.Load(configFile)
	if err != nil {
		return fmt.Errorf("failed to parse %q: %w", configFile, err)
	}

	userConfig.BlockMode = false

	for _, section := range userConfig.Sections() {
		for _, key := range section.Keys() {
			if key.Value() == "" {
				continue
			}

			defaultSec, err := masterFile.GetSection(section.Name())
			if err != nil {
				defaultSec, _ = masterFile.NewSection(section.Name())
			}
			defaultKey, err := defaultSec.GetKey(key.Name())
			if err != nil {
				defaultKey, _ = defaultSec.NewKey(key.Name(), key.Value())
			}
			defaultKey.SetValue(key.Value())
		}
	}

	configFiles = append(configFiles, configFile)
	return nil
}

func (cfg *Cfg) loadConfiguration(args *CommandLineArgs) (*ini.File, error) {
	// load config defaults
	defaultConfigFile := path.Join(HomePath, "conf/defaults.ini")
	configFiles = append(configFiles, defaultConfigFile)

	// check if config file exists
	if _, err := os.Stat(defaultConfigFile); os.IsNotExist(err) {
		fmt.Println("Plutono-server Init Failed: Could not find config defaults, make sure homepath command line parameter is set or working directory is homepath")
		os.Exit(1)
	}

	// load defaults
	parsedFile, err := ini.Load(defaultConfigFile)
	if err != nil {
		fmt.Printf("Failed to parse defaults.ini, %v\n", err)
		os.Exit(1)
		return nil, err
	}

	parsedFile.BlockMode = false

	// command line props
	commandLineProps := getCommandLineProperties(args.Args)
	// load default overrides
	applyCommandLineDefaultProperties(commandLineProps, parsedFile)

	// load specified config file
	err = loadSpecifiedConfigFile(args.Config, parsedFile)
	if err != nil {
		err2 := cfg.initLogging(parsedFile)
		if err2 != nil {
			return nil, err2
		}
		log.Fatalf(3, "%s", err.Error())
	}

	// apply environment overrides
	err = applyEnvVariableOverrides(parsedFile)
	if err != nil {
		return nil, err
	}

	// apply command line overrides
	applyCommandLineProperties(commandLineProps, parsedFile)

	// evaluate config values containing environment variables
	err = expandConfig(parsedFile)
	if err != nil {
		return nil, err
	}

	// update data path and logging config
	dataPath := valueAsString(parsedFile.Section("paths"), "data", "")

	cfg.DataPath = makeAbsolute(dataPath, HomePath)
	err = cfg.initLogging(parsedFile)
	if err != nil {
		return nil, err
	}

	return parsedFile, err
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func setHomePath(args *CommandLineArgs) {
	if args.HomePath != "" {
		HomePath = args.HomePath
		return
	}

	var err error
	HomePath, err = filepath.Abs(".")
	if err != nil {
		panic(err)
	}
	// check if homepath is correct
	if pathExists(filepath.Join(HomePath, "conf/defaults.ini")) {
		return
	}

	// try down one path
	if pathExists(filepath.Join(HomePath, "../conf/defaults.ini")) {
		HomePath = filepath.Join(HomePath, "../")
	}
}

var skipStaticRootValidation = false

func NewCfg() *Cfg {
	return &Cfg{
		Logger: log.New("settings"),
		Raw:    ini.Empty(),
	}
}

var theCfg *Cfg

// GetCfg gets the Cfg singleton.
// XXX: This is only required for integration tests so that the configuration can be reset for each test,
// as due to how the current DI framework functions, we can't create a new Cfg object every time (the services
// constituting the DI graph, and referring to a Cfg instance, get created only once).
func GetCfg() *Cfg {
	if theCfg != nil {
		return theCfg
	}

	theCfg = NewCfg()
	return theCfg
}

func (cfg *Cfg) validateStaticRootPath() error {
	if skipStaticRootValidation {
		return nil
	}

	if _, err := os.Stat(path.Join(StaticRootPath, "build")); err != nil {
		cfg.Logger.Error("Failed to detect generated javascript files in public/build")
	}

	return nil
}

func (cfg *Cfg) Load(args *CommandLineArgs) error {
	setHomePath(args)

	// Fix for missing IANA db on Windows
	_, zoneInfoSet := os.LookupEnv(zoneInfo)
	if runtime.GOOS == "windows" && !zoneInfoSet {
		if err := os.Setenv(zoneInfo, filepath.Join(HomePath, "tools", "zoneinfo.zip")); err != nil {
			cfg.Logger.Error("Can't set ZONEINFO environment variable", "err", err)
		}
	}

	iniFile, err := cfg.loadConfiguration(args)
	if err != nil {
		return err
	}

	cfg.Raw = iniFile

	// Temporarily keep global, to make refactor in steps
	Raw = cfg.Raw

	cfg.BuildVersion = BuildVersion
	cfg.BuildCommit = BuildCommit
	cfg.BuildStamp = BuildStamp
	cfg.BuildBranch = BuildBranch
	cfg.IsEnterprise = IsEnterprise
	cfg.Packaging = Packaging

	cfg.ErrTemplateName = "error"

	ApplicationName = "Plutono"

	Env = valueAsString(iniFile.Section(""), "app_mode", "development")
	cfg.Env = Env
	InstanceName = valueAsString(iniFile.Section(""), "instance_name", "unknown_instance_name")
	plugins := valueAsString(iniFile.Section("paths"), "plugins", "")
	cfg.PluginsPath = makeAbsolute(plugins, HomePath)
	cfg.BundledPluginsPath = makeAbsolute("plugins-bundled", HomePath)
	provisioning := valueAsString(iniFile.Section("paths"), "provisioning", "")
	cfg.ProvisioningPath = makeAbsolute(provisioning, HomePath)

	if err := cfg.readServerSettings(iniFile); err != nil {
		return err
	}

	// read data proxy settings
	dataproxy := iniFile.Section("dataproxy")
	DataProxyLogging = dataproxy.Key("logging").MustBool(false)
	DataProxyTimeout = dataproxy.Key("timeout").MustInt(30)
	DataProxyKeepAlive = dataproxy.Key("keep_alive_seconds").MustInt(30)
	DataProxyTLSHandshakeTimeout = dataproxy.Key("tls_handshake_timeout_seconds").MustInt(10)
	DataProxyExpectContinueTimeout = dataproxy.Key("expect_continue_timeout_seconds").MustInt(1)
	DataProxyMaxConnsPerHost = dataproxy.Key("max_conns_per_host").MustInt(0)
	DataProxyMaxIdleConns = dataproxy.Key("max_idle_connections").MustInt(100)
	DataProxyMaxIdleConnsPerHost = dataproxy.Key("max_idle_connections_per_host").MustInt(2)
	DataProxyIdleConnTimeout = dataproxy.Key("idle_conn_timeout_seconds").MustInt(90)
	cfg.SendUserHeader = dataproxy.Key("send_user_header").MustBool(false)

	if err := readSecuritySettings(iniFile, cfg); err != nil {
		return err
	}

	if err := readSnapshotsSettings(cfg, iniFile); err != nil {
		return err
	}

	// read dashboard settings
	dashboards := iniFile.Section("dashboards")
	DashboardVersionsToKeep = dashboards.Key("versions_to_keep").MustInt(20)
	MinRefreshInterval = valueAsString(dashboards, "min_refresh_interval", "5s")

	cfg.DefaultHomeDashboardPath = dashboards.Key("default_home_dashboard_path").MustString("")

	if err := readUserSettings(iniFile, cfg); err != nil {
		return err
	}
	if err := readAuthSettings(iniFile, cfg); err != nil {
		return err
	}
	if err := readRenderingSettings(iniFile, cfg); err != nil {
		return err
	}

	cfg.TempDataLifetime = iniFile.Section("paths").Key("temp_data_lifetime").MustDuration(time.Second * 3600 * 24)
	cfg.MetricsEndpointEnabled = iniFile.Section("metrics").Key("enabled").MustBool(true)
	cfg.MetricsEndpointBasicAuthUsername = valueAsString(iniFile.Section("metrics"), "basic_auth_username", "")
	cfg.MetricsEndpointBasicAuthPassword = valueAsString(iniFile.Section("metrics"), "basic_auth_password", "")
	cfg.MetricsEndpointDisableTotalStats = iniFile.Section("metrics").Key("disable_total_stats").MustBool(false)

	analytics := iniFile.Section("analytics")
	GoogleAnalyticsId = analytics.Key("google_analytics_ua_id").String()
	GoogleTagManagerId = analytics.Key("google_tag_manager_id").String()

	if err := readAlertingSettings(iniFile); err != nil {
		return err
	}

	explore := iniFile.Section("explore")
	ExploreEnabled = explore.Key("enabled").MustBool(true)

	panelsSection := iniFile.Section("panels")
	cfg.DisableSanitizeHtml = panelsSection.Key("disable_sanitize_html").MustBool(false)

	pluginsSection := iniFile.Section("plugins")
	cfg.PluginsEnableAlpha = pluginsSection.Key("enable_alpha").MustBool(false)
	cfg.PluginsAppsSkipVerifyTLS = pluginsSection.Key("app_tls_skip_verify_insecure").MustBool(false)
	cfg.PluginSettings = extractPluginSettings(iniFile.Sections())
	pluginsAllowUnsigned := pluginsSection.Key("allow_loading_unsigned_plugins").MustString("")
	for _, plug := range strings.Split(pluginsAllowUnsigned, ",") {
		plug = strings.TrimSpace(plug)
		cfg.PluginsAllowUnsigned = append(cfg.PluginsAllowUnsigned, plug)
	}
	cfg.MarketplaceURL = pluginsSection.Key("marketplace_url").MustString("https://grafana.com/grafana/plugins/")

	// Read and populate feature toggles list
	featureTogglesSection := iniFile.Section("feature_toggles")
	cfg.FeatureToggles = make(map[string]bool)
	featuresTogglesStr := valueAsString(featureTogglesSection, "enable", "")
	for _, feature := range util.SplitString(featuresTogglesStr) {
		cfg.FeatureToggles[feature] = true
	}

	// check old location for this option
	if panelsSection.Key("enable_alpha").MustBool(false) {
		cfg.PluginsEnableAlpha = true
	}

	cfg.readLDAPConfig()
	cfg.handleAWSConfig()
	cfg.readSessionConfig()
	cfg.readSmtpSettings()
	cfg.readQuotaSettings()
	cfg.readAnnotationSettings()
	cfg.readExpressionsSettings()
	if err := cfg.readPlutonoEnvironmentMetrics(); err != nil {
		return err
	}

	cfg.readDataSourcesSettings()

	if VerifyEmailEnabled && !cfg.Smtp.Enabled {
		log.Warnf("require_email_validation is enabled but smtp is disabled")
	}

	// check old key  name
	PlutonoComUrl = valueAsString(iniFile.Section("plutono_net"), "url", "")
	if PlutonoComUrl == "" {
		PlutonoComUrl = valueAsString(iniFile.Section("plutono_com"), "url", "https://grafana.com")
	}

	imageUploadingSection := iniFile.Section("external_image_storage")
	ImageUploadProvider = valueAsString(imageUploadingSection, "provider", "")

	cacheServer := iniFile.Section("remote_cache")
	dbName := valueAsString(cacheServer, "type", "database")
	connStr := valueAsString(cacheServer, "connstr", "")

	cfg.RemoteCacheOptions = &RemoteCacheOptions{
		Name:    dbName,
		ConnStr: connStr,
	}

	cfg.readDateFormats()
	cfg.readSentryConfig()

	return nil
}

func valueAsString(section *ini.Section, keyName string, defaultValue string) string {
	return section.Key(keyName).MustString(defaultValue)
}

type RemoteCacheOptions struct {
	Name    string
	ConnStr string
}

func (cfg *Cfg) readLDAPConfig() {
	ldapSec := cfg.Raw.Section("auth.ldap")
	LDAPConfigFile = ldapSec.Key("config_file").String()
	LDAPSyncCron = ldapSec.Key("sync_cron").String()
	LDAPEnabled = ldapSec.Key("enabled").MustBool(false)
	cfg.LDAPEnabled = LDAPEnabled
	LDAPActiveSyncEnabled = ldapSec.Key("active_sync_enabled").MustBool(false)
	LDAPAllowSignup = ldapSec.Key("allow_sign_up").MustBool(true)
	cfg.LDAPAllowSignup = LDAPAllowSignup
}

func (cfg *Cfg) handleAWSConfig() {
	awsPluginSec := cfg.Raw.Section("aws")
	cfg.AWSAssumeRoleEnabled = awsPluginSec.Key("assume_role_enabled").MustBool(true)
	allowedAuthProviders := awsPluginSec.Key("allowed_auth_providers").MustString("default,keys,credentials")
	for _, authProvider := range strings.Split(allowedAuthProviders, ",") {
		authProvider = strings.TrimSpace(authProvider)
		if authProvider != "" {
			cfg.AWSAllowedAuthProviders = append(cfg.AWSAllowedAuthProviders, authProvider)
		}
	}
	cfg.AWSListMetricsPageLimit = awsPluginSec.Key("list_metrics_page_limit").MustInt(500)
	// Also set environment variables that can be used by core plugins
	err := os.Setenv(awsds.AssumeRoleEnabledEnvVarKeyName, strconv.FormatBool(cfg.AWSAssumeRoleEnabled))
	if err != nil {
		cfg.Logger.Error(fmt.Sprintf("could not set environment variable '%s'", awsds.AssumeRoleEnabledEnvVarKeyName), err)
	}

	err = os.Setenv(awsds.AllowedAuthProvidersEnvVarKeyName, allowedAuthProviders)
	if err != nil {
		cfg.Logger.Error(fmt.Sprintf("could not set environment variable '%s'", awsds.AllowedAuthProvidersEnvVarKeyName), err)
	}
}

func (cfg *Cfg) readSessionConfig() {
	sec, _ := cfg.Raw.GetSection("session")

	if sec != nil {
		cfg.Logger.Warn(
			"[Removed] Session setting was removed in v6.2, use remote_cache option instead",
		)
	}
}

func (cfg *Cfg) initLogging(file *ini.File) error {
	logModeStr := valueAsString(file.Section("log"), "mode", "console")
	// split on comma
	logModes := strings.Split(logModeStr, ",")
	// also try space
	if len(logModes) == 1 {
		logModes = strings.Split(logModeStr, " ")
	}
	logsPath := valueAsString(file.Section("paths"), "logs", "")
	cfg.LogsPath = makeAbsolute(logsPath, HomePath)
	return log.ReadLoggingConfig(logModes, cfg.LogsPath, file)
}

func (cfg *Cfg) LogConfigSources() {
	var text bytes.Buffer

	for _, file := range configFiles {
		cfg.Logger.Info("Config loaded from", "file", file)
	}

	if len(appliedCommandLineProperties) > 0 {
		for _, prop := range appliedCommandLineProperties {
			cfg.Logger.Info("Config overridden from command line", "arg", prop)
		}
	}

	if len(appliedEnvOverrides) > 0 {
		text.WriteString("\tEnvironment variables used:\n")
		for _, prop := range appliedEnvOverrides {
			cfg.Logger.Info("Config overridden from Environment variable", "var", prop)
		}
	}

	cfg.Logger.Info("Path Home", "path", HomePath)
	cfg.Logger.Info("Path Data", "path", cfg.DataPath)
	cfg.Logger.Info("Path Logs", "path", cfg.LogsPath)
	cfg.Logger.Info("Path Plugins", "path", cfg.PluginsPath)
	cfg.Logger.Info("Path Provisioning", "path", cfg.ProvisioningPath)
	cfg.Logger.Info("App mode " + cfg.Env)
}

type DynamicSection struct {
	section *ini.Section
	Logger  log.Logger
}

// Key dynamically overrides keys with environment variables.
// As a side effect, the value of the setting key will be updated if an environment variable is present.
func (s *DynamicSection) Key(k string) *ini.Key {
	envKey := EnvKey(s.section.Name(), k)
	envValue := os.Getenv(envKey)
	key := s.section.Key(k)

	if len(envValue) == 0 {
		return key
	}

	key.SetValue(envValue)
	s.Logger.Info("Config overridden from Environment variable", "var", fmt.Sprintf("%s=%s", envKey, RedactedValue(envKey, envValue)))

	return key
}

// SectionWithEnvOverrides dynamically overrides keys with environment variables.
// As a side effect, the value of the setting key will be updated if an environment variable is present.
func (cfg *Cfg) SectionWithEnvOverrides(s string) *DynamicSection {
	return &DynamicSection{cfg.Raw.Section(s), cfg.Logger}
}

func readSecuritySettings(iniFile *ini.File, cfg *Cfg) error {
	security := iniFile.Section("security")
	SecretKey = valueAsString(security, "secret_key", "")
	DisableGravatar = security.Key("disable_gravatar").MustBool(true)
	cfg.DisableBruteForceLoginProtection = security.Key("disable_brute_force_login_protection").MustBool(false)

	CookieSecure = security.Key("cookie_secure").MustBool(false)
	cfg.CookieSecure = CookieSecure

	samesiteString := valueAsString(security, "cookie_samesite", "lax")

	if samesiteString == "disabled" {
		CookieSameSiteDisabled = true
		cfg.CookieSameSiteDisabled = CookieSameSiteDisabled
	} else {
		validSameSiteValues := map[string]http.SameSite{
			"lax":    http.SameSiteLaxMode,
			"strict": http.SameSiteStrictMode,
			"none":   http.SameSiteNoneMode,
		}

		if samesite, ok := validSameSiteValues[samesiteString]; ok {
			CookieSameSiteMode = samesite
			cfg.CookieSameSiteMode = CookieSameSiteMode
		} else {
			CookieSameSiteMode = http.SameSiteLaxMode
			cfg.CookieSameSiteMode = CookieSameSiteMode
		}
	}
	cfg.AllowEmbedding = security.Key("allow_embedding").MustBool(false)

	cfg.ContentTypeProtectionHeader = security.Key("x_content_type_options").MustBool(true)
	cfg.XSSProtectionHeader = security.Key("x_xss_protection").MustBool(true)
	cfg.StrictTransportSecurity = security.Key("strict_transport_security").MustBool(false)
	cfg.StrictTransportSecurityMaxAge = security.Key("strict_transport_security_max_age_seconds").MustInt(86400)
	cfg.StrictTransportSecurityPreload = security.Key("strict_transport_security_preload").MustBool(false)
	cfg.StrictTransportSecuritySubDomains = security.Key("strict_transport_security_subdomains").MustBool(false)
	cfg.CSPEnabled = security.Key("content_security_policy").MustBool(false)
	cfg.CSPTemplate = security.Key("content_security_policy_template").MustString("")

	// read data source proxy whitelist
	DataProxyWhiteList = make(map[string]bool)
	securityStr := valueAsString(security, "data_source_proxy_whitelist", "")

	for _, hostAndIP := range util.SplitString(securityStr) {
		DataProxyWhiteList[hostAndIP] = true
	}

	// admin
	cfg.DisableInitAdminCreation = security.Key("disable_initial_admin_creation").MustBool(false)
	cfg.AdminUser = valueAsString(security, "admin_user", "")
	cfg.AdminPassword = valueAsString(security, "admin_password", "")

	return nil
}

func readAuthSettings(iniFile *ini.File, cfg *Cfg) (err error) {
	auth := iniFile.Section("auth")

	cfg.LoginCookieName = valueAsString(auth, "login_cookie_name", "plutono_session")
	maxInactiveDaysVal := auth.Key("login_maximum_inactive_lifetime_days").MustString("")
	if maxInactiveDaysVal != "" {
		maxInactiveDaysVal = fmt.Sprintf("%sd", maxInactiveDaysVal)
		cfg.Logger.Warn("[Deprecated] the configuration setting 'login_maximum_inactive_lifetime_days' is deprecated, please use 'login_maximum_inactive_lifetime_duration' instead")
	} else {
		maxInactiveDaysVal = "7d"
	}
	maxInactiveDurationVal := valueAsString(auth, "login_maximum_inactive_lifetime_duration", maxInactiveDaysVal)
	cfg.LoginMaxInactiveLifetime, err = gtime.ParseDuration(maxInactiveDurationVal)
	if err != nil {
		return err
	}

	maxLifetimeDaysVal := auth.Key("login_maximum_lifetime_days").MustString("")
	if maxLifetimeDaysVal != "" {
		maxLifetimeDaysVal = fmt.Sprintf("%sd", maxLifetimeDaysVal)
		cfg.Logger.Warn("[Deprecated] the configuration setting 'login_maximum_lifetime_days' is deprecated, please use 'login_maximum_lifetime_duration' instead")
	} else {
		maxLifetimeDaysVal = "30d"
	}
	maxLifetimeDurationVal := valueAsString(auth, "login_maximum_lifetime_duration", maxLifetimeDaysVal)
	cfg.LoginMaxLifetime, err = gtime.ParseDuration(maxLifetimeDurationVal)
	if err != nil {
		return err
	}

	cfg.ApiKeyMaxSecondsToLive = auth.Key("api_key_max_seconds_to_live").MustInt64(-1)

	cfg.TokenRotationIntervalMinutes = auth.Key("token_rotation_interval_minutes").MustInt(10)
	if cfg.TokenRotationIntervalMinutes < 2 {
		cfg.TokenRotationIntervalMinutes = 2
	}

	DisableLoginForm = auth.Key("disable_login_form").MustBool(false)
	DisableSignoutMenu = auth.Key("disable_signout_menu").MustBool(false)
	OAuthAutoLogin = auth.Key("oauth_auto_login").MustBool(false)
	cfg.OAuthCookieMaxAge = auth.Key("oauth_state_cookie_max_age").MustInt(600)
	SignoutRedirectUrl = valueAsString(auth, "signout_redirect_url", "")

	// SigV4
	SigV4AuthEnabled = auth.Key("sigv4_auth_enabled").MustBool(false)
	cfg.SigV4AuthEnabled = SigV4AuthEnabled

	// SAML auth
	cfg.SAMLEnabled = iniFile.Section("auth.saml").Key("enabled").MustBool(false)
	cfg.SAMLSingleLogoutEnabled = iniFile.Section("auth.saml").Key("single_logout").MustBool(false)

	// anonymous access
	AnonymousEnabled = iniFile.Section("auth.anonymous").Key("enabled").MustBool(false)
	cfg.AnonymousEnabled = AnonymousEnabled
	cfg.AnonymousOrgName = valueAsString(iniFile.Section("auth.anonymous"), "org_name", "")
	cfg.AnonymousOrgRole = valueAsString(iniFile.Section("auth.anonymous"), "org_role", "")
	cfg.AnonymousHideVersion = iniFile.Section("auth.anonymous").Key("hide_version").MustBool(false)

	// basic auth
	authBasic := iniFile.Section("auth.basic")
	BasicAuthEnabled = authBasic.Key("enabled").MustBool(true)
	cfg.BasicAuthEnabled = BasicAuthEnabled

	authProxy := iniFile.Section("auth.proxy")
	AuthProxyEnabled = authProxy.Key("enabled").MustBool(false)
	cfg.AuthProxyEnabled = AuthProxyEnabled

	cfg.AuthProxyHeaderName = valueAsString(authProxy, "header_name", "")
	AuthProxyHeaderProperty = valueAsString(authProxy, "header_property", "")
	cfg.AuthProxyHeaderProperty = AuthProxyHeaderProperty
	cfg.AuthProxyAutoSignUp = authProxy.Key("auto_sign_up").MustBool(true)
	cfg.AuthProxyEnableLoginToken = authProxy.Key("enable_login_token").MustBool(false)

	ldapSyncVal := authProxy.Key("ldap_sync_ttl").MustInt()
	syncVal := authProxy.Key("sync_ttl").MustInt()

	if ldapSyncVal != authProxySyncTTL {
		cfg.AuthProxySyncTTL = ldapSyncVal
		cfg.Logger.Warn("[Deprecated] the configuration setting 'ldap_sync_ttl' is deprecated, please use 'sync_ttl' instead")
	} else {
		cfg.AuthProxySyncTTL = syncVal
	}

	cfg.AuthProxyWhitelist = valueAsString(authProxy, "whitelist", "")

	cfg.AuthProxyHeaders = make(map[string]string)
	headers := valueAsString(authProxy, "headers", "")

	for _, propertyAndHeader := range util.SplitString(headers) {
		split := strings.SplitN(propertyAndHeader, ":", 2)
		if len(split) == 2 {
			cfg.AuthProxyHeaders[split[0]] = split[1]
		}
	}

	return nil
}

func readUserSettings(iniFile *ini.File, cfg *Cfg) error {
	users := iniFile.Section("users")
	AllowUserSignUp = users.Key("allow_sign_up").MustBool(true)
	AllowUserOrgCreate = users.Key("allow_org_create").MustBool(true)
	cfg.AutoAssignOrg = users.Key("auto_assign_org").MustBool(true)
	AutoAssignOrg = cfg.AutoAssignOrg
	cfg.AutoAssignOrgId = users.Key("auto_assign_org_id").MustInt(1)
	AutoAssignOrgId = cfg.AutoAssignOrgId
	cfg.AutoAssignOrgRole = users.Key("auto_assign_org_role").In("Editor", []string{"Editor", "Admin", "Viewer"})
	AutoAssignOrgRole = cfg.AutoAssignOrgRole
	VerifyEmailEnabled = users.Key("verify_email_enabled").MustBool(false)

	LoginHint = valueAsString(users, "login_hint", "")
	PasswordHint = valueAsString(users, "password_hint", "")
	cfg.DefaultTheme = valueAsString(users, "default_theme", "")
	cfg.HomePage = valueAsString(users, "home_page", "")
	ExternalUserMngLinkUrl = valueAsString(users, "external_manage_link_url", "")
	ExternalUserMngLinkName = valueAsString(users, "external_manage_link_name", "")
	ExternalUserMngInfo = valueAsString(users, "external_manage_info", "")

	ViewersCanEdit = users.Key("viewers_can_edit").MustBool(false)
	cfg.EditorsCanAdmin = users.Key("editors_can_admin").MustBool(false)

	userInviteMaxLifetimeVal := valueAsString(users, "user_invite_max_lifetime_duration", "24h")
	userInviteMaxLifetimeDuration, err := gtime.ParseDuration(userInviteMaxLifetimeVal)
	if err != nil {
		return err
	}

	cfg.UserInviteMaxLifetime = userInviteMaxLifetimeDuration
	if cfg.UserInviteMaxLifetime < time.Minute*15 {
		return errors.New("the minimum supported value for the `user_invite_max_lifetime_duration` configuration is 15m (15 minutes)")
	}

	cfg.HiddenUsers = make(map[string]struct{})
	hiddenUsers := users.Key("hidden_users").MustString("")
	for _, user := range strings.Split(hiddenUsers, ",") {
		user = strings.TrimSpace(user)
		if user != "" {
			cfg.HiddenUsers[user] = struct{}{}
		}
	}

	return nil
}

func readRenderingSettings(iniFile *ini.File, cfg *Cfg) error {
	renderSec := iniFile.Section("rendering")
	cfg.RendererUrl = valueAsString(renderSec, "server_url", "")
	cfg.RendererCallbackUrl = valueAsString(renderSec, "callback_url", "")

	if cfg.RendererCallbackUrl == "" {
		cfg.RendererCallbackUrl = AppUrl
	} else {
		if cfg.RendererCallbackUrl[len(cfg.RendererCallbackUrl)-1] != '/' {
			cfg.RendererCallbackUrl += "/"
		}
		_, err := url.Parse(cfg.RendererCallbackUrl)
		if err != nil {
			// XXX: Should return an error?
			log.Fatalf(4, "Invalid callback_url(%s): %s", cfg.RendererCallbackUrl, err)
		}
	}

	cfg.RendererConcurrentRequestLimit = renderSec.Key("concurrent_render_request_limit").MustInt(30)
	cfg.ImagesDir = filepath.Join(cfg.DataPath, "png")

	return nil
}

func readAlertingSettings(iniFile *ini.File) error {
	alerting := iniFile.Section("alerting")
	AlertingEnabled = alerting.Key("enabled").MustBool(true)
	ExecuteAlerts = alerting.Key("execute_alerts").MustBool(true)
	AlertingRenderLimit = alerting.Key("concurrent_render_limit").MustInt(5)

	AlertingErrorOrTimeout = valueAsString(alerting, "error_or_timeout", "alerting")
	AlertingNoDataOrNullValues = valueAsString(alerting, "nodata_or_nullvalues", "no_data")

	evaluationTimeoutSeconds := alerting.Key("evaluation_timeout_seconds").MustInt64(30)
	AlertingEvaluationTimeout = time.Second * time.Duration(evaluationTimeoutSeconds)
	notificationTimeoutSeconds := alerting.Key("notification_timeout_seconds").MustInt64(30)
	AlertingNotificationTimeout = time.Second * time.Duration(notificationTimeoutSeconds)
	AlertingMaxAttempts = alerting.Key("max_attempts").MustInt(3)
	AlertingMinInterval = alerting.Key("min_interval_seconds").MustInt64(1)

	return nil
}

func readSnapshotsSettings(cfg *Cfg, iniFile *ini.File) error {
	snapshots := iniFile.Section("snapshots")

	ExternalSnapshotUrl = valueAsString(snapshots, "external_snapshot_url", "")
	ExternalSnapshotName = valueAsString(snapshots, "external_snapshot_name", "")

	ExternalEnabled = snapshots.Key("external_enabled").MustBool(true)
	SnapShotRemoveExpired = snapshots.Key("snapshot_remove_expired").MustBool(true)
	cfg.SnapshotPublicMode = snapshots.Key("public_mode").MustBool(false)

	return nil
}

func (cfg *Cfg) readServerSettings(iniFile *ini.File) error {
	server := iniFile.Section("server")
	var err error
	AppUrl, AppSubUrl, err = parseAppUrlAndSubUrl(server)
	if err != nil {
		return err
	}
	ServeFromSubPath = server.Key("serve_from_sub_path").MustBool(false)

	cfg.AppURL = AppUrl
	cfg.AppSubURL = AppSubUrl
	cfg.ServeFromSubPath = ServeFromSubPath
	cfg.Protocol = HTTPScheme

	protocolStr := valueAsString(server, "protocol", "http")

	if protocolStr == "https" {
		cfg.Protocol = HTTPSScheme
		CertFile = server.Key("cert_file").String()
		KeyFile = server.Key("cert_key").String()
	}
	if protocolStr == "h2" {
		cfg.Protocol = HTTP2Scheme
		CertFile = server.Key("cert_file").String()
		KeyFile = server.Key("cert_key").String()
	}
	if protocolStr == "socket" {
		cfg.Protocol = SocketScheme
		cfg.SocketPath = server.Key("socket").String()
	}

	cfg.Domain = valueAsString(server, "domain", "localhost")
	HttpAddr = valueAsString(server, "http_addr", DefaultHTTPAddr)
	HttpPort = valueAsString(server, "http_port", "3000")
	cfg.RouterLogging = server.Key("router_logging").MustBool(false)

	EnableGzip = server.Key("enable_gzip").MustBool(false)
	EnforceDomain = server.Key("enforce_domain").MustBool(false)
	staticRoot := valueAsString(server, "static_root_path", "")
	StaticRootPath = makeAbsolute(staticRoot, HomePath)
	cfg.StaticRootPath = StaticRootPath

	if err := cfg.validateStaticRootPath(); err != nil {
		return err
	}

	cdnURL := valueAsString(server, "cdn_url", "")
	if cdnURL != "" {
		cfg.CDNRootURL, err = url.Parse(cdnURL)
		if err != nil {
			return err
		}
	}

	cfg.ReadTimeout = server.Key("read_timeout").MustDuration(0)

	return nil
}

// GetContentDeliveryURL returns full content delivery URL with /<edition>/<version> added to URL
func (cfg *Cfg) GetContentDeliveryURL(prefix string) string {
	if cfg.CDNRootURL != nil {
		url := *cfg.CDNRootURL
		preReleaseFolder := ""

		if strings.Contains(cfg.BuildVersion, "pre") || strings.Contains(cfg.BuildVersion, "alpha") {
			preReleaseFolder = "pre-releases"
		}

		url.Path = path.Join(url.Path, prefix, preReleaseFolder, cfg.BuildVersion)
		return url.String() + "/"
	}

	return ""
}

func (cfg *Cfg) readDataSourcesSettings() {
	const MaxDataPointsLimitScale = 10 //for some reason we need to divide by 10 for the max data points limit. IE 100 means 1000. Not clear why that is the case
	datasources := cfg.Raw.Section("datasources")
	cfg.DataSourceLimit = datasources.Key("datasource_limit").MustInt(5000)
	cfg.OptionsLimit = datasources.Key("options_limit").MustInt64(1000)
	cfg.MaxDataPoints = cfg.OptionsLimit / MaxDataPointsLimitScale
}
