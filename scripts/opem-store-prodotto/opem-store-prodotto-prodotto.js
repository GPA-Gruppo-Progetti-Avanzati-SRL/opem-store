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
        "ns": "*",
        "name" : "Postepay Lunch DAY Dipendenti",
        "status": "active",
        "order": 1
    })

