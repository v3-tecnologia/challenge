#!/bin/sh

# Inicia o NATS no background
/usr/local/bin/nats-server -js -p 4222 --user admin --pass 123 &

# Inicia a api Go
/usr/local/bin/app
