#!/bin/bash

set -e  # Exit on error

# Ensure DATABASE_URL is set
if [ -z "$DATABASE_URL" ]; then
  echo "Error: DATABASE_URL is not set"
  exit 1
fi

# Run Goose migrations
goose -dir migrations postgres "$DATABASE_URL" up
