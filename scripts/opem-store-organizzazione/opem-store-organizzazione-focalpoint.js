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

db[opemCollectionName].insertOne(
    {
        "_bid": "FP-00001",
        "_et": "FOCAL-POINT",
        "domain": "card",
        "ns": "RUO",
        "officer_name" : "Giuseppe Celitti",
        "status": "active",
        "order": 1
    })

