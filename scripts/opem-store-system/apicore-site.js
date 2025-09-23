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
        "obj_type": "site",
        "name" : "Istituto Edenred",
        "description" : "Area Istituto Edenred",
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
                ,"description": "Gestione Approvigionamenti"
                ,"version": "latest"
                ,"path": "opem-fe-magazzino/latest/index.html"
            },
            {
                "id": "app-cards"
                ,"obj_type": "app-admin"
                ,"name": "opem-fe-magazzino"
                ,"description": "Carte"
                ,"version": "latest"
                ,"path": "opem-fe-magazzino/latest/index.html"
            },
            {
                "id": "app-flussi"
                ,"obj_type": "app-admin"
                ,"name": "opem-fe-magazzino"
                ,"description": "Flussi"
                ,"version": "latest"
                ,"path": "opem-fe-magazzino/latest/index.html"
            },
            {
                "id": "app-org"
                ,"obj_type": "app-admin"
                ,"name": "opem-fe-magazzino"
                ,"description": "Organizzazione"
                ,"version": "latest"
                ,"path": "opem-fe-magazzino/latest/index.html"
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
        "code" : "ruo",
        "domain": "card",
        "obj_type": "site",
        "name" : "Istituto RUO",
        "description" : "Area Istituto RUO",
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
                ,"description": "Gestione Approvigionamenti"
                ,"version": "latest"
                ,"path": "opem-fe-magazzino/latest/index.html"
            },
            {
                "id": "app-cards"
                ,"obj_type": "app-admin"
                ,"name": "opem-fe-magazzino"
                ,"description": "Carte"
                ,"version": "latest"
                ,"path": "opem-fe-magazzino/latest/index.html"
            },
            {
                "id": "app-flussi"
                ,"obj_type": "app-admin"
                ,"name": "opem-fe-magazzino"
                ,"description": "Flussi"
                ,"version": "latest"
                ,"path": "opem-fe-magazzino/latest/index.html"
            },
            {
                "id": "app-org"
                ,"obj_type": "app-admin"
                ,"name": "opem-fe-magazzino"
                ,"description": "Organizzazione"
                ,"version": "latest"
                ,"path": "opem-fe-magazzino/latest/index.html"
            }
        ],
        "sys_info": {
            "created_at": new Date(),
            "status": "active",
            "modified_at": new Date()
        }
    })
