const r3ds9DbName = "r3ds9"
const r3ds9CollectionName = "apicore_site"
let conn = new Mongo();
let db = conn.getDB(r3ds9DbName);

let c = db[r3ds9CollectionName]
if (!c)  {
    db.createCollection(r3ds9CollectionName)
}
else
{
    c.deleteMany({});
}

db[r3ds9CollectionName].insertOne(
    {
        "code" : "champ42",
        "domain": "cvf",
        "objType": "site",
        "name" : "CVF Campionato 42",
        "description" : "Il campionato piu' affollato d'Italia",
        "langs": "it",
        "apps": [
            {
                "id": "app-home"
                ,"objType": "app-admin"
                ,"name": "Applicazione Home"
                ,"description": "Applicazione Home"
                ,"path": "r3ds9-admin-app-home/index.tmpl"
            },
            {
                "id": "app-sys"
                ,"objType": "app-admin"
                ,"name": "Applicazione Sys Admin"
                ,"description": "Applicazione  Sys Admin"
                ,"path": "r3ds9-admin-app-sys/index.tmpl"
            }
        ]
        ,"sysInfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    })

db[r3ds9CollectionName].insertOne(
    {
        "code" : "champ43",
        "domain": "cvf",
        "objType": "site",
        "name" : "CVF Campionato 43",
        "description" : "Il campionato piu' affollato d'Italia",
        "langs": "it",
        "apps": [
            {
                "id": "app-home"
                ,"objType": "app-admin"
                ,"name": "Applicazione Home"
                ,"description": "Applicazione Home"
                ,"path": "r3ds9-admin-app-home/index.tmpl"
            },
            {
                "id": "app-sys"
                ,"objType": "app-admin"
                ,"name": "Applicazione Sys Admin"
                ,"description": "Applicazione  Sys Admin"
                ,"path": "r3ds9-admin-app-sys/index.tmpl"
            }
        ]
        ,"sysInfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    })
