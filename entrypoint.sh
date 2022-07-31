#!/bin/sh
echo "start container entry point"
set -e

if [ -f ./todo.db ]; then
  echo "delete $DB_PATH"
  rm $DB_PATH
fi

litestream restore -v -if-replica-exists -o $DB_PATH s3://$REPLICATE_BUCKET_NAME/replica
exec litestream replicate -exec "/usr/local/bin/app"
