#!/bin/bash
if [ "$1" == "0" ]
then
  SERVICE_NAME="kitex-client"
  TARGET_SERVICE="kitex-server"
else
  SERVICE_NAME="kitex-server"
  TARGET_SERVICE=" "
fi

# build
go mod tidy
go build -o output/bin/${SERVICE_NAME}

# run
if [ "$1" == "0" ]
then
  ./output/bin/${SERVICE_NAME} -serviceName ${SERVICE_NAME} -targetServiceName ${TARGET_SERVICE}
else
  ./output/bin/${SERVICE_NAME} -serviceName ${SERVICE_NAME}
fi