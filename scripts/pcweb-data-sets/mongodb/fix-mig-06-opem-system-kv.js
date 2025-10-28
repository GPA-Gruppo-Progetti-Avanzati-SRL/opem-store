print(new Date() + " #################### - fix-mig-06-opem-system-kv.js")

const opemCollectionName = "opem_system"


// Root level packages.
db[opemCollectionName].insertOne(
    {
        "scope": "root",
        "_bid": "theme",
        "_et": "kv",
        "category": "admin-app-general",
        "description": "Aspetto UI di amministrazione",
        "inherit": false
        ,"properties": [
            { "key": "theme", "value": "pi", "order": 0, "kind": "string" },
        ],
        "sys_info": {
            "created_at": new Date(),
            "status": "active",
            "modified_at": new Date()
        }
    });

db[opemCollectionName].insertOne(
    {
        "scope": "root",
        "_bid": "apps",
        "_et": "kv",
        "category": "admin-app-general",
        "description": "Apps disponibili",
        "inherit": false
        ,"properties": [
            { "key": "app-home",        "sys_name": "opem-fe-home", "status": "active",   "name": "Home",           "description": "description of Home",           "icon": "home",              "value": "/ui-admin/{0}/{1}/{2}/app-home",       "order": 0, "kind": "string" },
            { "key": "app-cards",       "sys_name": "opem-fe-cards", "status": "disabled", "name": "Carte",          "description": "description of Carte",          "icon": "credit_cards",      "value": "/ui-admin/{0}/{1}/{2}/app-cards",      "order": 0, "kind": "string" },
            { "key": "app-magazzino",   "sys_name": "opem-fe-magazzino", "status": "active",   "name": "Magazzino",      "description": "description of Magazzino",      "icon": "warehouse",         "value": "/ui-admin/{0}/{1}/{2}/app-magazzino",  "order": 0, "kind": "string" },
            { "key": "app-flussi",      "sys_name": "opem-fe-flussi", "status": "disabled", "name": "Flussi",         "description": "description of Flussi",         "icon": "folder_open",       "value": "/ui-admin/{0}/{1}/{2}/app-flussi",     "order": 0, "kind": "string" },
            { "key": "app-org",         "sys_name": "opem-fe-org", "status": "disabled", "name": "Organizzazione", "description": "description of Organizzazione", "icon": "groups",            "value": "/ui-admin/{0}/{1}/{2}/app-org",        "order": 0, "kind": "string" },
            { "key": "app-products",    "sys_name": "opem-fe-prodotto", "status": "disabled", "name": "Prodotti",       "description": "description of Prodotti",       "icon": "credit_card_gears", "value": "/ui-admin/{0}/{1}/{2}/app-products",   "order": 0, "kind": "string" },
            { "key": "app-system",      "sys_name": "opem-fe-system", "status": "disabled", "name": "Sistema",       "description": "description of Sistema",       "icon": "settings", "value": "/ui-admin/{0}/{1}/{2}/app-system",   "order": 0, "kind": "string" }
        ],
        "sys_info": {
            "created_at": new Date(),
            "status": "active",
            "modified_at": new Date()
        }
    });

// Override for CVF
db[opemCollectionName].insertOne(
    {
        "scope": "root/card",
        "_bid": "theme",
        "_et": "kv",
        "category": "admin-app-general",
        "description": "Aspetto UI di amministrazione",
        "inherit": false
        ,"properties": [
            { "key": "theme", "value": "pi", "order": 0, "kind": "string" },
        ],
        "sys_info": {
            "created_at": new Date(),
            "status": "active",
            "modified_at": new Date()
        }
    });
