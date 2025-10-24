print(new Date() + " #################### - fix-mig-12-addresses.js")
let crs = db["opem_addresses"].aggregate(
[
    {
        $lookup: {
            from: "opem_territorio",
            let: {
                et: "comune",
                name: "$townhall.text",
                prov: "$county.bid"
            },
            pipeline: [
                {
                    $match: {
                        $expr: {
                            $and: [
                                {
                                    $eq: ["$name", "$$name"]
                                },
                                {
                                    $eq: [
                                        "$provincia.bid",
                                        "$$prov"
                                    ]
                                },
                                {
                                    $eq: ["$_et", "$$et"]
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
            "doc_comune": {
                "$arrayElemAt": ["$doc_comune", 0]
            }
        }
    }
]);

let totalDocs = 0
let totals = { insertedCount: 0, matchedCount: 0, modifiedCount: 0, upsertedCount: 0}
let writeSize = 2000
let batch = []
while ( crs.hasNext() ) {
    doc = crs.next()
    if (doc.doc_comune && doc.doc_comune._bid ) {
        batch.push(
            { updateOne: {
                filter: {"_et": "address", "_bid": doc._bid, "domain": doc.domain, "site": doc.site},
                update: { "$set": { "townhall.bid": doc.doc_comune._bid } }
            }}
        )

        if (batch.length == writeSize) {
            try {
                resp = db["opem_addresses"].bulkWrite(batch)
                totals.insertedCount += resp.insertedCount
                totals.matchedCount += resp.matchedCount
                totals.modifiedCount += resp.modifiedCount
                totals.upsertedCount += resp.upsertedCount
                totalDocs += batch.length
                print (new Date() + " ######### ", totalDocs)
                batch = []

            } catch (error) {
                print(error)
            }
        }
    } else {
        print(" missing comune for address: ", doc.domain, doc.site, doc._bid)
    }
}

if (batch.length > 0) {
    try {
        let resp = db["opem_addresses"].bulkWrite(batch)
        totals.insertedCount += resp.insertedCount
        totals.matchedCount += resp.matchedCount
        totals.modifiedCount += resp.modifiedCount
        totals.upsertedCount += resp.upsertedCount
        totalDocs += batch.length
        print (new Date() + " ######### ", totalDocs)
        batch = []
    } catch (error) {
        print(error)
    }
}

print(totals)
let missedItems = db["opem_addresses"].countDocuments({"townhall.bid": {$exists: false}});
print("... cursor eof addresses: missed-items = ", missedItems);
