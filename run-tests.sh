#!/bin/bash

for test_dir in "$@"; do
    echo "Running tests in ${test_dir}"
    
    # Setup environment
    if [ -f "${test_dir}/setup.sh" ]; then
        (cd "${test_dir}" && ./setup.sh up)
    fi
    
    # Run tests
    (cd "${test_dir}" && go test ./...)
    
    # Teardown
    if [ -f "${test_dir}/setup.sh" ]; then
        (cd "${test_dir}" && ./setup.sh down)
    fi
done
