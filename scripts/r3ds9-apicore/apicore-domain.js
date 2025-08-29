const r3ds9DbName = "r3ds9"
const r3ds9CollectionName = "apicore_domain"

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
                ,"path": "r3ds9-admin-app-home/index.tmpl"
            },
            {
                "id": "app-sys-admin"
                ,"objType": "app-admin"
                ,"name": "Applicazione Sys Admin"
                ,"description": "Applicazione  Sys Admin"
                ,"path": "r3ds9-app-sys-admin/index.tmpl"
            }
        ],
        "members": [
            {
               "code": "cvf", "objType": "domain"
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
        "code" : "cvf",
        "objType": "domain",
        "name" : "Circolo Velico Fiumicino ASD",
        "description" : "Darsena Fiumicino",
        "langs": "it",
        "apps": [
            {
                "id": "app-home"
                ,"objType": "app-admin"
                ,"name": "Applicazione Home"
                ,"description": "Applicazione Home"
                ,"path": "r3ds9-admin-app-home/index.tmpl"
            }
        ],
        "members": [
            {
                "code": "champ42", "objType": "site"
            }
        ],
        "sysinfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    })
