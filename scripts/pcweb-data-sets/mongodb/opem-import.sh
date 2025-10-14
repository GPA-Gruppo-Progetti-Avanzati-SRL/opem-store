#!/bin/sh
# export mongoServer=mongodb://localhost:27017/opem
# export mongoServer=mongodb+srv://marioimperato:Y7csjKg4l4c1egiA@maidevfundamentalsclust.x5i5iii.mongodb.net/opem
export mongoServer=$1

mongosh $mongoServer --file drop-collections.js

echo ##### opem_system
mongoimport --collection=opem_system       --db=opem  --uri=$mongoServer      --file=exports/opem_system.json
echo ##### apicore_session
mongoimport --collection=apicore_session    --db=opem  --uri=$mongoServer      --file=exports/apicore_session.json
echo ##### apicore_user
mongoimport --collection=apicore_user       --db=opem  --uri=$mongoServer      --file=exports/apicore_user.json
echo ##### mag_boxes
mongoimport --collection=mag_boxes          --db=opem  --uri=$mongoServer      --file=exports/mag_boxes.json
echo ##### mag_magazzini
mongoimport --collection=mag_magazzini      --db=opem  --uri=$mongoServer      --file=exports/mag_magazzini.json
echo ##### org_organizzazione
mongoimport --collection=org_organizzazione --db=opem  --uri=$mongoServer      --file=exports/org_organizzazione.json
echo ##### opem_prodotto
mongoimport --collection=opem_prodotto       --db=opem  --uri=$mongoServer      --file=exports/opem_prodotto.json

mongosh $mongoServer --file create-indexes.js