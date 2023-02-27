#!/bin/bash
###
# @Author: licat
# @Date: 2023-01-15 14:05:33
 # @LastEditors: licat
 # @LastEditTime: 2023-02-22 00:58:57
# @Description: licat233@gmail.com
###

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

cd $current_path


name="luckydraw"
filename="luckydraw.api"

sql2rpc -api -db_schema=${name} -filename=${filename} -service_name=${name}-api -api_jwt="Auth" -api_middleware="AuthMiddleware" -api_prefix="api" --api_multiple=true --api_style=sqlRpc

if [ $? -ne 0 ]; then
    exit 1
fi

# goctl api go -api luckydraw.api -dir=../ -style goZero
goctl api go -api $filename -dir=../ -style goZero

goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api $filename -dir ../static

goctl api ts -api $filename -dir ../static