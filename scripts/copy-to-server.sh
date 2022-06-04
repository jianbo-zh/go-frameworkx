#!/bin/bash

remote_host=$1
remote_port=$2
remote_user=$3
remote_key=$4

upload_dir=$5
app_name=$6
app_version=$7

scp -i ${remote_key} -P ${remote_port} build/${app_name}-${app_version}.tgz ${remote_user}@${remote_host}:${upload_dir}