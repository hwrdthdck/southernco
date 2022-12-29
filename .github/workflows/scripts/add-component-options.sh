#!/usr/bin/env sh
#
#   Copyright The OpenTelemetry Authors.
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.
#
# Takes the list of components from the CODEOWNERS file and inserts them
# as a YAML list in a GitHub issue template, then prints out the resulting
# contents.
#
# Note that this is script is intended to be POSIX-compliant since it is
# intended to also be called from the Makefile on developer machines,
# which aren't guaranteed to have Bash or a GNU userland installed.

if [ -z "${FILE}" ]; then
  echo 'FILE is empty, please ensure it is set.'
  exit 1
fi

CUR_DIRECTORY=$(dirname "$0")

# Get the line number for text within a file
get_line_number() {
  text=$1
  file=$2
  
  grep -n "${text}" "${file}" | awk '{ print $1 }' | grep -oE '[0-9]+'
}

LABELS=""

START_LINE=$(get_line_number '# Start Collector components list' "${FILE}")
END_LINE=$(get_line_number '# End Collector components list' "${FILE}")
TOTAL_LINES=$(wc -l "${FILE}" | awk '{ print $1 }')

head -n "${START_LINE}" "${FILE}"
for COMPONENT in $(sh "${CUR_DIRECTORY}/get-components.sh"); do
  TYPE=$(echo "${COMPONENT}" | cut -f1 -d'/')
  REST=$(echo "${COMPONENT}" | cut -f2- -d'/' | sed "s%${TYPE}/%/%" | sed "s%${TYPE}\$%%")
  LABEL=""

  if [ -z "${TYPE}" ] || [ -z "${REST}" ]; then
    LABEL="${COMPONENT}"
  else
    LABEL="${TYPE}/${REST}"
  fi

  LABELS="${LABELS}${LABEL}\n"
done
printf "%s" "$LABELS" | sort | awk '{ printf "      - %s\n",$1 }'
tail -n $((TOTAL_LINES-END_LINE+1)) "${FILE}"

