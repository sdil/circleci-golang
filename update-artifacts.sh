#!/bin/bash

# Required vars: CIRCLECI_TOKEN, WORK_DIR

# Get the latest artifact
curl https://circleci.com/api/v1.1/project/github/sdil/circleci-golang/latest/artifacts?circle-token=$CIRCLECI_TOKEN | jq .[0].url | xargs -n 1 curl -o $WORK_DIR

# Backup database
# Restart service
# Test healthz endpoint
curl -i localhost:8000/healthz
