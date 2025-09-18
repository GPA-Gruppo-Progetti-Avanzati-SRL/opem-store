const opemDbName = "opem"
const opemCollectionName = "trt_territorio"
let conn = db.getMongo();
let db = conn.getDB(opemDbName);

let c = db[opemCollectionName]
if (!c)  {
    db.createCollection(opemCollectionName)
}
// It gets deleted by the first script of nazione
// else
// {
//     c.deleteMany({});
// }

db[opemCollectionName].insertOne(
    {
        "_et": "PROVINCIA",
        "code": "RM",
        "name" : "Roma",
        "code_uic_nazione":"IT",
        "status": "active",
        "order": 1
    })

