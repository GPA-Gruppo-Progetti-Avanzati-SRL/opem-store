
print("drop index: opem_boxes::_RM_id")
db["opem_boxes"           ].dropIndex("_RM_id");
print("drop index: opem_cards::_RM_id")
db["opem_cards"        ].dropIndex("_RM_id");
print("drop index: opem_magazzini::_RM_id")
db["opem_magazzini"       ].dropIndex("_RM_id");
print("drop index: opem_organizzazione::_RM_id")
db["opem_organizzazione"  ].dropIndex("_RM_id");
print("drop index: opem_persons::_RM_id")
db["opem_persons"        ].dropIndex("_RM_id");
print("drop index: opem_system::_RM_id")
db["opem_system"        ].dropIndex("_RM_id");
print("drop index: opem_territorio::_RM_idComune")
db["opem_territorio"      ].dropIndex("_RM_idComune");
print("drop index: opem_territorio::_RM_sigla")
db["opem_territorio"      ].dropIndex("_RM_sigla");
print("drop index: opem_territorio::_RM_code_uic")
db["opem_territorio"      ].dropIndex("_RM_code_uic");

