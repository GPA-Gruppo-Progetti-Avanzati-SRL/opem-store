#!/bin/sh
# export mongoServer=mongodb://localhost:27017/opem
export mongoServer=mongodb+srv://marioimperato:Y7csjKg4l4c1egiA@maidevfundamentalsclust.x5i5iii.mongodb.net/opem

## apicore_domain
mongoexport --collection=apicore_domain     --db=opem  --uri=$mongoServer      --out=apicore_domain.json
## apicore_kv
mongoexport --collection=apicore_kv         --db=opem  --uri=$mongoServer      --out=apicore_kv.json
## apicore_session
mongoexport --collection=apicore_session    --db=opem  --uri=$mongoServer      --out=apicore_session.json
## apicore_site
mongoexport --collection=apicore_site       --db=opem  --uri=$mongoServer      --out=apicore_site.json
## apicore_user
mongoexport --collection=apicore_user       --db=opem  --uri=$mongoServer      --out=apicore_user.json
## mag_boxes
mongoexport --collection=mag_boxes          --db=opem  --uri=$mongoServer      --out=mag_boxes.json
## mag_magazzini
mongoexport --collection=mag_magazzini      --db=opem  --uri=$mongoServer      --out=mag_magazzini.json
## org_organizzazione
mongoexport --collection=org_organizzazione --db=opem  --uri=$mongoServer      --out=org_organizzazione.json
## prd_prodotto
mongoexport --collection=prd_prodotto       --db=opem  --uri=$mongoServer      --out=prd_prodotto.json
## sequence
mongoexport --collection=sequence           --db=opem  --uri=$mongoServer      --out=sequence.json
## trt_territorio
mongoexport --collection=trt_territorio    --db=opem  --uri=$mongoServer      --out=trt_territorio.json

