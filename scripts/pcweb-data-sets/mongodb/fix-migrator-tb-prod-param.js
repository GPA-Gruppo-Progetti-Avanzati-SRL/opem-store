let crs = db["opem_prodotto_prod_param"].aggregate(
    [
        {
            "$lookup":
                {
                    "from": "opem_prodotto",
                    "let": {
                        "et": "PRODOTTO",
                        "idCompany": { "$toString": "$idCompany"},
                        "productCode": "$productCode",
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
                                            "$eq": ["$_bid", { "$concat" : [ "ISTITUTO-", "$$idCompany", "_", "$$productCode"] }]
                                        }
                                    ]
                                }
                            }
                        }
                    ],
                    as: "doc_prodotto"
                }
        },
        {
            "$project": {
                "progParam": 1,
                "valParam": 1,
                "func": 1,
                "prodotto": {
                    "$arrayElemAt": ["$doc_prodotto", 0]
                }
            }
        }
    ]);

while ( crs.hasNext() ) {
    doc = crs.next()
    db.opem_prodotto.updateOne({ "_et": "PRODOTTO", "_bid": doc.prodotto._bid}, { "$push": { "params.func" : doc.func, "params.value": doc.valParam, "params.mysql.prog_param": doc.progParam } })
}

db.opem_prodotto_prod_param.drop();