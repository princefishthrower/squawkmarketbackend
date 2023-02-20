#!/bin/bash

# Check that the first argument exists, and if not, exit
if [ -z "$1" ]; then
  echo "Need migration name!, i.e., usage: /bin/bash dev/draft-migration.sh <migration_name>"
  exit 1
fi

migrate create -ext sql -dir db/migrations $1