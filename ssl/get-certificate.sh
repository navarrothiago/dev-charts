#!/bin/bash
# This script checks whether an SMTP service is running at the given address
# The script uses the openssl s_client command to check for the STARTTLS
# response from the SMTP service
#
# Usage: $0 <url:port>

url=$1
if [ -z "$url" ]; then
    echo "Usage: $0 <url:port>"
    exit 1
fi

# starttls smtp
openssl s_client -connect "${url}" -starttls smtp
