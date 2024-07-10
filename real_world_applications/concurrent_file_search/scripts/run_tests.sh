#!/bin/bash

# Run the test with verbose output
go test -v ../test/search

# Check if the test passed
if [ $? -eq 0 ]; then
    echo "Tests passed!"
else
    echo "Tests failed."
    # Handle failure (e.g., send a notification, exit with failure)
    exit 1
fi