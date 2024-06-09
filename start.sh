#!/bin/bash

FILE=$1
CLEAR=$2

if [ -f $FILE ]
then
    export $(cat $FILE | tr -d '\r' | sed 's/#.*//g' | xargs)
fi

if [[ $ENV == 'dev' ]]; then
    docker-compose up -d minio postgres prometheus grafana
else
    export CONFIG_PATH=$FILE
    docker-compose up -d
fi

if [[ $CLEAR == '-c' ]]; then
    read
    docker-compose down -v && docker rmi urfu-radio-journal-radiojournal
fi
