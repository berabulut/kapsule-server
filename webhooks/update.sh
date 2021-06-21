#!/bin/bash

set -e


if [ $1 == "kapsule-server" ]; then
    cd ..

    PULL=`git pull`
    if [ "$PULL" != "Already up to date." ]; then
        sh deploy.sh
    fi

else
    #cd ../$1
    cd .. 
    sh deploy.sh
fi


