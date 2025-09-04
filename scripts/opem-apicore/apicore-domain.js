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
                ,"name": "opem-fe-magazzino"
                ,"description": "Applicazione Home"
                ,"version": "latest"
                ,"path": "opem-fe-magazzino/snapshot/index.html"
            },
            {
                "id": "app-magazzino"
                ,"objType": "app-admin"
                ,"name": "opem-fe-magazzino"
                ,"description": "Applicazione Home"
                ,"version": "latest"
                ,"path": "opem-fe-magazzino/latest/index.html"
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
                ,"name": "opem-fe-magazzino"
                ,"description": "Applicazione Home"
                ,"version": "latest"
                ,"path": "opem-fe-magazzino/latest/index.html"
            },
            {
                "id": "app-magazzino"
                ,"objType": "app-admin"
                ,"name": "opem-fe-magazzino"
                ,"description": "Applicazione Home"
                ,"version": "latest"
                ,"path": "opem-fe-magazzino/latest/index.html"
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
