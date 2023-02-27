#!/bin/bash -e

exec > >(tee -a /var/log/web/app.log|logger -t server -s 2>/dev/console) 2>&1

APP_ENV=${APP_ENV:-test}

echo "[`date`] Running entrypoint script in the '${APP_ENV}' environment..."

CONFIG_FILE=./configs/app.yml

echo "[`date`] Starting server..."
./server -config ${CONFIG_FILE} >> /var/log/web/app.log 2>&1
