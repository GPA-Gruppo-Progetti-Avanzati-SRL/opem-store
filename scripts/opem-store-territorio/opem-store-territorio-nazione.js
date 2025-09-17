const opemDbName = "opem"
const opemCollectionName = "trt_territorio"
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
        "_et": "NAZIONE",
        "code": "IT",
        "name" : "Italia",
        "status": "active",
        "order": 1
    })

