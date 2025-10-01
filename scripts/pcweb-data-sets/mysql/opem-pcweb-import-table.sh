#!/bin/sh
# tail: brew install coreutils
# on mac-os head doesn't support negative numbers

tail -n +16 $4  > tmp1.sql
ghead -n -6  tmp1.sql >  tmp2.sql
rm tmp1.sql
echo "doing.. $4"
mysql --user=$1 --password=$2 pcweb --host $3 --skip_ssl < tmp2.sql
rm tmp2.sql
