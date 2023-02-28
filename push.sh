#!/bin/bash
###
 # @Author: licat
 # @Date: 2023-02-21 11:00:26
 # @LastEditors: licat
 # @LastEditTime: 2023-02-21 11:01:35
 # @Description: licat233@gmail.com
###

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

cd $current_path

git rm -r --cached .
git add .
git commit -m "update"
git push origin main