#!/bin/bash

set -e


if [ $1 == "kapsule-server" ]; then
    cd ..

    echo "Stopping containers"
    docker-compose down  

    echo "Pulling updates for kapsule-server"   
    PULL=`git pull`
    
    if [ "$PULL" != "Already up to date." ]; then
        echo "Deploying with updates"
        sh deploy.sh
    fi

else
    #cd ../$1
    cd .. 
    sh deploy.sh
fi


