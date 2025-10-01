#!/bin/sh
# tail: brew install coreutils
# on mac-os head doesn't support negative numbers

mysql --user=$1 --password=$2 pcweb --host $3 --skip-ssl < delete-tables.sql

./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/tb_nazioni.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/tb_province.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/tb_comuni.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/companies.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/tb_dom_param.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/tb_branch.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/users.sql

./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/tb_prod.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/tb_prod_funct.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/tb_prod_param.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/tb_mag_card.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/tb_mag_param.sql

./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/groups.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/group_members.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/authorities.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/group_authorities.sql

# Big stuff need storage
#
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/tb_addresses.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/tb_applications.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/tb_cards.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/tb_file_err.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/tb_file_header.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/tb_file_header_errors.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/tb_file_header_records.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/tb_file_records.sql
./opem-pcweb-import-table.sh $1 $2 $3  ./sql-tables/tb_persons.sql

# Not migrated stuff
#
# ./opem-pcweb-import-table.sh $1 $2 $3  company_grouping.sql
# ./opem-pcweb-import-table.sh $1 $2 $3  company_grouping_companies.sql
# ./opem-pcweb-import-table.sh $1 $2 $3  tb_data_log.sql
# ./opem-pcweb-import-table.sh $1 $2 $3  tb_operation_log.sql
#./opem-pcweb-import-table.sh $1 $2 $3  tb_prog_address.sql
#./opem-pcweb-import-table.sh $1 $2 $3  tb_prog_appnum.sql
#./opem-pcweb-import-table.sh $1 $2 $3  tb_prog_card.sql
#./opem-pcweb-import-table.sh $1 $2 $3  tb_prog_mag.sql
#./opem-pcweb-import-table.sh $1 $2 $3  tb_prog_pers.sql
#./opem-pcweb-import-table.sh $1 $2 $3  tb_sub_companies.sql


