print(new Date() + " #################### - fix-mig-08-companies.js")

db["opem_system"].updateMany(
    { "_et": "site" },
    { $set: {
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
                    ,"description": "Gestione Approvigionamenti"
                    ,"version": "latest"
                    ,"path": "opem-fe-magazzino/latest/index.html"
                },
                {
                    "id": "app-cards"
                    ,"obj_type": "app-admin"
                    ,"name": "opem-fe-cards"
                    ,"description": "Carte"
                    ,"version": "latest"
                    ,"path": "opem-fe-cards/latest/index.html"
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
                },
                {
                    "id": "app-products"
                    ,"obj_type": "app-admin"
                    ,"name": "opem-fe-prodotto"
                    ,"description": "Prodotti"
                    ,"version": "latest"
                    ,"path": "opem-fe-prodotto/latest/index.html"
                },
                {
                    "id": "app-system"
                    ,"obj_type": "app-admin"
                    ,"name": "opem-fe-system"
                    ,"description": "Sistema"
                    ,"version": "latest"
                    ,"path": "opem-fe-system/latest/index.html"
                }
            ]
        }
    }
);





