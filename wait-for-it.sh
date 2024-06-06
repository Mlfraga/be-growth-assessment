#!/usr/bin/env bash
# Use this script to test if a given TCP host/port are available

HOST="$1"
PORT="$2"
shift 2
cmd="$@"

while ! nc -z $HOST $PORT; do
  >&2 echo "Waiting for $HOST:$PORT..."
  sleep 1
done

>&2 echo "$HOST:$PORT is available - executing command"
exec $cmd
