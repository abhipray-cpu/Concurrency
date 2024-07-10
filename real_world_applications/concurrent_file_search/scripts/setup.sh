#!/bin/bash

# Ensure the script is run as root
if [ "$(id -u)" != "0" ]; then
   echo "This script must be run as root" 1>&2
   exit 1
fi

# Check if Go is installed
if ! command -v go &> /dev/null
then
    echo "Go could not be found. Please install Go."
    exit 1
fi

echo "Go is installed. Proceeding with the setup."
# Proceed with any additional setup steps here