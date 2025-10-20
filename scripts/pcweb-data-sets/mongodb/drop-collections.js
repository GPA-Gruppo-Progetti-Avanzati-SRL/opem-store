const opemDbName = "opem"
let conn = db.getMongo();
let db = conn.getDB(opemDbName);

db["opem_addresses"].drop();
db["opem_boxes"].drop();
db["opem_cards"].drop();
db["opem_organizzazione"].drop();
db["opem_persons"].drop();
db["opem_prodotto"].drop();
db["opem_prodotto_funct"].drop();
db["opem_system"].drop();
db["opem_system_dom_param"].drop();
db["opem_territorio"].drop();
db["opem_jobs"].drop();
