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
        "_bid": "RM-Roma",
        "_et": "COMUNE",
        "code_provincia": "RM",
        "code_uic_nazione": "IT",
        "name" : "Roma",
        "sys_info": {
            "created_at": new Date(),
            "status": "active",
            "modified_at": new Date()
        }
    })

