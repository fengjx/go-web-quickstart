#!/bin/bash -e

echo "[`date`] WORK_DIR: ${WORK_DIR}"

APP_ENV=${APP_ENV:-test}

CUSTOM_CONFIG_FILE=${WORK_DIR}/configs/app-${APP_ENV}.yaml

echo "[`date`] Starting server..."
echo "[`date`] Custom config: ${CUSTOM_CONFIG_FILE}"
./${APP_NAME} ${CUSTOM_CONFIG_FILE}
