#!/bin/sh
echo $ECHO_ENV
echo $DBHOST
sql-migrate up -dryrun
sql-migrate up
/go/bin/app
