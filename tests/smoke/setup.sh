#!/bin/bash

case "$1" in
  up)
    docker-compose up -d
    ;;
  down)
    docker-compose down
    ;;
  *)
    echo "Usage: $0 {up|down}"
    exit 1
esac
