#!/bin/bash

# add --with-bak if needed
# commons
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name SysInfo --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name UserRole --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name App --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name FileVariant --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name FileReference --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name Address --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name BidTextPair --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name ValueTextPair --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name ValueRange --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name Note --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name Event --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name StatusCodeTextPair --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./commons/schema.yml --name Activity --out-dir ../.. --with-format

# api-cms
tpm-morphia-cli gen entity --schema-file ./apicms/schema.yml --name entRefStruct --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./apicms/schema.yml --name file --out-dir ../.. --with-format

# api-core
tpm-morphia-cli gen entity --schema-file ./apicore/schema.yml --name sequence --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./apicore/schema.yml --name member --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./apicore/schema.yml --name domain --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./apicore/schema.yml --name site --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./apicore/schema.yml --name user --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./apicore/schema.yml --name session --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./apicore/schema.yml --name keyValue --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./apicore/schema.yml --name keyValuePackage --out-dir ../.. --with-format

# territorio
tpm-morphia-cli gen entity --schema-file ./territorio/schema.yml --name nazione --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./territorio/schema.yml --name provincia --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./territorio/schema.yml --name comune --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./territorio/schema.yml --name territorio --out-dir ../.. --with-format

# prodotto
tpm-morphia-cli gen entity --schema-file ./prodotto/schema.yml --name prodotto --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./prodotto/schema.yml --name appDefinition --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./prodotto/schema.yml --name appNumberDefinition --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./prodotto/schema.yml --name HostProduct --out-dir ../.. --with-format

# organizzazione
tpm-morphia-cli gen entity --schema-file ./organizzazione/schema.yml --name focalPoint --out-dir ../.. --with-format

# magazzino
tpm-morphia-cli gen entity --schema-file ./magazzino/schema.yml --name magazzino --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./magazzino/schema.yml --name Info --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./magazzino/schema.yml --name Status --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./magazzino/schema.yml --name box --out-dir ../.. --with-format

# card
tpm-morphia-cli gen entity --schema-file ./card/schema.yml --name card --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./card/schema.yml --name CardApp --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./card/schema.yml --name CardHolder --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./card/schema.yml --name Person --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./card/schema.yml --name BirthInfo --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./card/schema.yml --name IdentityCard --out-dir ../.. --with-format

# file
tpm-morphia-cli gen entity --schema-file ./file/schema.yml --name File --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./file/schema.yml --name Stat --out-dir ../.. --with-format
tpm-morphia-cli gen entity --schema-file ./file/schema.yml --name FileRow --out-dir ../.. --with-format

# S£ Events
tpm-morphia-cli gen entity --schema-file ./s3-events/schema.yml --name ObjectStoreEvent --out-dir ../.. --with-format
