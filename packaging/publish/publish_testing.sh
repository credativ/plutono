#! /usr/bin/env bash
deb_ver=5.1.0-beta1
rpm_ver=5.1.0-beta1

wget https://s3-us-west-2.amazonaws.com/plutono-releases/release/plutono_${deb_ver}_amd64.deb

package_cloud push plutono/testing/debian/jessie plutono_${deb_ver}_amd64.deb
package_cloud push plutono/testing/debian/wheezy plutono_${deb_ver}_amd64.deb
package_cloud push plutono/testing/debian/stretch plutono_${deb_ver}_amd64.deb

wget https://s3-us-west-2.amazonaws.com/plutono-releases/release/plutono-${rpm_ver}.x86_64.rpm

package_cloud push plutono/testing/el/6 plutono-${rpm_ver}.x86_64.rpm
package_cloud push plutono/testing/el/7 plutono-${rpm_ver}.x86_64.rpm

rm plutono*.{deb,rpm}
