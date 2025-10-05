#!/bin/sh
# export mongoServer=mongodb://localhost:27017/opem
export mongoServer=mongodb+srv://marioimperato:Y7csjKg4l4c1egiA@maidevfundamentalsclust.x5i5iii.mongodb.net/opem

## apicore_session
mongoexport --collection=apicore_session    --db=opem  --uri=$mongoServer      --out=exports/apicore_session.json
## opem_system
mongoexport --collection=opem_system       --db=opem  --uri=$mongoServer      --out=exports/opem_system.json
## apicore_user
mongoexport --collection=apicore_user       --db=opem  --uri=$mongoServer      --out=exports/apicore_user.json
## mag_boxes
mongoexport --collection=mag_boxes          --db=opem  --uri=$mongoServer      --out=exports/mag_boxes.json
## mag_magazzini
mongoexport --collection=mag_magazzini      --db=opem  --uri=$mongoServer      --out=exports/mag_magazzini.json
## org_organizzazione
mongoexport --collection=org_organizzazione --db=opem  --uri=$mongoServer      --out=exports/org_organizzazione.json
## prd_prodotto
mongoexport --collection=opem_prodotto       --db=opem  --uri=$mongoServer      --out=exports/opem_prodotto.json



