# Run all test suites
test-all: 
	@./run-tests.sh tests/*/

# Run specific test suite (e.g. make test-dex)
test-%: 
	@./run-tests.sh tests/$*/

# Environment targets
kind-up:
	kind create cluster --name oauth2-proxy-e2e --config=infra/kind/config.yaml

kind-down:
	kind delete cluster --name oauth2-proxy-e2e
