const r3ds9DbName = "opem"
const r3ds9CollectionName = "apicore_domain"

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

db[r3ds9CollectionName].insertOne(
    {
        "code" : "root",
        "objType": "domain",
        "name" : "Master Domain",
        "description" : "the master of domains",
        "langs": "it",
        "apps": [
            {
                "id": "app-home"
                ,"objType": "app-admin"
                ,"name": "Applicazione Home"
                ,"description": "Applicazione Home"
                ,"path": "opem-fe-magazzino/browser/index.html"
            },
            {
                "id": "app-magazzino"
                ,"objType": "app-admin"
                ,"name": "Applicazione Gestione Magazzino"
                ,"description": "Applicazione  Gestione Magazzino"
                ,"path": "opem-fe-magazzino/browser/index.html"
            }
        ],
        "members": [
            {
               "code": "card", "objType": "domain"
            }
        ],
        "sysInfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    })

db[r3ds9CollectionName].insertOne(
    {
        "code" : "card",
        "objType": "domain",
        "name" : "OpeM Front Door",
        "description" : "OPeM FrontDoor",
        "langs": "it",
        "apps": [
            {
                "id": "app-home"
                ,"objType": "app-admin"
                ,"name": "Applicazione Home"
                ,"description": "Applicazione Home"
                ,"path": "opem-fe-magazzino/browser/index.html"
            },
            {
                "id": "app-magazzino"
                ,"objType": "app-admin"
                ,"name": "Applicazione Gestione Magazzino"
                ,"description": "Applicazione  Gestione Magazzino"
                ,"path": "opem-fe-magazzino/browser/index.html"
            }
        ],
        "members": [
            {
                "code": "edenred", "objType": "site"
            }
        ],
        "sysinfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    })
