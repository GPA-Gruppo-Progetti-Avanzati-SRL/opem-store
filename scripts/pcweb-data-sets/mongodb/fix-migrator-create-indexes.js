const opemDbName = "opem"
let conn = db.getMongo();
let db = conn.getDB(opemDbName);

db["mag_boxes"].createIndex({ "site": 1, "_bid": 1, "_et": 1 }, { unique: true });
db["mag_magazzini"].createIndex({  "site": 1, "_bid": 1, "_et": 1 }, { unique: true });
db["org_organizzazione"].createIndex({  "site": 1, "_bid": 1, "_et": 1 }, { unique: true });
db["opem_prodotto"].createIndex({  "site": 1, "_bid": 1, "_et": 1 }, { unique: true });
db["opem_system"].createIndex({ "_bid": 1, "_et": 1 }, { unique: false });

