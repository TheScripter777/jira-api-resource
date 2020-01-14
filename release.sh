#!/bin/bash
# Utility script to simplify release procedure.
# Need to have credentials to both github and dockerhub to use.

# Building docker image
docker image build -t turnscoffeeintoscripts/jira-api-issue-resource:$1 .

if [[ $? -eq "0" ]]; then
    # Pushing docker image
    docker image push turnscoffeeintoscripts/jira-api-issue-resource:$1
else
    exit 1
fi

if [[ $? -eq "0" ]]; then
    # Tag in git when the image was successfully push
    git tag -a $1 -m "Tagging version $1"
    git push origin $1
else
    exit 1
fi
