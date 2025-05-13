#!/bin/bash

DB_STRING="$DB_DRIVER://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME"

goose postgres "$DB_STRING" up