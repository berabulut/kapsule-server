# kapsule-server

Deployment project and server of kapsule.

## Set up

### Install Python

```
sudo apt update
sudo apt install software-properties-common
sudo add-apt-repository ppa:deadsnakes/ppa
sudo apt update
sudo apt install python3.8
python3 --version
sudo apt install python3-pip
```

### Install AWS CLI

`sudo python3 -m pip install awscli`

### Create .env file

`sudo nano .env`

Set this variables

``` 
AWS_ACCESS_KEY_ID=
AWS_SECRET_ACCESS_KEY=
AWS_ACCOUNT_ID=
AWS_DEFAULT_REGION=eu-central-1
OUTPUT_FORMAT=json
```

### Configure AWS CLI

`sh aws-config.sh`

## Create SSL Certificate

```
sudo ./init-letsencrypt.sh
```


## Deploy

```
sudo sh build.sh
```

## Automatic Update & Build 

#### Set a cron job on Ubuntu instance

`crontab -e`

Add this to end of file. 

```
*/5 * * * *    cd /home/ubuntu/kapsule-server/webhooks ; ./alive.sh
```

This cron job executes a script (every five minutes) that checks our webhooks server is alive or not. If it's dead it restarts it.  



