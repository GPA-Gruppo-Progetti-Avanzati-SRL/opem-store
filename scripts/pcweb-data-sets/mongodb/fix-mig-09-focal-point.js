print(new Date() + " #################### - fix-mig-09-focal-point.js")
let crs = db["opem_organizzazione"].aggregate(
[
    {
        "$lookup":
            {
                "from": "opem_territorio",
                "let": {
                    "et": "comune",
                    "bid": "$address.townhall.text",
                    "county": "$address.county.bid"
                },
                "pipeline": [
                    {
                        "$match": {
                            "$expr": {
                                "$and": [
                                    {
                                        "$eq": ["$_et", "$$et"]
                                    },
                                    {
                                        $eq: [
                                            "$provincia.bid",
                                            "$$county"
                                        ]
                                    },
                                    {
                                        "$eq": ["$name", "$$bid"]
                                    }
                                ]
                            }
                        }
                    }
                ],
                as: "doc_comune"
            }
    },
    {
        "$project": {
            "_bid": 1,
            "_et": 1,
            "domain": 1,
            "site": 1,
            "comune": {
                "$arrayElemAt": ["$doc_comune", 0]
            }
        }
    },
]);

let totals = { insertedCount: 0, matchedCount: 0, modifiedCount: 0, upsertedCount: 0}
let writeSize = 300
let batch = []
while ( crs.hasNext() ) {
    doc = crs.next()

    if (!doc.comune) {
        print("missing comune for focal-point: ", doc.domain, doc.site, doc._bid)
    } else {
        batch.push(
            {updateOne: {
                filter: {"_et": "focal-point", "_bid": doc._bid},
                update:  { "$set": { "address.townhall.bid": doc.comune._bid } }
            }}
        )

        if (batch.length == writeSize) {
            try {
                let resp = db["opem_organizzazione"].bulkWrite(batch)
                totals.insertedCount += resp.insertedCount
                totals.matchedCount += resp.matchedCount
                totals.modifiedCount += resp.modifiedCount
                totals.upsertedCount += resp.upsertedCount
                batch = []
            } catch( error ) {
                print( error )
            }
        }
    }
}

if (batch.length > 0) {
    try {
        let resp = db["opem_organizzazione"].bulkWrite(batch)
        totals.insertedCount += resp.insertedCount
        totals.matchedCount += resp.matchedCount
        totals.modifiedCount += resp.modifiedCount
        totals.upsertedCount += resp.upsertedCount
        batch = []
    } catch( error ) {
        print( error )
    }
}

print(totals)
