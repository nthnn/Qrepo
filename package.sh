#!/bin/bash
PROJECT_NAME="qrepo"
VERSION="1.0.0"
ARCHITECTURES=("amd64" "riscv64" "armhf")
BUILD_DIR="dist"
INSTALL_DIR="/usr/bin"
PACKAGE_DIR="package"

mkdir -p $BUILD_DIR
for ARCH in "${ARCHITECTURES[@]}"; do
    rm -rf $PACKAGE_DIR
    mkdir -p $PACKAGE_DIR/DEBIAN
    mkdir -p $PACKAGE_DIR/$INSTALL_DIR

    GOARCH=$ARCH go build -ldflags "-s -w" -o $PACKAGE_DIR/$INSTALL_DIR/$PROJECT_NAME github.com/nthnn/Qrepo

    cat <<EOF >$PACKAGE_DIR/DEBIAN/control
Package: $PROJECT_NAME
Version: $VERSION
Section: base
Priority: optional
Architecture: $ARCH
Maintainer: Nathanne Isip <nathanne@example.com>
Description: Simple and easy-to-use all-for-one build tool for AIX, Android, Darwin, Dragonfly, FreeBSD, illumos, iOS, Linux, NetBSD, OpenBSD, Plan9, Solaris, and Windows. 
EOF

    dpkg-deb --build $PACKAGE_DIR $BUILD_DIR/${PROJECT_NAME}_${VERSION}_${ARCH}.deb
done

rm -rf $PACKAGE_DIR
echo "Debian packages have been created in the $BUILD_DIR directory."
