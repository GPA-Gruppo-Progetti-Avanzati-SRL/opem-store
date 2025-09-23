const opemDbName = "opem"
const opemCollectionName = "apicore_kv"

let conn = db.getMongo();
let db = conn.getDB(opemDbName);

let c = db[opemCollectionName]
if (!c)  {
    db.createCollection(opemCollectionName)
}
else
{
    c.deleteMany({});
}

// Root level packages.
c.insertOne(
    {
        "name": "theme",
        "scope": "root",
        "obj_type": "kv",
        "category": "admin-app-general",
        "description": "Aspetto UI di amministrazione",
        "inherit": false
        ,"properties": [
            { "key": "theme", "value": "theme-12", "order": 0, "kind": "string" },
        ],
        "sys_info": {
            "created_at": new Date(),
            "status": "active",
            "modified_at": new Date()
        }
    });

c.insertOne(
    {
        "name": "apps",
        "scope": "root",
        "obj_type": "kv",
        "category": "admin-app-general",
        "description": "Apps disponibili",
        "inherit": false
        ,"properties": [
            { "key": "app-home",        "value": "home,/ui-admin/{0}/{1}/{2}/app-home,Home",                "order": 0, "kind": "record:icon,url,hint" },
            { "key": "app-cards",       "value": "credit_cards,/ui-admin/{0}/{1}/{2}/app-cards,Carte",  "order": 0, "kind": "record:icon,url,hint" },
            { "key": "app-magazzino",   "value": "warehouse,/ui-admin/{0}/{1}/{2}/app-magazzino,Magazzino",  "order": 0, "kind": "record:icon,url,hint" },
            { "key": "app-flussi",      "value": "folder_open,/ui-admin/{0}/{1}/{2}/app-flussi,Flussi",  "order": 0, "kind": "record:icon,url,hint" },
            { "key": "app-org",         "value": "groups,/ui-admin/{0}/{1}/{2}/app-org,Organizzazione",  "order": 0, "kind": "record:icon,url,hint" },
            { "key": "app-products",    "value": "credit_card_gears,/ui-admin/{0}/{1}/{2}/app-products,Organizzazione",  "order": 0, "kind": "record:icon,url,hint" },

        ],
        "sys_info": {
            "created_at": new Date(),
            "status": "active",
            "modified_at": new Date()
        }
    });

// Override for CVF
c.insertOne(
    {
        "name": "theme",
        "scope": "root/card",
        "obj_type": "kv",
        "category": "admin-app-general",
        "description": "Aspetto UI di amministrazione",
        "inherit": false
        ,"properties": [
            { "key": "theme", "value": "theme-12", "order": 0, "kind": "string" },
        ],
        "sys_info": {
            "created_at": new Date(),
            "status": "active",
            "modified_at": new Date()
        }
    });
