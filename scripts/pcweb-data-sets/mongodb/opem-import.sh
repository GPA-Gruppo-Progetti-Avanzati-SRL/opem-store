#!/bin/sh
# export mongoServer=mongodb://localhost:27017/opem
# export mongoServer=mongodb+srv://marioimperato:Y7csjKg4l4c1egiA@maidevfundamentalsclust.x5i5iii.mongodb.net/opem
export mongoServer=$1
mongosh $mongoServer --file drop-collections.js

## opem_boxes
mongoimport --collection=opem_boxes          --db=opem  --uri=$mongoServer      --file=exports/opem_boxes.json
## opem_cards
mongoimport --collection=opem_cards       --db=opem  --uri=$mongoServer      --file=exports/opem_cards.json
## opem_file
mongoimport --collection=opem_file       --db=opem  --uri=$mongoServer      --file=exports/opem_file.json
## opem_file_rows
mongoimport --collection=opem_file_rows       --db=opem  --uri=$mongoServer      --file=exports/opem_file_rows.json
## opem_magazzini
mongoimport --collection=opem_magazzini      --db=opem  --uri=$mongoServer      --file=exports/opem_magazzini.json
## opem_organizzazione
mongoimport --collection=opem_organizzazione --db=opem  --uri=$mongoServer      --file=exports/opem_organizzazione.json
## opem_persons
mongoimport --collection=opem_persons       --db=opem  --uri=$mongoServer      --file=exports/opem_persons.json
## opem_prodotto
mongoimport --collection=opem_prodotto       --db=opem  --uri=$mongoServer      --file=exports/opem_prodotto.json
## opem_system
mongoimport --collection=opem_system       --db=opem  --uri=$mongoServer      --file=exports/opem_system.json
## opem_territorio
mongoimport --collection=opem_territorio       --db=opem  --uri=$mongoServer      --file=exports/opem_territorio.json
## opem_user
mongoimport --collection=opem_user       --db=opem  --uri=$mongoServer      --file=exports/opem_user.json

mongosh $mongoServer --file fix-mig-02-create-indexes.js