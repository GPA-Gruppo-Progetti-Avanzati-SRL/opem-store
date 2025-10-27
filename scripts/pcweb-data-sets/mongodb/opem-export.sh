#!/bin/sh
# export mongoServer=mongodb://localhost:27017/opem
export mongoServer=mongodb+srv://marioimperato:Y7csjKg4l4c1egiA@maidevfundamentalsclust.x5i5iii.mongodb.net/opem

## opem_boxes
mongoexport --collection=opem_boxes          --db=opem  --uri=$mongoServer      --out=exports/opem_boxes.json
## opem_cards
mongoexport --collection=opem_cards       --db=opem  --uri=$mongoServer      --out=exports/opem_cards.json
## opem_file
mongoexport --collection=opem_file       --db=opem  --uri=$mongoServer      --out=exports/opem_file.json
## opem_file_rows
mongoexport --collection=opem_file_rows       --db=opem  --uri=$mongoServer      --out=exports/opem_file_rows.json
## opem_magazzini
mongoexport --collection=opem_magazzini      --db=opem  --uri=$mongoServer      --out=exports/opem_magazzini.json
## opem_organizzazione
mongoexport --collection=opem_organizzazione --db=opem  --uri=$mongoServer      --out=exports/opem_organizzazione.json
## opem_persons
mongoexport --collection=opem_persons       --db=opem  --uri=$mongoServer      --out=exports/opem_persons.json
## opem_prodotto
mongoexport --collection=opem_prodotto       --db=opem  --uri=$mongoServer      --out=exports/opem_prodotto.json
## opem_system
mongoexport --collection=opem_system       --db=opem  --uri=$mongoServer      --out=exports/opem_system.json
## opem_territorio
mongoexport --collection=opem_territorio       --db=opem  --uri=$mongoServer      --out=exports/opem_territorio.json
## opem_user
mongoexport --collection=opem_user       --db=opem  --uri=$mongoServer      --out=exports/opem_user.json