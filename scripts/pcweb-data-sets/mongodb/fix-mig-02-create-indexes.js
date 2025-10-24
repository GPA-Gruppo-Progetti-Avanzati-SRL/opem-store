print(new Date() + " #################### - fix-mig-02-create-indexes.js")

db["opem_boxes"].createIndex({"domain": 1, "site": 1, "_bid": 1, "_et": 1 }, { unique: true });
db["opem_magazzini"].createIndex({ "domain": 1, "site": 1, "_bid": 1, "_et": 1 }, { unique: true });
db["opem_organizzazione"].createIndex({ "domain": 1, "site": 1, "_bid": 1, "_et": 1 }, { unique: true });
db["opem_prodotto"].createIndex({"domain": 1,  "site": 1, "_bid": 1, "_et": 1 }, { unique: true });
db["opem_system"].createIndex({"domain": 1, "site": 1, "_bid": 1, "_et": 1 }, { unique: false });
db["opem_territorio"].createIndex({ "_bid": 1, "_et": 1 }, { unique: true });
db["opem_territorio"].createIndex({ "name": 1, "_et": 1 }, { unique: false });

db["opem_cards"].createIndex({ "_bid": 1,  "domain": 1, "site": 1, "_et": 1 }, { unique: true });
db["opem_persons"].createIndex({ "_bid": 1, "domain": 1, "site": 1, "_et": 1 }, { unique: true });
db["opem_addresses"].createIndex({ "_bid": 1, "domain": 1, "site": 1, "_et": 1 }, { unique: true });
db["opem_cards_apps"].createIndex({ "_bid": 1, "domain": 1, "site": 1, "_et": 1 }, { unique: false });
db["opem_user"].createIndex({ "nickname": 1, "_et": 1 }, { unique: true });
db["opem_file"].createIndex({ "_bid": 1, "domain": 1, "site": 1, "_et": 1 }, { unique: true });
db["opem_file_rows"].createIndex({ "_bid": 1, "domain": 1, "site": 1, "_et": 1 }, { unique: true });