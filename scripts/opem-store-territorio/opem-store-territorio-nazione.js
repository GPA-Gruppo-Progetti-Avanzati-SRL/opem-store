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
        "sys_info": {
            "created_at": new Date(),
            "status": "active",
            "modified_at": new Date()
        }

    })

