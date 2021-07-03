#!/bin/bash


if [ -f ../.env ]; then
    # Load Environment Variables
    export $(cat ../.env | grep -v '#' | awk '/=/ {print $1}')
fi

ALIVE=`sudo lsof -i -P -n | grep ${WEBHOOKS_SERVER_PORT}`
echo ALIVE

if [ ${#ALIVE} == 0 ]; then
    echo "Executing run.sh"
    sh run.sh
fi
