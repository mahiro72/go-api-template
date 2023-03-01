#!/bin/bash
 
ENV_FILE="./.env/.env"
SAMPLE_ENV_FILE="./.env/.sample.env"

echo "=== check env file ==="
 
if [ ! -e $ENV_FILE ]; then
  if [ ! -e $SAMPLE_ENV_FILE]; then
    echo "Error: Sample env file not found."
    exit 1
  fi
  cat $SAMPLE_ENV_FILE > $ENV_FILE
  echo ".env file created. (from .sample.env)"
fi

echo "=== .env file found ==="

