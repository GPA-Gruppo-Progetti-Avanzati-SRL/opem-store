const opemDbName = "opem"
const opemCollectionName = "org_organizzazione"
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

db[opemCollectionName].insertMany([
    {
        "_bid": "FP-RUO-00001",
        "_et": "FOCAL-POINT",
        "domain": "card",
        "site": "RUO",
        "officer_name" : "Giuseppe Celitti (RUO)",
        "status": "active",
        "order": 1
    },
    {
        "_bid": "FP-ERD-00001",
        "_et": "FOCAL-POINT",
        "domain": "card",
        "site": "edenred",
        "officer_name" : "Giuseppe Celitti (Edenred)",
        "status": "active",
        "order": 1
    }
    ]
)

