package licensing

import (
	"github.com/credativ/plutono/pkg/services/hooks"
	"github.com/credativ/plutono/pkg/setting"
)

const (
	openSource = "Open Source"
)

type OSSLicensingService struct {
	Cfg          *setting.Cfg        `inject:""`
	HooksService *hooks.HooksService `inject:""`
}

func (*OSSLicensingService) HasLicense() bool {
	return false
}

func (*OSSLicensingService) Expiry() int64 {
	return 0
}

func (*OSSLicensingService) Edition() string {
	return openSource
}

func (*OSSLicensingService) StateInfo() string {
	return ""
}

func (*OSSLicensingService) ContentDeliveryPrefix() string {
	return "plutono-oss"
}

func (l *OSSLicensingService) Init() error {
	return nil
}

func (*OSSLicensingService) HasValidLicense() bool {
	return false
}
