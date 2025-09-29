#!/bin/sh
# export mongoServer=mongodb://localhost:27017/opem
export mongoServer=mongodb+srv://marioimperato:Y7csjKg4l4c1egiA@maidevfundamentalsclust.x5i5iii.mongodb.net/opem

mongosh $mongoServer --file drop-collections.js

echo ##### apicore_site
mongoimport --collection=apicore_site       --db=opem  --uri=$mongoServer      --file=apicore_site.json
echo ##### apicore_domain
mongoimport --collection=apicore_domain     --db=opem  --uri=$mongoServer      --file=apicore_domain.json
echo ##### apicore_kv
mongoimport --collection=apicore_kv         --db=opem  --uri=$mongoServer      --file=apicore_kv.json
echo ##### apicore_session
mongoimport --collection=apicore_session    --db=opem  --uri=$mongoServer      --file=apicore_session.json
echo ##### apicore_user
mongoimport --collection=apicore_user       --db=opem  --uri=$mongoServer      --file=apicore_user.json
echo ##### mag_boxes
mongoimport --collection=mag_boxes          --db=opem  --uri=$mongoServer      --file=mag_boxes.json
echo ##### mag_magazzini
mongoimport --collection=mag_magazzini      --db=opem  --uri=$mongoServer      --file=mag_magazzini.json
echo ##### org_organizzazione
mongoimport --collection=org_organizzazione --db=opem  --uri=$mongoServer      --file=org_organizzazione.json
echo ##### prd_prodotto
mongoimport --collection=prd_prodotto       --db=opem  --uri=$mongoServer      --file=prd_prodotto.json
echo ##### sequence
mongoimport --collection=sequence           --db=opem  --uri=$mongoServer      --file=sequence.json
echo ##### trt_territorio
mongoimport --collection=trt_territorio    --db=opem   --uri=$mongoServer      --file=trt_territorio.json

mongosh $mongoServer --file create-indexes.js