#!/bin/sh

# Source the secret file
source /etc/ssm/.env

# Remove .env file
rm /etc/ssm/.env

# Execute the main command
exec /app/main
