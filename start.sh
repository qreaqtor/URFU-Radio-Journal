#!/bin/bash

if [ -f .env ]
then
  export $(cat .env | tr -d '\r' | sed 's/#.*//g' | xargs)
fi

docker-compose up -d

read
docker-compose down -v && docker rmi urfu-radio-journal-radiojournal
