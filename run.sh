#!/bin/sh
sql-migrate up -dryrun
sql-migrate up
/go/bin/app
