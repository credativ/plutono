package notifications

import (
	"testing"

	"github.com/credativ/plutono/pkg/bus"
	"github.com/credativ/plutono/pkg/models"
	"github.com/credativ/plutono/pkg/setting"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNotifications(t *testing.T) {
	Convey("Given the notifications service", t, func() {
		setting.StaticRootPath = "../../../public/"

		ns := &NotificationService{}
		ns.Bus = bus.New()
		ns.Cfg = setting.NewCfg()
		ns.Cfg.Smtp.Enabled = true
		ns.Cfg.Smtp.TemplatesPattern = "emails/*.html"
		ns.Cfg.Smtp.FromAddress = "from@address.com"
		ns.Cfg.Smtp.FromName = "Plutono Admin"

		err := ns.Init()
		So(err, ShouldBeNil)

		Convey("When sending reset email password", func() {
			err := ns.sendResetPasswordEmail(&models.SendResetPasswordEmailCommand{User: &models.User{Email: "asd@asd.com"}})
			So(err, ShouldBeNil)

			sentMsg := <-ns.mailQueue
			So(sentMsg.Body, ShouldContainSubstring, "body")
			So(sentMsg.Subject, ShouldEqual, "Reset your Plutono password - asd@asd.com")
			So(sentMsg.Body, ShouldNotContainSubstring, "Subject")
		})
	})
}
