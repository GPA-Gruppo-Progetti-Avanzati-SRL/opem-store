const opemDbName = "opem"
const opemCollectionName = "mag_boxes"
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
        "_bid": "BOX-1",
        "_et": "BOX",
        "domain": "card",
        "site": "edenred",
        "magazzino": { "bid": "MAG-FP-ERD-00001" },
        "prodotto": { "bid":"lunch-day-dip" },
        "info": {
            "box_code": "BOX-CODE-1",
            "package_code": "PKG-CODE-1"
        },
        "status" : {
            "num_in": 20,
            "status": "CX"
        },
        "recipient" : {
            "county":  { "bid": "RM" },
            "townhall": { "bid": "RM-Roma" },
            "zipcode": "00100",
            "attn_to": "PAOLINO PAPERINO"
        },
        "sys_info": {
            "created_at": new Date(),
            "status": "active",
            "modified_at": new Date()
        }
    })

