#!/bin/bash

# 设置错误时退出
set -e
cd /data/code

# 启动 usergrowth_server
./server > /data/code/logs/server_run.log 2>&1 &

# 启动 usergrowth_api
./gin_app > /data/code/logs/gin_app_run.log 2>&1 &

# 等待所有后台进程
wait