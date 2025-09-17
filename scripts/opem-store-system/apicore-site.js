const opemDbName = "opem"
const opemCollectionName = "apicore_site"
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

