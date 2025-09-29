const opemDbName = "opem"
let conn = db.getMongo();
let db = conn.getDB(opemDbName);

db["apicore_site"].updateMany({},
    { $set: {
            "langs": "it",
            "apps": [
                {
                    "id": "app-home"
                    ,"obj_type": "app-admin"
                    ,"name": "opem-fe-magazzino"
                    ,"description": "Applicazione Home"
                    ,"version": "latest"
                    ,"path": "opem-fe-magazzino/latest/index.html"
                },
                {
                    "id": "app-magazzino"
                    ,"obj_type": "app-admin"
                    ,"name": "opem-fe-magazzino"
                    ,"description": "Gestione Approvigionamenti"
                    ,"version": "latest"
                    ,"path": "opem-fe-magazzino/latest/index.html"
                },
                {
                    "id": "app-cards"
                    ,"obj_type": "app-admin"
                    ,"name": "opem-fe-magazzino"
                    ,"description": "Carte"
                    ,"version": "latest"
                    ,"path": "opem-fe-magazzino/latest/index.html"
                },
                {
                    "id": "app-flussi"
                    ,"obj_type": "app-admin"
                    ,"name": "opem-fe-magazzino"
                    ,"description": "Flussi"
                    ,"version": "latest"
                    ,"path": "opem-fe-magazzino/latest/index.html"
                },
                {
                    "id": "app-org"
                    ,"obj_type": "app-admin"
                    ,"name": "opem-fe-magazzino"
                    ,"description": "Organizzazione"
                    ,"version": "latest"
                    ,"path": "opem-fe-magazzino/latest/index.html"
                }
            ]
        }
    }
);

db["apicore_site"        ].dropIndex("_RM_id");
// db["apicore_domain"      ].dropIndex();
// db["apicore_kv"          ].dropIndex();
// db["apicore_session"     ].dropIndex();
// db["apicore_user"        ].dropIndex();
db["mag_boxes"           ].dropIndex("_RM_id");
db["mag_magazzini"       ].dropIndex("_RM_id");
db["org_organizzazione"  ].dropIndex("_RM_id");
db["prd_prodotto"        ].dropIndex("_RM_id");
// db["sequence"            ].dropIndex();
db["trt_territorio"      ].dropIndex("_RM_idComune");
db["trt_territorio"      ].dropIndex("_RM_code");
db["trt_territorio"      ].dropIndex("_RM_code_uic");


