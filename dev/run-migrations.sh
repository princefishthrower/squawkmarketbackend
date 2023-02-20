#!/bin/bash

# source .env.sh
source .env.sh

# run migrations
migrate -path db/migrations -database "$SQLITE_CONNECTION_STRING" up