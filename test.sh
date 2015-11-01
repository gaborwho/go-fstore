#!/bin/sh

bin/server &
SERVER="$!"

curl -X POST --data 'hello' localhost:8081/send/
echo

curl -X POST --data 'hello2' localhost:8081/send/
echo

curl -X GET localhost:8081
echo

echo "Finished"

kill $SERVER

