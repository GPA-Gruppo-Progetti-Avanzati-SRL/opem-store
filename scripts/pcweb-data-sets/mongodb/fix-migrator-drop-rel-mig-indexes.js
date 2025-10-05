

db["mag_boxes"           ].dropIndex("_RM_id");
db["mag_magazzini"       ].dropIndex("_RM_id");
db["org_organizzazione"  ].dropIndex("_RM_id");
db["opem_prodotto"       ].dropIndex("_RM_id");

db["opem_system"      ].dropIndex("_RM_idComune");
db["opem_system"      ].dropIndex("_RM_sigla");
db["opem_system"      ].dropIndex("_RM_code_uic");
// it might not exist
db["opem_system"        ].dropIndex("_RM_id");
