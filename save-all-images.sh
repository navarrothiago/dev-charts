#!/usr/bin/env bash
main() {

  local -r __dirname="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
  local -r __filename="${__dirname}/$(basename "${BASH_SOURCE[0]}")"

  local -r filepath=images.txt

  rm -f "${__dirname}"/*.tar

  #kubectl get events -n <namespace> | grep image | awk -F\" '{print $2}' | sort | uniq > "${filepath}"

  while read line; do
    output=$(echo "${line}" | cut -d: -f 1 | rev | cut -d/ -f 1 | rev).tar
    docker pull "${line}"
    docker save "${line}" -o "${output}"
  done < "${filepath}"


  exit 0
}
main "$@"
