# run
#if [ "$1" == "0" ]
#then
#  SERVICE_NAME="kitex-client"
#  TARGET_SERVICE="kitex-server"
#  ./output/bin/kitex-proxyless -serviceName ${SERVICE_NAME} -targetServiceName ${TARGET_SERVICE}
#else
#  SERVICE_NAME="kitex-server"
#  TARGET_SERVICE=" "
#fi
#
#if [ "$1" == "0" ]
#then
#  ./output/bin/${SERVICE_NAME} -serviceName ${SERVICE_NAME} -targetServiceName ${TARGET_SERVICE}
#else
#  ./output/bin/${SERVICE_NAME} -serviceName ${SERVICE_NAME}
#fi

./output/bin/kitex -serviceName "kitex-server" -targetServiceName "NA"