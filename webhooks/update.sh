#!/bin/bash

set -e


if [ $1 == "kapsule-server" ]; then
    cd ..

    PULL=`git pull`
    if [ "$PULL" != "Already up to date." ]; then
        ./deploy.sh
    fi

else
    #cd ../$1
    cd .. 
    ./deploy.sh
fi


