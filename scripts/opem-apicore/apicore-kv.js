const r3ds9DbName = "opem"
const r3ds9CollectionName = "apicore_kv"

let conn = db.getMongo();
let db = conn.getDB(r3ds9DbName);

let c = db[r3ds9CollectionName]
if (!c)  {
    db.createCollection(r3ds9CollectionName)
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
        "objType": "kv",
        "category": "admin-app-general",
        "description": "Aspetto UI di amministrazione",
        "inherit": false
        ,"properties": [
            { "key": "theme", "value": "theme-12", "order": 0, "kind": "string" },
        ]
        ,"sysInfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    });

c.insertOne(
    {
        "name": "apps",
        "scope": "root",
        "objType": "kv",
        "category": "admin-app-general",
        "description": "Apps disponibili",
        "inherit": false
        ,"properties": [
            { "key": "app-home",    "value": "home,/ui-admin/{0}/{1}/{2}/app-home,Home",         "order": 0, "kind": "record:icon,url,hint" },
            { "key": "app-magazzino",  "value": "Magazzino,/ui-admin/{0}/{1}/{2}/app-magazzino,Magazzino", "order": 0, "kind": "record:icon,url,hint" }
        ]
        ,"sysinfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    });

// Override for CVF
c.insertOne(
    {
        "name": "theme",
        "scope": "root/card",
        "objType": "kv",
        "category": "admin-app-general",
        "description": "Aspetto UI di amministrazione",
        "inherit": false
        ,"properties": [
            { "key": "theme", "value": "theme-12", "order": 0, "kind": "string" },
        ]
        ,"sysinfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    });
