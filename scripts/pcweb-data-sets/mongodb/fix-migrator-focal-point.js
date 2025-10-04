let crs = db["org_organizzazione"].aggregate(
[
    {
        "$lookup":
            {
                "from": "opem_system",
                "let": {
                    "et": "COMUNE",
                    "bid": "$address.townhall.text"
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
            "comune": {
                "$arrayElemAt": ["$doc_comune", 0]
            }
        }
    },
    {
        "$project": {
            "_bid": 1,
            "_et": 1,
            "bid_comune": "$comune._bid",
        }
    },
]);


while ( crs.hasNext() ) {
    doc = crs.next()
    db["org_organizzazione"].updateOne({"_et": "FOCAL-POINT", "_bid": doc._bid}, { "$set": { "address.townhall.bid": doc.bid_comune } })
}
print("....cursor eof");