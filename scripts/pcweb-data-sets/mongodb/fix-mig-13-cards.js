print(new Date() + " #################### - fix-mig-13-cards.js")
let crs = db["opem_cards"].aggregate(
[
    {
        "$lookup":
            {
                "from": "opem_addresses",
                "let": {
                    "domain": "$domain",
                    "site": "$site",
                    "et": "address",
                    "bid": {$ifNull:["$address.address","NULL"]},
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
                                        "$eq": ["$_bid", "$$bid"]
                                    }
                                ]
                            }
                        }
                    }
                ],
                as: "doc_address"
            }
    },
    {
        "$lookup":
            {
                "from": "opem_cards_apps",
                "let": {
                    "domain": "$domain",
                    "site": "$site",
                    "et": "card-app",
                    "bid": "$_bid",
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
                                        "$eq": ["$_bid", "$$bid"]
                                    }
                                ]
                            }
                        }
                    }
                ],
                as: "doc_apps"
            }
    },
    {
        "$project": {
            "_bid": 1,
            "_et": 1,
            "domain": 1,
            "site": 1,
            "address": 1,
            "doc_apps": 1,
            "doc_address": {
                "$arrayElemAt": ["$doc_address", 0]
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

    let udoc = {}
    let isEmpty = true
    if (doc.doc_address) {
        udoc["$push"] = { "addresses": { type: doc.address.type, zip_code: doc.doc_address.zip_code, townhall: doc.doc_address.townhall, county: doc.doc_address.county, country: doc.doc_address.country, address: doc.doc_address.address } }
        udoc["$unset"] = { "address": "" }
        isEmpty = false
    }

    if (doc.doc_apps && doc.doc_apps.length > 0) {
        let appsArray = doc.doc_apps.map(app => {
            if (app.app_number) {
                return { "app_type": app.appl_type, "app_number": app.app_number }
            } else {
                return { "app_type": app.appl_type }
            }
        })
        isEmpty = false
        udoc["$set"] = { "apps": appsArray }
    }

    if (!isEmpty) {
        batch.push(
            {
                updateOne: {
                    filter: {"_et": "card", "_bid": doc._bid, "domain": doc.domain, "site": doc.site},
                    update: udoc
                }
            }
        )

        if (batch.length == writeSize) {
            try {
                let resp = db["opem_cards"].bulkWrite(batch)
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
        print("update document empty for ", doc.domain, doc.site, doc._bid)
    }
}

if (batch.length > 0) {
    try {
        let resp = db["opem_cards"].bulkWrite(batch)
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
