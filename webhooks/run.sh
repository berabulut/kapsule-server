#!/bin/bash

sudo kill -9 `sudo lsof -t -i:4043`

echo "Building webhooks-server"
go build -o webhooks-server
echo "Starting webhooks-server"
./webhooks-server &disown

