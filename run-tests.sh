#!/bin/bash

# Exit manually to ensure teardown
set +e

overall_status=0
for test_dir in "$@"; do
    echo "Running tests in ${test_dir}"
    test_status=0
    
    # Setup environment
    if [ -f "${test_dir}/setup.sh" ]; then
        (cd "${test_dir}" && ./setup.sh up) || {
            echo "Setup failed for ${test_dir}"
            test_status=1
        }
    fi
    
    # Run tests if setup succeeded
    if [ $test_status -eq 0 ]; then
        (cd "${test_dir}" && go test ./...) || test_status=1
    else
        echo "Skipping tests due to setup failure in ${test_dir}"
    fi
    
    # Teardown - always execute if setup.sh exists
    if [ -f "${test_dir}/setup.sh" ]; then
        (cd "${test_dir}" && ./setup.sh down) || {
            echo "Teardown failed for ${test_dir}"
            # Don't override test status if tests already failed
            [ $test_status -eq 0 ] && test_status=1
        }
    fi

    # Update overall status
    if [ $test_status -ne 0 ]; then
        overall_status=1
    fi
done

exit $overall_status
