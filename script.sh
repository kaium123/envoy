#!/usr/bin/env bash

# Create kong.yml file if it doesn't exist
touch ./app/kong.yml

# Modify config.yml and write to kong.yml
cat ./app/config.yml | sed 's/^/"/' | sed 's/$/"/' | sed 's/^/echo /' | sh > ./app/kong.yml
# cat config.yml | sed 's/^/"/' | sed 's/$/"/' | sed 's/^/echo /' | sh > kong.yml

# Start Kong
kong start
