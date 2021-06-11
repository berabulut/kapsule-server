#!/bin/bash

set -e


if [ $1 == "kapsule-server" ]; then
    cd ../

    PULL=`git pull`
    if [ "$PULL" != "Already up to date." ]; then
        sh build.sh
    fi

else
    cd ../$1

    PULL=`git pull`
    if [ "$PULL" != "Already up to date." ]; then
        cd .. 
        sh build.sh
    fi
fi


