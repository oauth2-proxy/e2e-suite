#!/bin/bash

case "$1" in
  up)
    docker compose up --detach --wait --wait-timeout 120
    sleep 10
    ;;
  down)
    docker compose down
    ;;
  *)
    echo "Usage: $0 {up|down}"
    exit 1
esac
