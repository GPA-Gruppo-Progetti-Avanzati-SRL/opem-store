print(new Date() + " #################### - fix-mig-04-domain.js")

const opemCollectionName = "opem_system"

db[opemCollectionName].insertOne(
    {
        "_bid" : "root",
        "_et": "domain",
        "name" : "Master Domain",
        "description" : "the master of domains",
        "langs": "it",
        "apps": [
            {
                "id": "app-home"
                ,"obj_type": "app-admin"
                ,"name": "opem-fe-home"
                ,"description": "Applicazione Home"
                ,"version": "latest"
                ,"path": "opem-fe-home/snapshot/index.html"
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
        "_bid" : "card",
        "_et": "domain",
        "name" : "PCWeb Reloaded",
        "description" : "PCWeb Reloaded",
        "langs": "it",
        "apps": [
            {
                "id": "app-home"
                ,"obj_type": "app-admin"
                ,"name": "opem-fe-home"
                ,"description": "Applicazione Home"
                ,"version": "latest"
                ,"path": "opem-fe-home/latest/index.html"
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
