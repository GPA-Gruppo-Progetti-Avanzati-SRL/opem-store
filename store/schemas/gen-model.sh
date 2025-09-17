#!/bin/bash

# commons
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name SysInfo --out-dir ../.. --with-format --with-bak
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name UserRole --out-dir ../.. --with-format --with-bak
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name App --out-dir ../.. --with-format --with-bak
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name FileVariant --out-dir ../.. --with-format --with-bak
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name FileReference --out-dir ../.. --with-format --with-bak

# api-cms
tpm-morphia-cli gen entity --schema-file ./apicms/schema.yml --name entRefStruct --out-dir ../.. --with-format --with-bak
tpm-morphia-cli gen entity --schema-file ./apicms/schema.yml --name file --out-dir ../.. --with-format --with-bak

# api-core
tpm-morphia-cli gen entity --schema-file ./apicore/schema.yml --name member --out-dir ../.. --with-format --with-bak
tpm-morphia-cli gen entity --schema-file ./apicore/schema.yml --name domain --out-dir ../.. --with-format --with-bak
tpm-morphia-cli gen entity --schema-file ./apicore/schema.yml --name site --out-dir ../.. --with-format --with-bak
tpm-morphia-cli gen entity --schema-file ./apicore/schema.yml --name user --out-dir ../.. --with-format --with-bak
tpm-morphia-cli gen entity --schema-file ./apicore/schema.yml --name session --out-dir ../.. --with-format --with-bak
tpm-morphia-cli gen entity --schema-file ./apicore/schema.yml --name keyValue --out-dir ../.. --with-format --with-bak
tpm-morphia-cli gen entity --schema-file ./apicore/schema.yml --name keyValuePackage --out-dir ../.. --with-format --with-bak

tpm-morphia-cli gen entity --schema-file ./territorio/schema.yml --name nazione --out-dir ../.. --with-format --with-bak
tpm-morphia-cli gen entity --schema-file ./territorio/schema.yml --name provincia --out-dir ../.. --with-format --with-bak
tpm-morphia-cli gen entity --schema-file ./territorio/schema.yml --name comune --out-dir ../.. --with-format --with-bak

tpm-morphia-cli gen entity --schema-file ./prodotto/schema.yml --name prodotto --out-dir ../.. --with-format --with-bak

tpm-morphia-cli gen entity --schema-file ./organizzazione/schema.yml --name focalPoint --out-dir ../.. --with-format --with-bak
