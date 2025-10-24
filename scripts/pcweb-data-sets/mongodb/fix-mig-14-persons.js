print(new Date() + " #################### - fix-mig-14-persons.js")
let crs = db["opem_persons"].aggregate(
[
    {
        "$lookup":
            {
                "from": "opem_addresses",
                "let": {
                    "domain": "$domain",
                    "site": "$site",
                    "et": "address",
                    "bid1": {$ifNull:["$address1.bid","NULL"]},
                    "bid2": {$ifNull:["$address2.bid","NULL"]},
                },
                "pipeline": [
                    {
                        "$match": {
                            "$expr": {
                                "$and": [
                                    {
                                        "$eq": ["$domain", "$$domain"]
                                    },
                                    {
                                        "$eq": ["$site", "$$site"]
                                    },
                                    {
                                        "$eq": ["$_et", "$$et"]
                                    },
                                    {
                                        "$or": [{"$eq": ["$_bid", "$$bid1"]}, {"$eq": ["$_bid", "$$bid2"]}]
                                    }
                                ]
                            }
                        }
                    }
                ],
                as: "doc_addresses"
            }
    },
    {
        "$lookup":
            {
                "from": "opem_territorio",
                "let": {
                    "et": "comune",
                    "name": "$birth.townhall.text",
                    "prov": "$birth.county.bid",
                },
                "pipeline": [
                    {
                        "$match": {
                            "$expr": {
                                "$and": [
                                    {
                                        "$eq": ["$name", "$$name"]
                                    },
                                    {
                                        "$eq": ["$county.bid", "$$prov"]
                                    },
                                    {
                                        "$eq": ["$_et", "$$et"]
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
            "address1": 1,
            "address2": 1,
            "doc_addresses": 1,
            "doc_comune": 1,
        }
    }
]);

let totalDocs = 0
let totals = { insertedCount: 0, matchedCount: 0, modifiedCount: 0, upsertedCount: 0}
let writeSize = 2000
let batch = []
while ( crs.hasNext() ) {
    doc = crs.next()

    let udoc = {}
    if (doc.doc_addresses && doc.doc_addresses.length > 0) {
        let addressesArray = doc.doc_addresses.map(addr => {
            // check for 2 and then 1. COR seems always in 2... they might be equal....
            if (doc.address2 && addr._bid === doc.address2.bid) {
                addr.type = doc.address2.type
            } else {
                if (doc.address1 && addr._bid === doc.address1.bid) {
                    addr.type = doc.address1.type
                }
            }
            return { type: addr.type, zip_code: addr.zip_code, townhall: addr.townhall, county: addr.county, country: addr.country, address: addr.address }
        });

        batch.push(
            { updateOne: {
                    filter: {"_et": "person", "_bid": doc._bid, "domain": doc.domain, "site": doc.site},
                    update:  {
                        $set: { "addresses": addressesArray } ,
                        $unset: { "address1": "", "address2": "" }
                    }
                }}
        )

        if (batch.length == writeSize) {
            try {
                let resp = db["opem_persons"].bulkWrite(batch)
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
    }
}

if (batch.length > 0) {
    try {
        let resp = db["opem_persons"].bulkWrite(batch)
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
