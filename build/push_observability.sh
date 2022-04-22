#!/bin/bash

AUTOPOL_HOME=`dirname $(realpath "$0")`/../..
[[ "$REPO" == "" ]] && REPO="accuknox/observability"

# check version
VERSION=`git rev-parse --abbrev-ref HEAD`

if [ ! -z $1 ]; then
    VERSION=$1
fi

echo "[INFO] Pushing $REPO:$VERSION"
docker push $REPO:$VERSION

if [ $? != 0 ]; then
    echo "[FAILED] Failed to push $REPO:$VERSION"
    exit 1
fi
echo "[PASSED] Pushed $REPO:$VERSION"