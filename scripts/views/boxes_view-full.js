const opemDbName = "opem"
const opemViewName = "mag_boxes_view"
const opemViewSource = "mag_boxes"
let conn = db.getMongo();
let db = conn.getDB(opemDbName);

db.createView(
    opemViewName,
    opemViewSource,
    [
      {
        "$match": {
          "_et": "BOX"
        }
      },
      {
        "$project": {
          "domain": 1,
          "site": 1,
          "magazzino.bid": "$bid_magazzino",
          "prodotto.bid": "$bid_prodotto",
          "info": 1,
          "status": 1,
          "recipient.county.bid": "$recipient.county",
          "recipient.townhall.bid": "$recipient.townhall",
          "recipient.zipcode": 1,
          "recipient.attn_to": 1
        }
      },
      {
        "$lookup": {
          "from": "mag_magazzini",
          "let": {
            "et": "MAGAZZINO",
            "bid": "$magazzino.bid"
          },
          "pipeline": [
            {
              "$match": {
                "$expr": {
                  "$and": [
                    {
                      "$eq": [
                        "$_et",
                        "$$et"
                      ]
                    },
                    {
                      "$eq": [
                        "$_bid",
                        "$$bid"
                      ]
                    }
                  ]
                }
              }
            }
          ],
          "as": "doc_magazzino"
        }
      },
      {
        "$lookup": {
          "from": "prd_prodotto",
          "let": {
            "et": "PRODOTTO",
            "bid": "$prodotto.bid"
          },
          "pipeline": [
            {
              "$match": {
                "$expr": {
                  "$and": [
                    {
                      "$eq": [
                        "$_et",
                        "$$et"
                      ]
                    },
                    {
                      "$eq": [
                        "$_bid",
                        "$$bid"
                      ]
                    }
                  ]
                }
              }
            }
          ],
          "as": "doc_prodotto"
        }
      },
      {
        "$project": {
          "_bid": 1,
          "_et": 1,
          "domain": 1,
          "site": 1,
          "info": 1,
          "status": 1,
          "recipient": 1,
          "doc_magazzino": {
            "$arrayElemAt": [
              "$doc_magazzino",
              0
            ]
          },
          "doc_prodotto": {
            "$arrayElemAt": [
              "$doc_prodotto",
              0
            ]
          }
        }
      },
      {
        "$project": {
          "_bid": 1,
          "_et": 1,
          "domain": 1,
          "site": 1,
          "info": 1,
          "status": 1,
          "recipient": 1,
          "magazzino.bid": "$doc_magazzino._bid",
          "prodotto.bid": "$doc_prodotto._bid",
          "prodotto.text": "$doc_prodotto.name",
          "focal_point.bid": "$doc_magazzino.bid_focal_point"
        }
      }
    ],
    {

    }
)