#!/bin/sh

set -e


if [ $1 == "kapsule-server" ]; then
    cd ../
    git pull
    sh build.sh

else
    echo "cd ../$1"
    cd ../$1
    git pull
    cd .. 
    sh build.sh
fi


