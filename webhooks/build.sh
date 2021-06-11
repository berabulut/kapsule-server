sudo kill -9 `sudo lsof -t -i:4043`

chmod 777 update.sh

go build -o webhooks-server
./webhooks-server &disown

