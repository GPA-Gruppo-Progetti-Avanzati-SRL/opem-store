const opemDbName = "opem"
const opemCollectionName = "mag_magazzini"
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
        "_bid": "MAG-FP-ERD-00001",
        "_et": "MAGAZZINO",
        "domain": "card",
        "site": "edenred",
        "focal_point": { "bid": "FP-ERD-00001" },
        "sys_info": {
            "created_at": new Date(),
            "status": "active",
            "modified_at": new Date()
        }
    })

