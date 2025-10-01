const opemDbName = "opem"
let conn = db.getMongo();
let db = conn.getDB(opemDbName);

db["opem_system"].drop();
db["apicore_site"].drop();
db["apicore_domain"].drop();
db["apicore_kv"].drop();
db["apicore_session"].drop();
db["apicore_site"].drop();
db["apicore_user"].drop();
db["mag_boxes"].drop();
db["mag_magazzini"].drop();
db["org_organizzazione"].drop();
db["prd_prodotto"].drop();
db["sequence"].drop();
db["trt_territorio"].drop();


