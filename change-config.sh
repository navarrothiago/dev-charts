#!/bin/bash

# create bash script to toggle symbolic links between two files.
current=$(readlink kube.config)

if [ "${current}" == "kube.config.remote" ]; then
    ln -sf kube.config.local kube.config
    echo "Switched to local config"
else
    ln -sf kube.config.remote kube.config
    echo "Switched to remote config"
fi
