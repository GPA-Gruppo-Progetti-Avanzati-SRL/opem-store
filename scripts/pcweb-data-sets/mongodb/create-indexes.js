const opemDbName = "opem"
let conn = db.getMongo();
let db = conn.getDB(opemDbName);

db["mag_boxes"].createIndex({ "_bid": 1, "_et": 1 }, { unique: true });
db["mag_magazzini"].createIndex({ "_bid": 1, "_et": 1 }, { unique: true });
db["org_organizzazione"].createIndex({ "_bid": 1, "_et": 1 }, { unique: true });
db["prd_prodotto"].createIndex({ "_bid": 1, "_et": 1 }, { unique: true });

