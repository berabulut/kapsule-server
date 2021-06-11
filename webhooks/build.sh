sudo kill -9 `sudo lsof -t -i:4043`

go build
./webhooks &disown

