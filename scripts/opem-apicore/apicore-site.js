const r3ds9DbName = "opem"
const r3ds9CollectionName = "apicore_site"
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
        "code" : "edenred",
        "domain": "card",
        "objType": "site",
        "name" : "Istituto Edenred",
        "description" : "Area Istituto Edenred",
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
                ,"description": "Gestione Approvigionamenti"
                ,"version": "latest"
                ,"path": "opem-fe-magazzino/latest/index.html"
            },
            {
                "id": "app-cards"
                ,"objType": "app-admin"
                ,"name": "opem-fe-magazzino"
                ,"description": "Carte"
                ,"version": "latest"
                ,"path": "opem-fe-magazzino/latest/index.html"
            },
            {
                "id": "app-flussi"
                ,"objType": "app-admin"
                ,"name": "opem-fe-magazzino"
                ,"description": "Flussi"
                ,"version": "latest"
                ,"path": "opem-fe-magazzino/latest/index.html"
            },
            {
                "id": "app-org"
                ,"objType": "app-admin"
                ,"name": "opem-fe-magazzino"
                ,"description": "Organizzazione"
                ,"version": "latest"
                ,"path": "opem-fe-magazzino/latest/index.html"
            }
        ]
        ,"sysInfo": {
            "createdat": new Date(),
            "status": "active",
            "modifiedat": new Date()
        }
    })

