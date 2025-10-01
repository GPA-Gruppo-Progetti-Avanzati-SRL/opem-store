#!/bin/sh
# tail: brew install coreutils
# on mac-os head doesn't support negative numbers

mysql --user=$1 --password=$2 pcweb --host $3 --skip-ssl < delete-tables.sql
