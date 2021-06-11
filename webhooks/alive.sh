#!/bin/bash

ALIVE=`sudo lsof -i -P -n | grep 4043`

if [ ${#ALIVE} == 0 ]; then
    ./build.sh
fi
