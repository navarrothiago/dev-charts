# Air gapped machine

The following steps are used to create a tarball of all images used by the helm chart and push them to an air gapped machine:

- Deploy helm chart
- Extract images: `kubectl get events -n <namespace> | grep image | awk -F\" '{print $2}' | sort | uniq > "${filepath}"`
- Run `save-all-images.sh` script
- Push `*.tar` to air gapped machine
- Run `import.sh` inside the air gapped machine
