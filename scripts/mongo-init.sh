#!/bin/sh
# cd r3ds9-apigtw
# ./apigtw-mongo-init.sh

cd r3ds9-apicore
./apicore-mongo-init.sh

cd ../r3ds9-apicms
./apicms-mongo-init.sh
