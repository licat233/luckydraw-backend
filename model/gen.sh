#!/bin/bash
###
 # @Author: licat
 # @Date: 2023-02-20 23:28:30
 # @LastEditors: licat
 # @LastEditTime: 2023-02-20 23:50:13
 # @Description: licat233@gmail.com
###
current_path=$(
    cd $(dirname $0)
    pwd
)

cd $current_path

goctl model mysql ddl --src="../deploy/mysql/luckydraw.sql" --dir="."

name="luckydraw"
filename="luckydraw.api"

sql2rpc -model -db_schema=${name}
