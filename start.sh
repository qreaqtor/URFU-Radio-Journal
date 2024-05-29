#!/bin/bash

CLEAR=$1

if [ -f .env ]
then
    export $(cat .env | tr -d '\r' | sed 's/#.*//g' | xargs)
fi

docker-compose up -d

read
if [[ $CLEAR == '-c' ]]; then
    docker-compose down -v && docker rmi urfu-radio-journal-radiojournal
fi
