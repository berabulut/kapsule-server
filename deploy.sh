#!/bin/bash

set -e

if [ -f .env ]; then
    # Load Environment Variables
    export $(cat .env | grep -v '#' | awk '/=/ {print $1}')
fi

sh aws-config.sh

# get ECR credentials for docker pull
aws --region ${AWS_DEFAULT_REGION} ecr get-login-password \
    | docker login \
        --password-stdin \
        --username AWS \
        "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com"


# pull latest images
docker pull $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/kapsule/kapsule-ui:latest
docker pull $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/kapsule/kapsule-service:latest

# rename them
docker image tag $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/kapsule/kapsule-ui:latest kapsule-ui:latest
docker image tag $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/kapsule/kapsule-service:latest kapsule-service:latest

# remove long named image
docker rmi $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/kapsule/kapsule-ui
docker rmi $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/kapsule/kapsule-service

docker-compose down  
docker-compose build  
docker-compose up -d