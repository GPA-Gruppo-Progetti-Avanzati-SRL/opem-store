const opemDbName = "opem"
let conn = db.getMongo();
let db = conn.getDB(opemDbName);

db["apicore_domain"].drop();
db["apicore_kv"].drop();
db["apicore_site"].drop();
db["apicore_session"].drop();
db["apicore_user"].drop();
db["mag_boxes"].drop();
db["mag_magazzini"].drop();
db["opem_person"].drop();
db["opem_persons"].drop();
db["opem_prodotto"].drop();
db["opem_prodotto_prod_param"].drop();
db["opem_prodotto_funct"].drop();
db["opem_system"].drop();
db["opem_system_dom_param"].drop();
db["org_organizzazione"].drop();
db["prd_prodotto"].drop();
db["sequence"].drop();
db["trt_territorio"].drop();

