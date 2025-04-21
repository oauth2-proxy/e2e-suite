#!/bin/bash

CLUSTER_NAME="oauth2-proxy-e2e"
KIND_EXISTS=$(kind get clusters | grep -w "${CLUSTER_NAME}")
CURRENT_CTX=$(kubectl config current-context)
EXPECTED_CTX="kind-${CLUSTER_NAME}"

if [ -z "${KIND_EXISTS}" ]; then
    echo "Error: Kind cluster '${CLUSTER_NAME}' is not running."
    kind create cluster --name "${CLUSTER_NAME}" --config=$(dirname -- "${BASH_SOURCE[0]}")/config.yaml
else
    echo "Kind cluster '${CLUSTER_NAME}' is running."
fi

if [ "${CURRENT_CTX}" != "${EXPECTED_CTX}" ]; then
    echo "Warning: kubectl context is '${CURRENT_CTX}', but expected '${EXPECTED_CTX}'."
    kubectl config use-context "${EXPECTED_CTX}"
else
    echo "kubectl context is correctly set to '${EXPECTED_CTX}'."
fi

echo "Info: Ensure nginx is installed as default ingress"
kubectl apply -f https://kind.sigs.k8s.io/examples/ingress/deploy-ingress-nginx.yaml
