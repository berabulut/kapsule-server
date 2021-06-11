# kapsule-server

Deployment project and server of kapsule.

## Set up

``` 
git clone https://github.com/berabulut/kapsule-ui.git 
git clone https://github.com/berabulut/kapsule.git
```

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
*/5 * * * *    cd /home/ubuntu/kapsule-server/webhooks ; ./cron.sh
```

This cron job executes a script (every five minutes) that checks our webhooks server is alive or not. If it's dead it restarts it.  


#### Run it manually

```
cd webhooks
./build.sh
```

