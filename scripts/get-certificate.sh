#!/bin/bash

url=$1
if [ -z "$url" ]; then
    echo "Usage: $0 <url:port>"
    exit 1
fi
openssl s_client -connect "${url}" -starttls smtp
