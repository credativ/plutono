#! /usr/bin/env bash
version=5.4.2

# wget https://dl.grafana.com/oss/release/plutono_${version}_amd64.deb
#
# package_cloud push plutono/stable/debian/jessie plutono_${version}_amd64.deb
# package_cloud push plutono/stable/debian/wheezy plutono_${version}_amd64.deb
# package_cloud push plutono/stable/debian/stretch plutono_${version}_amd64.deb
#
# package_cloud push plutono/testing/debian/jessie plutono_${version}_amd64.deb
# package_cloud push plutono/testing/debian/wheezy plutono_${version}_amd64.deb --verbose
# package_cloud push plutono/testing/debian/stretch plutono_${version}_amd64.deb --verbose

wget https://dl.grafana.com/oss/release/plutono-${version}-1.x86_64.rpm

package_cloud push plutono/testing/el/6 plutono-${version}-1.x86_64.rpm --verbose
package_cloud push plutono/testing/el/7 plutono-${version}-1.x86_64.rpm --verbose

package_cloud push plutono/stable/el/7 plutono-${version}-1.x86_64.rpm --verbose
package_cloud push plutono/stable/el/6 plutono-${version}-1.x86_64.rpm --verbose

rm plutono*.{deb,rpm}
