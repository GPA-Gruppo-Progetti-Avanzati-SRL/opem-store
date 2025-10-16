
print("drop index: mag_boxes::_RM_id")
db["mag_boxes"           ].dropIndex("_RM_id");
print("drop index: mag_magazzini::_RM_id")
db["mag_magazzini"       ].dropIndex("_RM_id");
print("drop index: org_organizzazione::_RM_id")
db["org_organizzazione"  ].dropIndex("_RM_id");

// non piu' presente inquanto il mapping parte da una view.
// db["opem_prodotto"       ].dropIndex("_RM_id");

print("drop index: opem_system::_RM_idComune")
db["opem_system"      ].dropIndex("_RM_idComune");
print("drop index: opem_system::_RM_sigla")
db["opem_system"      ].dropIndex("_RM_sigla");
print("drop index: opem_system::_RM_code_uic")
db["opem_system"      ].dropIndex("_RM_code_uic");
print("drop index: opem_system::_RM_id")
db["opem_system"        ].dropIndex("_RM_id");
