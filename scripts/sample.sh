version=`cat VERSION`

PACKAGE_ROOT="./tmp"
DIST_DIR="./dist"
HOME_DIR="${PACKAGE_ROOT}/usr/share/alertmanager"
echo $HOME_DIR
BIN_DIR="${PACKAGE_ROOT}/usr/sbin"
CONF_DIR="${PACKAGE_ROOT}/etc/alertmanager"
SYSCONFIG_DIR="${PACKAGE_ROOT}/etc/sysconfig"
INITD_DIR="${PACKAGE_ROOT}/etc/init.d"
SYSTEMD_DIR="${PACKAGE_ROOT}/usr/lib/systemd/system/"
LOGROTATE_DIR="${PACKAGE_ROOT}/etc/logrotate.d"


if [ -d "$DIST_DIR" ] ; then 
	rm -r $DIST_DIR
fi
mkdir -p $DIST_DIR
mkdir -p $PACKAGE_ROOT
mkdir -p $HOME_DIR
mkdir -p $BIN_DIR
mkdir -p $CONF_DIR
mkdir -p $SYSCONFIG_DIR
mkdir -p $INITD_DIR
mkdir -p $SYSTEMD_DIR
mkdir -p $LOGROTATE_DIR

cp -p .build/linux-amd64/* $BIN_DIR
cp -p packaging/rpm/conf/* $HOME_DIR
cp -p packaging/rpm/init.d/* $INITD_DIR
cp -p packaging/rpm/sysconfig/* $SYSCONFIG_DIR
cp -p packaging/rpm/systemd/* $SYSTEMD_DIR
cp -p packaging/rpm/log/* $LOGROTATE_DIR