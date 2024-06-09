#!/bin/bash

FILE=$1
CLEAR=$2

extract_value() {
  local key=$1
  local file=$2
  echo $(grep "^$key:" $file | awk '{ print $2 }' | tr -d '\r')
}

extract_value_by_context() {
  local context=$1
  local key=$2
  local lines_count=$3
  local file=$4
  echo $(grep -A$lines_count "^$context:" $file | grep "^  $key:" | awk '{ print $2 }' | tr -d '\r')
}

if [ -f "$FILE" ]; then
    export MINIO_USER=$(extract_value_by_context minio user 2 $FILE)
    export MINIO_PASSWORD=$(extract_value_by_context minio password 2 $FILE)
    export POSTGRES_USER=$(extract_value_by_context postgres user 3 $FILE)
    export POSTGRES_PASSWORD=$(extract_value_by_context postgres password 3 $FILE)
    export POSTGRES_DATABASE=$(extract_value_by_context postgres database 3 $FILE)
fi

ENV=$(extract_value env $FILE)

if [[ $ENV == 'prod' ]]; then
    export CONFIG_PATH=$FILE
    docker-compose up -d
else
    docker-compose up -d minio postgres prometheus grafana
fi

if [[ $CLEAR == '-c' ]]; then
    read
    docker-compose down -v && docker rmi urfu-radio-journal-radiojournal
fi
