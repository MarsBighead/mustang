#/bin/bash

version=`psql -U usgmtr -h localhost -d usgmtr << EOF
    SELECT version FROM "SchemaInfo";
EOF`
 
version=$(echo ${version} | awk -F ' ' '{print $3}')
echo "$version"
echo "version is $version"
