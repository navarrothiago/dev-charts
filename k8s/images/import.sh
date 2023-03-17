#!/usr/bin/env bash
set -x
main() {
  # import tar files into registry
  local -r __dirname="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
  local -r __filename="${__dirname}/$(basename "${BASH_SOURCE[0]}")"

  local -r images="$(ls -1)"

  while read line; do
    sudo k3s ctr images import "${line}"
  done << EOF
$(ls -1)
EOF

}

main "$@"
