#!/bin/bash

upload_dir=$1
deploy_dir=$2
app_name=$3
app_version=$4

cd ${upload_dir}/${app_name}-${app_version}/deploy/${app_name}

export APP_NAME=${app_name}
export APP_VERSION=${app_version}
export DEPLOY_DIR=${deploy_dir}

echo export APP_NAME=$APP_NAME
echo export APP_VERSION=$APP_VERSION
echo export DEPLOY_DIR=$DEPLOY_DIR

# docker-compose restart
docker-compose down; docker-compose up -d;

