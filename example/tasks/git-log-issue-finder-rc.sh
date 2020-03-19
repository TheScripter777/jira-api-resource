#!/bin/bash

set -e

glifResult=$(glif --directory="$GIT_REPO_DIRECTORY" --tickets="$TICKETS_FILTER" --diff-tags="rc/@(LATEST-1)==>rc/@(LATEST)")

resultFile="${ISSUES_DIRECTORY}/${ISSUES_FILE}"

if [[ -f "${resultFile}" ]]; then
    rm -f ${resultFile}
fi

echo ${glifResult} >> ${resultFile}
echo ${glifResult}