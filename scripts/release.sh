#!/usr/bin/env bash

usage() {
    echo "USAGE: ./release.sh [version] [msg...]"
    exit 1
}

REVISION=$(git rev-parse HEAD)
GIT_TAG=$(git name-rev --tags --name-only $REVISION)
if [ "$GIT_TAG" = "" ]; then
    GIT_TAG="devel"
fi


VERSION=$1
if [ "$VERSION" = "" ]; then
    echo "Need to specify a version! Perhaps '$GIT_TAG'?"
    usage
fi

set -u -e

rm -rf /tmp/untilitworks_build/

mkdir -p /tmp/untilitworks_build/linux
GOOS=linux go build -ldflags "-X main.version=$VERSION" -o /tmp/untilitworks_build/linux/untilitworks ../
pushd /tmp/untilitworks_build/linux/
tar cvzf /tmp/untilitworks_build/untilitworks_linux.tar.gz untilitworks
popd

mkdir -p /tmp/untilitworks_build/darwin
GOOS=darwin go build -ldflags "-X main.version=$VERSION" -o /tmp/untilitworks_build/darwin/untilitworks ../
pushd /tmp/untilitworks_build/darwin/
tar cvzf /tmp/untilitworks_build/untilitworks_darwin.tar.gz untilitworks
popd

temple file < README.tmpl.md > ../README.md -var "version=$VERSION"
git add ../README.md
git commit -m 'release bump'

hub release create \
    -a /tmp/untilitworks_build/untilitworks_linux.tar.gz \
    -a /tmp/untilitworks_build/untilitworks_darwin.tar.gz \
    $VERSION

git push origin master
