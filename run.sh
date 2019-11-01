#!/bin/bash
./wait-for-it.sh -h $DBHOST -p $DBPORT -t 90
if [ $? -eq 0 ]; then
  sql-migrate up -dryrun
  sql-migrate up
  /go/bin/app
else
  exit 1
fi
