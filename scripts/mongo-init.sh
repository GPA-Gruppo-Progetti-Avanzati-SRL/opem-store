#!/bin/sh
# cd r3ds9-apigtw
# ./apigtw-mongo-init.sh

# export mongoServer=mongodb://localhost:27017/opem
export mongoServer=mongodb+srv://marioimperato:Y7csjKg4l4c1egiA@maidevfundamentalsclust.x5i5iii.mongodb.net/opem

cd opem-store-system
./opem-store-system-mongo-init.sh $mongoServer

cd ../opem-store-cms
./opem-store-cms-mongo-init.sh $mongoServer

cd ../opem-store-territorio
./opem-store-territorio-mongo-init.sh $mongoServer

cd ../opem-store-prodotto
./opem-store-prodotto-mongo-init.sh $mongoServer

cd ../opem-store-organizzazione
./opem-store-organizzazione-mongo-init.sh $mongoServer

cd ../opem-store-magazzino
./opem-store-magazzino-mongo-init.sh $mongoServer
