#!/bin/bash

MODE=$1

if [ -f .env ]
then
  export $(cat .env | tr -d '\r' | sed 's/#.*//g' | xargs)
fi

if [[ $MODE == "debug" ]]
then
  docker-compose up -d postgres
else
  docker-compose up -d
fi

read
docker-compose down && docker volume prune -f

if [[ $MODE != "debug" ]]
then
  docker rmi urfu-radio-journal-radiojournal
fi