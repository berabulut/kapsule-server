#!/bin/bash

set -e

if [ -f .env ]; then
    # Load Environment Variables
    export $(cat .env | grep -v '#' | awk '/=/ {print $1}')
fi

./aws-config.sh

# get ECR credentials for docker pull
echo "Logging into AWS ECR"
aws --region ${AWS_DEFAULT_REGION} ecr get-login-password \
    | docker login \
        --password-stdin \
        --username AWS \
        "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_DEFAULT_REGION}.amazonaws.com"


# pull latest images
echo "Pulling images!"
docker pull $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/kapsule/kapsule-ui:latest
docker pull $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/kapsule/kapsule-redirect:latest
docker pull $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/kapsule/kapsule-api:latest

# rename them
docker image tag $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/kapsule/kapsule-ui:latest kapsule-ui:latest
docker image tag $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/kapsule/kapsule-redirect:latest kapsule-redirect:latest
docker image tag $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/kapsule/kapsule-api:latest kapsule-api:latest

# remove long named image
docker rmi $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/kapsule/kapsule-ui
docker rmi $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/kapsule/kapsule-redirect
docker rmi $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/kapsule/kapsule-api

docker-compose down  
docker-compose build  
docker-compose up -d