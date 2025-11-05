# OAuth2 Proxy End-to-End Test Suite
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Foauth2-proxy%2Fe2e-suite.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Foauth2-proxy%2Fe2e-suite?ref=badge_shield)


This repository contains end-to-end tests for OAuth2 Proxy, verifying authentication flows with different identity providers like (Dex, Keycloak) or the integration with different environments and loadbalancers and mesh systems like nginx, traefik, istio, etc.

## Project Structure

```
├── infra/                  # Infra scripts and configuration files for more complex environments
│ ├── azure/                # Azure setup files
│ └── kind/                 # Kind cluster spec files
├── internal/               # Shared test utilities and page objects
│ ├── pages/                # Page object models for test pages
│ └── utils/                # Test utilities (Playwright, Testcontainers)
├── tests/                  # Test suites
│ ├── 01_smoke/             # Basic login flow test with dex
│ ├── 02_keycloak/          # Keycloak specific tests
│ └── 03_nginx/             # Nginx integration tests
│ └── 04_ingress-nginx/     # Ingress-Nginx integration tests
└── Makefile                # Makefile for triggering setup and test cases
```

## Prerequisites

- Go 1.20+
- Docker
- Playwright browser dependencies
    - `go run github.com/playwright-community/playwright-go/cmd/playwright@latest install --with-deps chromium`
- Kind (for local Kubernetes testing)
- Helm (for Kubernetes dependencies)

## Running Test locally

```sh
# Run all tests
make test-all

# Run specific test suite
make test-01_smoke
make test-02_keycloak

# For debugging
cd tests/01_smoke/
./setup.sh up # Configure environment (docker-compose in this case)

# Run the test files with your favourite debugger and IDE
# After you are done. Don't forget to cleanup
./setup.sh down
```

## Running Test workflow

Currently the e2e test suite is only triggered on demand by the maintainers of oauth2-proxy and can be triggered from within PRs in the upstream [oauth2-proxy repository](https://github.com/oauth2-proxy/oauth2-proxy) by writing comments.


### How to trigger?

If you are a maintainer / member of the [oauth2-proxy organization](https://github.com/oauth2-proxy), you can simple write a comment with the content `/e2e` on any PR which will automatically trigger the necessary workflow from within the PR. Which builds a temporary version of oauth2-proxy using the contents of the PR, and execute the test cases against it. Furthermore it attaches a commit status and labels to the PR.


| Commands             | Description                                                                                                                                        |
| -------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `/e2e`               | Executes the tests from the main branch of the e2e test suite repository.                                                                          |
| `/e2e <branch-name>` | Executes the tests from a custom branch of the e2e test suite repository.                                                                          |
| `/e2e PR-<number>`   | Executes the tests from an open PR of the e2e test suite repository. (Useful if new test cases are being added while an upstream PR is still open) |


## Contribution Guide

- Keep test cases focused and independent. 
- Use page object pattern for UI interactions whenever possible
- Include necessary setup scripts for new dependencies. Ensure they work locally, with a debugger and in the test workflow


Before starting to write a new test case. Try to identify if a test suite already exists.
1. Consider what you are testing for: Provider specifics or environment integration?
2. Check if we already have a test suite for the provider or integration

### Adding Test Cases for existing provider or integration

1. Locate the corresponding test suite folder (tests/01_smoke, tests/02_keycloak or tests/03_nginx)
2. Extend the page objects `internal/pages/` if needed
3. Create new test cases by adding new It() blocks in the existing test files and follow the established pattern of:
    - Container setup
    - Page interactions
    - Assertions

### Adding Test Cases for new provider or integration

1. Create a new test suite:
    - Make a new numbered folder in `tests/` (e.g.: `99_myprovider`)
    - Copy the structure from existing suites:
      - `compose.yaml / kustomize.yaml / values.yaml` - For setting up the provider and infrastructure
      - `setup.sh` - Initialization script for starting and stopping the provider and infrastructure
      - `configs/` - OAuth2 Proxy configurations for test cases
      - `e2e_test.go` - Test cases
2. Implement provider-specific pages if necessary
3. Implement your actual test flows and assertions


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Foauth2-proxy%2Fe2e-suite.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Foauth2-proxy%2Fe2e-suite?ref=badge_large)