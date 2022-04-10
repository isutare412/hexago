#!/bin/bash

ROOTDIR=$(dirname $(dirname $0))

PROJECT_NAME=hexago
MODE=dev

COMPOSE_FILE=$ROOTDIR/deployments/docker-compose.infra.yaml
ENV_FILE=$ROOTDIR/deployments/$MODE.env

COMPOSE_CMD="docker-compose -f $COMPOSE_FILE --env-file $ENV_FILE -p $PROJECT_NAME"

PS3='Please enter your choice: '
options=("up" "down" "ps" "logs" "quit")
select opt in "${options[@]}"
do
  case $opt in
    "up")
      CMD="$COMPOSE_CMD up -d"
      echo $CMD && eval $CMD
      break
      ;;
    "down")
      CMD="$COMPOSE_CMD down"
      echo $CMD && eval $CMD
      break
      ;;
    "ps")
      CMD="$COMPOSE_CMD ps"
      echo $CMD && eval $CMD
      break
      ;;
    "logs")
      CMD="$COMPOSE_CMD logs -f"
      echo $CMD && eval $CMD
      break
      ;;
    "quit")
      exit 0
      ;;
    *) echo "invalid option $REPLY";;
  esac
done
