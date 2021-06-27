# kapsule-server

Deployment project and server of kapsule.

## Local Development

```
docker-compose -f local.yaml build
docker-compose -f local.yaml up
```

## Set up

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

`./aws-config.sh`

## Create SSL Certificate

```
sudo ./init-letsencrypt.sh
```

## Deploy

```
sh deploy.sh
```

## Automatic Update & Build

### Set a cron job on Ubuntu instance

`sh add-cronjob.sh`

### Manual way

`crontab -e`

Add this to end of file.

```
*/5 * * * *  cd /home/ubuntu/kapsule-server/webhooks ; sh alive.sh
```

This cron job executes a script (every five minutes) that checks our webhooks server is alive or not. If it's dead it restarts it.
