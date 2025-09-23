const opemDbName = "opem"
const opemCollectionName = "prd_prodotto"
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
        "_bid": "lunch-day-dip",
        "_et": "PRODOTTO",
        "domain": "card",
        "site": "*",
        "name" : "Postepay Lunch DAY Dipendenti",
        "sys_info": {
            "created_at": new Date(),
            "status": "active",
            "modified_at": new Date()
        }
    })

