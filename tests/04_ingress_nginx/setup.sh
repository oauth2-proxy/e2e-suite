#!/bin/bash
set -euo pipefail

case "$1" in
  up)
    ../../infra/kind/check-or-start.sh 
    helm install ingress-nginx-test .
    sleep 5
    ;;
  down)
    helm uninstall ingress-nginx-test
    ;;
  *)
    echo "Usage: $0 {up|down}"
    exit 1
esac
