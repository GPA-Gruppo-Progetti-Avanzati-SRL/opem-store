const opemDbName = "opem"
let conn = db.getMongo();
let db = conn.getDB(opemDbName);

db["opem_boxes"].createIndex({ "site": 1, "_bid": 1, "_et": 1 }, { unique: true });
db["opem_magazzini"].createIndex({  "site": 1, "_bid": 1, "_et": 1 }, { unique: true });
db["opem_organizzazione"].createIndex({  "site": 1, "_bid": 1, "_et": 1 }, { unique: true });
db["opem_prodotto"].createIndex({  "site": 1, "_bid": 1, "_et": 1 }, { unique: true });
db["opem_system"].createIndex({ "_bid": 1, "_et": 1 }, { unique: false });
db["opem_territorio"].createIndex({ "_bid": 1, "_et": 1 }, { unique: true });
