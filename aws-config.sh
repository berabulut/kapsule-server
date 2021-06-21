#!/bin/bash

set -e

if [ -f .env ]; then
    # Load Environment Variables
    export $(cat .env | grep -v '#' | awk '/=/ {print $1}')
fi


# install python
PYTHON_VERSION=`python3 --version`

if [[ $PYTHON_VERSION == *"Python 3.8"* ]]; then
	echo "Python is already installed"
else
    echo "Installing Python"
    sudo apt update
    sudo apt install software-properties-common -y
    sudo add-apt-repository ppa:deadsnakes/ppa -y
    sudo apt update
    sudo apt install python3.8 -y
    python3 --version
    sudo apt install python3-pip -y
fi

# install aws-cli
echo "Installing AWS CLI"
sudo python3 -m pip install awscli -y

# configure aws-cli 
echo "[default]" >|  ~/.aws/credentials
echo "aws_access_key_id=${AWS_ACCESS_KEY_ID}" >> ~/.aws/credentials 
echo "aws_secret_access_key=${AWS_SECRET_ACCESS_KEY}" >> ~/.aws/credentials

echo "[default]" >|  ~/.aws/config
echo "region=${AWS_DEFAULT_REGION}" >> ~/.aws/config
echo "output=${OUTPUT_FORMAT}" >> ~/.aws/config


