#!/bin/bash

app_name=$1
app_version=$2

git archive --format=tar.gz --prefix=${app_name}-${app_version}/ --output="build/${app_name}-${app_version}.tgz" ${app_version}