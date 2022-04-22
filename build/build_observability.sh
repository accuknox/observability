#!/bin/bash


AUTOPOL_HOME=`dirname $(realpath "$0")`/../..
cd $AUTOPOL_HOME
AUTOPOL_SRC_HOME=$AUTOPOL_HOME
[[ "$REPO" == "" ]] && REPO="accuknox/observability"

# check version

# VERSION=`git rev-parse --abbrev-ref HEAD`
VERSION="1"
# if [ ! -z $1 ]; then
#     VERSION=$1
# fi

# remove old images
sudo docker images | grep observability | awk '{print $3}' | xargs -I {} docker rmi -f {} 2> /dev/null
echo "[INFO] Removed existing $REPO images"

# build image
echo "[INFO] Building $REPO:$VERSION"
sudo docker build -t $REPO:$VERSION .

if [ $? != 0 ]; then
    echo "[FAILED] Failed to build $REPO:$VERSION"
    exit 1
fi
echo "[PASSED] Built $REPO:$VERSION"