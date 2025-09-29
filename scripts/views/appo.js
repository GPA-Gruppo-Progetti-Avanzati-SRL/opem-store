
db.apicore_site.updateMany({}, { $set: {
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
            ]
        }
    }
);