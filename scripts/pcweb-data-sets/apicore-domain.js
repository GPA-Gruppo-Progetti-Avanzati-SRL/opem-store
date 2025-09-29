const opemDbName = "opem"
const opemCollectionName = "apicore_domain"

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

db[opemCollectionName].insertOne(
    {
        "code" : "root",
        "obj_type": "domain",
        "name" : "Master Domain",
        "description" : "the master of domains",
        "langs": "it",
        "apps": [
            {
                "id": "app-home"
                ,"obj_type": "app-admin"
                ,"name": "opem-fe-magazzino"
                ,"description": "Applicazione Home"
                ,"version": "latest"
                ,"path": "opem-fe-magazzino/snapshot/index.html"
            },
            {
                "id": "app-magazzino"
                ,"obj_type": "app-admin"
                ,"name": "opem-fe-magazzino"
                ,"description": "Applicazione Home"
                ,"version": "latest"
                ,"path": "opem-fe-magazzino/latest/index.html"
            }
        ],
        "members": [
            {
               "code": "card", "obj_type": "domain"
            }
        ],
        "sys_info": {
            "created_at": new Date(),
            "status": "active",
            "modified_at": new Date()
        }
    })

db[opemCollectionName].insertOne(
    {
        "code" : "card",
        "obj_type": "domain",
        "name" : "OpeM Front Door",
        "description" : "OPeM FrontDoor",
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
                ,"description": "Applicazione Home"
                ,"version": "latest"
                ,"path": "opem-fe-magazzino/latest/index.html"
            }
        ],
        "members": [
            {
                "code": "edenred", "obj_type": "site"
            }
        ],
        "sys_info": {
            "created_at": new Date(),
            "status": "active",
            "modified_at": new Date()
        }
    })
